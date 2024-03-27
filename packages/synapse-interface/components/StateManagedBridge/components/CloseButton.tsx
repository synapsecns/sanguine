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
      className={`
        flex items-center justify-center
        absolute right-2 w-8 h-8
        hover:cursor-pointer
        ${className}
      `}
      onClick={onClick}
    >
      <XIcon className="inline w-5 text-white" />
    </button>
  )
}
