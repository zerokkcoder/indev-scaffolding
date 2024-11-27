import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/layouts/default.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/index.vue')
      },
      {
        path: 'users',
        name: 'Users',
        children: [
          {
            path: 'list',
            name: 'UserList',
            component: () => import('@/views/users/list.vue')
          }
        ]
      },
      {
        path: 'content',
        name: 'Content',
        children: [
          {
            path: 'articles',
            name: 'Articles',
            component: () => import('@/views/content/articles.vue')
          },
          {
            path: 'comments',
            name: 'Comments',
            component: () => import('@/views/content/comments.vue')
          }
        ]
      },
      {
        path: 'permissions',
        name: 'Permissions',
        component: () => import('@/views/permissions/index.vue')
      },
      {
        path: 'products',
        name: 'Products',
        component: () => import('@/views/products/index.vue')
      },
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('@/views/settings/index.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
