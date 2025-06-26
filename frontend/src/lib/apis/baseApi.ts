import { apiClient } from './apiClient'
import { APIResponse, RefreshTokenResponse } from '../../@types/api'

export abstract class BaseApi {
  protected static apiClient = apiClient

  static initInterceptors() {
    this.apiClient.interceptors.request.use(
      (config) => {
        const token = this.getAccessToken()
        if (token) config.headers.Authorization = `Bearer ${token}`
        return config
      },
      (error) => Promise.reject(error)
    )

    this.apiClient.interceptors.response.use(
      (res) => res,
      async (err) => {
        const originalRequest = err.config

        if (err.response?.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true

          const res = await apiClient.post<APIResponse<RefreshTokenResponse>>('/auth/refresh-token', null)

          const newAccessToken = res.data.data.access_token
          localStorage.setItem('accessToken', newAccessToken)

          // Set it for current retry request
          originalRequest.headers['Authorization'] = `Bearer ${newAccessToken}`
          return apiClient(originalRequest)
        }

        return Promise.reject(err)
      }
    )
  }

  protected static getAccessToken() {
    return localStorage.getItem('accessToken')
  }

  protected static clearAccessToken() {
    localStorage.removeItem('accessToken')
  }

  protected static setAccessToken(accessToken: string) {
    localStorage.setItem('accessToken', accessToken)
  }

  protected static async request<T>(method: 'get' | 'post' | 'put' | 'delete' | 'patch', url: string, data?: object) {
    try {
      const response = await (method === 'get' ? this.apiClient.get<T>(url) : this.apiClient[method]<T>(url, data))
      return response
    } catch (error) {
      console.error(error)
      throw error
    }
  }
}
