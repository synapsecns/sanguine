import React from 'react'
import { useTranslations } from 'use-intl'

import { joinClassNames } from '@/utils/joinClassNames'

export const MaxButton = ({ onClick, isHidden }) => {
  const buttonClassNames = {
    display: `${isHidden ? 'hidden' : 'block'}`,
    spacing: 'px-1.5',
    text: 'text-fuchsia-400 text-xxs md:text-xs',
    animation: 'transition-all duration-150 transform-gpu',
    hover: 'hover:opacity-70 cursor-pointer',
  }

  const t = useTranslations('Bridge')

  return (
    <button onClick={onClick} className={joinClassNames(buttonClassNames)}>
      {t('Max')}
    </button>
  )
}
