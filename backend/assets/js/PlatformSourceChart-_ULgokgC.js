import{_ as o}from"./index.vue_vue_type_script_setup_true_name_ECharts_lang-j7KPP7lj.js";import{d as n,o as s,c as f,b as i}from"./index-mkXm0lbT.js";import{_ as c}from"./_plugin-vue_export-helper-x3n3nnut.js";import"./index-Z215cjZG.js";const m={class:"echarts"},p=n({__name:"PlatformSourceChart",setup(u){const t=[{value:40,name:"智慧文旅平台",percentage:"40%"},{value:10,name:"携程",percentage:"10%"},{value:20,name:"飞猪",percentage:"20%"},{value:30,name:"其他渠道",percentage:"30%"}],l={grid:{top:"0%",left:"2%",right:"2%",bottom:"0%"},tooltip:{trigger:"item",formatter:"{b} :  {c}人"},legend:{show:!0,top:"middle",left:"20px",icon:"circle",orient:"vertical",align:"auto",itemWidth:10,textStyle:{color:"#fff"},itemGap:20,formatter:function(a){let e="";return t.forEach(r=>{r.name===a&&(e=a+" --- "+r.percentage)}),e},data:t.map(a=>a.name)},series:[{type:"pie",radius:["60%","85%"],center:["68%","45%"],color:["#0E7CE2","#FF8352","#E271DE","#F8456B","#00FFFF","#4AEAB0"],itemStyle:{borderColor:"#031845",borderWidth:10},data:t,labelLine:{show:!1},label:{show:!1}},{type:"pie",radius:["20%","28%"],center:["68%","45%"],color:["#ffffff","red"],startAngle:105,data:[{value:30,name:"",itemStyle:{color:"transparent"}},{value:5,name:"",itemStyle:{color:"transparent"}},{value:65,name:"ddd",itemStyle:{color:"#ffffff"}}],silent:!0,labelLine:{show:!1},label:{show:!1}},{type:"pie",radius:[0,"30%"],center:["68%","45%"],startAngle:90,data:[{value:25,name:"1",itemStyle:{color:"transparent",borderWidth:4,borderColor:"#ffffff"}},{value:75,name:"2",itemStyle:{color:"transparent"}}],selectedOffset:10,silent:!0,labelLine:{show:!1},label:{show:!1}},{type:"pie",radius:["96%","97%"],center:["68%","45%"],color:["#007afe","transparent","#007afe","transparent","#007afe","transparent"],data:[{value:17,name:"11"},{value:17,name:"22"},{value:17,name:"33"},{value:17,name:"44"},{value:17,name:"55"},{value:17,name:"66"}],silent:!0,labelLine:{show:!1},label:{show:!1}},{type:"pie",zlevel:0,silent:!0,radius:["45%","46%"],center:["68%","45%"],z:10,label:{show:!1},labelLine:{show:!1},data:new Array(150).fill("").map((a,e)=>e%3===0?{name:(e+1).toString(),value:10,itemStyle:{color:"#fff",borderWidth:0,borderColor:"rgba(0,0,0,0)"}}:{name:(e+1).toString(),value:25,itemStyle:{color:"rgba(0,0,0,0)",borderWidth:0,borderColor:"rgba(0,0,0,0)"}})},{type:"pie",zlevel:0,silent:!0,radius:["58%","60%"],center:["68%","45%"],z:10,startAngle:90,label:{show:!1},color:["red","blue","red","blue"],labelLine:{show:!1},data:[{name:"r1",value:25,itemStyle:{color:{type:"linear",x:0,y:0,x2:0,y2:1,colorStops:[{offset:0,color:"rgba(51,149,191,0.5)"},{offset:1,color:"rgba(51,149,191,0)"}],global:!1}}},{name:"r2",value:25,itemStyle:{color:{type:"linear",x:0,y:0,x2:0,y2:1,colorStops:[{offset:0,color:"rgba(0,0,0,0)"},{offset:1,color:"rgba(51,149,191,0.5)"}],global:!1}}},{name:"r3",value:25,itemStyle:{color:{type:"linear",x:0,y:0,x2:0,y2:1,colorStops:[{offset:0,color:"rgba(51,149,191,0)"},{offset:1,color:"rgba(51,149,191,0.5)"}],global:!1}}},{name:"r4",value:25,itemStyle:{color:{type:"linear",x:0,y:0,x2:0,y2:1,colorStops:[{offset:0,color:"rgba(51,149,191,0.5)"},{offset:1,color:"rgba(0,0,0,0)"}],global:!1}}}]}]};return(a,e)=>(s(),f("div",m,[i(o,{option:l,resize:!1})]))}}),v=c(p,[["__scopeId","data-v-fd584460"]]);export{v as default};
