import{d as p,i as o,r as c,o as _,c as u,b as e,w as t,u as r,a7 as d,e as n,a8 as f}from"./index-mkXm0lbT.js";import{_ as m}from"./index.vue_vue_type_script_setup_true_name_ProTable_lang-rajnjMkI.js";import"./_plugin-vue_export-helper-x3n3nnut.js";import"./sortable.esm-qwEfw145.js";const b={class:"table-box"},x=p({name:"menuMange"}),C=p({...x,setup(y){const l=o(),s=o(),i=[{prop:"userGroup",label:"用户组名",align:"left"},{prop:"userGroupType",label:"类型(非积分会员默认无上限次数)"},{prop:"userGroupRate",label:"操作频率"},{prop:"userIfExpired",label:"是否可到期(可设定积分类到期)"},{prop:"operation",label:"操作",width:250,fixed:"right"}];return(h,k)=>{const a=c("el-button");return _(),u("div",b,[e(m,{ref_key:"proTable",ref:l,title:"菜单列表","row-key":"path",indent:20,columns:i,data:s.value},{operation:t(()=>[e(a,{type:"primary",link:"",icon:r(d)},{default:t(()=>[n(" 编辑 ")]),_:1},8,["icon"]),e(a,{type:"primary",link:"",icon:r(f)},{default:t(()=>[n(" 删除 ")]),_:1},8,["icon"])]),_:1},8,["data"])])}}});export{C as default};