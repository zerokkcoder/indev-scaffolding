<template>
  <div class="min-h-screen bg-gray-100">
    <!-- 侧边栏 -->
    <aside class="fixed left-0 top-0 z-40 h-screen bg-white transition-transform duration-300" :class="[isCollapsed ? 'w-16' : 'w-64']">
      <div class="flex h-16 items-center justify-center border-b px-4">
        <h1 class="text-xl font-bold truncate" :class="{ 'text-sm': isCollapsed }">{{ isCollapsed ? '后台' : '管理后台' }}</h1>
      </div>
      <nav class="space-y-1 p-4" :class="{ 'px-2': isCollapsed }">
        <router-link
          v-for="item in menuItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center rounded-lg py-2 text-sm font-medium text-gray-600 hover:bg-gray-100 relative group"
          :class="[
            isCurrentRoute(item.path) ? 'bg-gray-100' : '',
            isCollapsed ? 'justify-center px-2' : 'px-4'
          ]"
        >
          <component :is="item.icon" class="h-5 w-5" :class="{ 'mr-3': !isCollapsed }" />
          <span v-if="!isCollapsed">{{ item.name }}</span>
          <!-- 悬浮提示 -->
          <div
            v-if="isCollapsed"
            class="absolute left-full ml-2 px-2 py-1 bg-gray-800 text-white text-xs rounded invisible group-hover:visible whitespace-nowrap z-50"
          >
            {{ item.name }}
          </div>
        </router-link>
      </nav>
    </aside>

    <!-- 主要内容区域 -->
    <div :class="[isCollapsed ? 'pl-16' : 'pl-64']" class="transition-[padding] duration-300">
      <!-- 顶部导航栏 -->
      <header class="flex h-16 items-center justify-between border-b bg-white px-6">
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
      <main class="p-6">
        <router-view></router-view>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { LayoutDashboard, Users, Settings, Menu, Bell, User } from 'lucide-vue-next'
import { useLayoutStore } from '../stores/layout'
import { storeToRefs } from 'pinia'
import { ref } from 'vue'

const route = useRoute()
const layoutStore = useLayoutStore()
const { isCollapsed } = storeToRefs(layoutStore)

const menuItems = ref([
  {
    name: '仪表盘',
    path: '/dashboard',
    icon: LayoutDashboard
  },
  {
    name: '用户',
    path: '/users',
    icon: Users
  },
  {
    name: '设置',
    path: '/settings',
    icon: Settings
  }
])

const isCurrentRoute = (path: string) => {
  return route.path.startsWith(path)
}
</script>
