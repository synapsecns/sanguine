import { joinClassNames } from '@/utils/joinClassNames'

type MaxButtonTypes = {
  onClickBalance: () => void
  isDisabled: boolean
  isHidden: boolean
}

export default function MaxButton({
  onClickBalance,
  isDisabled,
  isHidden,
}: MaxButtonTypes) {
  const className = joinClassNames({
    space: 'px-4 py-1 -ml-1 mr-1 rounded',
    background: 'bg-zinc-100 dark:bg-separator',
    border: 'border border-zinc-200 dark:border-transparent',
    hover: 'enabled:hover:border-zinc-400 enabled:hover:dark:border-zinc-500',
    disabled: 'disabled:opacity-60 disabled:cursor-default',
  })

  if (isHidden) {
    return null
  } else {
    return (
      <button
        className={className}
        onClick={onClickBalance}
        disabled={isDisabled}
      >
        Max
      </button>
    )
  }
}
