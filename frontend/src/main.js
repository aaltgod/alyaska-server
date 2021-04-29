import Vue from 'vue'
import Router from 'vue-router'
import App from './App.vue'
import Vuesax from "vuesax"

import "vuesax/dist/vuesax.css"
import 'material-icons/iconfont/material-icons.css'

import {routes} from './routes'


Vue.config.productionTip = false

const router = new Router({
  routes,
  mode: 'history',
})


Vue.use(Router)
Vue.use(Vuesax, {

})

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
