const preEls = document.querySelectorAll("pre");

[...preEls].forEach((preEl) => {
  const btnCopy = document.createElement("button");
  btnCopy.innerText = "Copy";
  btnCopy.type = "button";

  // Обернём в шадоу рут, чтобы на кнопку распространялись только браузерные стили
  const divRoot = document.createElement("div");
  divRoot.style.position = "relative";
  const divShadowRoot = divRoot.attachShadow({ mode: "open" });
  const cssUrl = chrome.runtime.getURL("content-script.css");
  
  // UNSAFE_VAR_ASSIGNMENT   Unsafe assignment to innerHTML
  //divShadowRoot.innerHTML = "<link rel='stylesheet' href='" + cssUrl + "' />";
  const link = document.createElement("link");
  link.rel = "stylesheet";
  link.href = cssUrl;
  divShadowRoot.append(link);

  divShadowRoot.append(btnCopy);
  preEl.prepend(divRoot);

  const codeEl = preEl.querySelector("code");
  const code = codeEl.innerText;
  btnCopy.addEventListener("click", () => {
    navigator.clipboard.writeText(code).then(() => {
        chrome.runtime.sendMessage({action: 'send-code', code: code});
        notify();
    });
  });
});

chrome.runtime.onMessage.addListener((req, info, cb) => {
    if (req.action == "copy-all") {
        const allCode = getAllCode();
        navigator.clipboard.writeText(allCode).then(() => {
            notify();
            cb(allCode);
        });
        return true;
    }
})

function notify() {
    const scriptEl = document.createElement("script");
    scriptEl.src = chrome.runtime.getURL("execute.js");
    document.body.appendChild(scriptEl);
    scriptEl.onload = () => {
        scriptEl.remove();
    }
}

function getAllCode() {
    return [...preEls].map((preEl) => {
        return preEl.querySelector("code").innerText;
    }).join("");
}