import { createRouter, createWebHistory } from 'vue-router'
import StocksView from "@/views/StocksView.vue";
import RecommendationsView from "@/views/RecommendationsView.vue";

const routes = [
  {
    path: '/',
    name: 'stocks',
    component: StocksView,
  },
  {
    path: '/recommendations',
    name: 'recommendations',
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: RecommendationsView,
  },
]
const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
