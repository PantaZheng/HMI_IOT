webpackJsonp([18],{Wz6Y:function(e,t){},v2Qn:function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var n=i("mtWM"),a=i.n(n),r={name:"index",data:function(){return{projectList:[]}},methods:{getProjectList:function(){a.a.get(this.url+"/project/all").then(this.getProjectListSucc)},getProjectListSucc:function(e){(e=e.data)&&(this.projectList=e)},getPace:function(e,t){if(!e||!t)return"";e=new Date(e.replace(/-/g,"/")),t=new Date(t.replace(/-/g,"/")),console.log(e),console.log(t);var i=t.getTime()-e.getTime(),n=Math.floor(i/864e5),a=(new Date).getTime()-e.getTime(),r=(100*Math.floor(a/864e5)/n).toFixed(1);return console.log(r),parseFloat(r)},judge:function(){localStorage.getItem("id")||(this.$router.push("/"),alert("您还没有绑定用户哦！"))}},mounted:function(){this.getProjectList()},created:function(){this.judge()}},s={render:function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",e._l(e.projectList,function(t){return i("div",{key:t.id},[i("router-link",{staticClass:"weui-cell cell",attrs:{to:"/pace/"+t.id}},[i("div",{staticClass:"weui-cell__hd textColor"},[i("span",{staticClass:"iconMargin"},[i("svg",{staticClass:"icon",attrs:{"aria-hidden":"true"}},[i("use",{attrs:{"xlink:href":"#icon-"}})])]),e._v("\n        "+e._s(t.name)+"\n      ")]),e._v(" "),i("div",{staticClass:"weui-cell__bd textColor"},[i("label",{staticClass:"weui-label right"},[i("svg",{staticClass:"icon iconMargin iconColor",attrs:{"aria-hidden":"true"}},[i("use",{attrs:{"xlink:href":"#iconfuzeren3"}})]),e._v("\n          "+e._s(t.leader.name)+"\n        ")])])]),e._v(" "),i("el-progress",{staticClass:"paceMargin",attrs:{percentage:e.getPace(t.startTime,t.endTime)}})],1)}),0)},staticRenderFns:[]};var c=i("VU/8")(r,s,!1,function(e){i("Wz6Y")},"data-v-b1c175ce",null);t.default=c.exports}});
//# sourceMappingURL=18.6c9c65f0467d1dac8a7a.js.map