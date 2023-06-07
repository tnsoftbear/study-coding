const storageKeyPrefix = "tnsoftbear-gpt-extension-";
const promptInput = document.getElementById("prompt-input");
const apiKeyInput = document.getElementById("api-key-input");
const generateButton = document.getElementById("generate-button");
const apiKeySaveButton = document.getElementById("api-key-save-button");
const outputArea = document.getElementById("output-area");

generateButton.addEventListener("click", function (event) {
  generateResponse(event);
});

apiKeySaveButton.addEventListener("click", function (event) {
  saveApiKey();
});

window.addEventListener("load", loadApiKey);

// Add an event listener to the input element to listen for when the user presses enter
// input.addEventListener("keydown", function (event) {
//   if (event.key === "Enter") {
//     generateResponse(event);
//   }
// });

async function generateResponse(event) {
  event.preventDefault();
  const apiKey = apiKeyInput.value;
  if (apiKey === "") {
    outputArea.innerHTML = "Please enter an API key";
    return;
  }
  
  outputArea.innerHTML = "Loading...";

  const prompt = promptInput.value;
  const response = await fetch("https://api.openai.com/v1/chat/completions", {
    method: "POST",
    headers: new Headers({
      "Content-Type": "application/json",
      Authorization: `Bearer ${apiKey}`,
    }),
    body: JSON.stringify({
      model: "gpt-3.5-turbo",
      messages: [{ role: "user", content: prompt }],
      max_tokens: 3800,
      temperature: 0.8,
    }),
  });

  // Parse the response and display it in a new window
  let textFormatted = "";
  const data = await response.json();
  if (data.error) {
    textFormatted = data.error.message;
  } else {
    const text = data.choices[0].message.content;
    textFormatted = replaceCodeBlocks(text);
  }
  const output = document.getElementById("output-area");
  output.innerHTML = textFormatted;
}

function replaceCodeBlocks(text) {
  const regex = /```(\w*)\s([\s\S]+?)```/g;
  const replacedText = text.replace(regex, "<pre>$2</pre>").replace(/\n/g, "</br>");
  return replacedText;
}

function saveApiKey() {
  var apiKeyInput = document.getElementById("api-key-input");
  var apiKey = apiKeyInput.value;
  localStorage.setItem(storageKeyPrefix + "api_key", apiKey);
}

// Загрузка значения из localStorage
function loadApiKey() {
  var apiKeyInput = document.getElementById("api-key-input");
  var savedApiKey = localStorage.getItem(storageKeyPrefix + "api_key");
  if (savedApiKey) {
    apiKeyInput.value = savedApiKey;
  } else {
    outputArea.innerHTML = "Please enter an API key. You can generate it <a target='_blank' href='https://platform.openai.com/account/api-keys'>here</a>";
  }
}
