import api from './api'

export interface ChatMessage {
  role: 'user' | 'assistant'
  content: string
}

export interface ChatResponse {
  message: string
  error?: string
}

class ChatService {
  // Placeholder for AI integration
  async sendMessage(message: string): Promise<ChatResponse> {
    try {

      const response = await api.post('send-message', {
        "prompt": message,
      })

      return { message: response.data.data.response }

    } catch (error) {
      console.error('Chat service error:', error)
      return {
        message: 'Sorry, I encountered an error. Please try again.',
        error: error instanceof Error ? error.message : 'Unknown error'
      }
    }
  }
}

export default new ChatService()