<template>
  <nav class="navbar">
    <router-link to="/" class="nav-link">首頁</router-link>
    <router-link to="/my-decks" class="nav-link">我的牌組</router-link>
    <router-link to="/cards" class="nav-link">卡牌一覽</router-link>
    <span class="spacer"></span>
    
    <span v-if="token">嗨，{{ name }}！</span>
    <button v-if="token" @click="logout">登出</button>
    
    <template v-else>
      <router-link to="/login" class="nav-link">登入</router-link>
      <router-link to="/register" class="nav-link">註冊</router-link>
    </template>
  </nav>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'

const router = useRouter()
const token = ref(null)
const name = ref('')

onMounted(() => {
  token.value = localStorage.getItem('token')
  name.value = localStorage.getItem('name') || '使用者'
})

function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('name')
  router.push('/login')
}
</script>

<style scoped>
.navbar {
  display: flex;
  padding: 1rem;
  background-color: #333;
  color: white;
  align-items: center;
}
.nav-link {
  color: white;
  margin-right: 1rem;
  text-decoration: none;
}
.spacer {
  flex-grow: 1;
}
</style>