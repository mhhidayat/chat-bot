<script setup lang="ts">
import { ref, nextTick, onMounted } from 'vue'
import chatService from '../services/chatService'
import { marked } from 'marked'

interface Message {
  id: number
  text: string
  isUser: boolean
  timestamp: Date
}

const messages = ref<Message[]>([
  {
    id: 1,
    text: "Hello! I'm your AI assistant. How can I help you today?",
    isUser: false,
    timestamp: new Date()
  }
])

const newMessage = ref('')
const isTyping = ref(false)
const messagesContainer = ref<HTMLElement>()
let messageIdCounter = 2

const sendMessage = async () => {
  if (!newMessage.value.trim()) return

  // Add user message
  const userMessage: Message = {
    id: messageIdCounter++,
    text: newMessage.value.trim(),
    isUser: true,
    timestamp: new Date()
  }
  
  messages.value.push(userMessage)
  const userInput = newMessage.value.trim()
  newMessage.value = ''
  
  await scrollToBottom()
  
  // Show typing indicator
  isTyping.value = true
  await scrollToBottom()

  isTyping.value = false

  const response = await chatService.sendMessage(userInput)
  const botMessage: Message = {
    id: messageIdCounter++,
    text: response.message,
    isUser: false,
    timestamp: new Date()
  }
  
  messages.value.push(botMessage)
  await scrollToBottom()
}

const scrollToBottom = async () => {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const formatTime = (date: Date): string => {
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

function highlightResponse(message: string) {
  return marked.parse(message)
}

onMounted(() => {
  scrollToBottom()
})
</script>

<template>
  <div class="flex flex-col h-screen max-h-screen bg-gradient-to-br from-blue-50 via-white to-indigo-50">
    <!-- Header -->
    <div class="bg-white border-b border-gray-200 px-6 py-4 shadow-sm">
      <div class="flex items-center space-x-3">
        <div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-indigo-600 rounded-full flex items-center justify-center">
          <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
          </svg>
        </div>
        <div>
          <h1 class="text-lg font-semibold text-gray-900">AI Assistant</h1>
          <p class="text-sm text-gray-500">Online and ready to help</p>
        </div>
      </div>
    </div>

    <!-- Messages Container -->
    <div 
      ref="messagesContainer"
      class="flex-1 overflow-y-auto px-4 py-6 space-y-4 chat-container"
    >
      <div
        v-for="message in messages"
        :key="message.id"
        :class="['flex', message.isUser ? 'justify-end' : 'justify-start']"
        class="animate-slide-up"
      >
        <div class="flex flex-col space-y-1">
          <div :class="['message-bubble', message.isUser ? 'user-message' : 'bot-message']">
            <p class="text-sm leading-relaxed" v-html=highlightResponse(message.text)></p>
          </div>
          <div :class="['text-xs text-gray-400', message.isUser ? 'text-right' : 'text-left']">
            {{ formatTime(message.timestamp) }}
          </div>
        </div>
      </div>

      <!-- Typing indicator -->
      <div v-if="isTyping" class="flex justify-start animate-slide-up">
        <div class="message-bubble bot-message">
          <div class="typing-dots">
            <div class="typing-dot" style="animation-delay: 0ms"></div>
            <div class="typing-dot" style="animation-delay: 150ms"></div>
            <div class="typing-dot" style="animation-delay: 300ms"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Input Area -->
    <div class="bg-white border-t border-gray-200 px-4 py-4">
      <div class="flex items-end space-x-3 max-w-4xl mx-auto">
        <div class="flex-1 relative">
          <textarea
            v-model="newMessage"
            @keydown="handleKeyDown"
            placeholder="Type your message here..."
            class="w-full resize-none rounded-2xl border border-gray-300 px-4 py-3 pr-12 focus:border-primary-500 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-opacity-20 transition-colors text-sm"
            rows="1"
            style="max-height: 120px; min-height: 44px;"
          ></textarea>
        </div>
        <button
          @click="sendMessage"
          :disabled="!newMessage.trim() || isTyping"
          class="inline-flex items-center justify-center w-11 h-11 rounded-full bg-primary-500 hover:bg-primary-600 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:ring-offset-2"
        >
          <svg 
            class="w-5 h-5 text-white" 
            fill="none" 
            stroke="currentColor" 
            viewBox="0 0 24 24"
          >
            <path 
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
            ></path>
          </svg>
        </button>
      </div>
      <p class="text-xs text-gray-500 text-center mt-2">
        Press Enter to send, Shift+Enter for new line
      </p>
    </div>
  </div>
</template>