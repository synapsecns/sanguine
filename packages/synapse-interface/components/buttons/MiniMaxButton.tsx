import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
}: {
  onClickBalance: () => void
  disabled: boolean
}) {
  const baseClassName =
    'pl-3 pr-3 pt-2 pb-2 mr-2 rounded-md text-md font-light bg-bgLighter border border-transparent'

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
