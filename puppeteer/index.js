'use strict';

const puppeteer = require('puppeteer');

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

(async() => {

const browser = await puppeteer.launch({
  headless: true,
  // args: ["--allow-running-insecure-content", "--ignore-certificate-errors", "--unsafely-treat-insecure-origin-as-secure"],
  userDataDir: "/Users/akagi201/Downloads/tmp/puppeteer/Chrome",
});
const page = await browser.newPage();
await page.goto('https://webrtc.github.io/samples/src/content/getusermedia/record/', {waitUntil: 'networkidle'});
await sleep(5000);
page.on('dialog', dialog => {
  console.log("121312313212312");
  console.log(dialog.type);
  dialog.accept();
});
await sleep(1000);
console.log("Start recording");
await page.click("#record");
await sleep(5000);
console.log("Stop recording");
await page.click("#record");
await sleep(1000);
console.log("Download recording");
await page.click("#download");
await sleep(5000);

browser.close();
})();
