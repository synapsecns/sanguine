import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '@/components/icons/DropDownArrowSvg'

import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const ChainSelector = ({ chainId, label, onClick, ...props }) => {

  const chain = CHAINS_BY_ID[chainId]

  let buttonContent

  if (chainId) {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div>
          <img
            src={chain?.chainImg?.src}
            alt={chain?.name}
            className="w-6 h-6 rounded-sm"
          />
        </div>
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">{label}</div>
          <div className="text-md text-primaryTextColor">{chain?.name}</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  } else {
    buttonContent = (
      <div className="flex items-center space-x-3">
        <div className="text-left">
          <div className="text-xs text-secondaryTextColor">{label}</div>
          <div className="text-md text-primaryTextColor">Network</div>
        </div>
        <DropDownArrowSvg />
      </div>
    )
  }

  return (
    <button
      {...props}
      className={`
        bg-transparent
        p-md
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        border border-transparent
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
        rounded-sm
      `}
      onClick={onClick}
    >
      <div className="flex items-center">{buttonContent}</div>
    </button>
  )
}
