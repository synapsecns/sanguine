import { BigNumber } from '@ethersproject/bignumber'

import { WETH } from '@constants/tokens/swapMaster'
import { ETH } from '@constants/tokens/master'

import { CHAINS_BY_ID } from '@constants/chains'

import { formatBNToString } from '@bignumber/format'

import { getCoinTextColorCombined } from '@styles/tokens'
import { getNetworkLinkTextColor } from '@styles/chains'
import { AddToWalletMiniButton } from '@components/buttons/AddToWalletButton'
import ExplorerLink from '../pages/bridge/BridgeWatcher/ExplorerLink'
import { commify } from '@ethersproject/units'
import { BridgeWatcherTx } from '@types'
export function CheckingConfPlaceholder({ chainId }) {
  const { name, chainImg } = CHAINS_BY_ID[chainId]

  return (
    <div className="flex items-center p-1 max-w-[50%] rounded-lg ">
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chainId)}
                opacity-70 pr-2
              `}
            >
              Confirmations left on {name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}

export function PendingCreditTransactionItem({ chainId }) {
  const { name, chainImg } = CHAINS_BY_ID[chainId]

  return (
    <div className="flex items-center p-1 rounded-lg ">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chainImg.src}
        />
      </div>
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chainId)}
                opacity-70 pr-2
              `}
            >
              Waiting to be credited on
              <br />
              {name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}

export function EmptySubTransactionItem({ chainId }) {
  const { chainImg } = CHAINS_BY_ID[chainId]

  return (
    <div className="flex items-center p-1 rounded-lg ">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chainImg.src}
        />
      </div>
    </div>
  )
}

export function CreditedTransactionItem({ chainId }) {
  const { name, chainImg } = CHAINS_BY_ID[chainId]
  return (
    <div className="flex items-center p-1 border border-gray-700 rounded-lg">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chainImg}
        />
      </div>
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chainId)}
                opacity-70 pr-2
              `}
            >
              Bridging Completed on
              <br />
              {name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}
