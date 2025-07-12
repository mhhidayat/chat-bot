<script setup lang="ts">
import { ref, computed, inject } from 'vue'

// Props & Event
const emit = defineEmits(['promptSelected'])
const isDarkMode = inject('theme', false)

const selectedCategory = ref('📚 Learn')

const categories = [
  {
    name: '📚 Learn',
    items: [
      { name: 'What does a president do?', icon: '👔' },
      { name: 'Why do the seasons change?', icon: '🌞' },
      { name: 'What is democracy?', icon: '📖' },
      { name: 'Why is education important?', icon: '🎓' }
    ]
  },
  {
    name: '📑 News',
    items: [
      { name: 'What is the latest world news?', icon: '📰' },
      { name: 'How does inflation affect us?', icon: '📉' },
      { name: 'What are global warming updates?', icon: '🌍' },
      { name: 'Who won the recent election?', icon: '🗳️' }
    ]
  },
  {
    name: '👩‍🍳 Cooking',
    items: [
      { name: 'Rendang recipes', icon: '🍲' },
      { name: 'Important spices in cooking', icon: '🌶️' },
      { name: 'What is a healthy diet?', icon: '🥗' },
      { name: 'How to bake a cake?', icon: '🎂' }
    ]
  },
  {
    name: '📱 Electronics',
    items: [
      { name: 'What is the best phone in 2025?', icon: '📱' },
      { name: 'How does a CPU work?', icon: '🧠' },
      { name: 'What is 5G?', icon: '📶' },
      { name: 'How to build a PC?', icon: '🛠️' }
    ]
  }
]

function selectPrompt(prompt: string) {
  emit('promptSelected', prompt)
}

const currentItems = computed(() => {
  return categories.find(c => c.name === selectedCategory.value)?.items || []
})
</script>

<template>
  <div class="flex justify-center items-center w-full h-full">
    <div class="p-4 w-full max-w-xl">
      <h1 class="font-semibold text-3xl mb-4">How can I help you?</h1>

      <!-- Category Buttons -->
      <div class="flex flex-wrap gap-2 sm:gap-3 mb-6">
        <button
          v-for="category in categories"
          :key="category.name"
          @click="selectedCategory = category.name"
          :class="[
            isDarkMode ? 'bg-white/5 border-white/20 ring-white/50' : 'bg-black/5 norder-black/20 ring-black/30',
            'py-1 px-4 rounded-full transition duration-300 backdrop-blur-md border shadow-md hover:shadow-lg font-medium',
            selectedCategory === category.name ? 'ring-1' : ''
          ]"
        >
          {{ category.name }}
        </button>
      </div>

      <!-- Item List -->
      <ul class="flex flex-col gap-1 sm:gap-2">
        <li
          v-for="item in currentItems"
          v-on:click="selectPrompt(item.name)"
          :key="item.name"
          :class="[
            isDarkMode ? 'border-white' : 'border-black',
            'border-b-[0.5px] p-1 pb-2 border-opacity-30'
          ]"
        >
          <button
            :class="[
            isDarkMode ? 'hover:bg-gray-800/50' : 'hover:bg-gray-200',
            'hover:cursor-pointer py-2 px-4 rounded-lg w-full flex items-center gap-2 opacity-80 text-sm transition duration-200']"
          >
            <span class="text-lg">{{ item.icon }}</span>
            <span>{{ item.name }}</span>
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>