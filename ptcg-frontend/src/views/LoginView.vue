<!-- src/views/LoginView.vue -->
<template>
  <div class="p-4 max-w-md mx-auto">
    <h1 class="text-xl font-bold mb-4">登入</h1>
    <form @submit.prevent="login">
      <input v-model="account" placeholder="帳號" class="mb-2 p-2 border w-full" />
      <input v-model="password" type="password" placeholder="密碼" class="mb-2 p-2 border w-full" />
      <button type="submit" class="bg-blue-500 text-white px-4 py-2">登入</button>
    </form>
    <p v-if="error" class="text-red-500 mt-2">{{ error }}</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import api, { setAuthToken } from '../services/api'

const account = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const login = async () => {
  try {
    const res = await api.post('/login', {
      account: account.value,
      password: password.value,
    })
    const token = res.data.data.token
    console.log('✅ 登入成功，收到 token：', token)

    if (!token) throw new Error('伺服器未回傳 token')

    localStorage.setItem('token', 'Bearer ' + token)
    setAuthToken(token)

    error.value = ''
    router.push('/')
  } catch (err) {
    console.error('❌ 登入失敗', err)
    error.value = '登入失敗，請檢查帳號或密碼'
  }
}
</script>
