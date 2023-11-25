/**
node recursive_closure_mem_leak.js

Посмотреть как работает GC:
node --trace-gc recursive_closure_mem_leak.js

Эксперементируем с max-old-space-size и max-semi-space-size:
node --max-semi-space-size=1500 --trace-gc recursive_closure_mem_leak.js

Крешим с "heap out of memory":
node --max-old-space-size=15 --trace-gc recursive_closure_mem_leak.js
*/

const memory = [];
const bytesToMb = (bytes) => Math.round(bytes / 1024, 2) / 1024;
const recursiveClosure = (a) => (fn) => recursiveClosure(a.map((g) => fn(g)));
let f = recursiveClosure(new Array(1000).fill((x) => x * 2));
const timer = setInterval(() => {
    f = f((fn) => (x) => fn(x) * 2);
}, 5);
setInterval(() => {
    const usage = process.memoryUsage();
    const row = {
        rss: bytesToMb(usage.rss),
        heapTotal: bytesToMb(usage.heapTotal),
        heapUsed: bytesToMb(usage.heapUsed),
        external: bytesToMb(usage.external),
        stack: bytesToMb(usage.rss - usage.heapTotal),
    };
    memory.push(row);
    console.table(memory);
}, 1000);
setTimeout(() => {
    clearInterval(timer);
}, 10000);