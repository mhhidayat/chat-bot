<script setup lang="ts">
import { ref } from 'vue'
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

const messagesContainer = ref<HTMLElement>()
const bottomMarker = ref<HTMLElement | null>(null)

defineExpose({
  messagesContainer,
  bottomMarker
})
</script>

<template>
  <div 
    ref="messagesContainer"
    class="flex-1 px-4 pt-6 pb-40 space-y-4 chat-container max-w-4xl w-full mx-auto"
  >
    <MessageBubble
      v-for="message in messages"
      :key="message.id"
      :message="message"
    />

    <TypingIndicator v-if="isTyping" />
  </div>
  
  <div ref="bottomMarker"></div>
</template>