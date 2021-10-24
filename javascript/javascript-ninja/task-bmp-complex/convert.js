const { Transform } = require("stream");
const { InvalidImageError } = require("./errors");

const BYTE_PER_PIXEL = 3;

const FILE_HEADER_SIZE = 16;
const DIB_HEADER_SIZE = 40;

const FILE_HEADER = 1;
const DIB_HEADER = 2;
const OTHER_HEADER = 3;
const IMAGE_DATA = 4;
const TRAIL_DATA = 5;
const EOF_TIME = 6;

const RESULT_COMPLETE_OVERLOAD = 1;
const RESULT_COMPLETE_FULL = 2;
const RESULT_INCOMPLETE = 3;

// eslint-disable-next-line no-unused-vars
function ll(...args) {
  global.console.log(JSON.stringify(args));
}

class MirrorStream extends Transform {
  constructor(opt) {
    // opt = {highWaterMark: 15};
    super(opt);
    const maxBufLength = Math.max(
      this._readableState.highWaterMark,
      this._writableState.highWaterMark,
      DIB_HEADER_SIZE
    );
    this.outputBuf = Buffer.allocUnsafe(maxBufLength * 2);
    this.outputBufLength = 0;
    this.inputBuf = Buffer.allocUnsafe(maxBufLength);
    this.inputBufLength = 0;
    this.correctedBuf = null;
    this.header = null;
    this.readLength = 0;
    this.headerErrorMessage = null;
    this.encoding = null;
    // this.isPause = false;
    this.targetPixel = Buffer.alloc(3);
    this.readState = null;
    this.on("error", () => {});
    // this.on("drain", () => {
    //   if (this.isPause) {
    //     // ll("un-pause");
    //     this.isPause = false;
    //   }
    // });
  }

  transforming(inputBuf, encoding, doneCb) {
    // if (this.isPause) {
    //   return false;
    // }
    this.encoding = encoding;
    this.setInputBuf(inputBuf);
    this.traceReadLength();
    this.doTransform(doneCb);
    doneCb();
    return true;
  }

  /**
   * @param {Buffer} inputBuf
   * @param {String} encoding
   * @param {Function} doneCb
   * @private
   */
  _transform(inputBuf, encoding, doneCb) {
    // setImmediate(() => {
    if (!this.transforming(inputBuf, encoding, doneCb)) {
      this._transform(inputBuf, encoding, doneCb);
    }
    // });
  }

  _flush(doneCb) {
    if (!this.validateComplete()) {
      this.push(null);
      doneCb(new InvalidImageError(this.headerErrorMessage));
    }
    doneCb();
  }

  doTransform(doneCb) {
    do {
      // const [blockEndOffset, prePushCb] = this.detectOffsetAndHandler(doneCb);
      this.process(...this.detectOffsetAndHandler(doneCb));
      if (this.isPortionOverloaded()) {
        this.readStartOffset = this.readLength - this.getInputBufLength();
      }
    } while (this.getInputBufLength() > 0);
  }

  detectOffsetAndHandler(doneCb) {
    this.readState = this.detectReadState();

    if (this.readState === FILE_HEADER || this.readState === DIB_HEADER) {
      const handleDibHeader = () => {
        this.setHeader(this.getOutputBuf());
        if (!this.validateHeader()) {
          doneCb(new InvalidImageError(this.headerErrorMessage));
        }
      };
      return [DIB_HEADER_SIZE, handleDibHeader];
    }

    if (this.readState === OTHER_HEADER) {
      return [this.header.imageDataOffset, null];
    }

    if (this.readState === IMAGE_DATA) {
      // const lineCount = Math.floor(this.getInputBufLength() / this.header.lineWidthByte) || 1;
      const lineCount = 1;
      const linesEndOffset =
        this.readStartOffset -
        this.getOutputBufLength() +
        lineCount * this.header.lineWidthByte;
      const handleImageData = () => this.flipLines(lineCount);
      return [linesEndOffset, handleImageData];
    }

    if (this.readState === TRAIL_DATA) {
      return [this.header.fileSize, null];
    }

    return [];
  }

