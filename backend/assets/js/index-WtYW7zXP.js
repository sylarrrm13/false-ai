import{g as x,u as b}from"./config-FLsHmEoJ.js";import{d as y,V as S,Y as v,r as d,o as I,c as w,b as l,w as a,a as s,e as k,T as D,p as U,g as h}from"./index-mkXm0lbT.js";import{_ as C}from"./_plugin-vue_export-helper-x3n3nnut.js";const N=f=>(U("data-v-f59e4bb1"),f=f(),h(),f),A={style:{display:"flex","align-items":"center"}},B=N(()=>s("div",null,[s("div",{style:{"margin-bottom":"5px","font-size":"15px"}},"微信配置说明"),s("div",{style:{"font-size":"10px"}}," 系统微信登录通过关联公众号实现[请务必注册为服务号、个人公众号没有二维码等此类权限]、以便于后期联动小程序、开发文档前往微信公众平台 https://mp.weixin.qq.com/ 、拿到开发者配置信息即可、如果用户对公众号发送消息、我们将会从自定义回复管理当中的内容进行匹配自动回复、如果没有匹配到结果则回复下面设置的自定义回复默认信息、同时别忘记在微信公众号平台将自己的ip加为ip白名单、配置位置为公众号后台->基本配置：服务复制参考 https://域名/api/official/notify 将域名修改为您的域名 下方Token对应自己后台设置的Token、加密秘钥随机即可、当设置不指定首页并且配置了微信登录即可默认打开静默登录！ ")],-1)),T={style:{"text-align":"right"}},z=y({__name:"index",setup(f){S(async()=>{await x("wx_login").then(r=>{const{login:t,auto:u,name:c,appID:m,token:i,appSecret:n,ref1:p,ref2:_,ref3:V,ref4:o}=JSON.parse(r.data.value);JSON.parse(r.data.value),e.login=t=="1",e.auto=u=="1",e.name=c,e.appID=m,e.token=i,e.appSecret=n,e.ref1=p,e.ref2=_,e.ref3=V,e.ref4=o}).catch(r=>{})});const g=async function(){const r={login:e.login?"1":"0",auto:e.auto?"1":"0",name:e.name,appID:e.appID,token:e.token,appSecret:e.appSecret,ref1:e.ref1,ref2:e.ref2,ref3:e.ref3,ref4:e.ref4};await b("wx_login",r).then(t=>{const{data:u}=t;u.status==0&&D.success("更新成功")}).catch(t=>{})},e=v({login:!1,auto:!1,name:"",appID:"",token:"",appSecret:"",ref1:"",ref2:"",ref3:"",ref4:""});return(r,t)=>{const u=d("Apple"),c=d("el-icon"),m=d("el-card"),i=d("el-switch"),n=d("el-form-item"),p=d("el-input"),_=d("el-form"),V=d("el-button");return I(),w("div",null,[l(m,null,{default:a(()=>[s("div",A,[l(c,{size:"35",color:"blue",style:{"margin-right":"10px"}},{default:a(()=>[l(u)]),_:1}),B])]),_:1}),l(m,{style:{"margin-top":"10px"}},{footer:a(o=>[s("div",T,[l(V,{type:"primary",onClick:g},{default:a(()=>[k("保存设置")]),_:1})])]),default:a(()=>[l(_,{model:e,"label-width":"auto"},{default:a(()=>[l(n,{label:"启用微信登录 :"},{default:a(()=>[l(i,{modelValue:e.login,"onUpdate:modelValue":t[0]||(t[0]=o=>e.login=o)},null,8,["modelValue"])]),_:1}),l(n,{label:"启用静默登录 :"},{default:a(()=>[l(i,{modelValue:e.auto,"onUpdate:modelValue":t[1]||(t[1]=o=>e.auto=o)},null,8,["modelValue"])]),_:1}),l(n,{label:"公众号名 :"},{default:a(()=>[l(p,{modelValue:e.name,"onUpdate:modelValue":t[2]||(t[2]=o=>e.name=o)},null,8,["modelValue"])]),_:1}),l(n,{label:"APPID :"},{default:a(()=>[l(p,{modelValue:e.appID,"onUpdate:modelValue":t[3]||(t[3]=o=>e.appID=o),placeholder:""},null,8,["modelValue"])]),_:1}),l(n,{label:"AppSecret :"},{default:a(()=>[l(p,{modelValue:e.appSecret,"onUpdate:modelValue":t[4]||(t[4]=o=>e.appSecret=o),placeholder:""},null,8,["modelValue"])]),_:1}),l(n,{label:"订阅公众号欢迎信息 :"},{default:a(()=>[l(p,{type:"textarea",modelValue:e.ref1,"onUpdate:modelValue":t[5]||(t[5]=o=>e.ref1=o),placeholder:""},null,8,["modelValue"])]),_:1}),l(n,{label:"绑定账号回复信息 :"},{default:a(()=>[l(p,{type:"textarea",modelValue:e.ref2,"onUpdate:modelValue":t[6]||(t[6]=o=>e.ref2=o),placeholder:""},null,8,["modelValue"])]),_:1}),l(n,{label:"扫码登录回复信息 :"},{default:a(()=>[l(p,{type:"textarea",modelValue:e.ref3,"onUpdate:modelValue":t[7]||(t[7]=o=>e.ref3=o),placeholder:""},null,8,["modelValue"])]),_:1}),l(n,{label:"自定义回复默认信息 :"},{default:a(()=>[l(p,{type:"textarea",modelValue:e.ref4,"onUpdate:modelValue":t[8]||(t[8]=o=>e.ref4=o),placeholder:""},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1})])}}}),M=C(z,[["__scopeId","data-v-f59e4bb1"]]);export{M as default};
