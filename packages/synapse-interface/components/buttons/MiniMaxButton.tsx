import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
}) {
  const baseClassName =
    'px-2 py-0.5 mr-2 rounded-md text-sm font-light bg-bgBase/10 ring-1 ring-bgBase/10 border border-transparent'

  const disabledClassName = disabled
    ? `opacity-60 cursor-default`
    : `hover:border-[#AC8FFF]`

  return (
    <Button
      className={`${baseClassName} ${disabledClassName}`}
      onClick={disabled ? undefined : onClickBalance}
    >
      Max
    </Button>
  )
}
