var J=Object.defineProperty;var K=(t,e,n)=>e in t?J(t,e,{enumerable:!0,configurable:!0,writable:!0,value:n}):t[e]=n;var A=(t,e,n)=>(K(t,typeof e!="symbol"?e+"":e,n),n);import{s as Q}from"./storage-1f8884de.js";function x(){}function V(t){return t()}function z(){return Object.create(null)}function O(t){t.forEach(V)}function q(t){return typeof t=="function"}function W(t,e){return t!=t?e==e:t!==e||t&&typeof t=="object"||typeof t=="function"}function X(t){return Object.keys(t).length===0}function c(t,e){t.appendChild(e)}function D(t,e,n){t.insertBefore(e,n||null)}function P(t){t.parentNode&&t.parentNode.removeChild(t)}function p(t){return document.createElement(t)}function j(t){return document.createTextNode(t)}function S(){return j(" ")}function L(t,e,n,s){return t.addEventListener(e,n,s),()=>t.removeEventListener(e,n,s)}function w(t,e,n){n==null?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function Y(t){return Array.from(t.childNodes)}function G(t,e){e=""+e,t.data!==e&&(t.data=e)}let E;function k(t){E=t}function Z(){if(!E)throw new Error("Function called outside component initialization");return E}function mt(t){Z().$$.on_mount.push(t)}const y=[],F=[];let b=[];const R=[],tt=Promise.resolve();let B=!1;function et(){B||(B=!0,tt.then(H))}function I(t){b.push(t)}const M=new Set;let $=0;function H(){if($!==0)return;const t=E;do{try{for(;$<y.length;){const e=y[$];$++,k(e),nt(e.$$)}}catch(e){throw y.length=0,$=0,e}for(k(null),y.length=0,$=0;F.length;)F.pop()();for(let e=0;e<b.length;e+=1){const n=b[e];M.has(n)||(M.add(n),n())}b.length=0}while(y.length);for(;R.length;)R.pop()();B=!1,M.clear(),k(t)}function nt(t){if(t.fragment!==null){t.update(),O(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(I)}}function st(t){const e=[],n=[];b.forEach(s=>t.indexOf(s)===-1?e.push(s):n.push(s)),n.forEach(s=>s()),b=e}const N=new Set;let ot;function rt(t,e){t&&t.i&&(N.delete(t),t.i(e))}function gt(t,e,n,s){if(t&&t.o){if(N.has(t))return;N.add(t),ot.c.push(()=>{N.delete(t),s&&(n&&t.d(1),s())}),t.o(e)}else s&&s()}function $t(t){t&&t.c()}function it(t,e,n){const{fragment:s,after_update:o}=t.$$;s&&s.m(e,n),I(()=>{const a=t.$$.on_mount.map(V).filter(q);t.$$.on_destroy?t.$$.on_destroy.push(...a):O(a),t.$$.on_mount=[]}),o.forEach(I)}function ut(t,e){const n=t.$$;n.fragment!==null&&(st(n.after_update),O(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function ft(t,e){t.$$.dirty[0]===-1&&(y.push(t),et(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function ct(t,e,n,s,o,a,m,i=[-1]){const f=E;k(t);const r=t.$$={fragment:null,ctx:[],props:a,update:x,not_equal:o,bound:z(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(e.context||(f?f.$$.context:[])),callbacks:z(),dirty:i,skip_bound:!1,root:e.target||f.$$.root};m&&m(r.root);let d=!1;if(r.ctx=n?n(t,e.props||{},(l,_,...v)=>{const g=v.length?v[0]:_;return r.ctx&&o(r.ctx[l],r.ctx[l]=g)&&(!r.skip_bound&&r.bound[l]&&r.bound[l](g),d&&ft(t,l)),_}):[],r.update(),d=!0,O(r.before_update),r.fragment=s?s(r.ctx):!1,e.target){if(e.hydrate){const l=Y(e.target);r.fragment&&r.fragment.l(l),l.forEach(P)}else r.fragment&&r.fragment.c();e.intro&&rt(t.$$.fragment),it(t,e.target,e.anchor),H()}k(f)}class lt{constructor(){A(this,"$$");A(this,"$$set")}$destroy(){ut(this,1),this.$destroy=x}$on(e,n){if(!q(n))return x;const s=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return s.push(n),()=>{const o=s.indexOf(n);o!==-1&&s.splice(o,1)}}$set(e){this.$$set&&!X(e)&&(this.$$.skip_bound=!0,this.$$set(e),this.$$.skip_bound=!1)}}const at="4";typeof window<"u"&&(window.__svelte||(window.__svelte={v:new Set})).v.add(at);function U(t){let e,n;return{c(){e=p("span"),n=j(t[1]),w(e,"class","success svelte-s5p16k")},m(s,o){D(s,e,o),c(e,n)},p(s,o){o&2&&G(n,s[1])},d(s){s&&P(e)}}}function dt(t){let e,n,s,o,a,m,i,f,r,d,l,_,v,g,T,u=t[1]&&U(t);return{c(){e=p("div"),n=p("p"),s=j("Current count: "),o=p("b"),a=j(t[0]),m=S(),i=p("div"),f=p("button"),f.textContent="-",r=S(),d=p("button"),d.textContent="+",l=S(),_=p("button"),_.textContent="Save",v=S(),u&&u.c(),w(f,"class","svelte-s5p16k"),w(d,"class","svelte-s5p16k"),w(_,"class","svelte-s5p16k"),w(e,"class","container svelte-s5p16k")},m(h,C){D(h,e,C),c(e,n),c(n,s),c(n,o),c(o,a),c(e,m),c(e,i),c(i,f),c(i,r),c(i,d),c(i,l),c(i,_),c(i,v),u&&u.m(i,null),g||(T=[L(f,"click",t[3]),L(d,"click",t[2]),L(_,"click",t[4])],g=!0)},p(h,[C]){C&1&&G(a,h[0]),h[1]?u?u.p(h,C):(u=U(h),u.c(),u.m(i,null)):u&&(u.d(1),u=null)},i:x,o:x,d(h){h&&P(e),u&&u.d(),g=!1,O(T)}}}function _t(t,e,n){let{count:s}=e,o=null;function a(){n(0,s+=1)}function m(){n(0,s-=1)}function i(){Q.set({count:s}).then(()=>{n(1,o="Options saved!"),setTimeout(()=>{n(1,o=null)},1500)})}return t.$$set=f=>{"count"in f&&n(0,s=f.count)},[s,o,a,m,i]}class yt extends lt{constructor(e){super(),ct(this,e,_t,dt,W,{count:0})}}export{yt as O,lt as S,w as a,D as b,$t as c,gt as d,p as e,P as f,ut as g,ct as i,it as m,mt as o,W as s,rt as t};
