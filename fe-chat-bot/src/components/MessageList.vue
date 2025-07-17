<script setup lang="ts">
import { ref, onMounted, onUnmounted, inject } from 'vue'
import MessageBubble from './MessageBubble.vue'
import TypingIndicator from './TypingIndicator.vue'

interface Message {
  id: number
  text: string
  isUser: boolean
  timestamp: Date
}

interface Props {
  messages: Message[]
  isTyping: boolean
}

defineProps<Props>()

const isDarkMode = inject('theme')
const messagesContainer = ref<HTMLElement>()
const bottomMarker = ref<HTMLElement | null>(null)
const showScrollToBottom = ref(false)

const checkScrollPosition = () => {
  const scrollTop = window.pageYOffset || document.documentElement.scrollTop
  const scrollHeight = document.documentElement.scrollHeight
  const clientHeight = window.innerHeight
  const isNearBottom = scrollHeight - scrollTop - clientHeight < 200
  showScrollToBottom.value = !isNearBottom && scrollHeight > clientHeight
}

const scrollToBottom = () => {
  if (bottomMarker.value) {
    bottomMarker.value.scrollIntoView({ behavior: 'smooth' })
  }
}

onMounted(() => {
  window.addEventListener('scroll', checkScrollPosition)
  // Initial check
  checkScrollPosition()
})

onUnmounted(() => {
  window.removeEventListener('scroll', checkScrollPosition)
})

defineExpose({
  messagesContainer,
  bottomMarker,
  scrollToBottom
})
</script>

<template>
  <div class="relative flex-1">
    <div ref="messagesContainer"
      class="flex-1 px-4 pt-6 pb-60 space-y-4 chat-container max-w-4xl w-full mx-auto h-full overflow-y-auto">
      <MessageBubble v-for="message in messages" :key="message.id" :message="message" />

      <TypingIndicator v-if="isTyping" />
    </div>

    <!-- Scroll to Bottom Button -->
    <transition enter-active-class="transition-all duration-300 ease-out"
      leave-active-class="transition-all duration-300 ease-in" enter-from-class="opacity-0 translate-y-2 scale-95"
      enter-to-class="opacity-100 translate-y-0 scale-100" leave-from-class="opacity-100 translate-y-0 scale-100"
      leave-to-class="opacity-0 translate-y-2 scale-95">
      <div v-if="showScrollToBottom" class="fixed bottom-[10.5rem] left-1/2 transform -translate-x-1/2 z-10">
        <button @click="scrollToBottom" :class="[
          'flex items-center gap-2 px-4 py-2 rounded-full shadow-lg transition-all duration-200 hover:scale-105 text-xs font-medium',
          isDarkMode
            ? 'bg-zinc-700/80 hover:bg-zinc-600/80 text-zinc-100 border border-zinc-600/80'
            : 'bg-white/80 hover:bg-gray-50/80 text-gray-700/80 border border-gray-200/80'
        ]" title="Scroll to bottom">
          <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
          </svg>
          <span>Scroll to bottom</span>
        </button>
      </div>
    </transition>

    <div ref="bottomMarker"></div>
  </div>
</template>