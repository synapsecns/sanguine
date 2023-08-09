import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
}: {
  onClickBalance: () => void
  disabled: boolean
}) {
  const baseClassName =
    'pl-lg pr-lg pt-sm pb-sm mr-2 rounded-md text-md font-light bg-bgLighter border border-transparent'

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
