import{g as K,u as g}from"./config-FLsHmEoJ.js";import{d as h,V as x,Y as k,r as c,o as v,c as I,b as t,w as a,a as d,e as f,T as C,p as w,g as D}from"./index-mkXm0lbT.js";import{_ as P}from"./_plugin-vue_export-helper-x3n3nnut.js";const S=i=>(w("data-v-92cafbc3"),i=i(),D(),i),U={style:{display:"flex","align-items":"center"}},z=S(()=>d("div",null,[d("div",{style:{"margin-bottom":"5px","font-size":"15px"}},"R2存储设置说明"),d("div",{style:{"font-size":"13px"}},[f(" R2存储基于的是CloudFlare的免费对象存储系统,具体教程设置 "),d("a",{href:"https://zhuanlan.zhihu.com/p/658058503",target:"_blank"},"点我查看")])],-1)),B={style:{"text-align":"right"}},N=h({__name:"index",setup(i){x(async()=>{await K("R2_bucket").then(s=>{const{enable:o,accountID:u,accessKey:m,secretKey:p,bucket:_,endPoint:n}=JSON.parse(s.data.value);e.enable=o=="1",e.accountID=u,e.accessKey=m,e.secretKey=p,e.endPoint=n,e.bucket=_,JSON.parse(s.data.value)}).catch(s=>{})});const b=async function(){const s={enable:e.enable?"1":"0",accountID:e.accountID,accessKey:e.accessKey,secretKey:e.secretKey,bucket:e.bucket,endPoint:e.endPoint};await g("R2_bucket",s).then(o=>{const{data:u}=o;u.status==0&&C.success("更新成功")}).catch(o=>{})},e=k({enable:!1,accountID:"",accessKey:"",secretKey:"",bucket:"",endPoint:""});return(s,o)=>{const u=c("CoffeeCup"),m=c("el-icon"),p=c("el-card"),_=c("el-switch"),n=c("el-form-item"),r=c("el-input"),y=c("el-form"),V=c("el-button");return v(),I("div",null,[t(p,null,{default:a(()=>[d("div",U,[t(m,{size:"35",color:"blue",style:{"margin-right":"10px"}},{default:a(()=>[t(u)]),_:1}),z])]),_:1}),t(p,{style:{"margin-top":"10px"}},{footer:a(l=>[d("div",B,[t(V,{type:"primary",onClick:b},{default:a(()=>[f("保存设置")]),_:1})])]),default:a(()=>[t(y,{model:e,width:"200px","label-width":"auto"},{default:a(()=>[t(n,{label:"启用R2存储 :"},{default:a(()=>[t(_,{modelValue:e.enable,"onUpdate:modelValue":o[0]||(o[0]=l=>e.enable=l),"true-value":"1","false-value":"0"},null,8,["modelValue"])]),_:1}),t(n,{label:"EndPoint :"},{default:a(()=>[t(r,{modelValue:e.endPoint,"onUpdate:modelValue":o[1]||(o[1]=l=>e.endPoint=l),placeholder:""},null,8,["modelValue"])]),_:1}),t(n,{label:"AccountID :"},{default:a(()=>[t(r,{modelValue:e.accountID,"onUpdate:modelValue":o[2]||(o[2]=l=>e.accountID=l),placeholder:""},null,8,["modelValue"])]),_:1}),t(n,{label:"Bucket :"},{default:a(()=>[t(r,{modelValue:e.bucket,"onUpdate:modelValue":o[3]||(o[3]=l=>e.bucket=l),placeholder:""},null,8,["modelValue"])]),_:1}),t(n,{label:"AccessKey :"},{default:a(()=>[t(r,{modelValue:e.accessKey,"onUpdate:modelValue":o[4]||(o[4]=l=>e.accessKey=l),placeholder:""},null,8,["modelValue"])]),_:1}),t(n,{label:"SecretKey :"},{default:a(()=>[t(r,{modelValue:e.secretKey,"onUpdate:modelValue":o[5]||(o[5]=l=>e.secretKey=l),placeholder:""},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1})])}}}),J=P(N,[["__scopeId","data-v-92cafbc3"]]);export{J as default};