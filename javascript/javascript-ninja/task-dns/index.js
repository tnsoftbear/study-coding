const { resolve4, resolve6, setResolveServer } = require("./dns");

setResolveServer("1.1.1.1:53");
resolve4("cloudflare.com", (err, result) => {
  global.console.log("v4", err, result);
});

resolve6("www.google.com", (err, result) => {
  global.console.log("v6", err, result);
});
