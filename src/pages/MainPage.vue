<script setup lang="ts">
import MainHeader from "../widgets/MainHeader.vue";
</script>

<template>
<div class="wrapper-main-page">
  <div class="content">
    <MainHeader />
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
</style>
