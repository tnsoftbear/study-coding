// const OS = require('os')
// process.env.UV_THREADPOOL_SIZE = OS.cpus().length

process.env.UV_THREADPOOL_SIZE = 8;
require("./node_default_thread_pool.js");