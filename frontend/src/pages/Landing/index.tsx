import { useNavigate } from 'react-router-dom'

export default function Landing() {
  const navigate = useNavigate()

  return (
    <div>
      <button onClick={() => navigate('/student-management')}>Student Management</button>
    </div>
  )
}
