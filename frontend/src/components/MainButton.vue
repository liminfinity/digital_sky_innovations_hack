<script lang="ts" setup>
defineProps({
  isPrimary: Boolean,
  iconPath: String,
})
</script>

<template>
  <button :class="isPrimary ? 'primary' : 'secondary'">
    <img v-if="iconPath" :src="iconPath" alt="" />
    <slot />
  </button>
</template>

<style lang="scss" scoped>
button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around;
}

.primary {
  background-color: #4414ec;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 8px 18px;

  &:hover {
    filter: brightness(150%);
  }
}

.secondary {
  background-color: inherit;
  color: #4414ec;
  border: 1px solid #6842ec;
  border-radius: 4px;
  padding: 8px 18px;
  font-weight: 500;
  font-size: 16px;

  &:hover {
    background-color: rgba(68, 20, 236, 0.2);
  }
}
</style>
