function recursivePromise(promise) {
    promise.then(() => {
        console.log("promise");
        recursivePromise(Promise.resolve());
    });
}

recursivePromise(Promise.resolve());

// Следующий колбек никогда не вызовется, так как стек постоянно заполнен колбеками из очереди промисов
setTimeout(() => console.log("setTimeout"), 0);

// Альтернативный вариант бесконечного цикла заполнения очереди промисов, так что колбеки из очереди промисов никогда не будут вызваны
// for (let i = 0; i < 1000000000; i++) {
//     console.log(i);
//     Promise.resolve().then(() => console.log("promise"));
// }