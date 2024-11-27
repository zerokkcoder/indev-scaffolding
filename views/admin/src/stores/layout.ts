import { defineStore } from 'pinia'

export const useLayoutStore = defineStore('layout', {
  state: () => ({
    isCollapsed: localStorage.getItem('sidebarCollapsed') === 'true'
  }),
  actions: {
    toggleSidebar() {
      this.isCollapsed = !this.isCollapsed
      localStorage.setItem('sidebarCollapsed', this.isCollapsed.toString())
    }
  }
})
