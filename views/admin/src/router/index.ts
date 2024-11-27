import { createRouter, createWebHashHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
      layout: 'blank'
    }
  },
  {
    path: '/',
    component: () => import('@/layouts/default.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'dashboard',
        component: () => import('@/views/dashboard/index.vue'),
        meta: {
          title: '仪表盘',
          requiresAuth: true
        }
      },
      {
        path: 'users',
        name: 'users',
        children: [
          {
            path: 'list',
            name: 'usersList',
            component: () => import('@/views/users/list.vue'),
            meta: {
              title: '用户管理',
              requiresAuth: true
            }
          }
        ]
      },
      {
        path: 'content',
        name: 'content',
        children: [
          {
            path: 'articles',
            name: 'articles',
            component: () => import('@/views/content/articles.vue'),
            meta: {
              title: '文章管理',
              requiresAuth: true
            }
          },
          {
            path: 'comments',
            name: 'comments',
            component: () => import('@/views/content/comments.vue'),
            meta: {
              title: '评论管理',
              requiresAuth: true
            }
          }
        ]
      },
      {
        path: 'products',
        name: 'products',
        component: () => import('@/views/products/index.vue'),
        meta: {
          title: '产品管理',
          requiresAuth: true
        }
      },
      {
        path: 'permissions',
        name: 'permissions',
        component: () => import('@/views/permissions/index.vue'),
        meta: {
          title: '权限管理',
          requiresAuth: true
        }
      },
      {
        path: 'settings',
        name: 'settings',
        component: () => import('@/views/settings/index.vue'),
        meta: {
          title: '设置',
          requiresAuth: true
        }
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: '404',
    component: () => import('@/views/error/404.vue'),
    meta: {
      title: '404',
      layout: 'blank'
    }
  }
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = `${to.meta.title} - Admin Dashboard`

  // 检查是否需要认证
  if (to.meta.requiresAuth) {
    // 这里添加你的认证逻辑
    const isAuthenticated = localStorage.getItem('token')
    if (!isAuthenticated) {
      next({ name: 'login' })
      return
    }
  }

  next()
})

export default router
