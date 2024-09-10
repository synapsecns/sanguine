import { useTranslations } from 'next-intl'

import { joinClassNames } from '@/utils/joinClassNames'

export const NewTag = () => {
  const t = useTranslations('Bridge')
  const classNames = {
    space: 'px-2 py-[2px] rounded-md',
    border: 'border border-fuchsia-500',
    background: 'bg-gradient-to-r from-fuchsia-950 to-purple-900',
    font: 'text-sm ',
  }
  return <div className={joinClassNames(classNames)}>{t('New')}!</div>
}
