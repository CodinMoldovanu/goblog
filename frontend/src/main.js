import Vue from 'vue'
// import App from './App.vue'
import vuetify from './plugins/vuetify'
import axios from 'axios'
import router from './routes'
import store from './store'
import App from './layouts/App'
import Admin from './layouts/Admin'
import Default from './layouts/Default'
// import CKEditor from '@ckeditor/ckeditor5-vue'
import VueEditor from 'vue2-editor'

Vue.use(VueEditor)
Vue.use(axios)
Vue.component('default-layout', Default)
Vue.component('admin-layout', Admin)
// Vue.component('App')

// Vue.use(CKEditor)

Vue.prototype.$http = axios;
Vue.prototype.$api_url = "https://blog.codin.ro:1333"
const token = localStorage.getItem('accessToken');
if (token) {
  Vue.prototype.$http.defaults.headers.common['Authorization'] = 'Bearer ' + token;
  // console.log(axios.defaults.headers.common['Authorization'])
} else {
  Vue.prototype.$http.defaults.headers.common['Authorization'] = '';
}


Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
