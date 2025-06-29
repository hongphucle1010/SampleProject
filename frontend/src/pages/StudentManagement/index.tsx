import { useState } from 'react'
import { StudentService } from '../../lib/services/studentService'
import { Student } from '../../@types/student'
import { LoadingSpinner, ErrorDisplay, PageHeader, ActionButton, DataTable } from '../../components'

const StudentManagement = () => {
  const [students, setStudents] = useState<Student[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [hasLoaded, setHasLoaded] = useState(false)

  const fetchStudents = async () => {
    try {
      setLoading(true)
      setError(null)
      const response = await StudentService.getStudents()
      console.log('API Response:', response)

      // The API returns a paginated response, so we need to access response.data.data
      const studentData = response.data?.data || []
      console.log('Student data:', studentData)

      setStudents(Array.isArray(studentData) ? studentData : [])
      setHasLoaded(true)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch students')
      console.error('Error fetching students:', err)
    } finally {
      setLoading(false)
    }
  }

  const studentColumns = [
    {
      key: 'id',
      header: 'ID',
      render: (value: number) => value || 'N/A'
    },
    {
      key: 'name',
      header: 'Name',
      render: (value: string) => <span className='font-medium text-gray-900'>{value}</span>
    },
    {
      key: 'email',
      header: 'Email',
      render: (value: string) => <span className='text-gray-500'>{value}</span>
    },
    {
      key: 'dob',
      header: 'Date of Birth',
      render: (value: string) => <span className='text-gray-500'>{new Date(value).toLocaleDateString()}</span>
    },
    {
      key: 'gpa',
      header: 'GPA',
      render: (value: number) => (
        <span
          className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
            value >= 3.5
              ? 'bg-green-100 text-green-800'
              : value >= 3.0
                ? 'bg-yellow-100 text-yellow-800'
                : 'bg-red-100 text-red-800'
          }`}
        >
          {value.toFixed(2)}
        </span>
      )
    }
  ]

  // Show initial state with button if data hasn't been loaded yet
  if (!hasLoaded && !loading) {
    return (
      <div className='container mx-auto px-4 py-8'>
        <div className='bg-white shadow-lg rounded-lg overflow-hidden'>
          <PageHeader title='Student Management' subtitle='Manage and view student information' />

          <div className='p-8 text-center'>
            <p className='text-gray-600 mb-6'>Click the button below to load student data</p>
            <ActionButton onClick={fetchStudents} disabled={loading} size='large'>
              {loading ? 'Loading...' : 'Load Students'}
            </ActionButton>
          </div>
        </div>
      </div>
    )
  }

  if (loading) {
    return <LoadingSpinner size='large' className='min-h-screen' />
  }

  if (error) {
    return <ErrorDisplay error={error} onRetry={fetchStudents} />
  }

  return (
    <div className='container mx-auto px-4 py-8'>
      <div className='bg-white shadow-lg rounded-lg overflow-hidden'>
        <PageHeader title='Student Management' subtitle='Manage and view student information' />

        <DataTable data={students} columns={studentColumns} emptyMessage='No students found' />

        <div className='px-6 py-4 bg-gray-50 border-t border-gray-200'>
          <div className='flex justify-between items-center'>
            <span className='text-sm text-gray-700'>Total students: {students.length}</span>
            <ActionButton onClick={fetchStudents} size='small'>
              Refresh
            </ActionButton>
          </div>
        </div>
      </div>
    </div>
  )
}

export default StudentManagement
