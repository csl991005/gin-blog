import Vue from 'vue'

import App from './App.vue'
import router from './router'
import Antd from 'ant-design-vue'
import 'ant-design-vue/dist/antd.css'
import axios from 'axios'
import './assets/css/style.css'

axios.defaults.baseURL = "http://localhost:8080/api/v1"
Vue.prototype.$http = axios
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
