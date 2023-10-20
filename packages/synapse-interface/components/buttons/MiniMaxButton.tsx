import Button from '@tw/Button'
import { LoaderIcon } from 'react-hot-toast'

export default function MiniMaxButton({
  onClickBalance,
  disabled,
  loading,
}: {
  onClickBalance: () => void
  disabled: boolean
  loading?: boolean
}) {
  const baseClassName = `
    w-[89px] h-[32px]
    flex items-center mr-2 py-1 justify-center
    text-sm text-white
    border rounded-sm
    bg-bgLighter border border-transparent
  `

  const className = disabled
    ? `${baseClassName} opacity-60 cursor-default`
    : `${baseClassName} hover:border-[#AC8FFF]`

  return (
    <Button
      className={className}
      onClick={disabled ? undefined : onClickBalance}
    >
      {loading ? <LoaderIcon /> : 'Max'}
    </Button>
  )
}
