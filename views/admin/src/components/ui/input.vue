<template>
  <input
    :value="modelValue"
    @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
    :class="inputVariants({ variant, class: className })"
    v-bind="$attrs"
  />
</template>

<script setup lang="ts">
import { inputVariants } from './input'

interface Props {
  modelValue: string
  variant?: 'default' | 'error'
  className?: string
}

defineProps<Props>()

defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()
</script>
