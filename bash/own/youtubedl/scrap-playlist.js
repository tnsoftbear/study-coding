const videoElements = document.querySelectorAll('a#video-title');
const videoUrls = Array.from(videoElements).map(el => {
    let url = 'https://www.youtube.com' + el.getAttribute('href');
    const listIndex = url.indexOf('&list=');
    if (listIndex !== -1) {
        url = url.substring(0, listIndex);
    }
    return url;
});
console.log(videoUrls.join('\n'));