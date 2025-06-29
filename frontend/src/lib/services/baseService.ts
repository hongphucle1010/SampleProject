/* eslint-disable @typescript-eslint/no-explicit-any */
import { BaseApi } from '../apis/baseApi'

export class BaseService {
  protected static async callApi<C extends typeof BaseApi, M extends keyof C>(
    apiClass: C,
    method: M,
    ...args: C[M] extends (...args: any[]) => any ? Parameters<C[M]> : never
  ): Promise<C[M] extends (...args: any[]) => any ? ReturnType<C[M]> : never> {
    try {
      const methodName = `${apiClass.name}.${String(method)}`

      // If mocking is enabled, return mock data
      if (MockService.isMockEnabled()) {
        const mockResponse = MockService.getMockData(methodName)
        if (mockResponse !== undefined) {
          console.warn(`Mocking API: ${methodName}`)
          return mockResponse as any
        }
      }

      // Call the actual API method if no mock data is found
      if (typeof apiClass[method] !== 'function') {
        throw new Error(`Method ${String(method)} is not a valid function`)
      }

      const result = await (apiClass[method] as (...args: any[]) => any)(...args)
      return result as any
    } catch (error) {
      console.error(error)
      throw error
    }
  }
}

// Utility class for handling mocks
export class MockService {
  private static enabled = false
  private static mockData: Record<string, any> = {}

  static enableMock() {
    this.enabled = true
  }

  static disableMock() {
    this.enabled = false
  }

  static setMockData(apiMethod: string, data: any) {
    this.mockData[apiMethod] = data
  }

  static getMockData(apiMethod: string) {
    return this.mockData[apiMethod]
  }

  static isMockEnabled() {
    return this.enabled
  }
}
