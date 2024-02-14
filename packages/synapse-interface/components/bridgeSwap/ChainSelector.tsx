import { CHAINS_BY_ID } from '@/constants/chains'

import { ChevronDownIcon } from '@heroicons/react/outline'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const ChainSelector = ({ chainId, label, onClick, ...props }) => {

  const chain = CHAINS_BY_ID[chainId]

  return (
    <button
      {...props}
      className={`
        group
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
      <div className="flex items-center">
        <div className="flex items-center space-x-3">
            { chainId &&
              <div>
                <img
                  src={chain?.chainImg?.src}
                  alt={chain?.name}
                  className="w-6 h-6 rounded-sm"
                />
              </div>
            }
          <div className="text-left">
            <div className="text-xs text-secondaryTextColor">
                {label}
            </div>
            <div className="text-md text-primaryTextColor">
              {chain?.name ?? "Network"}
            </div>
          </div>
          <ChevronDownIcon
            className={`
                rotate-0
                text-white/30 group-hover:text-white/80 w-4 h-4
                active:text-white/80
                group-active:rotate-180 transition-all
            `}
          />
        </div>
      </div>
    </button>
  )
}
