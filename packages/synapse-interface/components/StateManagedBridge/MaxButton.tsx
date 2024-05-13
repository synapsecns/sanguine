import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'

export const MaxButton = ({ onClick, isHidden }) => {
  const buttonClassName = joinClassNames({
    display: `${isHidden ? 'hidden' : 'block'}`,
    spacing: 'px-1.5',
    text: 'text-fuchsia-400 text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  })

  return (
    <button onClick={onClick} className={buttonClassName}>
      Max
    </button>
  )
}
