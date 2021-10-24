/* eslint-disable no-bitwise */
const dgram = require("dgram");
const ip6addr = require("ip6addr");

const RESOLVE_TIMEOUT = 2000;

const CLASS_IN = 1;

const TYPE_A = 1;
const TYPE_AAAA = 28;
//
const QR_QUERY = 0;
const QR_RESPONSE = 1;
//
const RCODE_OK = 0;
const RCODE_SERVFAIL = 2;
const RCODE_NXDOMAIN = 3;
const RCODE_NOT_IMPLEMENTED = 4;
const RCODE_REFUSED = 5;
//
const OPCODE_QUERY = 0;

const HEADER_SIZE = 12;
// const RECURSION_DESIRED = 1;

let serverAddress = null;

const clients = [];

// eslint-disable-next-line no-unused-vars
function ll(...args) {
  if (args) {
    global.console.log(args);
  }
}

// eslint-disable-next-line no-unused-vars
function dec2bin(dec) {
  return Number(dec).toString(2);
}

function dec2hex(dec) {
  return Number(dec).toString(16);
}

function generateQueryId() {
  return Math.round(Math.random() * 1000);
}

function setFlagQr(flags, isResponse) {
  flags |= isResponse << 15;
  return flags;
}

function setOpCode(flags, opcode) {
  flags |= opcode << 11;
  return flags;
}

function isPointer(metaByte) {
  const type = (metaByte & 0b11000000) >> 6;
  return type === 0b11;
}

function readLabelSequence(response, offset) {
  const labels = [];
  let size = 0;
  do {
    if (response.length < offset + 1) {
      throw new Error("Incorrect offset to read label size byte");
    }
    size = response.readUInt8(offset);
    offset += 1;
    if (size > 0) {
      if (response.length < offset + size) {
        throw new Error("Incorrect label size to read it from response");
      }
      const label = response.toString("ascii", offset, offset + size);
      labels.push(label);
      offset += size;
    }
  } while (size);
  return labels.join(".");
}

function checkId(response, queryId) {
  if (response.length < 2) {
    return false;
  }
  const rId = response.readUInt16BE(0);
  return rId === queryId;
}

/**
 * @param {Buffer} response
 * @param {Number} queryId
 * @param {Function} readCb
 * @returns {*}
 */
