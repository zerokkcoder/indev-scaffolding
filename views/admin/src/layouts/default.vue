<template>
  <div class="min-h-screen bg-gray-100">
    <!-- 侧边栏 -->
    <aside class="fixed left-0 top-0 z-40 h-screen bg-white transition-transform duration-300" :class="[isCollapsed ? 'w-14' : 'w-56']">
      <div class="flex h-16 items-center justify-center border-b px-4">
        <h1 class="text-xl font-bold truncate" :class="{ 'text-sm': isCollapsed }">{{ isCollapsed ? '后台' : '管理后台' }}</h1>
      </div>
      <nav class="space-y-0.5 py-4" :class="{ 'px-2': isCollapsed, 'px-3': !isCollapsed }">
        <div
          v-for="item in menuItems"
          :key="item.name"
        >
          <div
            class="flex items-center rounded-lg py-2 text-sm font-medium text-gray-600 hover:bg-gray-100 relative group cursor-pointer"
            :class="{
              'bg-gray-100': item.children ? layoutStore.activeParentMenu === item.name : layoutStore.currentPath === item.path,
              'text-blue-600': layoutStore.activeParentMenu === item.name,
              'justify-center px-2': isCollapsed,
              'px-3': !isCollapsed
            }"
            @click="handleMenuClick(item)"
          >
            <component :is="item.icon" class="h-4 w-4" :class="{ 'mr-2': !isCollapsed }" />
            <span v-if="!isCollapsed" class="text-sm">
              {{ item.name }}
              <span v-if="item.children" class="ml-1 text-gray-400 text-xs">({{ item.children.length }})</span>
            </span>
            <!-- 悬浮提示 -->
            <div
              v-if="isCollapsed"
              class="absolute left-full ml-2 px-2 py-1 bg-gray-800 text-white text-xs rounded invisible group-hover:visible whitespace-nowrap z-50"
            >
              {{ item.name }}
              <span v-if="item.children" class="ml-1 text-gray-400">({{ item.children.length }})</span>
            </div>
          </div>
        </div>
      </nav>
    </aside>

    <!-- 主要内容区域 -->
    <div :class="[isCollapsed ? 'pl-14' : 'pl-56']" class="transition-[padding] duration-300">
      <!-- 顶部导航栏 -->
      <header class="fixed top-0 right-0 z-20 h-16 bg-white border-b flex items-center justify-between px-6"
        :class="[isCollapsed ? 'left-14' : 'left-56']">
        <div class="flex items-center">
          <button class="rounded-lg p-2 hover:bg-gray-100" @click="layoutStore.toggleSidebar">
            <Menu class="h-5 w-5" />
          </button>
        </div>
        <div class="flex items-center space-x-4">
          <button class="rounded-lg p-2 hover:bg-gray-100">
            <Bell class="h-5 w-5" />
          </button>
          <button class="rounded-lg p-2 hover:bg-gray-100">
            <User class="h-5 w-5" />
          </button>
        </div>
      </header>

      <!-- 页面内容 -->
      <main class="pt-16 min-h-screen bg-gray-100">
        <div class="flex min-h-[calc(100vh-4rem)]">
          <!-- 二级菜单 -->
          <div
            v-show="showSubmenu"
            class="w-40 bg-white border-r border-l border-gray-200 transition-all duration-300 min-h-full"
          >
            <div class="sticky top-0 p-3">
              <h2 class="text-sm font-medium text-gray-500 mb-3">{{ activeParentMenu }}</h2>
              <nav class="space-y-0.5">
                <router-link
                  v-for="subItem in currentSubmenuItems"
                  :key="subItem.path"
                  :to="subItem.path!"
                  class="flex items-center rounded-lg px-3 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
                  :class="{ 'bg-gray-100': layoutStore.isCurrentRoute(subItem.path!) }"
                >
                  <component :is="subItem.icon" class="h-3.5 w-3.5 mr-2" />
                  {{ subItem.name }}
                </router-link>
              </nav>
            </div>
          </div>

          <!-- 路由视图 -->
          <div class="flex-1 p-6">
            <router-view></router-view>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { 
  Menu, 
  LayoutDashboard,
  Users,
  Settings,
  Bell,
  User,
  FileText,
  UserPlus,
  Shield,
  Box,
  Mail,
  MessageSquare
} from 'lucide-vue-next'
import { useLayoutStore } from '@/stores/layout'
import type { MenuItem } from '@/types'

