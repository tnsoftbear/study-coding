// Listen for messages from the main page
onmessage = (event) => {
 console.log(event.data);
 postMessage({ message: 'Received: ' + event.data.message + '. Sent: Hello from the worker' });
}
