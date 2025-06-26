import { Link } from 'react-router-dom'
import { FiBookmark, FiCalendar, FiMapPin } from 'react-icons/fi'

export default function ProgramCard() {
  return (
    <div className="p-6 hover:bg-gray-50 transition">
      <div className="flex flex-col md:flex-row md:justify-between md:items-start gap-4">
        <div className="flex-1">
          <div className="flex items-center gap-3 mb-2">
            <span className="bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded">
              Стипендия
            </span>
            <span className="text-gray-500 text-sm">Назарбаев Университет</span>
          </div>
          
          <h3 className="text-xl font-bold mb-2 text-gray-900 hover:text-blue-600 transition">
            <Link to="/programs/1">Стипендия "Болашак" для магистратуры</Link>
          </h3>
          
          <p className="text-gray-600 mb-4 line-clamp-2">
            Полное финансирование обучения в ведущих мировых университетах для граждан Казахстана
          </p>
          
          <div className="flex flex-wrap gap-4">
            <div className="flex items-center text-gray-500">
              <FiMapPin className="mr-2" />
              <span>Казахстан, США, Великобритания</span>
            </div>
            <div className="flex items-center text-gray-500">
              <FiCalendar className="mr-2" />
              <span>До 15 декабря 2023</span>
            </div>
          </div>
        </div>
        
        <div className="flex md:flex-col gap-2">
          <button className="p-2 text-gray-400 hover:text-blue-600 transition">
            <FiBookmark />
          </button>
          <Link 
            to="/programs/1"
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition whitespace-nowrap"
          >
            Подробнее
          </Link>
        </div>
      </div>
    </div>
  )
}