import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
}: {
  onClickBalance: () => void
  disabled: boolean
}) {
  const space = 'px-4 py-1 mr-1 rounded'
  const bgColor = 'bg-[#565058]' // NEW: 'bg-zinc-100 dark:bg-zinc-700'
  const borderColor = 'border border-zinc-200 dark:border-transparent'
  const borderHover =
    'enabled:hover:border-zinc-400 enabled:hover:dark:border-zinc-500'
  const styleDisabled = 'disabled:opacity-60 disabled:cursor-default'

  return (
    <Button
      className={`${space} ${bgColor} ${borderColor} ${borderHover} ${styleDisabled}`}
      onClick={onClickBalance}
      disabled={disabled}
    >
      Max
    </Button>
  )
}