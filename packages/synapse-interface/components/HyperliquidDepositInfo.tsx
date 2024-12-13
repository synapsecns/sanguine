import Image from 'next/image'
import { CheckCircleIcon } from '@heroicons/react/outline'

import { ARBITRUM, HYPERLIQUID } from '@/constants/chains/master'

export const HyperliquidDepositInfo = ({
  fromChainId,
  isOnArbitrum,
  hasDepositedOnHyperliquid,
}) => {
  return (
    <div className="flex flex-col p-2 mb-2 space-y-1 text-sm border rounded border-zinc-300 dark:border-separator">
      <div className="flex items-center mb-1 space-x-2 ">
        <Image
          loading="lazy"
          src={HYPERLIQUID.chainImg}
          alt="Switch Network"
          width="16"
          height="16"
          className="w-4 h-4 max-w-fit"
        />

        <div>Hyperliquid Deposit</div>
      </div>
      <div className="flex items-center space-x-2">
        <CheckCircleIcon
          className={`w-3 h-3 ${
            fromChainId === ARBITRUM.id && isOnArbitrum
              ? 'text-green-500'
              : 'text-gray-500'
          }`}
        />
        <div>Bridge to Arbitrum</div>
      </div>
      <div className="flex items-center space-x-2">
        <CheckCircleIcon
          className={`w-3 h-3 ${
            hasDepositedOnHyperliquid ? 'text-green-500' : 'text-gray-500'
          }`}
        />
        <div>Deposit to Hyperliquid</div>
      </div>
    </div>
  )
}
