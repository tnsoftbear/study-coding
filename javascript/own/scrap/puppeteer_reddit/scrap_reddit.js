const puppeteer = require('puppeteer')

const url = 'https://old.reddit.com/r/webscraping/comments/18rzpgu/any_advice_on_how_to_scrape_a_150m_tweets/';
//const url = 'https://redd.it/18rzpgu';

(async () => {
    const browser = await puppeteer.launch({headless: "new"});
    const page = await browser.newPage();
    await page.goto(url);
    //const content = await page.content();
    //const comments = await page.$eval(".md", el => el.textContent);
    const comments = await page.$(".md");
    console.log(comments);
    await browser.close();
})();