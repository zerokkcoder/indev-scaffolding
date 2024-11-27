<template>
  <input
    :class="inputVariants({ variant, class: className })"
    v-bind="$attrs"
  />
</template>

<script setup lang="ts">
import { inputVariants } from './input'

interface Props {
  variant?: 'default' | 'error'
  className?: string
}

withDefaults(defineProps<Props>(), {
  variant: 'default',
  className: ''
})
</script>
