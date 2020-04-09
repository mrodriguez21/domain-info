import Vue from "vue";
import App from "./App.vue";
import router from "./routes";

import {
  TablePlugin,
  ButtonPlugin,
  FormPlugin,
  FormGroupPlugin,
  FormInputPlugin,
  InputGroupPlugin,
  TooltipPlugin,
} from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import "mdbvue/lib/css/mdb.min.css";
import "@fortawesome/fontawesome-free/css/all.min.css";

Vue.config.productionTip = false;
Vue.use(TablePlugin);
Vue.use(ButtonPlugin);
Vue.use(FormPlugin);
Vue.use(FormGroupPlugin);
Vue.use(FormInputPlugin);
Vue.use(InputGroupPlugin);
Vue.use(TooltipPlugin);

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");
