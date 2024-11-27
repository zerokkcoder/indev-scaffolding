<template>
  <div>
    <h1 class="mb-6 text-2xl font-bold">Dashboard</h1>
    
    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">Total Users</h3>
          <p class="mt-2 text-3xl font-bold">12,345</p>
          <p class="mt-2 text-sm text-green-500">+5.2% from last month</p>
        </div>
        <Users class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">Total Revenue</h3>
          <p class="mt-2 text-3xl font-bold">$54,321</p>
          <p class="mt-2 text-sm text-green-500">+3.7% from last month</p>
        </div>
        <DollarSign class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">Active Projects</h3>
          <p class="mt-2 text-3xl font-bold">48</p>
          <p class="mt-2 text-sm text-red-500">-2.3% from last month</p>
        </div>
        <FolderOpen class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">Tasks Completed</h3>
          <p class="mt-2 text-3xl font-bold">256</p>
          <p class="mt-2 text-sm text-green-500">+12.4% from last month</p>
        </div>
        <CheckSquare class="h-12 w-12 text-gray-400" />
      </Card>
    </div>

    <!-- 最近活动 -->
    <Card class="mt-6">
      <div class="border-b p-6">
        <h2 class="text-lg font-medium">Recent Activity</h2>
      </div>
      <div class="divide-y">
        <div v-for="(activity, index) in recentActivities" :key="index" class="flex items-center p-6">
          <div class="mr-4 rounded-full bg-blue-100 p-2">
            <component :is="activity.icon" class="h-5 w-5 text-blue-600" />
          </div>
          <div class="flex-1">
            <p class="font-medium">{{ activity.title }}</p>
            <p class="text-sm text-gray-500">{{ activity.time }}</p>
          </div>
        </div>
      </div>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Users, DollarSign, FolderOpen, CheckSquare, FileText, UserPlus, Settings } from 'lucide-vue-next'
import Card from '../../components/ui/card.vue'

const recentActivities = ref([
  {
    icon: FileText,
    title: 'New report has been generated',
    time: '5 minutes ago'
  },
  {
    icon: UserPlus,
    title: 'New user registered',
    time: '1 hour ago'
  },
  {
    icon: Settings,
    title: 'System update completed',
    time: '2 hours ago'
  }
])
</script>
