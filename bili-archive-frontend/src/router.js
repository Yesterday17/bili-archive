import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/step-00",
      name: "Welcome",
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/Welcome.vue")
    },
    {
      path: "/step-01",
      name: "Login",
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/Login.vue")
    },
    {
      path: "/step-02",
      name: "SelectFavorite",
      component: () =>
        import(/* webpackChunkName: "about" */ "./views/SelectFavorite.vue")
    }
  ]
});