  process(blockOffset, prePushCb) {
    this.collectOutputBuf(blockOffset);
    if (this.isPushTime()) {
      if (prePushCb) {
        prePushCb();
      }
      this.push(Buffer.from(this.getOutputBuf()));
      // this.push(this.getOutputBuf().slice(0, this.getOutputBufLength()));
      // this.isPause = !this.push(Buffer.from(this.getOutputBuf()));
      // if (this.isPause) {
      //   ll("paused");
      // }
      // this.isPause = false;
      this.dropOutputBuf();
    }
  }

  /**
   * @param blockOffset
   * @returns {null|*}
   */
  collectOutputBuf(blockOffset) {
    if (this.readLength > blockOffset) {
      this.outputBufStatus = RESULT_COMPLETE_OVERLOAD;
      const correctedEnd = blockOffset - this.readStartOffset;
      this.addOutputBuf(this.getInputBuf().slice(0, correctedEnd));
      this.setInputBuf(this.getInputBuf().slice(correctedEnd));
      return;
    }

    if (this.readLength === blockOffset) {
      this.outputBufStatus = RESULT_COMPLETE_FULL;
    } else {
      this.outputBufStatus = RESULT_INCOMPLETE;
    }

    this.addOutputBuf(this.getInputBuf());
    this.dropInputBuf();
  }

  /**
   * @param {Buffer} buf
   */
  addOutputBuf(buf) {
    buf.copy(this.outputBuf, this.outputBufLength, 0, buf.length);
    this.outputBufLength += buf.length;
  }

  /**
   * @returns {Buffer}
   */
  getOutputBuf() {
    return this.outputBuf.slice(0, this.outputBufLength);
  }

  getOutputBufLength() {
    return this.outputBufLength;
  }

  dropOutputBuf() {
    this.outputBufLength = 0;
  }

  /**
   * @param {Buffer} buf
   */
  setInputBuf(buf) {
    buf.copy(this.inputBuf, 0, 0, buf.length);
    this.inputBufLength = buf.length;
  }

  /**
   * @returns {Buffer}
   */
  getInputBuf() {
    return this.inputBuf.slice(0, this.inputBufLength);
  }

  getInputBufLength() {
    return this.inputBufLength;
  }

  dropInputBuf() {
    this.inputBufLength = 0;
  }

  isPortionOverloaded() {
    return this.outputBufStatus === RESULT_COMPLETE_OVERLOAD;
  }

  isPushTime() {
    return (
      this.outputBufStatus === RESULT_COMPLETE_FULL ||
      this.isPortionOverloaded()
    );
  }

  setHeader(headerBuf) {
    this.header = {
      // for processing
      width: MirrorStream.detectWidth(headerBuf),
      height: MirrorStream.detectHeight(headerBuf),
      fileSize: MirrorStream.detectFileSize(headerBuf),
      imageDataOffset: MirrorStream.detectImageDataOffset(headerBuf),
      dibHeaderLength: MirrorStream.detectDipHeaderLength(headerBuf),
      lineWidthByte: MirrorStream.detectLineWidthByte(headerBuf),
      trailDataOffset: MirrorStream.detectTrailDataOffset(headerBuf),
      // for validation
      format: MirrorStream.detectFormat(headerBuf),
      bitsPerPixel: MirrorStream.detectBidsPerPixel(headerBuf)
    };
    // ll(this.header);
  }

  traceReadLength() {
    this.readStartOffset = this.readLength;
    this.readLength += this.getInputBufLength();
    // ll('Start / Length', this.readStartOffset, this.readLength);
  }

  detectReadState() {
    if (this.readStartOffset < FILE_HEADER_SIZE) {
      return FILE_HEADER;
    }
    if (this.readStartOffset < DIB_HEADER_SIZE) {
      return DIB_HEADER;
    }
    if (this.readStartOffset < this.header.imageDataOffset) {
      return OTHER_HEADER;
    }
    if (this.readStartOffset < this.header.trailDataOffset) {
      return IMAGE_DATA;
    }
    if (this.readStartOffset < this.header.fileSize) {
      return TRAIL_DATA;
    }
    return EOF_TIME;
  }

