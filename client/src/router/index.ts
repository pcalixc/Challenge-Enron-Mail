import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/Home.vue'
import NotFound from '@/views/NotFound.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'Home',
      path: '/',
      component: HomeView,
    },
    {
      name: 'notFound',
      path: '/:pathMatch(.*)*',
      component: NotFound,
    },
  ]
})

export default router
