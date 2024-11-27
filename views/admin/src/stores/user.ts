import { defineStore } from 'pinia'
import { ref } from 'vue'
import { http } from '../utils/http'

interface UserInfo {
  id: number
  username: string
  created_at: string
  updated_at: string
}

interface LoginResponse {
  admin: UserInfo
  token: string
}

export const useUserStore = defineStore('user', () => {
  const userInfo = ref<UserInfo | null>(null)
  const token = ref<string>('')

  // 从 localStorage 恢复状态
  const initState = () => {
    const savedToken = localStorage.getItem('token')
    const savedUserInfo = localStorage.getItem('userInfo')
    if (savedToken) token.value = savedToken
    if (savedUserInfo) userInfo.value = JSON.parse(savedUserInfo)
  }

  // 登录
  const login = async (username: string, password: string) => {
    try {
      const response = await http.post<LoginResponse>('/auth/login', {
        username,
        password
      })
      
      // 保存登录状态
      token.value = response.data.token
      userInfo.value = response.data.admin
      
      // 保存到 localStorage
      localStorage.setItem('token', response.data.token)
      localStorage.setItem('userInfo', JSON.stringify(response.data.admin))
      
      return response
    } catch (error) {
      throw error
    }
  }

  // 登出
  const logout = () => {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  return {
    userInfo,
    token,
    initState,
    login,
    logout
  }
})
