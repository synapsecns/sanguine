import { joinClassNames } from '@/utils/joinClassNames'

type MaxButtonTypes = {
  disabled: boolean
  onClickBalance: () => void
}

export default function MaxButton({
  disabled,
  onClickBalance,
}: MaxButtonTypes) {
  const className = joinClassNames({
    space: 'px-4 py-1 -ml-1 mr-1 rounded',
    background: 'bg-zinc-100 dark:bg-separator', // TODO: Remove
    // background: 'bg-zinc-100 dark:bg-zinc-700',
    border: 'border border-zinc-200 dark:border-transparent',
    hover: 'enabled:hover:border-zinc-400 enabled:hover:dark:border-zinc-500',
    disabled: 'disabled:opacity-60 disabled:cursor-default',
  })

  return (
    <button className={className} onClick={onClickBalance} disabled={disabled}>
      Max
    </button>
  )
}
