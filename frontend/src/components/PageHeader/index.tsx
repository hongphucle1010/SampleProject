import React from 'react'

interface PageHeaderProps {
  title: string
  subtitle?: string
  className?: string
}

const PageHeader: React.FC<PageHeaderProps> = ({ 
  title, 
  subtitle, 
  className = '' 
}) => {
  return (
    <div className={`px-6 py-4 bg-gray-50 border-b border-gray-200 ${className}`}>
      <h1 className="text-2xl font-bold text-gray-900">{title}</h1>
      {subtitle && (
        <p className="text-gray-600 mt-1">{subtitle}</p>
      )}
    </div>
  )
}

export default PageHeader 