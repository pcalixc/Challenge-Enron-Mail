import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/Home.vue'
import NotFounded from '@/views/NotFounded.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'Home',
      path: '/mail-searcher',
      component: HomeView,
    },
    {
      name: 'notFound',
      path: '/:pathMatch(.*)*',
      component: NotFounded,
    },
  ]
})

export default router
