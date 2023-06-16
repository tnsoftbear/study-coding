const crypto = require("crypto");
const start = Date.now();
crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("1:", Date.now() - start);
});

crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("2:", Date.now() - start);
});

crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("3:", Date.now() - start);
});

crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("4:", Date.now() - start);
});

crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("5:", Date.now() - start);
});

crypto.pbkdf2("a", "b", 100000, 512, "sha512", () => {
  console.log("6:", Date.now() - start);
});

// Для 2х на 2х ядерном процессоре:
// 2: 1120
// 1: 1138

// Для 4х на 2х ядерном процессоре:
// 3: 2093
// 2: 2109
// 4: 2112
// 1: 2149

// Для 6х на 2х ядерном процессоре:
// 1: 2531
// 4: 2585
// 3: 2587
// 2: 2588
// 5: 3799
// 6: 3817
