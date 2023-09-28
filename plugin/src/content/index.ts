
// Content scripts
// https://developer.chrome.com/docs/extensions/mv3/content_scripts/

// Some global styles on the page
import { TranslatorService } from '@/proto/translator.pb';
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

async function healthCheck() {
    const resp = await TranslatorService.HealthCheck({}, { pathPrefix: "http://localhost:8080" })
    console.log("resp: ", resp)
}

async function recordStream() {
    const stream = await navigator.mediaDevices.getUserMedia({
        video: false,
        audio: true,
    })

    const resp = await TranslatorService.HealthCheck({}, { pathPrefix: "http://localhost:8080" })
    console.log("resp: ", resp)

    recorder = new MediaRecorder(stream)

    let chunks: Blob[] = []

    recorder.ondataavailable = async event => {
        if (event.data.size > 0) {
            console.log(await event.data)
            // console.log(await event.data.text())
            const buf = await event.data.arrayBuffer()
            const stream = await event.data.stream()
            console.log("buf ====> ", buf)
            // console.log(await event.data.arrayBuffer())
            // chunks.push()
            // console.log(chunks.length)
            // console.log(chunks)
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
    healthCheck()
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