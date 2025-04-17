<script setup lang="ts">

</script>

<template>
  <button class="primary">
    <slot/>
  </button>
</template>

<style scoped lang="scss">
.primary {
  background-color: #2060EB;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 8px 18px;
}
</style>
