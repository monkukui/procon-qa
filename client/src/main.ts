import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import mavonEditor from "mavon-editor";
import "mavon-editor/dist/css/index.css";
import VueGtag from "vue-gtag";

Vue.config.productionTip = false;
Vue.use(mavonEditor);

Vue.use(VueGtag, {
  config: { id: process.env.VUE_APP_GA_ID },
});

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount("#app");
