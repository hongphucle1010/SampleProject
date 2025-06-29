import React from 'react'

interface ActionButtonProps {
  children: React.ReactNode
  onClick?: () => void
  disabled?: boolean
  variant?: 'primary' | 'secondary' | 'danger'
  size?: 'small' | 'medium' | 'large'
  className?: string
  type?: 'button' | 'submit' | 'reset'
}

const ActionButton: React.FC<ActionButtonProps> = ({
  children,
  onClick,
  disabled = false,
  variant = 'primary',
  size = 'medium',
  className = '',
  type = 'button'
}) => {
  const baseClasses = 'font-bold rounded transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2'
  
  const variantClasses = {
    primary: 'bg-blue-500 hover:bg-blue-700 disabled:bg-blue-300 text-white focus:ring-blue-500',
    secondary: 'bg-gray-500 hover:bg-gray-700 disabled:bg-gray-300 text-white focus:ring-gray-500',
    danger: 'bg-red-500 hover:bg-red-700 disabled:bg-red-300 text-white focus:ring-red-500'
  }
  
  const sizeClasses = {
    small: 'py-1 px-3 text-sm',
    medium: 'py-2 px-4 text-sm',
    large: 'py-3 px-6 text-lg'
  }

  return (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled}
      className={`${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
    >
      {children}
    </button>
  )
}

export default ActionButton 