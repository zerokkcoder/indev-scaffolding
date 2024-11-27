<template>
  <div class="p-6">
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-gray-900">文章列表</h1>
      <p class="mt-2 text-sm text-gray-600">管理系统内的所有文章</p>
    </div>

    <!-- 搜索和过滤区域 -->
    <div class="mb-6 flex items-center justify-between">
      <div class="flex gap-4">
        <div class="relative">
          <input
            type="text"
            placeholder="搜索文章..."
            class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
          <Search class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
        </div>
        <select class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option value="">所有分类</option>
          <option value="news">新闻</option>
          <option value="tech">技术</option>
          <option value="guide">教程</option>
        </select>
      </div>
      <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2">
        新建文章
      </button>
    </div>

    <!-- 文章列表 -->
    <div class="space-y-4">
      <div v-for="article in articles" :key="article.id" class="bg-white rounded-lg shadow p-6">
        <div class="flex items-center justify-between">
          <div class="flex-1">
            <h3 class="text-lg font-medium text-gray-900">{{ article.title }}</h3>
            <div class="mt-1 flex items-center text-sm text-gray-500">
              <span>{{ article.author }}</span>
              <span class="mx-2">·</span>
              <span>{{ article.category }}</span>
              <span class="mx-2">·</span>
              <span>{{ article.date }}</span>
            </div>
          </div>
          <div class="flex items-center gap-4">
            <span
              class="px-2 py-1 text-xs font-semibold rounded-full"
              :class="{
                'bg-green-100 text-green-800': article.status === 'published',
                'bg-yellow-100 text-yellow-800': article.status === 'draft',
                'bg-red-100 text-red-800': article.status === 'archived'
              }"
            >
              {{ statusText[article.status] }}
            </span>
            <div class="flex items-center gap-2">
              <button class="p-2 text-gray-400 hover:text-blue-600">
                <Edit2 class="h-4 w-4" />
              </button>
              <button class="p-2 text-gray-400 hover:text-red-600">
                <Trash2 class="h-4 w-4" />
              </button>
            </div>
          </div>
        </div>
        <p class="mt-3 text-sm text-gray-600">{{ article.excerpt }}</p>
        <div class="mt-4 flex items-center gap-2">
          <span v-for="tag in article.tags" :key="tag" class="px-2 py-1 text-xs bg-gray-100 text-gray-600 rounded-full">
            {{ tag }}
          </span>
        </div>
      </div>
    </div>

    <!-- 分页 -->
    <div class="mt-6 flex items-center justify-between">
      <div class="text-sm text-gray-700">
        显示 <span class="font-medium">1</span> 到 <span class="font-medium">10</span> 条，共 <span class="font-medium">30</span> 条
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
import { Search, Edit2, Trash2 } from 'lucide-vue-next'

interface Article {
  id: number
  title: string
  author: string
  category: string
  date: string
  status: 'published' | 'draft' | 'archived'
  excerpt: string
  tags: string[]
}

const statusText = {
  published: '已发布',
  draft: '草稿',
  archived: '已归档'
}

const articles = ref<Article[]>([
  {
    id: 1,
    title: '如何提高代码质量：10个实用技巧',
    author: '张三',
    category: '技术',
    date: '2023-05-01',
    status: 'published',
    excerpt: '代码质量直接影响着项目的可维护性和可扩展性。本文将介绍10个提高代码质量的实用技巧，帮助你写出更好的代码。',
    tags: ['编程技巧', '代码质量', '最佳实践']
  },
  {
    id: 2,
    title: '2024年前端开发趋势展望',
    author: '李四',
    category: '新闻',
    date: '2023-05-15',
    status: 'draft',
    excerpt: '随着技术的不断发展，前端开发领域也在持续变革。本文将探讨2024年前端开发可能出现的新趋势。',
    tags: ['前端开发', '技术趋势', 'Web开发']
  },
  {
    id: 3,
    title: 'Vue.js 3.0完全指南',
    author: '王五',
    category: '教程',
    date: '2023-06-01',
    status: 'archived',
    excerpt: '本文将全面介绍Vue.js 3.0的新特性，包括Composition API、响应式系统的改进等内容。',
    tags: ['Vue.js', '前端框架', '教程']
  }
])
</script>
