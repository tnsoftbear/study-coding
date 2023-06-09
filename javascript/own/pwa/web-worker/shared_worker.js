var port = null;

self.onconnect = function(event) {
  port = event.ports[0];
  port.onmessage = function(event) {
    console.log('Received message from window:', event.data);
    port.postMessage('Hello from shared_worker.js!');
  };
};