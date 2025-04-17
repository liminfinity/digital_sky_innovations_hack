<script setup lang="ts">
import MainButton from "../components/MainButton.vue";
</script>

<template>
  <div class="wrapper-header">
    <h1 class="title">ПИД-регулятор</h1>
    <input type="text" class="search" placeholder="Поиск по ПИДу / параметру...">
    <div class="btns">
      <MainButton :is-primary="true">Сохранить</MainButton>
      <MainButton :is-primary="true">История изменений</MainButton>
    </div>
  </div>
</template>

<style scoped lang="scss">
.wrapper-header {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.btns {
  width: 30%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.title {
  font-size: 32px;
  font-weight: 800;
}

.search {
  width: 30%;
  background-color: #fff;
  border: 1px solid #2060EB;
  padding: 8px 12px;
  border-radius: 4px;
}
</style>
