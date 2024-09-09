import { joinClassNames } from '@/utils/joinClassNames'

export function BridgeAmountContainer({ children }) {
  const classNames = {
    space: 'flex items-center gap-4 p-2 rounded-md w-full',
    bgColor: 'bg-white dark:bg-inherit',
    borderColor: 'border border-zinc-200 dark:border-zinc-700',
  }

  return <div className={joinClassNames(classNames)}>{children}</div>
}
