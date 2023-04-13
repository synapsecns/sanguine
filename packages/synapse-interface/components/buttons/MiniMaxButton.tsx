import Button from '@tw/Button'

export default function MiniMaxButton({
  onClickBalance,
}: {
  onClickBalance: () => void
}) {
  return (
    <Button
      className={`
        pt-1 pb-1 pl-2 pr-2 mr-2
        rounded-md
        text-sm font-medium
        bg-bgLighter hover:bg-bgLightest active:bg-bgLightest
      `}
      onClick={onClickBalance}
    >
      Max
    </Button>
  )
}
