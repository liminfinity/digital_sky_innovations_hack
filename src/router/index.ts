import { createRouter, createWebHistory } from 'vue-router'
import AuthPage from "../pages/AuthPage.vue";
import MainPage from "../pages/MainPage.vue";

const routes = [
    { path: '/', component: AuthPage },
    { path: '/main', component: MainPage },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
