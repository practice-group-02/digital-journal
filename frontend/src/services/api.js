import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  timeout: 5000,
})

// Добавляем интерцептор для обработки ошибок
api.interceptors.response.use(
  response => response,
  error => {
    if (error.response) {
      // Обработка HTTP ошибок
      console.error('API Error:', error.response.status, error.response.data)
    } else if (error.request) {
      // Запрос был сделан, но ответ не получен
      console.error('API Error: No response received')
    } else {
      // Ошибка при настройке запроса
      console.error('API Error:', error.message)
    }
    
    return Promise.reject(error)
  }
)

export default api