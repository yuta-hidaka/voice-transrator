<script lang="ts">
    import { TranslatorService } from '@/proto/translator.pb';
    import { storage } from "@/storage";

    export let count: number;
    export let err: any;
    export let resp: any;
    let successMessage: string | null = null;

    function increment() {
        count += 1;
    }

    function decrement() {
        count -= 1;
    }

    function save() {
        storage.set({ count }).then(() => {
            successMessage = "Options saved!";

            setTimeout(() => {
                successMessage = null;
            }, 1500);
        });
    }
    async function start() {
        console.log(chrome.tabs);
        const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
        console.log("send message start");
        const response = await chrome.tabs.sendMessage(tab.id ?? 0, {start: true});
       
        // do something with response here, not outside the function
        console.log(response);
    }

   async function stop() {
        const [tab] = await chrome.tabs.query({active: true, lastFocusedWindow: true});
        console.log("send message stop");
        const response = await chrome.tabs.sendMessage(tab.id ?? 0, {start: false});
       
        // do something with response here, not outside the function
        console.log(response);
    }

   async function healthCheck() {
    // const transport = new GrpcWebFetchTransport({
    //     baseUrl: 'http://localhost:8080',
    // });
    // const client = new TranslatorService(transport);
    const resp = await TranslatorService.HealthCheck({},{ pathPrefix: "http://localhost:8080" })
    console.log("resp: ", resp)
    
    }

</script>

<div class="container">
    <p>Current count: <b>{count}</b></p>
    <div>
        <button on:click={decrement}>-</button>
        <button on:click={increment}>+</button>
        <button on:click={save}>Save</button>
        <button  on:click={start}>Start</button>
        <button  on:click={stop}>Stop</button>
        <button  on:click={healthCheck}>healthCheck</button>
        <div>
            aaaaaaa
            {#if err}
                <p>Error: {err.message}</p>
            {/if}
            {#if resp}
                <p>Response: {resp.message}</p>
            {/if}
        </div>
        {#if successMessage}<span class="success">{successMessage}</span>{/if}
    </div>
</div>

<style>
    .container {
        min-width: 250px;
    }

    button {
        border-radius: 2px;
        box-shadow: 0 1px 4px rgba(0, 0, 0, 0.6);
        background-color: #2ecc71;
        color: #ecf0f1;
        transition: background-color 0.3s;
        padding: 5px 10px;
        border: none;
    }

    button:hover,
    button:focus {
        background-color: #27ae60;
    }

    .success {
        color: #2ecc71;
        font-weight: bold;
    }
</style>
