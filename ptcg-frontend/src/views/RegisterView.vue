<template>
  <div class="p-4 max-w-md mx-auto">
    <h1 class="text-xl font-bold mb-4">註冊</h1>
    <form @submit.prevent="register">
      <input v-model="account" placeholder="帳號" class="mb-2 p-2 border w-full" />
      <input v-model="password" type="password" placeholder="密碼" class="mb-2 p-2 border w-full" />
      <input v-model="name" placeholder="暱稱" class="mb-2 p-2 border w-full" />
      <button type="submit" class="bg-green-500 text-white px-4 py-2">註冊</button>
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
const name = ref('')
const error = ref('')
const router = useRouter()

const register = async () => {
  try {
    const res = await api.post('/register', {
      account: account.value,
      password: password.value,
      name: name.value,
    })
    const token = res.data.data.token
    console.log('✅ 註冊成功，收到 token：', token)

    if (!token) throw new Error('伺服器未回傳 token')

    localStorage.setItem('token', 'Bearer ' + token)
    setAuthToken(token)

    error.value = ''
    router.push('/')
  } catch (err) {
    console.error('❌ 註冊失敗', err)
    error.value = '註冊失敗，請確認帳號是否重複或資料格式正確'
  }
}
</script>
