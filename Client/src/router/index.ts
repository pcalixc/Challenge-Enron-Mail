import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      name: 'Home',
      path: '/',
      component: HomeView
      },
    {
      name: 'Login',
      path: '/login',
      component: LoginView,
    
    }
    ,
    {
      name: 'Register',
      path: '/register',
      component: RegisterView,
    }
  ]
})



export default router
