import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import LoginView from '../views/LoginView.vue'
import MyDecks from '../views/MyDecks.vue'
import NewDeck from '../views/NewDeck.vue'
import CardOverview from '../views/CardOverview.vue'

const routes = [
  { path: '/', name: 'Home', component: Home },
  { path: '/login', name: 'Login', component: LoginView },
  {
    path: '/my-decks',
    name: 'MyDecks',
    component: MyDecks,
    meta: { requiresAuth: true },
  },
  {
     path: '/new-deck',
  name: 'NewDeck',
  component: NewDeck,
  meta: { requiresAuth: true }
  },
  {
  path: "/cards",
  name: "CardOverview",
  component: CardOverview,
},
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router