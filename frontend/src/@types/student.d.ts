export interface Student {
  id?: number
  name: string
  email: string
  dob: string
  gpa: number
}

export interface CreateStudentRequest {
  name: string
  email: string
  dob: string
  gpa: number
}

export interface StudentQuery {
  name?: string
  email?: string
  minGpa?: number
  maxGpa?: number
  sortBy?: string
  sortOrder?: 'asc' | 'desc'
  page?: number
  pageSize?: number
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  pageSize: number
}
