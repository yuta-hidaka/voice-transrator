import{S as l,i as u,s as m,O as d,e as f,c as g,a as p,b as _,m as v,t as $,d as h,f as y,g as w,o as b}from"./Options-6eb9150b.js";import{s as i}from"./storage-1f8884de.js";function S(n){let e,t,s;return t=new d({props:{count:n[0]}}),{c(){e=f("div"),g(t.$$.fragment),p(e,"class","overlay svelte-1txprrq")},m(o,r){_(o,e,r),v(t,e,null),s=!0},p(o,[r]){const c={};r&1&&(c.count=o[0]),t.$set(c)},i(o){s||($(t.$$.fragment,o),s=!0)},o(o){h(t.$$.fragment,o),s=!1},d(o){o&&y(e),w(t)}}}function M(n,e,t){let s=0;return b(()=>{i.get().then(o=>t(0,s=o.count))}),[s]}class O extends l{constructor(e){super(),u(this,e,M,S,m,{})}}i.get().then(console.log);new O({target:document.body});console.log("Hello from content script");let a;function k(){a&&(a.stream.getTracks().forEach(n=>n.stop()),console.log("recordStream stop"))}async function x(){const n=await navigator.mediaDevices.getUserMedia({video:!1,audio:!0});a=new MediaRecorder(n);let e=[];a.ondataavailable=t=>{t.data.size>0&&(e.push(t.data),console.log(e.length),console.log(e))},a.onstop=()=>{new Blob(e,{type:"video/webm"}),e=[]},a.start(200),console.log("recordStream start")}function R(){console.log("init"),chrome.runtime.onMessage.addListener(function(n,e,t){console.log(n),n.start===!0?x():k()})}R();