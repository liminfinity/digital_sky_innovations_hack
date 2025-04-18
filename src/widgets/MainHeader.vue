<script lang="ts" setup>
import MainButton from '../components/MainButton.vue'
import router from '../router/index.js'
import { usePidsStore } from '../store/pidsStore.ts'
import { savePidsApi } from '../api.ts'
import {useStoriesStore} from "../store/storiesStore.ts";

const store = usePidsStore()
const storiesStore = useStoriesStore()

async function savePIDs() {
  const res = await savePidsApi(store.data)
  if (res) storiesStore.stories.push(res)
  console.log(storiesStore.stories)
}
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
    <input class="search" placeholder="Поиск по ПИДу..." type="text" />
    <div class="btns">
      <MainButton
        :is-primary="true"
        icon-path="src/assets/save.svg"
        @click="savePIDs"
        >Сохранить
      </MainButton>
      <MainButton
        :is-primary="true"
        icon-path="src/assets/history.svg"
        @click="router.push('/history')">
        История изменений
      </MainButton>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.wrapper-header {
  padding: 30px 100px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;
  position: fixed;
  background-color: #f9f7fc;
  z-index: 10;
  width: 80%;
  left: 0;
}

.btns {
  width: 30%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

.title {
  color: #3510b8;
  font-size: 32px;
  font-weight: 800;
}

.search {
  width: 30%;
  background-color: #fff;
  border: 1px solid #4414ec;
  padding: 8px 12px;
  border-radius: 4px;
}

.second-btns {
  display: flex;
  flex-direction: row;
  gap: 12px;
}
</style>
