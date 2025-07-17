<script setup lang="ts">
import { ref, inject, nextTick } from 'vue'
import DropDownModel from './DropDownModel.vue'

interface Props {
  isTyping: boolean
}

interface Emits {
  sendMessage: [message: string]
  modelSelected: [model: string, provider: string]
}

defineProps<Props>()
const emit = defineEmits<Emits>()

const isDarkMode = inject('theme')
const newMessage = ref('')
const textareaRef = ref<HTMLTextAreaElement>()

const sendMessage = () => { 
  if (!newMessage.value.trim()) return
  
  const message = newMessage.value.trim()
  newMessage.value = ''
  emit('sendMessage', message)
  
  nextTick(() => {
    autoResize()
  })
}

const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

const autoResize = () => {
  if (textareaRef.value) {
    textareaRef.value.style.height = 'auto'
    textareaRef.value.style.height = `${textareaRef.value.scrollHeight}px`
  }
}

const handleSelectedModel = (val: Array<string>) => {
  emit('modelSelected', val[0], val[1])
}

defineExpose({
  newMessage,
  textareaRef,
  autoResize
})
</script>

<template>
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
        />
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
</template>