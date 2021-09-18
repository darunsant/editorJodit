import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Vue2Editor from "vue2-editor";
import JoditVue from 'jodit-vue'
import 'jodit/build/jodit.min.css'

Vue.config.productionTip = false
Vue.use(Vue2Editor);
Vue.use(JoditVue)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
