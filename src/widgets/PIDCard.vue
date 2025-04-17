<template>
  <div class="sound-card">
    <h2 class="sound-card-title">Название пида</h2>

    <div class="slider-container" v-for="(slider, index) in sliders" :key="index">
      <label class="slider-label">{{ slider.label }}</label>
      <div class="slider-control">
        <!-- Changed to input field -->
        <input
            type="text"
            class="left-value"
            v-model.number="slider.leftValue"
            @change="updateSliderFromInput(index)"
        >
        <div class="slider-wrapper">
          <input
              type="range"
              :min="slider.min"
              :max="slider.max"
              v-model.number="slider.value"
              class="slider"
              @input="updateInputFromSlider(index)"
          >
          <!-- The blue progress bar -->
          <div class="slider-progress" :style="{ width: `${(slider.value - slider.min) / (slider.max - slider.min) * 100}%` }"></div>
          <div class="slider-progress-back"></div>
          <!-- The gray background track is handled in CSS -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

// Define reactive sliders array with initial values
const sliders = ref([
  {
    label: 'Параметр P',
    min: 0,
    max: 100,
    value: 80,
    leftValue: 64
  },
  {
    label: 'Параметр P',
    min: 0,
    max: 100,
    value: 80,
    leftValue: 64
  },
  {
    label: 'Параметр P',
    min: 0,
    max: 100,
    value: 80,
    leftValue: 64
  }
]);

// Update the input value when the slider changes
const updateInputFromSlider = (index) => {
  // For example: if slider is at 80, left value shows 64 (80% of 80)
  sliders.value[index].leftValue = Math.round(sliders.value[index].value * 0.8);
};

// Update the slider value when the input changes
const updateSliderFromInput = (index) => {
  let inputValue = sliders.value[index].leftValue;

  // Validate input
  if (isNaN(inputValue)) {
    inputValue = 0;
  }

  // Convert from left value to slider value (reverse of the formula above)
  // If left value is 64, slider should be 80 (64 / 0.8)
  let sliderValue = Math.round(inputValue / 0.8);

  // Clamp to min/max
  sliderValue = Math.max(sliders.value[index].min,
      Math.min(sliders.value[index].max, sliderValue));

  // Update slider value
  sliders.value[index].value = sliderValue;

  // Update left value to ensure it's valid
  sliders.value[index].leftValue = Math.round(sliderValue * 0.8);
};
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
  color: #000;
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
  background-color: #2563eb;
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
  background-color: #f0f0f0;
  border-radius: 3px;
}
</style>
