import{d as f,i as p,r as d,o,c as u,b as c,h as a,w as l,F as x,f as y,k as v,D as r,u as s,B as n,aF as I}from"./index-mkXm0lbT.js";import{_ as w}from"./_plugin-vue_export-helper-x3n3nnut.js";const k={class:"card content-box"},F=f({name:"proForm"}),N=f({...F,setup(V){let t=p({});const m=p({form:{inline:!1,labelPosition:"right",labelWidth:"80px",size:"default",disabled:!1,labelSuffix:" :"},columns:[{formItem:{label:"用户名",prop:"username",labelWidth:"80px",required:!0},attrs:{typeName:"input",clearable:!0,placeholder:"请输入用户名",disabled:!0}},{formItem:{label:"密码",prop:"password",class:"data"},attrs:{typeName:"input",clearable:!0,autofocus:!0,placeholder:"请输入密码",type:"password"}},{formItem:{label:"邮箱",prop:"email"},attrs:{typeName:"input",placeholder:"请输入邮箱",clearable:!0,style:"width:500px"}}]});return(i,g)=>{const _=d("el-alert"),b=d("el-form-item");return o(),u("div",k,[c(_,{title:"通过 component :is 组件属性 && v-bind 属性透传，可以将 template 中的 html 代码全部改变为 columns 配置项，具体配置请看代码。",type:"warning",closable:!1}),(o(),a(n("el-form"),r(m.value.form,{ref:"proFormRef",model:s(t)}),{default:l(()=>[(o(!0),u(x,null,y(m.value.columns,e=>(o(),a(n("el-form-item"),I(r({key:e.prop},e.formItem)),{default:l(()=>[(o(),a(n(`el-${e.attrs.typeName}`),r(e.attrs,{modelValue:s(t)[e.formItem.prop],"onUpdate:modelValue":h=>s(t)[e.formItem.prop]=h}),null,16,["modelValue","onUpdate:modelValue"]))]),_:2},1040))),128)),c(b,null,{default:l(()=>[v(i.$slots,"operation",{},void 0,!0)]),_:3})]),_:3},16,["model"]))])}}}),P=w(N,[["__scopeId","data-v-69441d13"]]);export{P as default};
