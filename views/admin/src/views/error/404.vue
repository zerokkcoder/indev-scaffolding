<template>
  <div class="min-h-screen bg-gradient-to-b from-slate-50 to-slate-100/50 flex items-center justify-center p-4">
    <div class="max-w-2xl w-full text-center">
      <!-- 404 数字动画 -->
      <div class="relative">
        <h1 class="text-[180px] font-bold text-slate-900/5">404</h1>
        <div class="absolute inset-0 flex items-center justify-center">
          <h2 class="text-8xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-sky-500 to-indigo-500 animate-pulse">
            404
          </h2>
        </div>
      </div>

      <!-- 错误信息 -->
      <div class="mt-8 space-y-4">
        <h3 class="text-2xl font-semibold text-slate-900">页面未找到</h3>
        <p class="text-slate-600 max-w-md mx-auto">
          抱歉，您访问的页面不存在或已被移除。请检查 URL 是否正确，或返回首页继续浏览。
        </p>
      </div>

      <!-- 按钮组 -->
      <div class="mt-12 flex items-center justify-center gap-4">
        <router-link 
          to="/" 
          class="px-6 py-3 bg-gradient-to-r from-sky-500 to-indigo-500 text-white rounded-lg hover:from-sky-600 hover:to-indigo-600 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:ring-offset-2 transition-all duration-300 shadow-lg hover:shadow-xl"
        >
          返回首页
        </router-link>
        <button 
          @click="$router.go(-1)"
          class="px-6 py-3 border border-slate-200 text-slate-700 rounded-lg hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-slate-500 focus:ring-offset-2 transition-all duration-300"
        >
          返回上一页
        </button>
      </div>

      <!-- 装饰元素 -->
      <div class="mt-16 flex items-center justify-center gap-8 text-slate-400">
        <FileSearch class="w-16 h-16 animate-bounce" />
        <AlertCircle class="w-12 h-12 animate-pulse" />
        <FileQuestion class="w-16 h-16 animate-bounce" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { FileSearch, AlertCircle, FileQuestion } from 'lucide-vue-next'
</script>
