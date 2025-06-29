import React from 'react'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Landing from '../pages/Landing'
import StudentManagement from '../pages/StudentManagement'

const Router: React.FC = () => {
  const userRoutes = [
    {
      path: '/',
      element: <Landing />,
      errorElement: (
        <div>
          <h1>Home Error</h1>
        </div>
      )
    },
    {
      path: '/student-management',
      element: <StudentManagement />,
      errorElement: (
        <div>
          <h1>Student Management Error</h1>
        </div>
      )
    }
  ]
  const router = createBrowserRouter(userRoutes)
  return <RouterProvider router={router} />
}

export default Router
