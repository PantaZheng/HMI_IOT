webpackJsonp([21],{"+Qb5":function(t,e){},"5RC8":function(t,e,i){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var a={name:"gainHeader",data:function(){return{name:sessionStorage.getItem("gainName")}}},s={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"header"},[e("div",{staticClass:"header-left back-icon"},[e("router-link",{attrs:{to:"/missionDetail"}},[e("svg",{staticClass:"icon arrow",attrs:{"aria-hidden":"true"}},[e("use",{attrs:{"xlink:href":"#iconleft"}})])])],1),this._v(" "),e("div",{staticClass:"header-title v-center"},[this._v("\n    "+this._s(this.name)+"\n  ")]),this._v(" "),e("div",{staticClass:"header-right"})])},staticRenderFns:[]};var n=i("VU/8")(a,s,!1,function(t){i("5yNg")},"data-v-0e16e5f6",null).exports,l=i("OMN4"),r=i.n(l),c={name:"gainDetail",data:function(){return{gainName:"",okGain:!1,gain:"",file:"",gainID:sessionStorage.getItem("gainID")}},methods:{getGain:function(){r.a.get(this.url+"/gain/id/"+sessionStorage.getItem("gainID")).then(this.getGainSucc)},getGainSucc:function(t){200===t.status&&(t=t.data,this.gain=t,this.gainName=t.name)},downFile:function(){window.location.href="http://bci.renjiwulian.com/gain/file/"+this.gainID},submit:function(){var t=this;this.$axios({method:"put",url:this.url+"/gain/",data:{id:parseInt(sessionStorage.getItem("gainID")),state:this.GainYes}}).then(function(e){200===e.status?(alert("提交成功！"),t.$router.push("/check")):alert(e.data)}).catch(function(t){alert(t)}),this.okGain=!1}},mounted:function(){this.getGain()}},d={render:function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"page-bg"},[i("el-divider"),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(0),t._v(" "),i("div",{staticClass:"weui-cell__bd"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[t._v("\n        "+t._s(t.gain.moduleName)+"\n      ")])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(1),t._v(" "),i("div",{staticClass:"weui-cell__bd"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[t._v("\n        "+t._s(t.gain.name)+"\n      ")])])]),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(2),t._v(" "),i("div",{staticClass:"weui-cell__bd"},[i("div",{staticClass:"weui_cell_bd weui_cell_primary"},[t._v("\n        "+t._s(t.gain.ownerName)+"\n      ")])])]),t._v(" "),i("el-divider"),t._v(" "),i("div",{staticClass:"weui-cell"},[t._m(3),t._v(" "),i("el-button",{attrs:{slot:"trigger",size:"small",type:"primary",download:"file.txt"},on:{click:t.downFile},slot:"trigger"},[t._v("查看附件")])],1),t._v(" "),i("el-divider"),t._v(" "),i("div",{staticClass:"row"},[i("span",{staticClass:"v-center"},[i("svg",{staticClass:"icon taskIcon",attrs:{"aria-hidden":"true"}},[i("use",{attrs:{"xlink:href":"#iconrenwu"}})]),t._v("\n        成果简介\n      ")])]),t._v(" "),i("el-divider"),t._v(" "),i("div",{staticClass:"weui-cell"},[t._v(t._s(t.gain.remark))]),t._v(" "),i("el-divider")],1)},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("所属课题")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("成果名称")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("成果发表者")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"weui-cell__hd"},[e("label",{staticClass:"weui-label"},[this._v("成果附件")])])}]};var u=i("VU/8")(c,d,!1,function(t){i("+Qb5")},"data-v-2742fca0",null).exports,_={name:"gain",components:{GainHeader:n,GainDetail:u}},v={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[e("gain-header"),this._v(" "),e("gain-detail")],1)},staticRenderFns:[]};var o=i("VU/8")(_,v,!1,function(t){i("Etzi")},"data-v-3e28529a",null);e.default=o.exports},"5yNg":function(t,e){},Etzi:function(t,e){}});
//# sourceMappingURL=21.7d6100118cbefaec9c3b.js.map