import Vue from 'vue';
import App from './App.vue';
import router from './router';
import mavonEditor from 'mavon-editor';
import 'mavon-editor/dist/css/index.css';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.use(mavonEditor);
new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
