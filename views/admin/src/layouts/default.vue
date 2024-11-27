<template>
  <div class="min-h-screen bg-gray-100">
    <!-- 侧边栏 -->
    <aside class="fixed left-0 top-0 z-40 h-screen bg-white transition-transform duration-300" :class="[isCollapsed ? '-translate-x-full' : '', 'w-64']">
      <div class="flex h-16 items-center justify-center border-b px-4">
        <h1 class="text-xl font-bold">Admin Panel</h1>
      </div>
      <nav class="space-y-1 p-4">
        <router-link
          v-for="item in menuItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center rounded-lg px-4 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
          :class="{ 'bg-gray-100': isCurrentRoute(item.path) }"
        >
          <component :is="item.icon" class="mr-3 h-5 w-5" />
          {{ item.name }}
        </router-link>
      </nav>
    </aside>

    <!-- 主要内容区域 -->
    <div :class="[isCollapsed ? 'pl-0' : 'pl-64']" class="transition-[padding] duration-300">
      <!-- 顶部导航栏 -->
      <header class="flex h-16 items-center justify-between border-b bg-white px-6">
        <div class="flex items-center">
          <button class="rounded-lg p-2 hover:bg-gray-100" @click="toggleSidebar">
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
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import { LayoutDashboard, Users, Settings, Menu, Bell, User } from 'lucide-vue-next'

const route = useRoute()
const isCollapsed = ref(false)

const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
}

const menuItems = ref([
  {
    name: 'Dashboard',
    path: '/dashboard',
    icon: LayoutDashboard
  },
  {
    name: 'Users',
    path: '/users',
    icon: Users
  },
  {
    name: 'Settings',
    path: '/settings',
    icon: Settings
  }
])

const isCurrentRoute = (path: string) => {
  return route.path.startsWith(path)
}
</script>
