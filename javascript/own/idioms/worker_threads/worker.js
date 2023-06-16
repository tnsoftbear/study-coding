// Воркер-файл (worker.js)

const { parentPort, workerData } = require('worker_threads');

// Обработка сообщений от основного потока
parentPort.on('message', message => {
  console.log('Получено сообщение от основного потока:', message);
  
  // Отправка сообщения обратно в основной поток
  parentPort.postMessage('Hello from the worker thread!');
});

// Вывод данных, переданных воркеру при создании
console.log('Данные, переданные воркеру:', workerData);
