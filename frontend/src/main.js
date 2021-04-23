import Vue from 'vue'
import Router from 'vue-router'
import App from './App.vue'

import {routes} from './routes'


Vue.config.productionTip = false


const router = new Router({
  routes,
  mode: 'history',
})


Vue.use(Router)

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
