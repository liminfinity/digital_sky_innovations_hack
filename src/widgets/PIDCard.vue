<template>
  <div class="sound-card">
    <h2 class="sound-card-title">Название пида</h2>

    <div
      v-for="(slider, index) in sliders"
      :key="index"
      class="slider-container">
      <label class="slider-label">{{ slider.label }}</label>
      <div class="slider-control">
        <!-- Changed to input field -->
        <input
          v-model.number="slider.value"
          class="left-value"
          type="text"
          @change="updateSliderFromInput(index)" />
        <div class="slider-wrapper">
          <input
            v-model.number="slider.value"
            :max="slider.max"
            :min="slider.min"
            class="slider"
            type="range" />
          <!-- The blue progress bar -->
          <div
            :style="{
              width: `${((slider.value - slider.min) / (slider.max - slider.min)) * 100}%`,
            }"
            class="slider-progress"></div>
          <div class="slider-progress-back"></div>
          <!-- The gray background track is handled in CSS -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const sliders = ref([
  {
    label: 'Параметр P',
    min: 0,
    max: 100,
    value: 50,
  },
  {
    label: 'Параметр I',
    min: 0,
    max: 100,
    value: 50,
  },
  {
    label: 'Параметр D',
    min: 0,
    max: 100,
    value: 50,
  },
])

const updateSliderFromInput = (index) => {
  // Validate input
  if (
    sliders.value[index].value < sliders.value[index].min ||
    sliders.value[index].value > sliders.value[index].max
  ) {
    alert('поменяй')
  }

  sliderValue = Math.max(
    sliders.value[index].min,
    Math.min(sliders.value[index].max, sliderValue)
  )

  sliders.value[index].value = sliderValue
}
</script>

<style scoped>
.sound-card {
  width: 100%;
  max-width: 500px;
  background-color: white;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  padding: 24px;
}

.sound-card-title {
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 24px;
  color: #3510b8;
}

.slider-container {
  margin-bottom: 24px;
  position: relative;
}

.slider-label {
  display: block;
  margin-bottom: 12px;
  font-weight: 500;
  color: #000;
  font-size: 16px;
}

.slider-control {
  display: flex;
  align-items: center;
  gap: 12px;
  position: relative;
}

.left-value {
  width: 48px;
  height: 36px;
  background-color: #f0f0f0;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 500;
  color: #333;
  text-align: center;
  border: none;
  outline: none;
  font-size: 16px;
}

.left-value:focus {
  box-shadow: 0 0 0 2px rgba(37, 99, 235, 0.3);
}

.slider-wrapper {
  flex-grow: 1;
  position: relative;
  height: 36px;
}

.slider {
  position: absolute;
  top: 15px;
  width: 100%;
  height: 6px;
  -webkit-appearance: none;
  appearance: none;
  background: transparent;
  outline: none;
  z-index: 2;
  cursor: pointer;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 20px;
  height: 20px;
  background: white;
  border-radius: 50%;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  margin-top: -7px;
  z-index: 3;
}

.slider::-moz-range-thumb {
  width: 20px;
  height: 20px;
  background: white;
  border-radius: 50%;
  cursor: pointer;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  border: none;
  z-index: 3;
}

/* Blue progress bar */
.slider-progress {
  position: absolute;
  top: 15px;
  left: 0;
  height: 6px;
  background-color: #4414ec;
  border-radius: 3px;
  pointer-events: none;
  z-index: 1;
}

/* Gray background track */
.slider-progress-back {
  content: '';
  position: absolute;
  top: 15px;
  left: 0;
  width: 100%;
  height: 6px;
  background-color: #cbc2ec;
  border-radius: 3px;
}
</style>
