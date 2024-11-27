<template>
  <div :class="cardVariants({ variant, class: className })">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { cardVariants } from './card'

interface Props {
  variant?: 'default' | 'destructive'
  className?: string
}

withDefaults(defineProps<Props>(), {
  variant: 'default',
  className: ''
})
</script>

<script lang="ts">
import { defineComponent, h } from 'vue'
import {
  cardHeaderVariants,
  cardTitleVariants,
  cardDescriptionVariants,
  cardContentVariants,
  cardFooterVariants
} from './card'

export const CardHeader = defineComponent({
  name: 'CardHeader',
  setup(_, { slots }) {
    return () => h('div', { class: cardHeaderVariants() }, slots.default?.())
  }
})

export const CardTitle = defineComponent({
  name: 'CardTitle',
  setup(_, { slots }) {
    return () => h('h3', { class: cardTitleVariants() }, slots.default?.())
  }
})

export const CardDescription = defineComponent({
  name: 'CardDescription',
  setup(_, { slots }) {
    return () => h('p', { class: cardDescriptionVariants() }, slots.default?.())
  }
})

export const CardContent = defineComponent({
  name: 'CardContent',
  setup(_, { slots }) {
    return () => h('div', { class: cardContentVariants() }, slots.default?.())
  }
})

export const CardFooter = defineComponent({
  name: 'CardFooter',
  setup(_, { slots }) {
    return () => h('div', { class: cardFooterVariants() }, slots.default?.())
  }
})
</script>
