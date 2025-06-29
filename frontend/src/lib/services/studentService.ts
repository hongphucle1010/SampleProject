import { BaseService } from './baseService'
import { StudentApi } from '../apis/studentApi'

export class StudentService extends BaseService {
  static async getStudents() {
    return this.callApi(StudentApi, 'getStudents')
  }
}
