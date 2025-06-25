import { useParams } from 'react-router-dom'
import { getProgramById } from '../services/programs'
import { useEffect, useState } from 'react'
import LoadingSpinner from '../components/LoadingSpinner'
import { FaCalendarAlt, FaMapMarkerAlt, FaUniversity, FaExternalLinkAlt } from 'react-icons/fa'

export default function ProgramDetail() {
  const { id } = useParams()
  const [program, setProgram] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await getProgramById(id)
        setProgram(data)
      } catch (error) {
        console.error('Error:', error)
      } finally {
        setLoading(false)
      }
    }

    fetchData()
  }, [id])

  if (loading) return <LoadingSpinner />

  if (!program) return (
    <div className="container py-8 text-center">
      <h2 className="text-2xl font-bold mb-4">Программа не найдена</h2>
      <p className="text-gray-600">Попробуйте выбрать другую программу из списка</p>
    </div>
  )

  return (
    <div className="container py-8">
      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <div className="p-8">
          <h1 className="text-3xl font-bold mb-4">{program.title}</h1>
          
          <div className="flex flex-wrap gap-4 mb-6">
            <span className="bg-blue-100 text-blue-800 px-3 py-1 rounded-full text-sm">
              {program.type}
            </span>
            <span className="bg-green-100 text-green-800 px-3 py-1 rounded-full text-sm">
              {program.country}
            </span>
          </div>
          
          <div className="grid md:grid-cols-2 gap-6 mb-8">
            <div className="space-y-4">
              <div className="flex items-start">
                <FaUniversity className="mt-1 mr-3 flex-shrink-0 text-gray-500" />
                <div>
                  <h3 className="font-semibold">Организация</h3>
                  <p>{program.organization}</p>
                </div>
              </div>
              
              <div className="flex items-start">
                <FaMapMarkerAlt className="mt-1 mr-3 flex-shrink-0 text-gray-500" />
                <div>
                  <h3 className="font-semibold">Страна</h3>
                  <p>{program.country}</p>
                </div>
              </div>
            </div>
            
            <div className="space-y-4">
              <div className="flex items-start">
                <FaCalendarAlt className="mt-1 mr-3 flex-shrink-0 text-gray-500" />
                <div>
                  <h3 className="font-semibold">Дедлайн</h3>
                  <p>{new Date(program.deadline).toLocaleDateString()}</p>
                </div>
              </div>
              
              {program.link && (
                <div className="flex items-start">
                  <FaExternalLinkAlt className="mt-1 mr-3 flex-shrink-0 text-gray-500" />
                  <div>
                    <h3 className="font-semibold">Сайт программы</h3>
                    <a 
                      href={program.link} 
                      target="_blank" 
                      rel="noopener noreferrer"
                      className="text-blue-600 hover:underline"
                    >
                      Перейти на сайт
                    </a>
                  </div>
                </div>
              )}
            </div>
          </div>
          
          <div className="prose max-w-none">
            <h2 className="text-xl font-bold mb-4">Описание</h2>
            <p className="whitespace-pre-line">{program.description}</p>
          </div>
        </div>
        
        <div className="bg-gray-50 px-8 py-4 border-t border-gray-200">
          <a 
            href={program.link} 
            target="_blank" 
            rel="noopener noreferrer"
            className="inline-block bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-medium transition"
          >
            Подать заявку
          </a>
        </div>
      </div>
    </div>
  )
}