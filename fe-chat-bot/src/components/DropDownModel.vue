<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, inject, computed } from 'vue'

// Props & Event
const emit = defineEmits(['selectedModel'])
const isDarkMode = inject('theme')

// Dropdown & Search Logic
const showDropdown = ref(false)
const search = ref('')
const showAll = ref(false)

// Model Data
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

// Default selected model
const selectedModel = ref({
  name: models.value[2].name,
  value: models.value[2].value
})

// Computed model list
const filteredModels = computed(() => {
  return models.value.filter(model =>
    model.name.toLowerCase().includes(search.value.toLowerCase())
  )
})

const displayedModels = computed(() => {
  return showAll.value ? filteredModels.value : filteredModels.value.slice(0, 5)
})

// Handle Model Selection
function selectModel(name: string, value: string) {
  selectedModel.value = { name, value }
  showDropdown.value = false
  search.value = ''
  showAll.value = false
  emit('selectedModel', value)
}

// Handle Dropdown Toggle
function toggleDropdown() {
  showDropdown.value = !showDropdown.value
  if (!showDropdown.value) {
    search.value = ''
    showAll.value = false
  }
}

// Click outside close logic
const dropdownRef = ref<HTMLElement | null>(null)

function handleClickOutside(event: MouseEvent) {
  if (
    dropdownRef.value &&
    !dropdownRef.value.contains(event.target as Node)
  ) {
    showDropdown.value = false
    search.value = ''
    showAll.value = false
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
  <div class="relative inline-block m-2 max-w-80 w-full" ref="dropdownRef">
    <!-- Button -->
    <button
      @click="toggleDropdown"
      class="flex justify-between items-center max-w-2xl px-2 py-1 rounded-lg border text-xs transition-colors"
      :class="isDarkMode
        ? 'bg-gray-800 text-white border-gray-700 hover:bg-gray-700'
        : 'bg-white text-gray-800 border-gray-300 hover:bg-gray-100'"
    >
      {{ selectedModel.name }}
      <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" stroke-width="2"
        viewBox="0 0 24 24">
        <path d="M19 9l-7 7-7-7"></path>
      </svg>
    </button>

    <!-- Dropdown -->
    <transition name="fade">
      <div
        v-if="showDropdown"
        class="absolute z-50 mt-2 w-full rounded-lg shadow-xl border overflow-hidden bottom-8"
        :class="isDarkMode
          ? 'bg-gray-900 border-gray-700 text-white'
          : 'bg-white border-gray-300 text-gray-800'"
      >
        <!-- Search -->
        <div class="p-2 border-b" :class="isDarkMode ? 'border-gray-700' : 'border-gray-200'">
          <input
            v-model="search"
            type="text"
            placeholder="Search models..."
            class="w-full px-3 py-2 text-sm rounded-md border focus:outline-none"
            :class="isDarkMode
              ? 'bg-gray-800 border-gray-600 text-white placeholder-gray-400'
              : 'bg-gray-50 border-gray-300 text-gray-800 placeholder-gray-500'"
            @click.stop
          />
        </div>

        <!-- Model List -->
        <ul class="max-h-64 overflow-y-auto text-sm"
          :class="isDarkMode ? 'divide-gray-700' : 'divide-gray-200'">
          <li
            v-for="model in displayedModels"
            :key="model.value"
            @click="selectModel(model.name, model.value)"
            class="flex justify-between px-4 py-2 cursor-pointer hover:font-semibold transition"
            :class="[
              model.value === selectedModel.value
                ? (isDarkMode
                  ? 'bg-gray-700 text-white'
                  : 'bg-gray-200 text-black')
                : (model.is_paid
                  ? (isDarkMode ? 'text-yellow-400' : 'text-yellow-600')
                  : (isDarkMode ? 'text-gray-300' : 'text-gray-700')),
              isDarkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-100'
            ]"
          >
            <span>
              {{ model.name }}
              <span v-if="model.is_paid" class="ml-1" title="Model berbayar">⚙️</span>
            </span>
            <span>{{ model.is_paid ? '💵' : '🚀' }}</span>
          </li>
        </ul>

        <!-- "Lihat Semua" -->
        <div
          v-if="!showAll && filteredModels.length > 5"
          @click.stop="showAll = true"
          class="text-center px-4 py-2 text-xs font-medium cursor-pointer"
          :class="isDarkMode ? 'text-blue-400 hover:bg-gray-800' : 'text-blue-600 hover:bg-gray-100'"
        >
          Show all ({{ filteredModels.length }})
        </div>

      </div>
    </transition>
  </div>
</template>

<style scoped>
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.15s ease;
  }
  .fade-enter-from,
  .fade-leave-to {
    opacity: 0;
  }
</style>
