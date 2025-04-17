<script lang="ts" setup>
import MainButton from '../components/MainButton.vue'
import { ref } from 'vue'
import { loginApi } from '../api.ts'
import router from '../router'

const userName = ref('')
const password = ref('')

async function loginUser() {
  const res = await loginApi(userName.value, password.value)
  await router.push('/main')
}
</script>

<template>
  <div class="wrapper-auth">
    <div class="form">
      <h1 class="title">Авторизация</h1>
      <div class="form-values">
        <div class="inputs">
          <label for="login">Логин</label>
          <input id="login" v-model="userName" />
          <label for="password">Пароль</label>
          <input id="password" v-model="password" />
        </div>
        <MainButton :is-primary="true" class="btn" @click="loginUser"
          >Войти
        </MainButton>
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.wrapper-auth {
  width: 100vw;
  height: 100vh;
  background-color: #f9f7fc;
  display: flex;
  align-items: center;
  justify-content: center;
}

.form {
  background-color: #fff;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 32px;
  padding: 5%;
  border-radius: 16px;
}

.title {
  color: #280c8a;
  font-weight: 800;
  font-size: 32px;
}

.form-values {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 24px;
}

.inputs {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.btn {
  width: 100%;
  text-align: center;
}

#login,
#password {
  font-weight: 400;
  font-size: 16px;
}

input {
  border: 1px solid #896dec;
  border-radius: 8px;
  height: 30px;
}
</style>
