// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
// at.alicdn.com/t/font_1164867_8tpi7xlyslo.js
import Vue from 'vue'
import App from './App'
import router from './router'
import './assets/styles/reset.css'
import './assets/styles/border.css'
import './assets/styles/weui.css'
import './assets/styles/weuix.css'
import './assets/styles/iconfont.css'
import './assets/styles/weui.min.css'
import ElementUI from 'element-ui'
import '../node_modules/element-ui/lib/theme-chalk/index.css'
import fastClick from 'fastclick'

Vue.config.productionTip = false
Vue.use(ElementUI)
fastClick.attach(document.body)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
