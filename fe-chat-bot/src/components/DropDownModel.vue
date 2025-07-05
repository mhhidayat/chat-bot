<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, defineEmits } from 'vue'

const emit = defineEmits(['selectedModel'])

const showDropdown = ref(false)
const models = ref([
    { name: 'Gemini 2.5 Pro', value: 'gemini-2.5-pro', is_paid: false },
    { name: 'Gemini 2.5 Flash', value: 'gemini-2.5-flash', is_paid: false },
    { name: 'Gemini 2.5 Flash-Lite Preview 06-17', value: 'gemini-2.5-flash-lite-preview-06-17', is_paid: false },
    // { name: 'Gemini 2.5 Flash Preview TTS', value: 'gemini-2.5-flash-preview-tts', is_paid: true }, //Hanya mendukung output AUDIO
    { name: 'Gemini 2.0 Flash', value: 'gemini-2.0-flash', is_paid: false },
    // { name: 'Gemini 2.0 Flash Preview Image Generation', value: 'gemini-2.0-flash-preview-image-generation', is_paid: true }, //Hanya mendukung output IMAGE
    { name: 'Gemini 2.0 Flash-Lite', value: 'gemini-2.0-flash-lite', is_paid: false },
    { name: 'Gemini 1.5 Flash', value: 'gemini-1.5-flash', is_paid: false },
    { name: 'Gemini 1.5 Flash-8B', value: 'gemini-1.5-flash-8b', is_paid: false },
    // { name: 'Gemma 3 & 3n', value: 'Gemma 3 & 3n', is_paid: true }, // Ini bagus tapi masih belum di rilis google
])

const selectedModel = ref({
    name: models.value[2].name,
    value: models.value[2].value
})

function selectModel(name: string, value: string) {
  selectedModel.value = { name, value }
  showDropdown.value = false
  emit('selectedModel', value)
}

// ref ke wrapper
const dropdownRef = ref<HTMLElement | null>(null)

function handleClickOutside(event: MouseEvent) {
  if (
    dropdownRef.value &&
    !dropdownRef.value.contains(event.target as Node)
  ) {
    showDropdown.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <div class="relative inline-block m-2" ref="dropdownRef">
    <button
      type="button"
      @click="showDropdown = !showDropdown"
      class="px-2 py-1 rounded-lg focus:outline-none text-xs text-gray-400 bg-gray-800 border border-gray-700 transition-colors duration-200 hover:bg-gray-700 hover:text-white"
    >
      {{ selectedModel.name }}
    </button>

    <div
      v-if="showDropdown"
      class="absolute z-10 bottom-full mb-2 w-80 bg-gray-900 border border-gray-700 rounded-lg shadow-lg overflow-hidden"
    >
      <ul>
        <li
          v-for="model in models"
          :key="model.value"
          @click="selectModel(model.name, model.value)"
          :class="[
            'px-4 py-2 hover:bg-gray-800 cursor-pointer text-sm flex justify-between',
            model.value === selectedModel.value ? 'bg-gray-700 text-white font-semibold' : '',
            model.is_paid ? 'text-yellow-400' : 'text-gray-300'
          ]"
        >
          <span>{{ model.name }} <span v-if="model.is_paid" title="Dalam Pengembangan">⚙️</span></span>
          <span>{{ model.is_paid ? '💵' : '🚀' }}</span>
        </li>
        <span>
          <li class="px-4 py-2 text-xs text-gray-500">
            <span class="text-[0.70rem] text-red-500/50">⚠️ Akan ada model-model baru yang lebih canggih, jadi tetap pantau pembaruan kami! 🍗</span>
          </li>
        </span>
      </ul>
    </div>
  </div>
</template>
