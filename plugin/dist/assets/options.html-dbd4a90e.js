import"./modulepreload-polyfill-3cfb730f.js";import{O as n}from"./Options-d98f4a8b.js";import{s as o}from"./storage-1f8884de.js";function r(){const t=document.getElementById("app");t&&o.get().then(({count:e})=>{new n({target:t,props:{count:e}})})}document.addEventListener("DOMContentLoaded",r);
