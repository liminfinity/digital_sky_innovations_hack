<script setup lang="ts">
import MainHeader from '../widgets/MainHeader.vue'
import PIDCard from '../widgets/PIDCard.vue'
</script>

<template>
  <div class="wrapper-main-page">
    <div class="content">
      <MainHeader />
      <div class="cards-grid">
        <PIDCard v-for="i in 8" :key="i" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
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
}
</style>
