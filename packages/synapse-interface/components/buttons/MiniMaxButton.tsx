type MaxButtonTypes = {
  disabled: boolean
  onClickBalance: () => void
}

const join = (a) => Object.values(a).join(' ')

export default function MaxButton({ disabled, onClickBalance }: MaxButtonTypes) {
  const className = join({
    space: 'px-4 py-1 mr-1 rounded',
    bgColor: 'bg-[#565058]', // NEW: 'bg-zinc-100 dark:bg-zinc-700',
    borderColor: 'border border-zinc-200 dark:border-transparent',
    borderHover:
      'enabled:hover:border-zinc-400 enabled:hover:dark:border-zinc-500',
    styleDisabled: 'disabled:opacity-60 disabled:cursor-default',
  })

  return (
    <button className={className} onClick={onClickBalance} disabled={disabled}>
      Max
    </button>
  )
}
