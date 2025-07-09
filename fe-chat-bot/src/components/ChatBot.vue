<script setup lang="ts">
import { ref, inject } from 'vue'
import chatService from '../services/chatService'
import DropDownModel from './DropDownModel.vue';
import { marked } from 'marked'
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-dark.css'

// custom renderer
const renderer = new marked.Renderer()

renderer.code = function({ text, lang }) {
  const validLang = lang && hljs.getLanguage(lang) ? lang : 'plaintext'
  const highlighted = hljs.highlight(text, { language: validLang, ignoreIllegals: true }).value
  return `<pre><code class="hljs ${validLang}">${highlighted}</code></pre>`
}

marked.use({ renderer })

interface Message {
  id: number
  text: string
  isUser: boolean
  timestamp: Date
}

const messages = ref<Message[]>([
  {
    id: 1,
    text: "Halo! Saya asisten AI Anda. Ada yang bisa saya bantu hari ini?",
    isUser: false,
    timestamp: new Date()
  }
])

const newMessage = ref('')
const isTyping = ref(false)
const isDarkMode = inject('theme')
const messagesContainer = ref<HTMLElement>()
const bottomMarker = ref<HTMLElement | null>(null)
let selectedModel = ref('gemini-2.5-flash-lite-preview-06-17') // Default model
let selectedProvider = ref('gemini') // Default provider
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
  
  // Show typing indicator
  isTyping.value = true
  setTimeout(() => {
    if (bottomMarker.value) {
      bottomMarker.value.scrollIntoView({ behavior: 'smooth' })
    }
    autoResize()
  }, 5);

  const response = await chatService.sendMessage(userInput, selectedModel.value, selectedProvider.value)
    const botMessage: Message = {
    id: messageIdCounter++,
    text: response.message,
    isUser: false,
    timestamp: new Date()
  }
  messages.value.push(botMessage)
  isTyping.value = false
  setTimeout(() => {
    if (bottomMarker.value) {
      bottomMarker.value.scrollIntoView({ behavior: 'smooth' })
    }
  }, 5);
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

const textareaRef = ref<HTMLTextAreaElement>()

const autoResize = () => {
  if (textareaRef.value) {
    textareaRef.value.style.height = 'auto'
    textareaRef.value.style.height = `${textareaRef.value.scrollHeight}px`
  }
}

function handleSelectedModel(val: Array<string>) {
  selectedModel.value = val[0]
  selectedProvider.value = val[1]
}

</script>

<template>
  <div 
    class="flex flex-col h-screen max-h-screen"
  >

    <!-- Messages Container -->
    <div 
      ref="messagesContainer"
      class="flex-1 px-4 pt-6 pb-40 space-y-4 chat-container max-w-4xl w-full mx-auto"
    >
      <div
        v-for="message in messages"
        :key="message.id"
        :class="['flex', message.isUser ? 'justify-end' : 'justify-start']"
        class="animate-slide-up"
      >
        <div class="flex flex-col space-y-1">
          <div 
            :class="[
              'message-bubble',
              message.isUser 
                ? (isDarkMode ? 'user-message-dark' : 'user-message') 
                : (isDarkMode ? 'bot-message-dark' : 'bot-message')
            ]"
          >
            <p class="text-sm leading-relaxed" v-html="highlightResponse(message.text)"></p>
          </div>
          <div :class="['text-xs text-gray-500', message.isUser ? 'text-right' : 'text-left']">
            {{ formatTime(message.timestamp) }}
          </div>
        </div>
      </div>

      <!-- Typing indicator -->
      <div v-if="isTyping" class="flex justify-start animate-slide-up">
        <div :class="['message-bubble', isDarkMode ? 'bot-message-dark' : 'bot-message']">
          <div class="typing-dots">
            <div class="typing-dot" style="animation-delay: 0ms"></div>
            <div class="typing-dot" style="animation-delay: 150ms"></div>
            <div class="typing-dot" style="animation-delay: 300ms"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Input Area -->
    <div class="fixed bottom-0 left-0 right-0 mx-5 md:mx-0">
      <div 
        class="flex items-end space-x-3 max-w-3xl mx-auto p-3 rounded-3xl shadow"
        :class="isDarkMode ? 'bg-gray-800' : 'bg-white/80'"
      >
        <div 
          class="flex-1 relative rounded-2xl border"
          :class="isDarkMode ? 'bg-gray-900 border-gray-700' : 'bg-white border-gray-200'"
        >
           <textarea
            v-model="newMessage"
            @keydown="handleKeyDown"
            @input="autoResize"
            ref="textareaRef"
            placeholder="Type your message here…"
            :class="[
              'w-full resize-none min-h-16 max-h-60 rounded-2xl px-4 py-3 pr-12 border-0 outline-none focus:ring-0 transition-colors text-sm overflow-hidden',
              isDarkMode 
                ? 'bg-gray-900 text-gray-100' 
                : 'bg-white text-gray-900'
            ]"
            rows="1"
            ></textarea>
        <DropDownModel @selectedModel="handleSelectedModel" />
        </div>
        <button
          @click="sendMessage"
          :disabled="!newMessage.trim() || isTyping"
          class="inline-flex items-center justify-center w-11 h-11 rounded-full transition-colors duration-200 focus:outline-none focus:ring-0 bg-gray-900 disabled:bg-gray-700"
        >
          <svg 
            class="w-5 h-5"
            :class="isDarkMode ? 'text-white' : 'text-white'"
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
      <p :class="['text-xs text-center my-1', isDarkMode ? 'text-gray-400' : 'text-gray-500']">
        Press Enter to send, Shift+Enter for a new line
      </p>
    </div>

    <div ref="bottomMarker"></div>
  </div>
</template>