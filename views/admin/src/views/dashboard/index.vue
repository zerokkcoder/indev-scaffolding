<template>
  <div>
    <Card>
      <CardHeader>
        <CardTitle>欢迎回来</CardTitle>
        <CardDescription>
          {{ userStore.userInfo.username }}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <p>这里是仪表盘页面，您可以在这里查看系统概览。</p>
      </CardContent>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import Card from '@/components/ui/card.vue'
import { CardHeader, CardTitle, CardDescription, CardContent } from '@/components/ui/card.vue'

const userStore = useUserStore()
</script>
