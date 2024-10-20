import { useContext } from 'react'
import { Chain } from 'types'

import { Web3Context } from '@/providers/Web3Provider'
import { ConnectedIndicator } from '@/components/icons/ConnectedIndicator'

export const ChainOption = ({
  option,
  isSelected,
  onSelect,
  isOrigin,
}: {
  option: Chain
  isSelected: boolean
  onSelect: (option: Chain) => void
  isOrigin?: boolean
}) => {
  const web3Context = useContext(Web3Context)

  const { web3Provider } = web3Context || {}

  const networkId = web3Provider?.networkId
  const connectedAddress = web3Provider?.connectedAddress

  return (
    <li
      key={option.id}
      className={`
      pl-2.5 pr-2.5 py-2.5 rounded-[.1875rem] border border-solid
      hover:border-[--synapse-focus] active:opacity-40
      cursor-pointer whitespace-nowrap group flex justify-between items-center
      ${
        isSelected
          ? 'border-[--synapse-focus] hover:opacity-70'
          : 'border-transparent'
      }
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
      {isOrigin &&
        (option.id === networkId ? (
          <ConnectedIndicator />
        ) : connectedAddress ? (
          <span className="opacity-0 text-sm text-[--synapse-secondary] group-hover:opacity-100">
            Switch
          </span>
        ) : null)}
    </li>
  )
}
