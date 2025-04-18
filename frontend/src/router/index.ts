import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  { path: '/', component: () => import('../pages/AuthPage.vue')},
  { path: '/main', component: () => import('../pages/MainPage.vue')},
  { path: '/history', component: () => import('../pages/HistoryPage.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
