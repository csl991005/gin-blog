import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugin/http'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import './assets/css/style.css'


Vue.use(Antd);
Vue.prototype.$message.config({
  top: `60px`,
  duration: 2,
  maxCount: 3,
});

new Vue({
  router,
  render: (h) => h(App)
}).$mount('#app')
