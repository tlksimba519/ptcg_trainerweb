<template>
  <div class="my-decks">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">我的牌組</h1>
      <router-link to="/new-deck" class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600">
        ＋新增牌組
      </router-link>
    </div>

    <div v-if="loading">載入中...</div>

<div v-else-if="decks.length === 0">
  尚未建立任何牌組。
</div>

<ul v-else>
  <li
    v-for="deck in decks"
    :key="deck.id"
    class="mb-2 border p-2 rounded shadow-sm"
  >
    <strong>{{ deck.title }}</strong>
    - {{ Array.isArray(deck.cards) ? deck.cards.length : 0 }} 張卡

    <div class="desc">{{ deck.description }}</div>
  </li>
</ul>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const decks = ref([])
const loading = ref(true)

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) {
    console.warn('⚠️ 無 Token，跳過資料載入')
    loading.value = false
    return
  }

  try {
    const res = await axios.get(`${import.meta.env.VITE_API_BASE_URL}/decks`, {
      headers: { Authorization: token },
    })

    const raw = res.data?.data
    decks.value = Array.isArray(raw) ? raw : []

  } catch (err) {
    console.error('❌ 無法取得牌組資料', err)
    decks.value = []
  } finally {
    loading.value = false
  }
})

</script>

<style scoped>
.my-decks {
  padding: 2rem;
  text-align: left;
}
.desc {
  font-size: 0.9rem;
  color: gray;
}
</style>
