# MicroExpress

Мы не знаем лучше способа разобраться с чем-то, чем попытаться воспроизвести это самостоятельно. Сегодня мы с вами напишем упрощённый аналог express.js

В файле `src/microexpress.js` в качестве экспорта реализовать функцию, создающую ваше приложение (аналог `const app = express()`); У вашего созданного приложения должны быть два метода:

- `use(middleware)` - регистрирует новую middleware для обработки. Если сигнатура функции содержит менее 3 аргументов, то эта функция будет получать на вход параметры `(req, res, next)`, где:

  - `req` - объект запроса для текущего HTTP-запроса
  - `res` - объект ответа для текущего HTTP-запроса
  - `next` - функция, которую необходимо вызывать для передачи управления следующей миддлвари. Функция `next` вызывается без аргументов в случае успешного завершения и с ошибкой - в случае неуспешного (классический node-style callback).

  Если же функция, переданная в `use` содержит более 3 аргументов, то в дополнение к 3 вышеперечисленным самым первым аргументом функции приходит `err` - объект ошибки

- `handler()` - возвращает обработчик, который принимает на вход `req` и `res` (за примером использования смотрите `src/index.js`)

Любая синхронная ошибка, возникшая в одной из middleware, должна быть корректно перехвачена и приводить к запуску цепочки middleware обработки ошибок. В остальном middleware ведут себя аналогично express - т.е. middleware должны вызываться строго в порядке добавления их через `use`, следующая middleware запускается только после того, как предыдущая вызвала `next`;

В файле `src/microrouter.js` в качестве экспорта реализовать функцию, создающую роутер. Роутер имеет три метода:

- `get(url, handler)` - регистрирует обработчик для GET-запроса на урл `url`
- `post(url, handler)` - регистрирует обработчик для POST-запроса на урл `url`
- `middleware()` - возвращает middleware, которое мы можем подключить в наше приложение

`handler` - функция. Если функция принимает менее трех параметров, то мы считаем эту функцию синхронной и запускаем обработчики роутов после выхода из нее. В случае если функция принимает три и более параметров - считаем что функция асинхронная и вызывает `next()`, который приходит третьим параметром аналогично поведению middleware

URL могут быть простой строкой вида `/foo/test`, так и содержать параметры вида `/foo/:some`, где `:some` это параметр. Значение таких параметров должны прийти в обработчик роутера в `req.params`. Так при урле `/foo/bar` содержимое `params` будет `{ some: "bar" }`. Обработка ошибок в роутах должна происходить аналогично обработке ошибок в middleware. Вы можете использовать библиотеку `url-pattern` для реализации разбора таких URL.

Вы можете быть уверены, что в качестве обработчиков и в `use` и в `get`, `post` всегда приходят корректные функции.
