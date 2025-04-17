<template>
  <div class="sound-card">
    <h2 class="sound-card-title">{{ pid.name }}</h2>
    <div
      v-for="(slider, index) in sliders"
      :key="slider.key"
      class="slider-container">
      <label class="slider-label">{{ slider.label }}</label>
      <div class="slider-control">
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
            type="range"
            step="0.0001"
            @change="updateSliderFromInput(index)" />
          <div
            :style="{
              width: `${((slider.value - slider.min) / (slider.max - slider.min)) * 100}%`,
            }"
            class="slider-progress"></div>
          <div class="slider-progress-back"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'

interface IPid {
  name: string
  Kp: number
  Ki: number
  Kd: number
  integral_min: number
  integral_max: number
  inp_rise_deriative: number
  inp_fall_deriative: number
  min: number
  max: number
  preset_allowed_at_low: number
  preset_allowed_at_high: number
}

interface Slider {
  key: string
  label: string
  min: number
  max: number
  value: number
  pidKey: keyof IPid
}

const props = defineProps<{
  pid: IPid
}>()

const emit = defineEmits(['update-pid'])

const sliders = ref<Slider[]>([])

const initializeSliders = () => {
  sliders.value = [
    {
      key: 'kp',
      label: 'Параметр P',
      min: props.pid.preset_allowed_at_low,
      max: props.pid.preset_allowed_at_high,
      value: props.pid.Kp,
      pidKey: 'Kp',
    },
    {
      key: 'ki',
      label: 'Параметр I',
      min: props.pid.integral_min,
      max: props.pid.integral_max,
      value: props.pid.Ki,
      pidKey: 'Ki',
    },
    {
      key: 'kd',
      label: 'Параметр D',
      min: props.pid.inp_fall_deriative,
      max: props.pid.inp_rise_deriative,
      value: props.pid.Kd,
      pidKey: 'Kd',
    },
  ]
}

const updateSliderFromInput = (index: number) => {
  const slider = sliders.value[index]

  if (slider.value < slider.min || slider.value > slider.max) {
    alert(
      'Значение должно быть в пределах от ' + slider.min + ' до ' + slider.max
    )
    slider.value = Math.max(slider.min, Math.min(slider.max, slider.value))
  }
  const updatedPid = { ...props.pid }
  // @ts-ignore
  updatedPid[slider.pidKey] = slider.value
  console.log(updatedPid)
  emit('update-pid', updatedPid)
}

watch(
  () => props.pid,
  () => {
    initializeSliders()
  },
  { deep: true }
)

onMounted(() => {
  initializeSliders()
})
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
  padding: 0 15px;
  width: 48px;
  height: 36px;
  background-color: #fff;
  border: 1px solid #4414ec;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 500;
  color: #000;
  text-align: center;
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
