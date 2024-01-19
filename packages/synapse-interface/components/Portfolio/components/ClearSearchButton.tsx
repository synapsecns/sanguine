import { XIcon } from '@heroicons/react/outline'

export const ClearSearchButton = ({
  show,
  onClick,
}: {
  show: boolean
  onClick: () => void
}) => {
  return (
    <button
      data-test-id="clear-search-button"
      className={`
        ${show ? 'visible' : 'invisible'}
        flex w-6 h-6 mr-2
        items-center justify-center
        border border-separator rounded-full
        hover:cursor-pointer hover:border-secondary
      `}
      onClick={onClick}
    >
      <XIcon strokeWidth={3} className="inline w-4 text-secondary" />
    </button>
  )
}
