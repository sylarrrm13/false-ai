import{d as N,J as T,aa as U,i as g,Y as w,ah as C,r as n,o as A,c as E,b as o,w as t,a as z,e as v,u as k,P as B,aP as L,F as h,aQ as q,aR as I,U as K,aS as M,E as P}from"./index-mkXm0lbT.js";import{u as $,a as D}from"./tabs-g61ySpkf.js";import{m as H}from"./md5-8ytPGeQ8.js";import{_ as J}from"./_plugin-vue_export-helper-x3n3nnut.js";const O={class:"login-btn"},Q=N({__name:"LoginForm",setup(Y){const y=T(),d=U(),b=$(),F=D(),l=g(),V=w({username:[{required:!0,message:"请输入用户名",trigger:"blur"}],password:[{required:!0,message:"请输入密码",trigger:"blur"}]}),u=g(!1),a=w({username:"",password:""}),c=e=>{e&&e.validate(async s=>{if(s){u.value=!0;try{const{data:i}=await q({...a,password:H(a.password)});d.setToken(i.access_token),d.setRefreshToken(i.refresh_token),d.setUserInfo({name:a.username}),await I(),b.setTabs([]),F.setKeepAliveName([]),y.push(K),P({title:M(),message:"欢迎登录 ",type:"success",duration:3e3})}finally{u.value=!1}}})},R=e=>{e&&e.resetFields()};return C(()=>{document.onkeydown=e=>{if(e=window.event||e,e.code==="Enter"||e.code==="enter"||e.code==="NumpadEnter"){if(u.value)return;c(l.value)}}}),(e,s)=>{const i=n("user"),m=n("el-icon"),p=n("el-input"),_=n("el-form-item"),S=n("lock"),x=n("el-form"),f=n("el-button");return A(),E(h,null,[o(x,{ref_key:"loginFormRef",ref:l,model:a,rules:V,size:"large"},{default:t(()=>[o(_,{prop:"username"},{default:t(()=>[o(p,{modelValue:a.username,"onUpdate:modelValue":s[0]||(s[0]=r=>a.username=r),placeholder:"用户名"},{prefix:t(()=>[o(m,{class:"el-input__icon"},{default:t(()=>[o(i)]),_:1})]),_:1},8,["modelValue"])]),_:1}),o(_,{prop:"password"},{default:t(()=>[o(p,{modelValue:a.password,"onUpdate:modelValue":s[1]||(s[1]=r=>a.password=r),type:"password",placeholder:"密码","show-password":"",autocomplete:"new-password"},{prefix:t(()=>[o(m,{class:"el-input__icon"},{default:t(()=>[o(S)]),_:1})]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["model","rules"]),z("div",O,[o(f,{icon:k(B),round:"",size:"large",onClick:s[2]||(s[2]=r=>R(l.value))},{default:t(()=>[v(" 重置 ")]),_:1},8,["icon"]),o(f,{icon:k(L),round:"",size:"large",type:"primary",loading:u.value,onClick:s[3]||(s[3]=r=>c(l.value))},{default:t(()=>[v(" 登录 ")]),_:1},8,["icon","loading"])])],64)}}}),Z=J(Q,[["__scopeId","data-v-52da98f8"]]);export{Z as default};