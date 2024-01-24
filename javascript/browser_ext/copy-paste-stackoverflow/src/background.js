chrome.commands.onCommand.addListener((command) => {
    if (command === "copy-all") {
        getCurrentTabId().then((tabId) => askTabToCopyAll(tabId))
    }
})

async function getCurrentTabId() {
    let [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
    return tab.id;
}

askTabToCopyAll = (tabId) => {
    chrome.tabs.sendMessage(tabId, { action: "copy-all" },
    (allCode) => {
        console.log(allCode);
    });
}
