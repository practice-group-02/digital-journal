export default function FilterPanel({ filters, setFilters }) {
  const countries = ['Казахстан', 'Россия', 'США', 'Германия', 'Великобритания']
  const programTypes = ['Стипендия', 'Грант', 'Стажировка', 'Конкурс']

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h3 className="font-bold text-lg mb-4">Фильтры</h3>
      
      <div className="space-y-6">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Страна</label>
          <select
            value={filters.country}
            onChange={(e) => setFilters({...filters, country: e.target.value})}
            className="w-full border rounded-lg p-2"
          >
            <option value="">Все страны</option>
            {countries.map(country => (
              <option key={country} value={country}>{country}</option>
            ))}
          </select>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Тип программы</label>
          <select
            value={filters.type}
            onChange={(e) => setFilters({...filters, type: e.target.value})}
            className="w-full border rounded-lg p-2"
          >
            <option value="">Все типы</option>
            {programTypes.map(type => (
              <option key={type} value={type}>{type}</option>
            ))}
          </select>
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Дедлайн до</label>
          <input
            type="date"
            value={filters.deadline}
            onChange={(e) => setFilters({...filters, deadline: e.target.value})}
            className="w-full border rounded-lg p-2"
          />
        </div>
        
        <button 
          onClick={() => setFilters({ country: '', type: '', deadline: '' })}
          className="w-full bg-gray-200 hover:bg-gray-300 py-2 rounded-lg transition"
        >
          Сбросить фильтры
        </button>
      </div>
    </div>
  )
}