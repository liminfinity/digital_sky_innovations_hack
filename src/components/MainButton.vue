<script setup lang="ts">
defineProps({
  isPrimary: Boolean
})
</script>

<template>
  <button :class="isPrimary ? 'primary' : 'secondary'">
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
