// В основном файле (main.js)

const { Worker } = require('worker_threads');

// Создание нового рабочего потока
const worker = new Worker('./worker.js', { workerData: 'some data' });

// Обработка сообщений от рабочего потока
worker.on('message', message => {
  console.log('Получено сообщение от рабочего потока:', message);
});

// Обработка ошибок рабочего потока
worker.on('error', error => {
  console.error('Ошибка в рабочем потоке:', error);
});

// Обработка завершения рабочего потока
worker.on('exit', exitCode => {
  console.log(`Рабочий поток завершился с кодом ${exitCode}`);
});

// Отправка сообщения в рабочий поток
worker.postMessage('Hello from the main thread!');
