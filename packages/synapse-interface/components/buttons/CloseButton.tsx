import { XIcon } from '@heroicons/react/outline'

export const CloseButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <button
      className={`
        flex items-center justify-center
        m-1
        w-6 h-6
        hover:cursor-pointer
        absolute right-2
        group hover:bg-slate-950/50 rounded-full
      `}
      onClick={onClick}
    >
      <XIcon className="inline w-4 text-white/50 group-hover:text-white" />
    </button>
  )
}
