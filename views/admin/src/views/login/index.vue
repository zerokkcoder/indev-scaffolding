<template>
  <div class="min-h-screen flex items-center justify-center bg-background">
    <div class="w-full max-w-md">
      <Card>
        <CardHeader>
          <CardTitle class="text-2xl">后台管理系统</CardTitle>
          <CardDescription>请输入您的账号和密码</CardDescription>
        </CardHeader>
        <CardContent>
          <form @submit.prevent="handleLogin" class="space-y-4">
            <div class="space-y-2">
              <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="username">
                用户名
              </label>
              <Input
                id="username"
                v-model="form.username"
                placeholder="请输入用户名"
                required
              />
            </div>
            <div class="space-y-2">
              <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="password">
                密码
              </label>
              <Input
                id="password"
                type="password"
                v-model="form.password"
                placeholder="请输入密码"
                required
              />
            </div>
            <Button class="w-full">
              登录
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import Button from '@/components/ui/button.vue'
import Input from '@/components/ui/input.vue'
import Card from '@/components/ui/card.vue'
import { CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card.vue'

const router = useRouter()
const userStore = useUserStore()

const form = ref({
  username: '',
  password: ''
})

const handleLogin = async () => {
  try {
    // 这里替换为实际的登录 API 调用
    // const response = await login(form.value)
    // userStore.setToken(response.token)
    // userStore.setUserInfo(response.userInfo)
    
    // 临时模拟登录成功
    userStore.setToken('fake-token')
    userStore.setUserInfo({
      username: form.value.username,
      avatar: ''
    })
    
    router.push('/')
  } catch (error) {
    console.error('Login failed:', error)
  }
}
</script>
