import Overlay from "../components/Overlay.svelte";
import { storage } from "../storage";

// Content scripts
// https://developer.chrome.com/docs/extensions/mv3/content_scripts/

// Some global styles on the page
import "./styles.css";

alert("Hello from content script");

// Some JS on the page
storage.get().then(console.log);

// Some svelte component on the page
new Overlay({ target: document.body });

console.log("Hello from content script");

function init() {
    const el = document.createElement("input")
    el.setAttribute("type", "checkbox")
    document.body.appendChild(el)
    el.addEventListener("click", (event) => {
        console.log("clicked", event)
    })
}

init()