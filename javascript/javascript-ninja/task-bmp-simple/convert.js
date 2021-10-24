const { InvalidImageError } = require("./errors");

/**
 * @param {Buffer} input
 * @returns {Number}
 */
function detectWidth(input) {
  const width = input.readInt32LE(18);
  return width;
}

/**
 * @param {Buffer} input
 * @returns {Number}
 */
function detectHeight(input) {
  const height = input.readInt32LE(22);
  return height;
}

/**
 * @param {Buffer} input
 * @returns {Number}
 */
function detectSize(input) {
  const size = input.readUInt32LE(2);
  return size;
}

/**
 * @param {Buffer} input
 * @returns {Number}
 */
function detectOffset(input) {
  const dataOffset = input.readInt32LE(10);
  return dataOffset;
}

/**
 * @param {Buffer} input
 * @returns {Boolean}
 */
function validateHeader(input) {
  if (input.length < 40) {
    throw new InvalidImageError();
  }

  const size = detectSize(input);
  if (size !== input.length) {
    throw new InvalidImageError();
  }

  const width = detectWidth(input);
  const height = detectHeight(input);
  if (width > 2000 || height > 2000 || width <= 0 || height <= 0) {
    throw new InvalidImageError();
  }

  const format = input.toString("ascii", 0, 2);
  if (format !== "BM") {
    throw new InvalidImageError();
  }

  const dibHeaderLength = input.readInt32LE(14);
  if (dibHeaderLength < 40) {
    throw new InvalidImageError();
  }

  const bitsPerPixel = input.readInt16LE(28);
  if (bitsPerPixel !== 24) {
    throw new InvalidImageError();
  }

  const dataOffset = detectOffset(input);
  if (dataOffset > size || dataOffset < dibHeaderLength) {
    throw new InvalidImageError();
  }

  /**
   const dataActualSize = size - dataOffset;
   const bytesPerPixel = 3;
   if (dataActualSize !== width * height * bytesPerPixel) {
    throw new InvalidImageError();
  }

   const dataSize = input.readInt32LE(34);
   if (dataSize !== 0 && dataSize !== dataActualSize) {
    throw new InvalidImageError();
  }
   */

  return true;
}

/**
 * @param {Buffer} source
 * @returns {Buffer}
 */
function convertImage(source) {
  validateHeader(source);
  const bytePerPixel = 3;
  const width = detectWidth(source);
  const height = detectHeight(source);
  const rowPadding = (4 - ((width * bytePerPixel) % 4)) % 4;
  const byteWidth = bytePerPixel * width + rowPadding;
  const size = detectSize(source);
  const dataStart = detectOffset(source);
  const target = Buffer.alloc(size, 0);
  source.copy(target, 0, 0, dataStart);
  for (let y = 0; y < height; y += 1) {
    const rowStart = dataStart + y * byteWidth;
    for (let x = 0; x < width; x += 1) {
      const sourceStart = rowStart + x * bytePerPixel;
      const targetStart =
        rowStart + width * bytePerPixel - (x + 1) * bytePerPixel;
      source.copy(target, targetStart, sourceStart, sourceStart + 3);
    }
  }
  return target;
}

module.exports = convertImage;
