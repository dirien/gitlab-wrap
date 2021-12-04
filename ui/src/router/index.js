import { createWebHistory, createRouter } from "vue-router";

import Home from "../views/Home.vue";
import PageNotFound from "../views/PageNotFound.vue";

const routes = [
  {
    path: "/",
    name: "WrapIt",
    component: Home,
  },
  {
    name: "WrapItResult",
    path: "/:username",
    component: Home,
    props: true,
  },
  {
    path: "/:catchAll(.*)*",
    name: "PageNotFound",
    component: PageNotFound,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
