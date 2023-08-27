
// Content scripts
// https://developer.chrome.com/docs/extensions/mv3/content_scripts/

// Some global styles on the page
import "./styles.css";

// Some JS on the page
// storage.get().then(console.log);

// Some svelte component on the page
// new Overlay({ target: document.body });

console.log("Hello from content script");

let recorder: MediaRecorder;

function stopRecording() {
    if (!recorder) return
    recorder.stream.getTracks().forEach(track => track.stop())
    console.log("recordStream stop")
}

async function recordStream() {
    const stream = await navigator.mediaDevices.getUserMedia({
        video: false,
        audio: true,
    })

    recorder = new MediaRecorder(stream)

    let chunks: Blob[] = []

    recorder.ondataavailable = event => {
        if (event.data.size > 0) {
            chunks.push(event.data)
            console.log(chunks.length)
            console.log(chunks)
        }
    }

    recorder.onstop = () => {
        const blob = new Blob(chunks, {
            type: 'video/webm'
        })

        chunks = []
    }

    recorder.start(200)
    console.log("recordStream start")
}

function init() {
    console.log("init")
    chrome.runtime.onMessage.addListener(
        function (request, _, __) {
            console.log(request)
            if (request.start === true)
                recordStream()
            else
                stopRecording()
        }
    );
}

init()