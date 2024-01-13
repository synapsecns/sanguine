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
      className="rounded-md p-2 border border-zinc-500 hover:bg-zinc-100 hover:dark:bg-zinc-800"
      {...props}
    >
      <DotsHorizontalIcon className="w-5 h-5 stroke-black dark:stroke-white" />
    </Button>
  )
}
