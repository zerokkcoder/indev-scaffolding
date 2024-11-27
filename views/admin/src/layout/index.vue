<template>
  <div class="min-h-screen bg-background">
    <!-- Header -->
    <header class="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div class="container flex h-14 items-center">
        <div class="mr-4 flex">
          <a class="mr-6 flex items-center space-x-2" href="/">
            <span class="font-bold">Admin</span>
          </a>
        </div>
        <div class="flex flex-1 items-center justify-between space-x-2 md:justify-end">
          <div class="flex items-center">
            <Button
              variant="ghost"
              @click="userStore.logout(); router.push('/login')"
            >
              <LogOut class="mr-2 h-4 w-4" />
              <span>退出登录</span>
            </Button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <div class="container py-6">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import Button from '@/components/ui/button.vue'
import { LogOut } from 'lucide-vue-next'

const router = useRouter()
const userStore = useUserStore()
</script>
