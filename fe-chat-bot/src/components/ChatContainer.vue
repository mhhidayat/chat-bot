<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import chatService from '../services/chatService'
import PromptOptions from './PromptOptions.vue'
import MessageList from './MessageList.vue'
import ChatInput from './ChatInput.vue'

interface Message {
  id: number
  text: string
  isUser: boolean
  timestamp: Date
}

const messages = ref<Message[]>([])
const isTyping = ref(false)
const showPromptOptions = ref<boolean>(true)

let selectedModel = ref('gemini-2.5-flash-lite-preview-06-17')
let selectedProvider = ref('gemini')
let messageIdCounter = 1

const messageListRef = ref<InstanceType<typeof MessageList>>()
const chatInputRef = ref<InstanceType<typeof ChatInput>>()

const sendMessage = async (message: string) => {
  // Add user message
  const userMessage: Message = {
    id: messageIdCounter++,
    text: message,
    isUser: true,
    timestamp: new Date()
  }
  
  messages.value.push(userMessage)
  
  // Show typing indicator
  isTyping.value = true
  setTimeout(() => {
    scrollToBottom()
    chatInputRef.value?.autoResize()
  }, 5)

  try {
    const response = await chatService.sendMessage(message, selectedModel.value, selectedProvider.value)
    const botMessage: Message = {
      id: messageIdCounter++,
      text: response.message,
      isUser: false,
      timestamp: new Date()
    }
    messages.value.push(botMessage)
  } catch (error) {
    const errorMessage: Message = {
      id: messageIdCounter++,
      text: 'Sorry, I encountered an error. Please try again.',
      isUser: false,
      timestamp: new Date()
    }
    messages.value.push(errorMessage)
  } finally {
    isTyping.value = false
    setTimeout(() => {
      scrollToBottom()
    }, 5)
  }
}

const scrollToBottom = () => {
  if (messageListRef.value?.bottomMarker) {
    messageListRef.value.bottomMarker.scrollIntoView({ behavior: 'smooth' })
  }
}

const handleSelectedModel = (model: string, provider: string) => {
  selectedModel.value = model
  selectedProvider.value = provider
}

const handleSelectedPrompt = (prompt: string) => {
  if (chatInputRef.value) {
    chatInputRef.value.newMessage = prompt
    nextTick(() => {
      chatInputRef.value?.textareaRef?.focus()
    })
  }
}

watch(() => chatInputRef.value?.newMessage, (newValue) => {
  if (newValue?.trim() || messages.value.length > 0) {
    showPromptOptions.value = false
  } else {
    showPromptOptions.value = true
  }
}, { deep: true })
</script>

<template>
  <div class="flex flex-col h-screen max-h-screen">
    <transition
      enter-active-class="transition-opacity duration-500 interpolate-ease-in-out"
      leave-active-class="transition-opacity duration-500 interpolate-ease-in-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <PromptOptions 
        v-if="showPromptOptions" 
        @promptSelected="handleSelectedPrompt" 
      />
    </transition>

    <MessageList 
      ref="messageListRef"
      :messages="messages"
      :isTyping="isTyping"
    />

    <ChatInput 
      ref="chatInputRef"
      :isTyping="isTyping"
      @sendMessage="sendMessage"
      @modelSelected="handleSelectedModel"
    />
  </div>
</template>