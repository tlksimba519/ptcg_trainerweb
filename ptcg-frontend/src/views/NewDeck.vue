<template>
  <div class="p-4 max-w-2xl mx-auto">
    <h1 class="text-xl font-bold mb-4">建立新牌組</h1>
    <form @submit.prevent="submitDeck">
      <input v-model="title" placeholder="牌組名稱" class="mb-2 p-2 border w-full" />
      <textarea v-model="description" placeholder="簡介" class="mb-2 p-2 border w-full" />

      <div class="mb-4">
        <h2 class="font-semibold mb-2">加入卡片</h2>
        <div class="flex items-center mb-2" v-for="(item, index) in cards" :key="index">
          <select v-model="item.card_id" class="border p-2 mr-2 flex-1">
            <option disabled value="">選擇卡片</option>
            <option v-for="card in allCards" :key="card.card_id" :value="card.card_id">
              {{ card.card_id }} - {{ card.name }}
            </option>
          </select>
          <input v-model.number="item.quantity" type="number" min="1" class="border p-2 w-20 mr-2" />
          <button type="button" @click="removeCard(index)" class="text-red-500">移除</button>
        </div>
        <button type="button" @click="addCard" class="bg-gray-200 px-2 py-1">＋ 增加卡片</button>
      </div>

      <button type="submit" class="bg-blue-500 text-white px-4 py-2">建立牌組</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const title = ref('')
const description = ref('')
const cards = ref([{ card_id: '', quantity: 1 }])
const allCards = ref([])

const addCard = () => {
  cards.value.push({ card_id: '', quantity: 1 })
}
const removeCard = (index) => {
  cards.value.splice(index, 1)
}

const submitDeck = async () => {
  try {
    const token = localStorage.getItem('token')
    await axios.post('http://localhost:8080/decks', {
      title: title.value,
      description: description.value,
      cards: cards.value,
    }, {
      headers: { Authorization: token },
    })
    alert('✅ 建立成功！')
  } catch (err) {
    console.error('❌ 建立失敗', err)
    alert('建立失敗')
  }
}

// 🧠 初始化時載入所有卡片清單
onMounted(async () => {
  try {
    const res = await axios.get('http://localhost:8080/cards')
    allCards.value = res.data.data
  } catch (err) {
    console.error('❌ 無法取得卡片資料', err)
  }
})
</script>
