import{H as S,a4 as g,C as N,d as k,r as f,a5 as T,o as a,c as C,b as s,u as e,w as t,e as o,h as i,a6 as y,A as p,a7 as h,a8 as x,l as v,j as B,a9 as d,p as O,g as A,a as I}from"./index-mkXm0lbT.js";import{_ as V}from"./_plugin-vue_export-helper-x3n3nnut.js";const U=()=>{const r=S(),b=g().authButtonListGet[r.name]||[];return{BUTTONS:N(()=>{let _={};return b.forEach(u=>_[u]=!0),_})}},j=r=>(O("data-v-34a32b33"),r=r(),A(),r),D={class:"card content-box"},E=j(()=>I("span",{class:"text"}," 按钮权限 🍓🍇🍈🍉",-1)),H=k({name:"authButton"}),G=k({...H,setup(r){const{BUTTONS:l}=U();return(b,w)=>{const _=f("el-alert"),u=f("el-divider"),n=f("el-button"),m=f("el-row"),c=T("auth");return a(),C("div",D,[E,s(_,{class:"mb20",title:`当前用户按钮权限：${JSON.stringify(Object.keys(e(l)))}`,type:"success",closable:!1},null,8,["title"]),s(u,{"content-position":"left"},{default:t(()=>[o(" 使用 Hooks 方式绑定权限 ")]),_:1}),s(m,{class:"mb20"},{default:t(()=>[e(l).add?(a(),i(n,{key:0,type:"primary",icon:e(y)},{default:t(()=>[o(" 新增 ")]),_:1},8,["icon"])):p("",!0),e(l).edit?(a(),i(n,{key:1,type:"warning",icon:e(h)},{default:t(()=>[o(" 编辑 ")]),_:1},8,["icon"])):p("",!0),e(l).delete?(a(),i(n,{key:2,type:"danger",plain:"",icon:e(x)},{default:t(()=>[o(" 删除 ")]),_:1},8,["icon"])):p("",!0),e(l).import?(a(),i(n,{key:3,type:"info",plain:"",icon:e(v)},{default:t(()=>[o(" 导入数据 ")]),_:1},8,["icon"])):p("",!0),e(l).export?(a(),i(n,{key:4,type:"info",plain:"",icon:e(B)},{default:t(()=>[o(" 导出数据 ")]),_:1},8,["icon"])):p("",!0)]),_:1}),s(u,{"content-position":"left"},{default:t(()=>[o(" 使用 v-auth 指令绑定单个权限 ")]),_:1}),s(m,{class:"mb20"},{default:t(()=>[d((a(),i(n,{type:"primary",icon:e(y)},{default:t(()=>[o(" 新增 ")]),_:1},8,["icon"])),[[c,"add"]]),d((a(),i(n,{type:"warning",icon:e(h)},{default:t(()=>[o(" 编辑 ")]),_:1},8,["icon"])),[[c,"edit"]]),d((a(),i(n,{type:"danger",plain:"",icon:e(x)},{default:t(()=>[o(" 删除 ")]),_:1},8,["icon"])),[[c,"delete"]]),d((a(),i(n,{type:"info",plain:"",icon:e(v)},{default:t(()=>[o(" 导入数据 ")]),_:1},8,["icon"])),[[c,"import"]]),d((a(),i(n,{type:"info",plain:"",icon:e(B)},{default:t(()=>[o(" 导出数据 ")]),_:1},8,["icon"])),[[c,"export"]])]),_:1}),s(u,{"content-position":"left"},{default:t(()=>[o(" 使用 v-auth 指令绑定多个权限 ")]),_:1}),s(m,null,{default:t(()=>[d((a(),i(n,{type:"primary",icon:e(y)},{default:t(()=>[o(" 新增 ")]),_:1},8,["icon"])),[[c,["add","edit","delete","import","export"]]]),d((a(),i(n,{type:"warning",icon:e(h)},{default:t(()=>[o(" 编辑 ")]),_:1},8,["icon"])),[[c,["add","edit","delete","import","export"]]]),d((a(),i(n,{type:"danger",plain:"",icon:e(x)},{default:t(()=>[o(" 删除 ")]),_:1},8,["icon"])),[[c,["add","edit","delete","import","export"]]]),d((a(),i(n,{type:"info",plain:"",icon:e(v)},{default:t(()=>[o(" 导入数据 ")]),_:1},8,["icon"])),[[c,["add","edit","delete","import","export"]]]),d((a(),i(n,{type:"info",plain:"",icon:e(B)},{default:t(()=>[o(" 导出数据 ")]),_:1},8,["icon"])),[[c,["add","edit","delete","import","export"]]])]),_:1})])}}}),P=V(G,[["__scopeId","data-v-34a32b33"]]);export{P as default};