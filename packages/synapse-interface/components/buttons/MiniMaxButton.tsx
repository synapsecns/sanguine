import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
}: {
  onClickBalance: () => void
  disabled: boolean
}) {
  const baseClassName =
    'px-4 py-1.5 rounded bg-zinc-200 dark:bg-zinc-600 border border-transparent text-inherit'

  const className = disabled
    ? `${baseClassName} opacity-60 cursor-default`
    : `${baseClassName} hover:border-[#AC8FFF]`

  return (
    <Button
      className={className}
      onClick={disabled ? undefined : onClickBalance}
    >
      Max
    </Button>
  )
}
