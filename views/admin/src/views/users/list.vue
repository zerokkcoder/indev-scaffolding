<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">用户列表</h1>
      <p class="mt-2 text-sm text-gray-600">管理系统用户账号</p>
    </div>

    <!-- 搜索和过滤区域 -->
    <div class="mb-6 flex items-center justify-between">
      <div class="flex gap-4">
        <div class="relative">
          <input
            type="text"
            placeholder="搜索用户..."
            class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <Search class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
        </div>
        <select class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">所有状态</option>
          <option value="active">活跃</option>
          <option value="inactive">未激活</option>
          <option value="blocked">已封禁</option>
        </select>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
        添加用户
      </button>
    </div>

    <!-- 用户列表表格 -->
    <div class="bg-white rounded-lg shadow overflow-hidden">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">用户</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">角色</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">状态</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">注册时间</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">操作</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="user in users" :key="user.id">
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="flex items-center">
                <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center">
                  <User class="h-6 w-6 text-gray-400" v-if="!user.avatar" />
                  <img v-else :src="user.avatar" :alt="user.name" class="h-10 w-10 rounded-full" />
                </div>
                <div class="ml-4">
                  <div class="text-sm font-medium text-gray-900">{{ user.name }}</div>
                  <div class="text-sm text-gray-500">{{ user.email }}</div>
                </div>
              </div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <div class="text-sm text-gray-900">{{ user.role }}</div>
            </td>
            <td class="px-6 py-4 whitespace-nowrap">
              <span
                class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                :class="{
                  'bg-green-100 text-green-800': user.status === 'active',
                  'bg-yellow-100 text-yellow-800': user.status === 'inactive',
                  'bg-red-100 text-red-800': user.status === 'blocked'
                }"
              >
                {{ statusText[user.status] }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
              {{ user.createdAt }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button class="text-blue-600 hover:text-blue-900 mr-3">编辑</button>
              <button class="text-red-600 hover:text-red-900">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- 分页 -->
    <div class="mt-6 flex items-center justify-between">
      <div class="text-sm text-gray-700">
        显示 <span class="font-medium">1</span> 到 <span class="font-medium">10</span> 条，共 <span class="font-medium">50</span> 条
      </div>
      <div class="flex gap-2">
        <button class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50">上一页</button>
        <button class="px-3 py-1 border border-gray-300 rounded-lg hover:bg-gray-50">下一页</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { User, Search } from 'lucide-vue-next'

interface UserData {
  id: number
  name: string
  email: string
  role: string
  status: 'active' | 'inactive' | 'blocked'
  avatar?: string
  createdAt: string
}

const statusText = {
  active: '活跃',
  inactive: '未激活',
  blocked: '已封禁'
}

const users = ref<UserData[]>([
  {
    id: 1,
    name: '张三',
    email: 'zhangsan@example.com',
    role: '管理员',
    status: 'active',
    createdAt: '2023-01-01'
  },
  {
    id: 2,
    name: '李四',
    email: 'lisi@example.com',
    role: '编辑',
    status: 'inactive',
    createdAt: '2023-02-15'
  },
  {
    id: 3,
    name: '王五',
    email: 'wangwu@example.com',
    role: '用户',
    status: 'blocked',
    createdAt: '2023-03-20'
  }
])
</script>
