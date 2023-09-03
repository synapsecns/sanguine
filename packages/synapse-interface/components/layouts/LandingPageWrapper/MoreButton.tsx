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
      className={`
        w-full
        group rounded-md
        px-2.5 py-2.5
        border border-[#2F2F2F] hover:border-white
        bg-transparent hover:bg-transparent
        focus:bg-transparent active:bg-transparent
        ${className}
      `}
      {...props}
    >
      <div className="space-x-2">
        <div className="rounded-md">
          <DotsHorizontalIcon className="w-5 h-5 text-white" />
        </div>
      </div>
    </Button>
  )
}
