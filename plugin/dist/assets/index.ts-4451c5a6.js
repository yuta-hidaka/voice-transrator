import{s as o}from"./storage-1f8884de.js";console.log("Hello from background script");chrome.runtime.onInstalled.addListener(()=>{o.get().then(console.log)});
