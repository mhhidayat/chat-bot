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
  async sendMessage(message: string, model: string, provider: string): Promise<ChatResponse> {
    try {

      const response = await api.post('send-message', {
        "prompt": message,
        "model": model,
        "provider": provider
      })

      return { message: response.data.data.response }

    } catch (error) {
      console.error('Chat service error:', error)
      return {
        message: `Silahkan coba ganti model atau coba beberapa saat lagi, mohon maklum keterbatasannya. Versi berbayar masih dalam pengembangan. 🚀`,
        error: error instanceof Error ? error.message : 'Unknown error'
      }
    }
  }
}

export default new ChatService()