const user = {
    username: "Ivan",
    address: {
        street: "Lenina"
    }
}
// Полностью скопирует внутренний объект
// Альтернатива: lodash.cloneDeep(user)
const copy = JSON.parse(JSON.stringify(user))
copy.address.street = "Pushkina"
copy.username = "Michael"
console.log(user.address.street, user.username) // "Lenina Ivan"

// Копирует ссылку на внутренний объект, поэтому изменения внутри копии отразятся на оригинале
// Shallow copy of an object (surface copying) = Неглубокая копия объекта (копирование поверхности)
const wrong_copy1 = {...user}
const wrong_copy2 = Object.assign({}, user)
wrong_copy1.address.street = "Pushkina"
wrong_copy1.username = "Petr"
console.log(user.address.street, user.username) // "Pushkina Ivan"
wrong_copy2.address.street = "MKAD"
wrong_copy2.username = "Sergey"
console.log(user.address.street, user.username) // "MKAD Ivan"

// Копирует ссылку на объект user
const pointer_copy = user;
pointer_copy.address.street = "Kirova"
pointer_copy.username = "Vladimir"
console.log(user.address.street, user.username) // "Kirova Vladimir"

// Альтернативная реализация глубокого копирования
const clone_copy = clone(user);
clone_copy.address.street = "Krasina"
clone_copy.username = "Alexey"
console.log(user.address.street, user.username) // "Kirova Vladimir"

function clone(object) {
  let result = {};
  for (let key in object) {
    let value = object[key];
    if (typeof value === 'object' && value !== null) {
      value = clone(value);
    }
    result[key] = value;
  }
  return result;
}