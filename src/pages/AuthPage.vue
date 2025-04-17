<script setup lang="ts">
import MainButton from "../components/MainButton.vue";
import router from "../router/index.js";
</script>

<template>
  <div class="wrapper-auth">
    <div class="form">
      <h1 class="title">Авторизация</h1>
      <div class="form-values">
        <div class="inputs">
          <label for="login">Логин</label>
          <input id="login">
          <label for="password">Пароль</label>
          <input id="password">
        </div>
        <MainButton @click="router.push('/main')" class="btn" :is-primary="true">Войти</MainButton>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.wrapper-auth {
  width: 100vw;
  height: 100vh;
  background-color: #F9F7FC;
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
  color: #280C8A;
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

#login, #password {
  font-weight: 400;
  font-size: 16px;
}

input {
  border: 1px solid #896DEC;
  border-radius: 8px;
  height: 30px;
}
</style>
