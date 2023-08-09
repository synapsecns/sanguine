import { XIcon } from '@heroicons/react/outline'

export const CloseButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button
      className={`
        flex items-center justify-center
        w-8 h-8
        hover:cursor-pointer
      `}
      onClick={onClick}
    >
      <XIcon className="inline w-5 text-white" />
    </button>
  )
}
