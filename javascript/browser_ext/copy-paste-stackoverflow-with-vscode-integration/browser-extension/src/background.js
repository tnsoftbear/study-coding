chrome.commands.onCommand.addListener((command) => {
  if (command === "copy-all") {
    getCurrentTabId().then((tabId) => askTabToCopyAll(tabId));
  }
});

async function getCurrentTabId() {
  let [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
  return tab.id;
}

chrome.runtime.onMessage.addListener((req, info, cb) => {
  if (req.action == 'send-code') {
    sendCodeToVsCode(req.code).then(() => {
      console.log("Copied");
    });
    incrementCounter(calcLOC(req.code));
  }
  console.log(req);
  console.log(getCounter());
  if (req.action == 'get-count') {
    getCounter().then((counter) => cb(counter));
    return true;
  }
});

askTabToCopyAll = (tabId) => {
  chrome.tabs.sendMessage(tabId, { action: "copy-all" }, (allCode) => {
    sendCodeToVsCode(allCode).then(() => {
      console.log("Copied all");
    });
    incrementCounter(calcLOC(allCode));
  });
};

sendCodeToVsCode = (code) => {
  return fetch("http://localhost:4450/copypaste", {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
    body: JSON.stringify({ code }),
  }).catch(e => {
    console.error("VS Code not found");
    console.error(e);
  });
};

calcLOC = (code) => {
  return code.split("\n").length;
}

incrementCounter = (count) => {
    getCounter().then((counter) => {
        chrome.storage.local.set({ counter: counter + count });
    });
}

getCounter = () => {
    return chrome.storage.local.get('counter').then((data) => {
        return data.counter ?? 0
    });
};

chrome.runtime.onInstalled.addListener(({ reason }) => {
    if (reason == "install") {
        chrome.tabs.create({
            url: chrome.runtime.getURL("welcome.html")
        });
        chrome.runtime.setUninstallURL('http://localhost:4450/leave');
    }
})