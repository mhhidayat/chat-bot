<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'

const isDarkMode = ref(false)
const emit = defineEmits(['theme'])

onMounted(() => {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme !== null) {
    isDarkMode.value = savedTheme === 'true'
  } else {
    isDarkMode.value = true // Set default to dark mode
  }
  document.body.classList.toggle('bg-gray-900', isDarkMode.value)
  document.body.classList.toggle('bg-gray-100', !isDarkMode.value)
  document.body.classList.toggle('text-gray-100', isDarkMode.value)
  document.body.classList.toggle('text-gray-900', !isDarkMode.value)
  emit('theme', isDarkMode.value)
})

watch(isDarkMode, (val) => {
  localStorage.setItem('theme', val.toString())
  document.body.classList.toggle('bg-gray-900', val)
  document.body.classList.toggle('bg-gray-100', !val)
  document.body.classList.toggle('text-gray-100', val)
  document.body.classList.toggle('text-gray-900', !val)
  emit('theme', isDarkMode.value)
})

</script>

<template>
    <button
      @click="isDarkMode = !isDarkMode"
      class="fixed top-4 right-4 z-10 px-3 py-1 rounded-full border shadow text-xs font-medium transition-colors"
      :class="isDarkMode 
        ? 'bg-gray-800 border-gray-700 text-gray-100 hover:bg-gray-700' 
        : 'bg-white border-gray-200 text-gray-900 hover:bg-gray-200'"
      aria-label="Toggle dark mode"
      title="Theme"
    >
      <span>{{ isDarkMode ? '🌙' : '☀️' }}</span>
      <span class="hidden xl:inline ml-1">
        {{ isDarkMode ? 'Dark' : 'Light' }}
      </span>
    </button>
</template>