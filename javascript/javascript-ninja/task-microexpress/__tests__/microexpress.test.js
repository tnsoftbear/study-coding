const microexpress = require("../src/microexpress");

it("microexpress должен экспортировать функцию", () => {
  expect(typeof microexpress).toBe("function");
});
