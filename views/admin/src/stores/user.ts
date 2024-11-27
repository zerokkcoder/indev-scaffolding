import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref({
    username: '',
    avatar: ''
  })

  function setToken(value: string) {
    token.value = value
    localStorage.setItem('token', value)
  }

  function setUserInfo(info: typeof userInfo.value) {
    userInfo.value = info
  }

  function logout() {
    token.value = ''
    userInfo.value = {
      username: '',
      avatar: ''
    }
    localStorage.removeItem('token')
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    logout
  }
})
