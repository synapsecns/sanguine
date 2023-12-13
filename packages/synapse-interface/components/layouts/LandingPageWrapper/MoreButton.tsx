import Button from '@tw/Button'
import { DotsHorizontalIcon } from '@heroicons/react/outline'

export function MoreButton({
  open,
  onClick,
  className,
  ...props
}: {
  open: boolean
  onClick?: () => void
  className?: string
  props?: any
}) {
  return (
    <Button
      onClick={onClick ? onClick : () => {}}
      className="w-full rounded-md p-2 border border-[#2F2F2F] hover:border-white"
      {...props}
    >
      <DotsHorizontalIcon className="w-5 h-5" />
    </Button>
  )
}