  /**
   * Flip lines directly in source (mutate argument)
   * @param {Number} lineCount
   * @returns {Buffer}
   */
  flipLines(lineCount) {
    for (let y = 0; y < lineCount; y += 1) {
      this.flipLine(
        this.getOutputBuf().slice(
          y * this.header.lineWidthByte,
          (y + 1) * this.header.lineWidthByte
        )
      );
    }
  }

  /**
   * @param {Buffer} source
   * @returns {Buffer}
   */
  flipLine(source) {
    let sourceStart = 0;
    let targetStart = 0;
    for (let x = 0; x < Math.floor(this.header.width / 2); x += 1) {
      sourceStart = x * BYTE_PER_PIXEL;
      targetStart =
        this.header.width * BYTE_PER_PIXEL - (x + 1) * BYTE_PER_PIXEL;
      source.copy(
        this.targetPixel,
        0,
        targetStart,
        targetStart + BYTE_PER_PIXEL
      );
      source.copy(
        source,
        targetStart,
        sourceStart,
        sourceStart + BYTE_PER_PIXEL
      );
      this.targetPixel.copy(source, sourceStart, 0, BYTE_PER_PIXEL);
    }
  }

  /**
   * @returns {Boolean}
   */
  validateComplete() {
    if (!this.validateHeader()) {
      return false;
    }
    if (this.readLength < this.header.fileSize) {
      this.headerErrorMessage = '"File size" does not correspond real size';
      return false;
    }
    return true;
  }

  /**
   * @returns {Boolean}
   */
  validateHeader() {
    const headerJson = JSON.stringify(this.header);

    if (!this.header) {
      this.headerErrorMessage = `Cannot read header from file`;
      return false;
    }

    if (
      this.header.width > 8000 ||
      this.header.height > 8000 ||
      this.header.width <= 0 ||
      this.header.height <= 0
    ) {
      this.headerErrorMessage = `Incorrect "width" or "height" in image header ${headerJson}`;
      return false;
    }

    if (this.header.format !== "BM") {
      this.headerErrorMessage = `Incorrect "format" in image header ${headerJson}`;
      return false;
    }

    if (this.header.dibHeaderLength < 40) {
      this.headerErrorMessage = `Incorrect "dib header length" in image header ${headerJson}`;
      return false;
    }

    if (this.header.bitsPerPixel !== 24) {
      this.headerErrorMessage = `Incorrect "bits per pixel" in image header ${headerJson}`;
      return false;
    }

    if (
      this.header.imageDataOffset > this.header.fileSize ||
      this.header.imageDataOffset < this.header.dibHeaderLength
    ) {
      this.headerErrorMessage = `Incorrect "image data offset" in image header ${headerJson}`;
      return false;
    }

    return true;
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {String}
   */
  static detectFormat(headerBuf) {
    return headerBuf.toString("ascii", 0, 2);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectWidth(headerBuf) {
    return headerBuf.readInt32LE(18);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectHeight(headerBuf) {
    return headerBuf.readInt32LE(22);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectFileSize(headerBuf) {
    return headerBuf.readUInt32LE(2);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectImageDataOffset(headerBuf) {
    return headerBuf.readInt32LE(10);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectDipHeaderLength(headerBuf) {
    return headerBuf.readInt32LE(14);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectBidsPerPixel(headerBuf) {
    return headerBuf.readInt16LE(28);
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectTrailDataOffset(headerBuf) {
    const sizeByte = MirrorStream.detectImageDataSizeByte(headerBuf);
    const imageDataOffset = MirrorStream.detectImageDataOffset(headerBuf);
    return imageDataOffset + sizeByte;
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectImageDataSizeByte(headerBuf) {
    const height = MirrorStream.detectHeight(headerBuf);
    const lineWidthByte = MirrorStream.detectLineWidthByte(headerBuf);
    return lineWidthByte * height;
  }

  /**
   * @param {Buffer} headerBuf
   * @returns {Number}
   */
  static detectLineWidthByte(headerBuf) {
    const bytePerPixel = BYTE_PER_PIXEL;
    const width = MirrorStream.detectWidth(headerBuf);
    const rowPadding = (4 - ((width * bytePerPixel) % 4)) % 4;
    return bytePerPixel * width + rowPadding;
  }
}

module.exports = MirrorStream;
