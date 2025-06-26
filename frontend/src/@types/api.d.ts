// src/@types/api.d.ts

export interface APIResponse<T> {
  status: 'success'
  message: string
  data: T
}

export interface APIError {
  status: 'error'
  message: string
  details: string
}

export interface RefreshTokenResponse {
  access_token: string
}
