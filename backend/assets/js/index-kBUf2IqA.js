import{W as w}from"./index-PbbbS5Xx.js";import{d as u,i as _,r as d,o as h,c as g,b as e,w as t,e as o,a as i,p as b,g as x}from"./index-mkXm0lbT.js";import{_ as V}from"./_plugin-vue_export-helper-x3n3nnut.js";import"./upload-hbbC0aqT.js";const p=n=>(b("data-v-9dd82296"),n=n(),x(),n),B={class:"card content-box"},C=p(()=>i("span",{class:"text"},"富文本编辑器 🍓🍇🍈🍉",-1)),T=p(()=>i("a",{href:"https://www.wangeditor.com/v5/toolbar-config.html"}," https://www.wangeditor.com/v5/toolbar-config.html ",-1)),I=p(()=>i("a",{href:"https://www.wangeditor.com/v5/editor-config.html"}," https://www.wangeditor.com/v5/editor-config.html ",-1)),k=["innerHTML"],E=u({name:"wangEditor"}),N=u({...E,setup(n){const r=_(""),c=_(!1);return(S,a)=>{const m=d("el-button"),l=d("el-descriptions-item"),f=d("el-descriptions"),v=d("el-dialog");return h(),g("div",B,[C,e(w,{value:r.value,"onUpdate:value":a[0]||(a[0]=s=>r.value=s),height:"400px"},null,8,["value"]),e(m,{type:"primary",onClick:a[1]||(a[1]=s=>c.value=!0)},{default:t(()=>[o(" 内容预览 ")]),_:1}),e(f,{title:"配置项 📚",column:1,border:""},{default:t(()=>[e(l,{label:"value"},{default:t(()=>[o(' 双向绑定的 value 值，使用示例： v-model:value="content"> ')]),_:1}),e(l,{label:"toolbarConfig"},{default:t(()=>[o(" 富文本 ToolBar区域 配置： "),T]),_:1}),e(l,{label:"editorConfig"},{default:t(()=>[o(" 富文本 编辑区域 配置： "),I]),_:1}),e(l,{label:"height"},{default:t(()=>[o(" 富文本高度，默认为 500px ")]),_:1}),e(l,{label:"mode"},{default:t(()=>[o(' 富文本模式，默认为 default（"default" | "simple"） ')]),_:1}),e(l,{label:"hideToolBar"},{default:t(()=>[o(" 隐藏 ToolBar 区域，默认为 false ")]),_:1}),e(l,{label:"disabled"},{default:t(()=>[o(" 禁用富文本编辑器，默认为 false ")]),_:1})]),_:1}),e(v,{modelValue:c.value,"onUpdate:modelValue":a[2]||(a[2]=s=>c.value=s),title:"富文本内容预览",width:"1300px",top:"50px"},{default:t(()=>[i("div",{class:"view",innerHTML:r.value},null,8,k)]),_:1},8,["modelValue"])])}}}),U=V(N,[["__scopeId","data-v-9dd82296"]]);export{U as default};