import { Link } from 'react-router-dom'

export default function NotFound() {
  return (
    <div className="container py-8 text-center">
      <h1 className="text-4xl font-bold mb-4">404</h1>
      <p className="text-xl mb-6">Страница не найдена</p>
      <Link 
        to="/" 
        className="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition"
      >
        Вернуться на главную
      </Link>
    </div>
  )
}