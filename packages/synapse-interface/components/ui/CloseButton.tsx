import { XIcon } from '@heroicons/react/outline'

export const CloseButton = ({
  onClick,
  className,
}: {
  onClick: () => void
  className?: string
}) => {
  return (
    <button
      className={`absolute flex items-center justify-center w-8 h-8 hover:cursor-pointer right-1 ${className}`}
      onClick={onClick}
    >
      <XIcon className="inline w-5 text-white" />
    </button>
  )
}
