<template>
  <div>
    <h1 class="mb-6 text-2xl font-bold">仪表盘</h1>
    
    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">用户总数</h3>
          <p class="mt-2 text-3xl font-bold">12,345</p>
          <p class="mt-2 text-sm text-green-500">较上月增长5.2%</p>
        </div>
        <Users class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">总收入</h3>
          <p class="mt-2 text-3xl font-bold">¥54,321</p>
          <p class="mt-2 text-sm text-green-500">较上月增长3.7%</p>
        </div>
        <DollarSign class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">进行中的项目</h3>
          <p class="mt-2 text-3xl font-bold">48</p>
          <p class="mt-2 text-sm text-red-500">较上月下降2.3%</p>
        </div>
        <FolderOpen class="h-12 w-12 text-gray-400" />
      </Card>

      <Card class="flex items-center p-6">
        <div class="flex-1">
          <h3 class="text-sm font-medium text-gray-500">已完成任务</h3>
          <p class="mt-2 text-3xl font-bold">256</p>
          <p class="mt-2 text-sm text-green-500">较上月增长12.4%</p>
        </div>
        <CheckSquare class="h-12 w-12 text-gray-400" />
      </Card>
    </div>

    <!-- 最近活动 -->
    <Card class="mt-6">
      <div class="border-b p-6">
        <h2 class="text-lg font-medium">最近活动</h2>
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
    title: '新的报告已生成',
    time: '5分钟前'
  },
  {
    icon: UserPlus,
    title: '新用户注册',
    time: '1小时前'
  },
  {
    icon: Settings,
    title: '系统更新完成',
    time: '2小时前'
  }
])
</script>
