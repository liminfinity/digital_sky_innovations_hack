<script setup lang="ts">
import MainButton from "../components/MainButton.vue";
</script>

<template>
  <div class="wrapper-header">
    <div>
    <h1 class="title">ПИД-регулятор</h1>
      <div class="second-btns">
      <MainButton :is-primary="false">Undo</MainButton>
      <MainButton :is-primary="false">Reset to original</MainButton>
      </div>
    </div>
    <input type="text" class="search" placeholder="Поиск по ПИДу / параметру...">
    <div class="btns">
      <MainButton :is-primary="true" icon-path="src/assets/save.svg">Сохранить</MainButton>
      <MainButton :is-primary="true" icon-path="src/assets/history.svg">История изменений</MainButton>
    </div>
  </div>
</template>

<style scoped lang="scss">
.wrapper-header {
  padding: 60px 0;
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
  color: #3510B8;
  font-size: 32px;
  font-weight: 800;
}

.search {
  width: 30%;
  background-color: #fff;
  border: 1px solid #4414EC;
  padding: 8px 12px;
  border-radius: 4px;
}

.second-btns {
  display: flex;
  flex-direction: row;
  gap: 12px;
}
</style>
