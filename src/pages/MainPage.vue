<script lang="ts" setup>
import MainHeader from '../widgets/MainHeader.vue'
import PIDCard from '../widgets/PIDCard.vue'
import { usePidsStore } from '../store/pidsStore.ts'
import { onMounted } from 'vue'
import { getPids } from '../api.ts'

const store = usePidsStore()

function updatePids(v: any) {
  console.log(v)
  store.data = v.data
  console.log(store.data)
}

onMounted(async () => {
  const res = await getPids()
  //@ts-ignore
  if (res) store.data = res.data
  console.log(store.data)
})
</script>

<template>
  <div class="wrapper-main-page">
    <div class="content">
      <MainHeader />
      <div class="cards-grid">
        <PIDCard
          :pid="pid"
          v-for="pid in store.data"
          :key="pid.name"
          @update-pid="updatePids" />
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.wrapper-main-page {
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content {
  width: 80%;
  height: 100%;
}

.cards-grid {
  display: grid;
  grid-template-columns: repeat(2, 0.47fr);
  grid-template-rows: repeat(5, 1fr);
  grid-column-gap: 100px;
  grid-row-gap: 20px;
  margin-top: 150px;
}
</style>
