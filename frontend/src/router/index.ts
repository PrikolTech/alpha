import { createRouter, createWebHistory } from "vue-router";
import BaseLayout from "@/views/BaseLayout.vue";
import StartView from "@/views/StartView.vue";

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/:pathMatch(.*)*",
      redirect: { name: "start" },
    },
    {
      path: "/",
      name: "base",
      component: BaseLayout,
      redirect: { name: "start" },
      children: [
        {
          path: "/start",
          name: "start",
          component: StartView,
        },
      ],
    },
  ],
  scrollBehavior(to, _from, savedPosition) {
    if (to.hash) {
      return {
        el: to.hash,
        behavior: "smooth",
      };
    } else if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0 };
    }
  },
});

router.beforeEach(async (to, from) => {});

export default router;
