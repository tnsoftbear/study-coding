fetchDataFromServer()
  .then((data) => filterData(data))
  .then((filteredData) => sortData(filteredData))
  .then((sortedData) => printData(sortedData))
  .catch((err) => console.error("Catch error: ", err));

function fetchDataFromServer() {
  return new Promise((resolve, reject) => {
    if (true) resolve("From the Muddy Banks Of The Wishkah");
    else reject("Data fetching failed");
  });
}

function filterData(data) {
  if (true) return filterVowels(data);
  else throw new Error("Data filtering failed");
}

function sortData(data) {
  return new Promise((resolve, reject) => {
    if (true) resolve(data.split("").sort().join(""));
    else reject("Data sorting failed");
  });
}

function printData(data) {
  console.log("Success: ", data.toString());
}

function filterVowels(text) {
  let filtered = "";
  let vowels = "aeiouAEIOU ";
  for (let i = 0; i < text.length; i++) {
    if (!vowels.includes(text[i])) {
      filtered += text[i];
    }
  }
  return filtered;
}
