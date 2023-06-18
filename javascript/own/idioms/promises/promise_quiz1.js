let promise = new Promise((resolve, reject) => {
    resolve(1);

    setTimeout(() => resolve(2), 1000);
});

promise.then(console.log);

// Output: 1