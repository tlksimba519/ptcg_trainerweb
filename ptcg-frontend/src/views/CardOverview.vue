<template>
  <div class="card-overview p-6">
    <h1 class="text-4xl font-bold mb-10 text-center text-gray-800">📚 卡片總覽</h1>

    <div v-if="loading" class="text-center text-lg text-gray-600">載入中...</div>

    <div v-else-if="cards.length === 0" class="text-center text-gray-500 text-lg">
      沒有任何卡片資料
    </div>

    <div class="mx-auto max-w-6xl px-4">
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 gap-6">
        <div
          v-for="card in cards"
          :key="card.id"
          class="bg-white border border-gray-200 rounded-xl shadow-md hover:shadow-xl transition duration-300 overflow-hidden cursor-pointer"
        >
          <img
            :src="card.image_url || defaultImage"
            alt="card image"
            class="w-full h-52 object-cover"
          />
          <div class="p-4">
            <h2 class="text-lg font-bold text-gray-800 mb-1">{{ card.name }}</h2>
            <p class="text-sm text-gray-500 mb-2">{{ card.type }}</p>
            <p class="text-sm text-gray-700 line-clamp-3">{{ card.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const cards = ref([])
const loading = ref(true)

// 用漂亮一點的預設圖
const defaultImage = 'https://images.unsplash.com/photo-1526406915890-5cbac1b21984?fit=crop&w=600&q=80'

onMounted(async () => {
  try {
    const res = await axios.get(`${import.meta.env.VITE_API_BASE_URL}/cards`)
    cards.value = res.data.data || res.data || []
  } catch (err) {
    console.error('❌ 載入卡片資料失敗', err)
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.card-overview {
  text-align: center;
}

/* 限制描述最多顯示3行，多餘的以省略號顯示 */
.line-clamp-3 {
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}
</style>
