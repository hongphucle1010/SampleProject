import { APIResponse } from '../../@types/api'
import { Student, PaginatedResponse } from '../../@types/student'
import { BaseApi } from './baseApi'

export class StudentApi extends BaseApi {
  static async getStudents() {
    return (await this.apiClient.get<APIResponse<PaginatedResponse<Student>>>('/students')).data
  }
}
