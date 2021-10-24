const microrouter = require("../src/microrouter");

it("microrouter должен экспортировать функцию", () => {
  expect(typeof microrouter).toBe("function");
});