function readResponse(response, queryId, readCb) {
  if (response.length < HEADER_SIZE) {
    readCb(new Error(`Response size too small ${response.length}`));
    return;
  }

  const flags = response.readUInt16BE(2);
  const isInvalidFlag = flags & 0b0000000001110000;
  if (isInvalidFlag) {
    readCb(new Error("Invalid flag Z - must be zero"));
    return;
  }

  const qr = flags >>> 15;
  if (qr !== QR_RESPONSE) {
    readCb(new Error("Response QR should be 1"));
    return;
  }

  const opCode = (flags & 0b0111100000000000) >>> 11;
  if (opCode !== OPCODE_QUERY) {
    readCb(new Error("OpCode is not standard query"));
    return;
  }

  const rCode = flags & 0b1111;
  if (rCode === RCODE_NXDOMAIN) {
    // readCb(new Error("Non-existent Domain Name");
    readCb(null, []);
    return;
  }
  if (rCode === RCODE_REFUSED) {
    readCb(new Error("REFUSED"));
    return;
  }
  if (rCode === RCODE_NOT_IMPLEMENTED) {
    readCb(new Error("NOT IMPLEMENTED"));
    return;
  }
  if (rCode === RCODE_SERVFAIL) {
    readCb(new Error("SERVFAIL"));
    return;
  }
  if (rCode !== RCODE_OK) {
    readCb(new Error(`Response failed. Response code: ${rCode}`));
    return;
  }

  const questionCount = response.readUInt16BE(4);
  if (questionCount !== 1) {
    const errorMessage = `Response failed. Question count should be 1. It is ${questionCount}`;
    readCb(new Error(errorMessage));
    return;
  }
  const answerCount = response.readUInt16BE(6);
  // const authorityRr = response.readUInt16BE(8);
  // const additionalRr = response.readUInt16BE(10);
  // ll(questionCount, answerCount, authorityRr, additionalRr);

  let offset = HEADER_SIZE;

  const questionSizeMin = 9;
  const responseSizeMin = offset + questionCount * questionSizeMin;
  if (response.length < responseSizeMin) {
    const errorMessage = `Response question section size broke (Response length ${response.length}, assumed min size: ${questionSizeMin})`;
    readCb(new Error(errorMessage));
    return;
  }

  const questions = [];
  for (let i = 0; i < questionCount; i += 1) {
    const domainOffset = offset;

    let qDomainName = "";
    try {
      qDomainName = readLabelSequence(response, offset, readCb);
    } catch (e) {
      readCb(e);
      return;
    }

    const size = qDomainName.length + 2; // leading size byte and trailing zero byte
    offset += size;

    const qType = response.readUInt16BE(offset);
    offset += 2;
    const qClass = response.readUInt16BE(offset);
    offset += 2;
    questions.push({
      offset: domainOffset,
      name: qDomainName,
      type: qType,
      class: qClass
    });
  }

  // const addressSectionSize = clients[queryId].queryType === TYPE_A ? 14 : 26;
  // const assumedMinimalSize = offset + addressSectionSize * answerCount;
  // if (response.length < assumedMinimalSize) {
  //   const errorMessage = `Address section is broken. Response size is not enough (${response.length} < ${assumedMinimalSize})`;
  //   readCb(new Error(errorMessage));
  //   return;
  // }
  const addresses = [];
  for (let i = 0; i < answerCount; i += 1) {
    if (response.length < offset + 1) {
      readCb(new Error("Response answer section broken at metaByte field"));
      return;
    }
    const metaByte = response.readUInt8(offset);
    // offset += 1; // don't add offset
    let aDomainName;
    if (isPointer(metaByte)) {
      if (response.length < offset + 2) {
        readCb(new Error("Response answer section broken at POINTER field"));
        return;
      }
      let pointerOffset = response.readUInt16BE(offset);
      offset += 2;
      pointerOffset &= 0b0011111111111111;
      try {
        aDomainName = readLabelSequence(response, pointerOffset);
      } catch (e) {
        readCb(e);
        return;
      }
    } else {
      try {
        aDomainName = readLabelSequence(response, offset, readCb);
        offset += aDomainName.length + 2;
      } catch (e) {
        readCb(e);
        return;
      }
    }

    if (response.length < offset + 2) {
      readCb(new Error("Response answer section broken at TYPE field"));
      return;
    }
    const aType = response.readUInt16BE(offset);
    offset += 2;
    // if (aType !== TYPE_A && aType !== TYPE_AAAA) {
    //   // readCb(new Error(`Incorrect query type in response ${aType}`));
    //   readCb(null, []);
    //   return;
    // }

    if (response.length < offset + 2) {
      readCb(new Error("Response answer section broken at CLASS field"));
      return;
    }
    const aClass = response.readUInt16BE(offset);
    offset += 2;

    if (response.length < offset + 4) {
      readCb(new Error("Response answer section broken at TTL field"));
      return;
    }
    const aTtl = response.readUInt32BE(offset);
    offset += 4;

    if (response.length < offset + 2) {
      readCb(new Error("Response answer section broken at RDLENGTH field"));
      return;
    }
    const aDataLength = response.readUInt16BE(offset);
    offset += 2;
    if (aDataLength !== 4 && aDataLength !== 16) {
      const errorMessage = `RDLENGTH incorrect value ${aDataLength} - should be 4 or 16`;
      readCb(new Error(errorMessage));
      return;
    }
    if (response.length < offset + aDataLength) {
      const errorMessage = `Response actual size is smaller than assumed by RDLENGTH ${aDataLength}`;
      readCb(new Error(errorMessage));
      return;
    }

    const ipNums = [];
    if (aType === TYPE_A) {
      for (let j = 0; j < aDataLength; j += 1) {
        ipNums.push(response.readUInt8(offset));
        offset += 1;
      }
    } else if (aType === TYPE_AAAA) {
      for (let j = 0; j < aDataLength / 2; j += 1) {
        ipNums.push(dec2hex(response.readUInt16BE(offset)));
        offset += 2;
      }
    }

    let aIpRaw;
    let aIp;
    if (aType === TYPE_A) {
      aIpRaw = ipNums.join(".");
    } else {
      aIpRaw = ipNums.join(":");
    }
    try {
      aIp = ip6addr.parse(aIpRaw).toString();
    } catch (e) {
      readCb(new Error(`Incorrect ip address ${aIpRaw}`));
      return;
    }
    const address = {
      name: aDomainName,
      type: aType,
      class: aClass,
      ttl: aTtl,
      length: aDataLength,
      ip: aIp
    };
    addresses.push(address);
  }

  const ips = [];
  for (let i = 0; i < addresses.length; i += 1) {
    ips.push(addresses[i].ip);
  }
  ips.sort();

  clearTimeout(clients[queryId].resolveTimeoutId);
  readCb(null, ips);
}

