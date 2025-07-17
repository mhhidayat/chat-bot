<script setup lang="ts">
import { inject } from 'vue'

const isDarkMode = inject('theme')
</script>

<template>
  <div class="flex justify-start animate-slide-up">
    <div :class="['message-bubble', isDarkMode ? 'bot-message-dark' : 'bot-message']">
      <div class="typing-dots">
        <div class="typing-dot" style="animation-delay: 0ms"></div>
        <div class="typing-dot" style="animation-delay: 150ms"></div>
        <div class="typing-dot" style="animation-delay: 300ms"></div>
      </div>
    </div>
  </div>
</template>