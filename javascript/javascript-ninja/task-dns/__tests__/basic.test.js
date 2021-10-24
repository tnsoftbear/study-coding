const { promisify } = require("util");
const proxyquire = require("proxyquire");

/*
Illya Klymov: формально то что перед вами - это не честные
Unit-тесты. Представьте, что во во время тестов 1.1.1.1 не доступен.
Или просто возникли проблемы с соединением. Как следствие - тест провалится,
хотя проблема не в вашем коде.

Тем не менее давать вам полноценную тестовую инфраструктуру нам не хочется,
по очевидным причинам. Да и для ручной проверки, все ли ок таких тестов достаточно.
На CI эти тесты не запускаются, чтоб гарантированно корректно тестировать ваш код
*/

const { resolve4, resolve6, setResolveServer } = proxyquire
  .noCallThru()
  .load("../dns", {
    dns: null,
    child_process: null,
    http: null
  });

const promisifiedResolve4 = promisify(resolve4);
const promisifiedResolve6 = promisify(resolve6);

it("Should correctly resolve v4 of javascript.ninja", async () => {
  if (process.env.CI) {
    return;
  }

  setResolveServer("1.1.1.1");
  const result = await promisifiedResolve4("javascript.ninja");
  expect(result.length).toBe(1);
  expect(result[0]).toBe("185.203.72.17");
});

it("Should correctly resolve v6 of javascript.ninja", async () => {
  if (process.env.CI) {
    return;
  }

  setResolveServer("1.1.1.1");
  const result = await promisifiedResolve6("ipv6.test-ipv6.ams.vr.org");
  expect(result.length).toBe(1);
  expect(result[0]).toBe("2607:f740:d::f77");
});