function writeRequest(hostname, type, queryId) {
  const request = Buffer.alloc(512);
  request.writeUInt16BE(queryId, 0);
  let flags = 0;
  flags = setFlagQr(flags, QR_QUERY);
  flags = setOpCode(flags, OPCODE_QUERY);
  request.writeUInt16BE(flags, 2);
  request.writeUInt16BE(1, 4); // qdcount should be 1 always
  request.writeUInt16BE(0, 6); // ancount
  request.writeUInt16BE(0, 8); // nscount
  request.writeUInt16BE(0, 10); // arcount
  let offset = 12;

  const labels = hostname.split(".");
  for (let i = 0, cnt = labels.length; i < cnt; i += 1) {
    request.writeUInt8(labels[i].length, offset);
    offset += 1;
    request.write(labels[i], offset, "ascii");
    offset += labels[i].length;
  }
  request.writeUInt8(0b0, offset);
  offset += 1;
  request.writeUInt16BE(type, offset);
  offset += 2;
  request.writeUInt16BE(CLASS_IN, offset);
  offset += 2;

  const realRequest = request.slice(0, offset);
  clients[queryId].client.send(realRequest, err => {
    if (err) {
      throw new Error(`Error on send request: ${err}`);
    }
  });
}

function resolve(hostname, queryType, readCb) {
  if (!serverAddress) {
    readCb(new Error("Dns server unknown"));
    return;
  }
  const queryId = generateQueryId();
  let [serverIp, serverPort] = serverAddress.split(":");
  serverIp += ""; // TODO
  serverPort = serverPort || 53;
  const client = dgram.createSocket("udp4");

  clients[queryId] = { serverIp, serverPort, client, queryType };
  client.connect(serverPort, serverIp, err => {
    if (err) {
      readCb(new Error(`Cannot connect: ${err}`));
      return;
    }
    writeRequest(hostname, queryType, queryId);
    clients[queryId].resolveTimeoutId = setTimeout(() => {
      client.close();
      readCb(new Error(`Timeout expired: ${RESOLVE_TIMEOUT}`));
    }, RESOLVE_TIMEOUT);
  });
  client.on("message", response => {
    if (!checkId(response, queryId)) {
      return;
    }
    readResponse(response, queryId, readCb);
    clearTimeout(clients[queryId].resolveTimeoutId);
    client.close();
  });
  client.on("close", () => {
    clearTimeout(clients[queryId].resolveTimeoutId);
  });
}

function resolve4(hostname, readCb) {
  resolve(hostname, TYPE_A, readCb);
}

function resolve6(hostname, readCb) {
  resolve(hostname, TYPE_AAAA, readCb);
}

function validateAddrV4(addr) {
  const res = addr.split(":");
  if (res.length === 0 || res.length > 2) {
    throw new Error("Address incorrect format");
  }
  const [ip, portStr] = res;
  try {
    ip6addr.parse(ip);
  } catch (e) {
    throw new Error("Ip address incorrect format");
  }
  if (portStr) {
    if (!/^\d+$/.test(portStr)) {
      throw new Error("Port incorrect format");
    }
    const port = parseInt(portStr, 10);
    if (port < 1 || port > 65535) {
      throw new Error("Port incorrect value");
    }
  }
  return true;
}

function setResolveServer(addr) {
  if (!validateAddrV4(addr)) {
    throw new Error("Invalid address of dns server. It should be IPv4 format");
  }
  serverAddress = addr;
}

module.exports = {
  resolve4,
  resolve6,
  setResolveServer
};
