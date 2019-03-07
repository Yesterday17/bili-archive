import Vue from "vue";
import Router from "vue-router";

import Welcome from "./views/Welcome.vue";
import Environment from "./views/Environment.vue";
import Login from "./views/Login.vue";
import SelectFavorite from "./views/SelectFavorite.vue";
import Download from "./views/Download.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/step-00",
      name: "step-00",
      component: Welcome,
      meta: {
        keepAlive: true
      }
    },
    {
      path: "/step-01",
      name: "step-01",
        component: Environment,
      meta: {
        keepAlive: true
      }
    },
    {
      path: "/step-02",
      name: "step-02",
        component: Login,
      meta: {
        keepAlive: true
      }
    },
    {
      path: "/step-03",
      name: "step-03",
        component: SelectFavorite,
        meta: {
            keepAlive: true
        }
    },
      {
          path: "/step-04",
          name: "step-04",
          component: Download,
          meta: {
              keepAlive: false
          }
      }
  ]
});
