<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">评论管理</h1>
      <p class="mt-2 text-sm text-gray-600">管理文章评论和用户反馈</p>
    </div>

    <!-- 搜索和过滤区域 -->
    <div class="mb-6 flex items-center justify-between">
      <div class="flex gap-4">
        <div class="relative">
          <input
            type="text"
            placeholder="搜索评论..."
            class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <Search class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
        </div>
        <select class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">所有状态</option>
          <option value="approved">已通过</option>
          <option value="pending">待审核</option>
          <option value="spam">垃圾评论</option>
        </select>
      </div>
      <div class="flex gap-2">
        <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2">
          批量通过
        </button>
        <button class="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2">
          批量删除
        </button>
      </div>
    </div>

    <!-- 评论列表 -->
    <div class="space-y-4">
      <div v-for="comment in comments" :key="comment.id" class="bg-white rounded-lg shadow p-6">
        <div class="flex items-start gap-4">
          <div class="flex-shrink-0">
            <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center">
              <User class="h-6 w-6 text-gray-400" v-if="!comment.avatar" />
              <img v-else :src="comment.avatar" :alt="comment.author" class="h-10 w-10 rounded-full" />
            </div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-sm font-medium text-gray-900">{{ comment.author }}</h3>
                <p class="text-sm text-gray-500">
                  评论于 {{ comment.date }} · 文章：{{ comment.article }}
                </p>
              </div>
              <span
                class="px-2 py-1 text-xs font-semibold rounded-full"
                :class="{
                  'bg-green-100 text-green-800': comment.status === 'approved',
                  'bg-yellow-100 text-yellow-800': comment.status === 'pending',
                  'bg-red-100 text-red-800': comment.status === 'spam'
                }"
              >
                {{ statusText[comment.status] }}
              </span>
            </div>
            <div class="mt-2">
              <p class="text-sm text-gray-600">{{ comment.content }}</p>
            </div>
            <div class="mt-4 flex items-center gap-4">
              <button 
                v-if="comment.status !== 'approved'"
                class="text-sm text-green-600 hover:text-green-800 flex items-center gap-1"
              >
                <CheckCircle class="h-4 w-4" />
                通过
              </button>
              <button 
                v-if="comment.status !== 'spam'"
                class="text-sm text-yellow-600 hover:text-yellow-800 flex items-center gap-1"
              >
                <AlertTriangle class="h-4 w-4" />
                标记为垃圾评论
              </button>
              <button class="text-sm text-blue-600 hover:text-blue-800 flex items-center gap-1">
                <MessageSquare class="h-4 w-4" />
                回复
              </button>
              <button class="text-sm text-red-600 hover:text-red-800 flex items-center gap-1">
                <Trash2 class="h-4 w-4" />
                删除
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分页 -->
    <div class="mt-6 flex items-center justify-between">
      <div class="text-sm text-gray-700">
        显示 <span class="font-medium">1</span> 到 <span class="font-medium">10</span> 条，共 <span class="font-medium">25</span> 条
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
import { Search, User, CheckCircle, AlertTriangle, MessageSquare, Trash2 } from 'lucide-vue-next'

interface Comment {
  id: number
  author: string
  avatar?: string
  date: string
  content: string
  article: string
  status: 'approved' | 'pending' | 'spam'
}

const statusText = {
  approved: '已通过',
  pending: '待审核',
  spam: '垃圾评论'
}

const comments = ref<Comment[]>([
  {
    id: 1,
    author: '张三',
    date: '2023-06-15 10:30',
    content: '这篇文章写得非常好，对我帮助很大！期待更多类似的内容。',
    article: '如何提高代码质量：10个实用技巧',
    status: 'approved'
  },
  {
    id: 2,
    author: '李四',
    date: '2023-06-15 11:45',
    content: '文章中提到的第三点我有不同的看法，我认为...',
    article: 'Vue.js 3.0完全指南',
    status: 'pending'
  },
  {
    id: 3,
    author: '王五',
    date: '2023-06-15 14:20',
    content: 'Buy cheap watches! Visit our website...',
    article: '2024年前端开发趋势展望',
    status: 'spam'
  }
])
</script>
