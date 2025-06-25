import api from './api'

export const getPrograms = async (filters = {}) => {
  try {
    const response = await api.get('/programs', { 
      params: {
        country: filters.country,
        type: filters.type,
        deadline: filters.deadline
      }
    })
    return response.data
  } catch (error) {
    console.error('Error fetching programs:', error)
    throw error
  }
}

export const getProgramById = async (id) => {
  try {
    const response = await api.get(`/programs/${id}`)
    return response.data
  } catch (error) {
    console.error('Error fetching program:', error)
    throw error
  }
}

export const getFeaturedPrograms = async () => {
  try {
    const response = await api.get('/programs/featured')
    return response.data
  } catch (error) {
    console.error('Error fetching featured programs:', error)
    throw error
  }
}