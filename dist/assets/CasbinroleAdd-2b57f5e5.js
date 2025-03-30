import{u as p,d as _,o as d,c as r,e,a as i,n as h,t as f,w as g,g as l,v as n,k as m}from"./index-8eb781de.js";import{_ as v,a as b,F as x}from"./FooterAdmin-d2f58c87.js";import"./_plugin-vue_export-helper-c27b6911.js";const y={class:"leading-normal tracking-normal",id:"main-body"},w={class:"flex flex-wrap"},k={class:"p-6 bg-gray-100 mb-20"},M={class:"py-12"},P={key:0,class:"flex p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800",role:"alert"},E=e("svg",{class:"inline flex-shrink-0 mr-3 w-5 h-5",fill:"currentColor",viewBox:"0 0 20 20",xmlns:"http://www.w3.org/2000/svg"},[e("path",{"fill-rule":"evenodd",d:"M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z","clip-rule":"evenodd"})],-1),T={class:"font-medium"},B={key:1,class:"mt-10 sm:mt-0"},S={class:"md:grid md:grid-cols-3 md:gap-6"},V={class:"mt-5 md:mt-0 md:col-span-2"},z=["onSubmit"],C={class:"shadow overflow-hidden sm:rounded-md"},U={class:"px-4 py-5 bg-white sm:p-6"},A={class:"grid grid-cols-6 gap-6"},D={class:"col-span-6 sm:col-span-3"},q=e("label",{for:"first-name",class:"block text-sm font-medium text-gray-700"},"Role name",-1),F={class:"col-span-6 sm:col-span-3"},H=e("label",{for:"last-name",class:"block text-sm font-medium text-gray-700"},"Path (Example: /*)",-1),N={class:"col-span-6 sm:col-span-4"},O=e("label",{for:"email-address",class:"block text-sm font-medium text-gray-700"},"Method (GET,POST,PUT,PATCH,DELETE)",-1),R=e("div",{class:"px-4 py-3 bg-gray-50 text-right sm:px-6"},[e("button",{type:"submit",class:"inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"},"Save")],-1),J={__name:"CasbinroleAdd",setup($){const c=p(),o=_({session_id:"",session_name:"",error:"",form:{rolename:"",path:"",method:""}});(()=>{try{m.get("/role/api/add").then(s=>{s.data.error!=null?o.value.error=s.data.error:(o.value.session_id=s.data.session_id,o.value.session_name=s.data.session_name)})}catch(s){console.log(s)}})();const u=()=>{try{m.post("/role/post_add",o.value.form).then(s=>{s.rolename="",s.path="",s.method="",c.push({name:"role_list"})})}catch(s){o.value.error=s}};return(s,t)=>(d(),r("div",y,[e("div",w,[i(v),e("div",{class:h(["w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen",s.sideBarOpen?"overlay":""]),id:"main-content"},[i(b),e("div",k,[e("div",M,[o.value.error?(d(),r("div",P,[E,e("div",null,[e("span",T,f(o.value.error),1)])])):(d(),r("div",B,[e("div",S,[e("div",V,[e("form",{onSubmit:g(u,["prevent"])},[e("div",C,[e("div",U,[e("div",A,[e("div",D,[q,l(e("input",{type:"text",name:"first-name",id:"first-name","onUpdate:modelValue":t[0]||(t[0]=a=>o.value.form.rolename=a),required:"",autocomplete:"name",class:"mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"},null,512),[[n,o.value.form.rolename]])]),e("div",F,[H,l(e("input",{type:"text",name:"last-Path",id:"last-Path","onUpdate:modelValue":t[1]||(t[1]=a=>o.value.form.path=a),autocomplete:"Path",class:"mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"},null,512),[[n,o.value.form.path]])]),e("div",N,[O,l(e("input",{type:"text",name:"Method",id:"Method","onUpdate:modelValue":t[2]||(t[2]=a=>o.value.form.method=a),required:"",autocomplete:"Method",class:"mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"},null,512),[[n,o.value.form.method]])])])]),R])],40,z)])])]))])]),i(x)],2)])]))}};export{J as default};
