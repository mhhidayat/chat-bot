<script setup lang="ts">
import { ref, provide } from 'vue'
import Theme from './components/Theme.vue'
import ChatBot from './components/ChatBot.vue'

const isDarkMode = ref(false)
provide('theme', isDarkMode)
</script>

<template>
  <div class="h-screen">
    <Theme @theme="isDarkMode = $event" />
    <ChatBot />
  </div>
</template>