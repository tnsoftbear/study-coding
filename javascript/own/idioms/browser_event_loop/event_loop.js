function ll(...args) {
  console.log(...args);
}

setTimeout(() => {
  ll("setTimeout A");
  Promise.resolve().then(() => {
    ll("promise in setTimeout A");
  });
  queueMicrotask(() => {
    ll("queueMicrotask in setTimeout A");
  });
}, 0);

setTimeout(() => {
  ll("setTimeout B");
}, 0);

Promise.resolve().then(() => {
  ll("Promise X");
});

queueMicrotask(() => {
  ll("queueMicrotask");
});

Promise.resolve().then(() => {
  ll("Promise Y");
});

Promise.resolve("Resolved value Z").then((v) => ll("Promise Z", v));

Promise.resolve(ll("Console log in resolve W")).then(() => ll("Promise W"));

// Броузерный Event Loop возвращает результаты аналогично серверному Node.js (однако, тут нет setImmediate())

// Console log in resolve W
// Promise X
// queueMicrotask
// Promise Y
// Promise Z Resolved value Z
// Promise W
// setTimeout A
// promise in setTimeout A
// queueMicrotask in setTimeout A
// setTimeout B

// [Event Loop от А до Я. Архитектура браузера и Node JS. Движки и рендер.](https://www.youtube.com/watch?v=zDlg64fsQow)