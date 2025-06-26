import axios from 'axios'
import { BACKEND_URL } from '../../constants/api'

export const apiHost = `${BACKEND_URL}`

export const apiClient = axios.create({
  baseURL: apiHost,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true // Send cookies with requests
})
