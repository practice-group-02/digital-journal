import { useState } from 'react'
import { FiFilter, FiX, FiCalendar, FiMapPin, FiBookmark } from 'react-icons/fi'
import ProgramCard from '../components/ProgramCard'

export default function Programs() {
  const [filtersOpen, setFiltersOpen] = useState(false)
  const [filters, setFilters] = useState({
    country: '',
    type: '',
    deadline: ''
  })

  const countries = ['Казахстан', 'Россия', 'США', 'Германия', 'Великобритания', 'Франция']
  const programTypes = ['Стипендия', 'Грант', 'Стажировка', 'Конкурс']

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="container mx-auto px-4 py-8">
        <div className="flex flex-col md:flex-row gap-8">
          {/* Mobile Filters Button */}
          <button 
            onClick={() => setFiltersOpen(true)}
            className="md:hidden flex items-center gap-2 bg-white px-4 py-2 rounded-lg shadow-sm"
          >
            <FiFilter /> Фильтры
          </button>

          {/* Filters Sidebar */}
          <div className={`${filtersOpen ? 'block' : 'hidden'} md:block fixed md:static inset-0 z-50 bg-white md:bg-transparent p-4 md:p-0 overflow-y-auto`}>
            <div className="bg-white md:bg-transparent p-6 md:p-0 rounded-lg md:rounded-none md:w-64">
              <div className="flex justify-between items-center mb-6 md:mb-4">
                <h3 className="text-lg font-bold">Фильтры</h3>
                <button 
                  onClick={() => setFiltersOpen(false)}
                  className="md:hidden text-gray-500 hover:text-gray-700"
                >
                  <FiX size={20} />
                </button>
              </div>

              <div className="space-y-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">Страна</label>
                  <select
                    value={filters.country}
                    onChange={(e) => setFilters({...filters, country: e.target.value})}
                    className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-300 focus:border-blue-300"
                  >
                    <option value="">Все страны</option>
                    {countries.map(country => (
                      <option key={country} value={country}>{country}</option>
                    ))}
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">Тип программы</label>
                  <select
                    value={filters.type}
                    onChange={(e) => setFilters({...filters, type: e.target.value})}
                    className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-300 focus:border-blue-300"
                  >
                    <option value="">Все типы</option>
                    {programTypes.map(type => (
                      <option key={type} value={type}>{type}</option>
                    ))}
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">Дедлайн до</label>
                  <div className="relative">
                    <input
                      type="date"
                      value={filters.deadline}
                      onChange={(e) => setFilters({...filters, deadline: e.target.value})}
                      className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:ring-2 focus:ring-blue-300 focus:border-blue-300"
                    />
                    <FiCalendar className="absolute right-3 top-2.5 text-gray-400" />
                  </div>
                </div>

                <button 
                  onClick={() => setFilters({ country: '', type: '', deadline: '' })}
                  className="w-full bg-gray-100 hover:bg-gray-200 text-gray-800 py-2 rounded-lg transition"
                >
                  Сбросить фильтры
                </button>
              </div>
            </div>
          </div>

          {/* Programs List */}
          <div className="flex-1">
            <h1 className="text-2xl md:text-3xl font-bold mb-6">Все программы</h1>
            
            <div className="bg-white rounded-xl shadow-sm overflow-hidden">
              <div className="p-6 border-b border-gray-200 flex justify-between items-center">
                <div className="text-gray-600">
                  Найдено: <span className="font-medium">24 программы</span>
                </div>
                <div className="flex items-center gap-2">
                  <span className="text-gray-600">Сортировка:</span>
                  <select className="border-0 font-medium focus:ring-2 focus:ring-blue-300 rounded">
                    <option>По дате</option>
                    <option>По релевантности</option>
                  </select>
                </div>
              </div>

              <div className="divide-y divide-gray-200">
                {[1, 2, 3, 4, 5].map(item => (
                  <ProgramCard key={item} />
                ))}
              </div>

              <div className="p-6 border-t border-gray-200 flex justify-center">
                <button className="bg-white border border-gray-300 hover:bg-gray-50 px-6 py-2 rounded-lg font-medium transition">
                  Показать еще
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}