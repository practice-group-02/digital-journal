import { Link } from 'react-router-dom'
import { FiSearch, FiMail, FiArrowRight } from 'react-icons/fi'
import ProgramCard from '../components/ProgramCard'

export default function Home() {
  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="bg-gradient-to-r from-blue-600 to-indigo-800 text-white py-24">
        <div className="container mx-auto px-4 text-center">
          <h1 className="text-4xl md:text-5xl font-bold mb-6 leading-tight">
            Откройте мир образовательных возможностей
          </h1>
          <p className="text-xl md:text-2xl mb-8 max-w-3xl mx-auto opacity-90">
            Найдите идеальные стипендии, гранты и стажировки для студентов Казахстана
          </p>
          
          <div className="flex flex-col md:flex-row gap-4 justify-center max-w-2xl mx-auto">
            <div className="relative flex-grow">
              <input
                type="text"
                placeholder="Поиск программ..."
                className="w-full py-4 px-6 pr-12 rounded-full text-gray-800 focus:outline-none focus:ring-2 focus:ring-blue-300"
              />
              <FiSearch className="absolute right-6 top-1/2 transform -translate-y-1/2 text-gray-500 text-xl" />
            </div>
            <Link 
              to="/programs"
              className="bg-white text-blue-600 px-8 py-4 rounded-full font-bold hover:bg-blue-50 transition flex items-center justify-center gap-2"
            >
              Найти <FiArrowRight />
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-white">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                title: "Стипендии",
                desc: "Полное или частичное финансирование обучения",
                color: "bg-blue-100 text-blue-600"
              },
              {
                title: "Гранты",
                desc: "Финансовая поддержка исследовательских проектов",
                color: "bg-purple-100 text-purple-600"
              },
              {
                title: "Стажировки",
                desc: "Практический опыт в международных компаниях",
                color: "bg-green-100 text-green-600"
              }
            ].map((item, index) => (
              <div key={index} className="p-6 rounded-xl hover:shadow-lg transition">
                <span className={`${item.color} w-12 h-12 rounded-full flex items-center justify-center mb-4`}>
                  {index + 1}
                </span>
                <h3 className="text-xl font-bold mb-2">{item.title}</h3>
                <p className="text-gray-600">{item.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Popular Programs */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-4">
          <div className="flex justify-between items-center mb-8">
            <h2 className="text-2xl md:text-3xl font-bold">Популярные программы</h2>
            <Link to="/programs" className="text-blue-600 hover:underline flex items-center gap-1">
              Все программы <FiArrowRight />
            </Link>
          </div>
          
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map(item => (
              <ProgramCard key={item} />
            ))}
          </div>
        </div>
      </section>

      {/* Newsletter */}
      <section className="py-16 bg-white">
        <div className="container mx-auto px-4 text-center max-w-3xl">
          <h2 className="text-2xl md:text-3xl font-bold mb-4">Будьте в курсе новых возможностей</h2>
          <p className="text-gray-600 mb-8">
            Подпишитесь на нашу рассылку и получайте персональные подборки стипендий
          </p>
          
          <div className="flex flex-col sm:flex-row gap-2">
            <input
              type="email"
              placeholder="Ваш email"
              className="flex-grow py-3 px-6 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-300"
            />
            <button className="bg-blue-600 hover:bg-blue-700 text-white py-3 px-8 rounded-full font-medium transition flex items-center justify-center gap-2">
              Подписаться <FiMail />
            </button>
          </div>
        </div>
      </section>
    </div>
  )
}