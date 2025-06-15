<template>
  <div class="home-container">
    <h1>歡迎來到 PTCG Trainer Web</h1>
    <router-link to="/my-decks" class="deck-link">
      <img src="../assets/deck-icon.svg" alt="我的牌組" class="icon" />
      <span>我的牌組</span>
    </router-link>

    <section class="tournament-section bg-white shadow p-4 rounded border">
  <h2 class="text-xl font-semibold mb-2">🎉 賽事專區</h2>
  <div v-if="loading">載入中...</div>
  <div v-else-if="tournaments.length > 0">
    <div v-for="t in tournaments" :key="t.name" class="mb-4">
      <p><strong>賽事名稱：</strong>{{ t.name }}</p>
      <p><strong>地點：</strong>{{ t.location }}</p>
      <p><strong>級別：</strong>{{ t.level }}</p>
      <p><strong>注意事項：</strong>{{ t.notes }}</p>
      <p><strong>日期：</strong>{{ t.date }}</p>
      <hr/>
    </div>
  </div>
  <div v-else>目前沒有賽事資訊。</div>
</section>

  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const tournaments = ref([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await axios.get('http://localhost:8080/tournament')
    tournaments.value = res.data.data || []
  } catch (err) {
    console.error('❌ 載入賽事資訊失敗', err)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.home-container {
  text-align: center;
  margin-top: 5rem;
}
.deck-link {
  display: inline-flex;
  align-items: center;
  text-decoration: none;
  color: inherit;
  margin-top: 2rem;
}
.icon {
  width: 48px;
  height: 48px;
  margin-right: 8px;
}
.tournament-section {
  margin-top: 2rem;
  text-align: left;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}
</style>
