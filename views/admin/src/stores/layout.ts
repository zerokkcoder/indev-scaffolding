import { defineStore } from 'pinia'
import type { MenuItem } from '@/types'

export const useLayoutStore = defineStore('layout', {
  state: () => ({
    isCollapsed: localStorage.getItem('sidebarCollapsed') === 'true',
    activeParentMenu: localStorage.getItem('activeParentMenu') || '',
    currentPath: ''
  }),
  actions: {
    toggleSidebar() {
      this.isCollapsed = !this.isCollapsed
      localStorage.setItem('sidebarCollapsed', this.isCollapsed.toString())
    },
    setActiveParentMenu(menuName: string | null) {
      console.log('Setting active parent menu:', menuName)
      this.activeParentMenu = menuName || ''
      localStorage.setItem('activeParentMenu', this.activeParentMenu)
    },
    setCurrentPath(path: string) {
      console.log('Setting current path:', path)
      this.currentPath = path
    },
    isCurrentRoute(path: string) {
      if (!path) return false
      const isMatch = this.currentPath === path
      console.log('Route match check:', {
        currentPath: this.currentPath,
        menuPath: path,
        isMatch
      })
      return isMatch
    },
    isActiveParent(item: MenuItem) {
      // 如果是一级菜单且没有子菜单，只在路由完全匹配时才显示选中状态
      if (!item.children) {
        const isActive = this.isCurrentRoute(item.path || '')
        console.log('Menu active check (no children):', {
          item: item.name,
          path: item.path,
          isActive
        })
        return isActive
      }
      
      // 如果是带有子菜单的一级菜单，只在菜单展开时显示选中状态
      const isActive = this.activeParentMenu === item.name
      console.log('Menu active check (with children):', {
        item: item.name,
        activeParent: this.activeParentMenu,
        isActive
      })
      return isActive
    }
  }
})
