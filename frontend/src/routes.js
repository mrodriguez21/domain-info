import Vue from "vue";
import VueRouter from "vue-router";

import Search from "./components/Search";
import History from "./components/History";

Vue.use(VueRouter);

export default new VueRouter({
  mode: "history",
  routes: [
    { path: "/", component: Search },
    { path: "/history", component: History },
  ],
});
