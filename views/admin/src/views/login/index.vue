<template>
  <div class="min-h-screen bg-gradient-to-b from-slate-50 to-slate-100/50 flex items-center justify-center p-4">
    <div class="max-w-md w-full">
      <!-- Logo -->
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-sky-500 to-indigo-500">
          管理后台
        </h1>
        <p class="mt-2 text-slate-600">欢迎回来，请登录您的账号</p>
      </div>

      <!-- 登录表单 -->
      <div class="bg-white p-8 rounded-xl shadow-lg">
        <form @submit.prevent="handleLogin" class="space-y-6">
          <!-- 用户名 -->
          <div>
            <label for="username" class="block text-sm font-medium text-slate-700 mb-1">用户名</label>
            <div class="relative">
              <input
                id="username"
                v-model="form.username"
                type="text"
                required
                class="w-full px-4 py-2 border border-slate-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent"
                :class="{ 'border-red-300': errors.username }"
              />
              <User class="absolute right-3 top-2.5 h-5 w-5 text-slate-400" />
            </div>
            <p v-if="errors.username" class="mt-1 text-sm text-red-500">{{ errors.username }}</p>
          </div>

          <!-- 密码 -->
          <div>
            <label for="password" class="block text-sm font-medium text-slate-700 mb-1">密码</label>
            <div class="relative">
              <input
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                required
                class="w-full px-4 py-2 border border-slate-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent"
                :class="{ 'border-red-300': errors.password }"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-2.5 text-slate-400 hover:text-slate-600"
              >
                <Eye v-if="showPassword" class="h-5 w-5" />
                <EyeOff v-else class="h-5 w-5" />
              </button>
            </div>
            <p v-if="errors.password" class="mt-1 text-sm text-red-500">{{ errors.password }}</p>
          </div>

          <!-- 登录按钮 -->
          <button
            type="submit"
            class="w-full py-3 bg-gradient-to-r from-sky-500 to-indigo-500 text-white rounded-lg hover:from-sky-600 hover:to-indigo-600 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:ring-offset-2 transition-all duration-300"
            :disabled="loading"
          >
            <div class="flex items-center justify-center">
              <Loader2 v-if="loading" class="animate-spin -ml-1 mr-2 h-5 w-5" />
              <span>{{ loading ? '登录中...' : '登录' }}</span>
            </div>
          </button>

          <!-- 错误提示 -->
          <p v-if="loginError" class="text-center text-sm text-red-500">{{ loginError }}</p>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { User, Eye, EyeOff, Loader2 } from 'lucide-vue-next'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const showPassword = ref(false)
const loginError = ref('')

const form = reactive({
  username: '',
  password: ''
})

const errors = reactive({
  username: '',
  password: ''
})

const handleLogin = async () => {
  // 重置错误
  errors.username = ''
  errors.password = ''
  loginError.value = ''

  // 表单验证
  if (!form.username) {
    errors.username = '请输入用户名'
    return
  }
  if (!form.password) {
    errors.password = '请输入密码'
    return
  }

  try {
    loading.value = true
    // 调用 store 的登录方法
    await userStore.login(form.username, form.password)
    
    // 登录成功后跳转到首页
    router.push('/')
  } catch (error: any) {
    loginError.value = error?.response?.data?.message || '登录失败，请重试'
    console.error('Login failed:', error)
  } finally {
    loading.value = false
  }
}
</script>