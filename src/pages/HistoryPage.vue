<template>
  <div class="history-page">
    <header class="header">
      <h1 class="title">История изменений</h1>
      <input class="search" placeholder="Поиск по ПИДу..." type="text" />
      <MainButton :is-primary="true" @click="router.back()">
        <span class="icon">↩</span>
        Вернуться к редактору ПИД
      </MainButton>
    </header>

    <div class="history-container">
      <div class="history-list">
        <div
          v-for="story in storiesStore.stories"
          :key="story.id"
          class="history-item">
          <div class="item-info">
            <div class="timestamp">{{ convertISODateToTimeAndDate(story.created_at)[0] }} {{ convertISODateToTimeAndDate(story.created_at)[1] }}</div>
            <div class="user">{{ story.username }}</div>
          </div>
          <button class="apply-button" @click="selectStory(story.id)">Применить изменение</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MainButton from '../components/MainButton.vue'
import router from '../router/index.js'
import {useStoriesStore} from "../store/storiesStore.ts";
import {getPidById} from "../api.ts";
import {usePidsStore} from "../store/pidsStore.ts";

const store = usePidsStore()
const storiesStore = useStoriesStore()

async function selectStory(storyId: number) {
  const res = await getPidById(storyId)
  console.log(res.data.pids)
  if (res)
    store.data = res.data.pids
}

function convertISODateToTimeAndDate(isoDateString: string): [string, string] {
  const date = new Date(isoDateString);
  const hours = date.getHours();
  const minutes = date.getMinutes();
  const timeString = `${hours}:${minutes.toString().padStart(2, '0')}`;
  const day = date.getDate().toString().padStart(2, '0');
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const year = date.getFullYear();
  const dateString = `${day}.${month}.${year}`;
  return [timeString, dateString];
}

</script>

<style lang="scss" scoped>
.history-page {
  background-color: #f5f5f8;
  width: 80vw;
  padding: 20px;
  font-family: Arial, sans-serif;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
  position: fixed;
  top: 30px;
  width: 80%;
}

.title {
  font-weight: 800;
  font-size: 32px;
  line-height: 24px;
  letter-spacing: 0;
  vertical-align: middle;
  color: #3510b8;
}

.search {
  width: 30%;
  background-color: #fff;
  border: 1px solid #4414ec;
  padding: 8px 12px;
  border-radius: 4px;
}

.icon {
  font-size: 16px;
}

.history-container {
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  overflow: hidden;
  flex-grow: 1;
}

.history-list {
  max-height: calc(
    100vh - 150px
  ); /* Adjust based on header height and padding */
  overflow-y: auto;
  padding: 0;
}

.history-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #e0e0ff;
}

.history-item:last-child {
  border-bottom: none;
}

.item-info {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.timestamp {
  font-size: 18px;
  font-weight: bold;
}

.user {
  color: #777;
  font-size: 14px;
}

.apply-button {
  background-color: #5000ff;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 10px 15px;
  font-size: 14px;
  cursor: pointer;
}

/* Custom scrollbar styling */
.history-list::-webkit-scrollbar {
  width: 8px;
}

.history-list::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.history-list::-webkit-scrollbar-thumb {
  background: #ddd;
  border-radius: 4px;
}

.history-list::-webkit-scrollbar-thumb:hover {
  background: #ccc;
}
</style>