const router = useRouter()
const route = useRoute()
const layoutStore = useLayoutStore()

const isCollapsed = computed(() => layoutStore.isCollapsed)
const activeParentMenu = computed(() => layoutStore.activeParentMenu)
const showSubmenu = computed(() => activeParentMenu.value && !isCollapsed.value)

// 菜单配置
const menuItems = ref<MenuItem[]>([
  {
    name: '仪表盘',
    path: '/dashboard',
    icon: LayoutDashboard
  },
  {
    name: '用户管理',
    icon: Users,
    children: [
      {
        name: '用户列表',
        path: '/users/list',
        icon: User
      }
    ]
  },
  {
    name: '内容管理',
    icon: FileText,
    children: [
      {
        name: '文章列表',
        path: '/content/articles',
        icon: Mail
      },
      {
        name: '评论管理',
        path: '/content/comments',
        icon: MessageSquare
      }
    ]
  },
  {
    name: '权限管理',
    path: '/permissions',
    icon: Shield
  },
  {
    name: '产品管理',
    path: '/products',
    icon: Box
  },
  {
    name: '系统设置',
    path: '/settings',
    icon: Settings
  }
])

// 找到当前路由对应的菜单项
const findParentMenu = (path: string): MenuItem | null => {
  for (const item of menuItems.value) {
    if (item.children) {
      const hasMatchingChild = item.children.some(child => child.path === path)
      if (hasMatchingChild) {
        return item
      }
    }
  }
  return null
}

// 监听路由变化
watch(
  () => route.path,
  (newPath) => {
    try {
      console.log('Route changed:', {
        path: newPath,
        route: route
      })
      
      // 更新当前路径
      layoutStore.setCurrentPath(newPath)
      
      // 如果找到父菜单，则展开它
      const parentMenu = findParentMenu(newPath)
      if (parentMenu) {
        layoutStore.setActiveParentMenu(parentMenu.name)
      } else {
        // 如果没有找到父菜单，说明是一级菜单，清除父菜单状态
        layoutStore.setActiveParentMenu(null)
      }
    } catch (error) {
      console.error('Error in route watcher:', error)
    }
  },
  { immediate: true }
)

const handleMenuClick = (item: MenuItem) => {
  try {
    console.log('Menu clicked:', {
      item,
      currentPath: route.path
    })
    
    if (item.children) {
      // 如果侧边栏是收缩状态，先展开它
      if (isCollapsed.value) {
        layoutStore.toggleSidebar()
      }
      
      // 如果点击的是当前已经展开的菜单，则关闭它
      if (layoutStore.activeParentMenu === item.name) {
        layoutStore.setActiveParentMenu(null)
      } else {
        // 否则展开新点击的菜单，并导航到第一个子菜单
        layoutStore.setActiveParentMenu(item.name)
        if (item.children.length > 0 && item.children[0].path) {
          router.push(item.children[0].path)
        }
      }
    } else if (item.path) {
      // 如果是没有子菜单的项目，清除父菜单状态并导航
      layoutStore.setActiveParentMenu(null)
      router.push(item.path)
    }
  } catch (error) {
    console.error('Error in menu click handler:', error)
  }
}

const currentSubmenuItems = computed(() => {
  const parentItem = menuItems.value.find(item => item.name === activeParentMenu.value)
  return parentItem?.children || []
})
</script>
