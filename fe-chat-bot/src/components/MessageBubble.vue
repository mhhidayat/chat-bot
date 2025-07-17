<script setup lang="ts">
import { inject } from 'vue'
import { marked } from 'marked'
import hljs from 'highlight.js'

interface Message {
    id: number
    text: string
    isUser: boolean
    timestamp: Date
}

interface Props {
    message: Message
}

defineProps<Props>()

const isDarkMode = inject('theme')

// Custom renderer for markdown
const renderer = new marked.Renderer()
renderer.code = function ({ text, lang }) {
    const validLang = lang && hljs.getLanguage(lang) ? lang : 'plaintext'
    const highlighted = hljs.highlight(text, { language: validLang, ignoreIllegals: true }).value
    return `<pre><code class="hljs ${validLang}">${highlighted}</code></pre>`
}

// Configure marked options for better formatting
marked.use({
    renderer,
    breaks: true, // Convert \n to <br>
    gfm: true     // GitHub Flavored Markdown
})



const highlightResponse = (message: string) => {
    return marked.parse(message)
}
</script>

<template>
    <div :class="['flex', message.isUser ? 'justify-end' : 'justify-start']" class="animate-slide-up">
        <div class="flex flex-col space-y-1">
            <div :class="[
                'message-bubble',
                message.isUser
                    ? (isDarkMode ? 'user-message-dark' : 'user-message')
                    : (isDarkMode ? 'bot-message-dark' : 'bot-message')
            ]">
                <div class="text-sm leading-relaxed prose prose-sm max-w-none" v-html="highlightResponse(message.text)">
                </div>
            </div>

        </div>
    </div>
</template>