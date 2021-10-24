const path = require("path");
const fs = require("fs");

const convert = require("../convert");

it("корректно преобразовывает тестовое изображение", () => {
  const input = fs.readFileSync(path.resolve(__dirname, "assets", "input.bmp"));
  const output = fs.readFileSync(
    path.resolve(__dirname, "assets", "output.bmp")
  );
  expect(convert(input)).toEqual(output);
});
