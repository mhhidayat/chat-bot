<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, inject } from 'vue'

const emit = defineEmits(['selectedModel'])
const isDarkMode = inject('theme')

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
  <div
    class="relative inline-block m-2"
    ref="dropdownRef"
    :class="isDarkMode ? '' : ''"
  >
    <button
      type="button"
      @click="showDropdown = !showDropdown"
      :class="[
        'px-2 py-1 rounded-lg focus:outline-none text-xs border transition-colors duration-200',
        isDarkMode
          ? 'text-gray-400 bg-gray-800 border-gray-700 hover:bg-gray-700 hover:text-white'
          : 'text-gray-700 bg-white border-gray-300 hover:bg-gray-100 hover:text-black'
      ]"
    >
      {{ selectedModel.name }}
    </button>

    <div
      v-if="showDropdown"
      :class="[
        'absolute z-10 bottom-full mb-2 w-80 rounded-lg shadow-lg overflow-hidden border',
        isDarkMode
          ? 'bg-gray-900 border-gray-700'
          : 'bg-white border-gray-300'
      ]"
    >
      <ul>
        <li
          v-for="model in models"
          :key="model.value"
          @click="selectModel(model.name, model.value)"
          :class="[
            'px-4 py-2 hover:cursor-pointer text-sm flex justify-between',
            isDarkMode
              ? 'hover:bg-gray-800'
              : 'hover:bg-gray-100',
            model.value === selectedModel.value
              ? (isDarkMode
                  ? 'bg-gray-700 text-white font-semibold'
                  : 'bg-gray-200 text-black font-semibold')
              : (model.is_paid
                  ? (isDarkMode ? 'text-yellow-400' : 'text-yellow-600')
                  : (isDarkMode ? 'text-gray-300' : 'text-gray-700'))
          ]"
        >
          <span>
            {{ model.name }}
            <span v-if="model.is_paid" title="Dalam Pengembangan">⚙️</span>
          </span>
          <span>{{ model.is_paid ? '💵' : '🚀' }}</span>
        </li>
        <span>
          <li
            :class="[
              'px-4 py-2 text-xs',
              isDarkMode ? 'text-gray-500' : 'text-gray-400'
            ]"
          >
            <span
              :class="[
                'text-[0.70rem]',
                isDarkMode ? 'text-red-500/50' : 'text-red-500/80'
              ]"
            >⚠️ Akan ada model-model baru yang lebih canggih, jadi tetap pantau pembaruan kami! 🍗</span>
          </li>
        </span>
      </ul>
    </div>
  </div>
</template>
