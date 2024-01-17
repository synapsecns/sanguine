import { XIcon } from '@heroicons/react/outline'

export const CloseButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button
      className={`
        w-8 h-8
        hover:cursor-pointer
        hover:bg-zinc-300 hover:dark:bg-zinc-950 mr-1 rounded hover:text-black hover:dark:text-zinc-300
      `}
      onClick={onClick}
    >
      <XIcon className="inline w-5 text-inherit" />
    </button>
  )
}
