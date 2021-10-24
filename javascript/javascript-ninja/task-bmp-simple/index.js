/* global process */
const fs = require("fs");
const convert = require("./convert");

const inputFile = process.argv[2];
const outputFile = process.argv[3];
if (!inputFile) {
  global.console.error("CLI usage: node index.js input.bmp output.bmp");
  process.exit(1);
}

fs.readFile(inputFile, (err, data) => {
  if (err) {
    global.console.error(err);
    process.exit(1);
  }

  fs.writeFile(outputFile, convert(data), writeError => {
    if (writeError) {
      global.console.error(err);
      process.exit(1);
    }
  });
});
