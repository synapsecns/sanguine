import { type Chain } from '@/utils/types'

export const ChainOption = ({
  option,
  isSelected,
  onSelect,
  isOrigin,
}: {
  option: any
  isSelected: boolean
  onSelect: (option: Chain) => void
  isOrigin?: boolean
}) => {
  return (
    <li
      key={option.id}
      className={`
      pl-2.5 pr-2.5 py-2.5 rounded-[.1875rem] border border-solid
       active:opacity-40
      cursor-pointer whitespace-nowrap group flex justify-between items-center
      ${isSelected ? 'hover:opacity-70' : 'border-transparent'}
    `}
      onClick={() => onSelect(option)}
    >
      <div className="flex items-center">
        {option?.imgUrl && (
          <img
            src={option?.imgUrl}
            alt={`${option?.name} chain icon`}
            className="inline w-4 h-4 mr-2"
          />
        )}
        {option?.name}
      </div>
    </li>
  )
}
