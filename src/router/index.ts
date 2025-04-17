import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from '../pages/AuthPage.vue'
import MainPage from '../pages/MainPage.vue'
import HistoryPage from '../pages/HistoryPage.vue'

const routes = [
  { path: '/', component: AuthPage },
  { path: '/main', component: MainPage },
  { path: '/history', component: HistoryPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
