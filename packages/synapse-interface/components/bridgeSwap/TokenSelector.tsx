import { ChevronDownIcon } from '@heroicons/react/solid'

import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export const TokenSelector = ({ token, label, onClick, ...props }) => {
  return (
    <button
      {...props}
      className={`
        group
        p-md rounded-sm min-w-[80px]
        bg-slate-400/10 hover:bg-slate-400/20
        border border-white/10
        ${getMenuItemHoverBgForCoin(token?.color)}
        ${getBorderStyleForCoinHover(token?.color)}
      `}
      onClick={onClick}
    >
      <div className="flex items-center justify-items-end space-x-2">
        { token &&
          <div className="flex-none hidden md:inline-block">
            <img src={token?.icon?.src} alt={token?.symbol} className="size-6" />
          </div>
        }
        <div className="text-left flex-grow">
          <div className="text-lg text-primaryTextColor">
            {token?.symbol ?? label}
          </div>
        </div>
        <div className='flex-shrink'>
          <ChevronDownIcon
            className={`
                flex-none
                rotate-0
                text-white/30 group-hover:text-white/80 size-4
                active:text-white/80
                group-active:rotate-180 transition-all
            `}
          />
        </div>
      </div>
    </button>
  )
}
