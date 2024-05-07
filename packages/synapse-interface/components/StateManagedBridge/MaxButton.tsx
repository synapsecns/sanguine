import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

export const MaxButton = ({ onClick, isHidden }) => {
  const buttonClassName = joinClassNames({
    display: `${isHidden ? 'hidden' : 'block'}`,
    text: 'text-synapsePurple text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  return (
    <div onClick={onClick} className={buttonClassName}>
      Max
    </div>
  )
}
