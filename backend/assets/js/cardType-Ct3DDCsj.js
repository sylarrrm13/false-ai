import{Z as t,_ as r}from"./index-mkXm0lbT.js";const d=e=>(e.pageSize==null&&(e.pageSize=15),t.get(r+"/cardType",{filter:e.filter,pageNum:e.pageNum,pageSize:e.pageSize})),p=e=>t.delete(r+"/cardType/",e,{loading:!0}),i=e=>t.put(r+"/cardType/",e),g=e=>t.post(r+"/cardType/",e);export{g as a,p as d,d as g,i as u};
