(window.webpackJsonp=window.webpackJsonp||[]).push([[23],{278:function(e,t,a){"use strict";function r(e,t){if(!document)return;let a=document.getElementsByTagName("head")[0];for(let r of a.childNodes)if(r.src===e)return void t();let r=document.createElement("script");r.type="text/javascript",r.src=e,a.appendChild(r),r.onload=r.onreadystatechange=function(){this.readyState&&"loaded"!==this.readyState&&"complete"!==this.readyState||(r.onload=r.onreadystatechange=null,t())}}a.d(t,"a",(function(){return r}))},309:function(e,t,a){"use strict";a.r(t);var r=a(278),n={name:"swagger",props:["src"],data:()=>({image:""}),mounted(){Object(r.a)("https://unpkg.com/swagger-ui-dist@5.9.1/swagger-ui-bundle.js",()=>{SwaggerUIBundle({url:this.src,domNode:this.$refs.swagger,presets:[SwaggerUIBundle.presets.apis]})})}},s=a(10),d=Object(s.a)(n,(function(){var e=this._self._c;return e("div",{staticStyle:{border:"1px solid #eeeeee",overflow:"hidden",width:"100%"}},[e("div",{ref:"swagger"})])}),[],!1,null,"2025ceae",null);t.default=d.exports}}]);