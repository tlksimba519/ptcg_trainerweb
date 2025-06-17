import axios from 'axios'

const api = axios.create({
  baseURL: 'https://ptcgtrainerweb-production.up.railway.app/',
})

// 統一掛上 token
export function setAuthToken(token) {
  api.defaults.headers.common['Authorization'] = `Bearer ${token}`
}

// 初始化時，如果 localStorage 有 token 就直接掛上
const token = localStorage.getItem('token')
if (token) {
  setAuthToken(token)
}

export default api
