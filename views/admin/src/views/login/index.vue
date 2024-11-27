<template>
  <div class="flex min-h-screen items-center justify-center bg-gray-100">
    <Card class="w-full max-w-md">
      <form @submit.prevent="handleSubmit" class="space-y-6 p-6">
        <div class="text-center">
          <h1 class="text-2xl font-bold">Admin Login</h1>
          <p class="mt-2 text-sm text-gray-600">Please sign in to continue</p>
        </div>

        <div class="space-y-4">
          <div class="space-y-2">
            <label class="text-sm font-medium" for="username">Username</label>
            <Input
              id="username"
              v-model="formData.username"
              :variant="errors.username ? 'error' : 'default'"
              placeholder="Enter your username"
              type="text"
            />
            <p v-if="errors.username" class="text-sm text-red-500">{{ errors.username }}</p>
          </div>

          <div class="space-y-2">
            <label class="text-sm font-medium" for="password">Password</label>
            <Input
              id="password"
              v-model="formData.password"
              type="password"
              :variant="errors.password ? 'error' : 'default'"
              placeholder="Enter your password"
            />
            <p v-if="errors.password" class="text-sm text-red-500">{{ errors.password }}</p>
          </div>
        </div>

        <Button type="submit" class="w-full" :disabled="isLoading">
          {{ isLoading ? 'Signing in...' : 'Sign in' }}
        </Button>

        <p v-if="loginError" class="text-center text-sm text-red-500">{{ loginError }}</p>
      </form>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../stores/user'
import Card from '../../components/ui/card.vue'
import Button from '../../components/ui/button.vue'
import Input from '../../components/ui/input.vue'

const router = useRouter()
const userStore = useUserStore()

const formData = ref({
  username: '',
  password: ''
})

const errors = ref({
  username: '',
  password: ''
})

const isLoading = ref(false)
const loginError = ref('')

// 当用户输入时清除对应的错误信息
watch(() => formData.value.username, () => {
  errors.value.username = ''
})

watch(() => formData.value.password, () => {
  errors.value.password = ''
})

const validateForm = () => {
  let isValid = true
  errors.value.username = ''
  errors.value.password = ''

  if (!formData.value.username.trim()) {
    errors.value.username = 'Username is required'
    isValid = false
  }

  if (!formData.value.password.trim()) {
    errors.value.password = 'Password is required'
    isValid = false
  }

  return isValid
}

const handleSubmit = async () => {
  loginError.value = ''
  
  if (!validateForm()) {
    return
  }

  try {
    isLoading.value = true
    await userStore.login(formData.value.username.trim(), formData.value.password.trim())
    router.push('/dashboard')
  } catch (error: any) {
    loginError.value = error.message || 'Login failed'
  } finally {
    isLoading.value = false
  }
}
</script>
