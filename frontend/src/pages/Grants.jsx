import { FiAlertCircle } from 'react-icons/fi'

export default function Grants() {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="container mx-auto px-4 py-12">
        <div className="max-w-3xl mx-auto bg-white rounded-xl shadow-sm overflow-hidden">
          <div className="p-8 text-center">
            <div className="mx-auto w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mb-4">
              <FiAlertCircle className="text-blue-600 text-2xl" />
            </div>
            <h1 className="text-2xl md:text-3xl font-bold mb-4">Раздел грантов в разработке</h1>
            <p className="text-gray-600 mb-6">
              Мы активно работаем над наполнением этого раздела. Скоро здесь появится полная информация о доступных грантах.
            </p>
            <div className="bg-blue-50 border border-blue-200 rounded-lg p-4 text-left max-w-md mx-auto">
              <h3 className="font-medium text-blue-800 mb-2">Что вы можете сделать сейчас:</h3>
              <ul className="list-disc list-inside text-blue-700 space-y-1">
                <li>Посмотреть доступные программы</li>
                <li>Подписаться на уведомления</li>
                <li>Связаться с нашей поддержкой</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}