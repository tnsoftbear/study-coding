const crypto = require("crypto");
const start = Date.now();

ll("Start cycle");
for (i = 0; i < 20; i++) {
  let profileCb = profileClosure(i);
  crypto.pbkdf2(
    Date.now().toString(),        // password
    (Date.now() + 5).toString(),  // salt
    100000,                       // iterations
    512,                          // keylen
    "sha512",                     // digest
    profileCb                     // callback
  );
}
ll("End cycle");

function profileClosure(i) {
  return function profile() {
    ll("call#: %d,\ttime: %d", i, Date.now() - start);
  };
}

function ll(...args) {
  console.log(...args);
}

// Пул потоков (Worker pool) предоставляет планировщику потоков 4 потока, поэтому первые 4 вызова будут выполнены сразу, а остальные 16 будут ждать освобождения потоков

// Start cycle
// End cycle
// call#: 0,       time: 2324
// call#: 1,       time: 2350
// call#: 2,       time: 2370
// call#: 3,       time: 2408
// call#: 4,       time: 4871
// call#: 5,       time: 4918
// call#: 6,       time: 4999
// call#: 7,       time: 5010
// call#: 9,       time: 8096
// call#: 10,      time: 8114
// call#: 8,       time: 8186
// call#: 11,      time: 8228
// call#: 13,      time: 10422
// call#: 12,      time: 10439
// call#: 14,      time: 10471
// call#: 15,      time: 10512
// call#: 16,      time: 12869
// call#: 17,      time: 12935
// call#: 19,      time: 12995
// call#: 18,      time: 13008

// https://www.youtube.com/watch?v=zDlg64fsQow 57 :00
// https://nodejs.org/en/docs/guides/dont-block-the-event-loop