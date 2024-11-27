import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '../layouts/default.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: [
        {
          path: '',
          redirect: '/dashboard'
        },
        {
          path: '/dashboard',
          component: () => import('../views/dashboard/index.vue')
        }
      ]
    },
    {
      path: '/login',
      component: () => import('../views/login/index.vue')
    }
  ]
})

export default router
