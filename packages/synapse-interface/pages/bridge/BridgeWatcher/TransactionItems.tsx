import { BigNumber } from '@ethersproject/bignumber'

import { WETH } from '@constants/tokens/swapMaster'
import { ETH } from '@constants/tokens/master'

import { CHAINS_BY_ID } from '@constants/chains'

import { formatBNToString } from '@bignumber/format'

import { getCoinTextColorCombined } from '@styles/tokens'
import { getNetworkLinkTextColor } from '@styles/chains'
import { AddToWalletMiniButton } from '@components/buttons/AddToWalletButton'
import ExplorerLink from './ExplorerLink'
import { commify } from '@ethersproject/units'

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

export function SubTransactionItem({
  transactionHash,
  currentChainId,
  blockNumber,
  chainId,
  timestamp,
  event,
  args,
  token,
  tokenAmount,
  ...rest
}) {
  // console.log({ event, args, token, tokenAmount })
  const isCurrentChain = chainId == currentChainId

  const { to } = args ?? {}
  const { name, chainImg } = CHAINS_BY_ID[chainId]

  if (token?.symbol == WETH.symbol) {
    token = ETH
  }

  let decimalsToShow
  if (token?.swapableType == 'ETH') {
    decimalsToShow = 3
  } else if (['NFD', 'DOG', 'L2DAO', 'PLS'].includes(token?.swapableType)) {
    decimalsToShow = 0
  } else if (['OHM'].includes(token?.swapableType)) {
    decimalsToShow = 4
  } else {
    decimalsToShow = 2
  }

  let showAddBtn
  if (token?.symbol == ETH.symbol) {
    showAddBtn = false
  } else {
    showAddBtn = true
  }

  let formattedTokenAmount
  if (tokenAmount) {
    try {
      formattedTokenAmount = commify(
        formatBNToString(
          tokenAmount.mul(
            BigNumber.from(10).pow(18 - (token?.decimals[chainId] ?? 18))
          ),
          18,
          decimalsToShow
        )
      )
    } catch (err) {}
  }
  // console.log({token,tokenAmount})

  return (
    <div className="flex items-center p-1 border border-gray-700 rounded-lg">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chainImg.src}
        />
      </div>
      <div className="flex-grow">
        <div>
          <div className="w-full text-sm">
            <ExplorerLink
              overrideExistingClassname={true}
              className={`
                ${getNetworkLinkTextColor(chainId)}
                opacity-70 hover:opacity-100
              `}
              chainId={chainId}
              transactionHash={transactionHash}
            />
          </div>
        </div>
        <div className="w-full">
          <div className="w-full text-sm">
            {formattedTokenAmount && (
              <span className="font-medium text-gray-400">
                {formattedTokenAmount}
              </span>
            )}
            {token && (
              <>
                <span
                  className={`font-medium ${getCoinTextColorCombined(token)}`}
                >
                  {' '}
                  {token.symbol}{' '}
                </span>
                <img
                  src={token.icon.src}
                  className="w-4 h-4 inline -mt-0.5 rounded-sm"
                />
              </>
            )}
          </div>
        </div>
      </div>
      <div className="flex-shrink">
        {isCurrentChain && showAddBtn && (
          <AddToWalletMiniButton
            token={token}
            icon={token?.icon?.src}
            chainId={chainId}
            className={`float-right inline-block `}
          />
        )}
      </div>
    </div>
  )
}
