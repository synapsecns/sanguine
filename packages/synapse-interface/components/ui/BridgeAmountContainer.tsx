import { joinClassNames } from '@/utils/joinClassNames'

export function BridgeAmountContainer({ children }) {
  const className = joinClassNames({
    space: 'flex items-center gap-4 p-2 rounded-md w-full',
    bgColor: 'bg-white dark:bg-inherit',
    borderColor: 'border border-zinc-200 dark:border-zinc-700',
  })

  return <div className={className}>{children}</div>
}
