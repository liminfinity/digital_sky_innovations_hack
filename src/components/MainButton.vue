<script setup lang="ts">
defineProps({
  isPrimary: Boolean,
  iconPath: String,
})
</script>

<template>
  <button :class="isPrimary ? 'primary' : 'secondary'">
    <img :src="iconPath" alt="">
    <slot/>
  </button>
</template>

<style scoped lang="scss">
button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;
}
.primary {
  background-color: #4414EC;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 8px 18px;
}

.secondary {
  background-color: inherit;
  color: #4414EC;
  border: 1px solid #6842EC;
  border-radius: 4px;
  padding: 8px 18px;
  font-weight: 500;
  font-size: 16px;
}
</style>
