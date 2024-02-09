import numeral from 'numeral'
import { useCallback, useMemo } from 'react'
import Image from 'next/image'
import { useAppDispatch } from '@/store/hooks'
import {
  setFromChainId,
  setToChainId,
  setFromToken,
  setToToken,
} from '@/slices/bridge/reducer'
import { Chain, Token } from '@/utils/types'
import { formatBigIntToString } from '@/utils/bigint/format'
import { trimTrailingZeroesAfterDecimal } from '@/utils/trimTrailingZeroesAfterDecimal'

function isObject(object): boolean {
  return typeof object === 'object' && object !== null
}

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenAmount,
  isOrigin,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: string | number
  isOrigin: boolean
}) => {
  const dispatch = useAppDispatch()

  const handleSelectChainCallback = useCallback(() => {
    if (isOrigin) {
      dispatch(setFromChainId(chain?.id as number))
    } else {
      dispatch(setToChainId(chain?.id as number))
    }
  }, [isOrigin, chain])

  const handleSelectTokenCallback = useCallback(() => {
    if (isOrigin && chain && token) {
      dispatch(setFromChainId(chain?.id as number))
      dispatch(setFromToken(token as Token))
    } else {
      dispatch(setToChainId(chain?.id as number))
      dispatch(setToToken(token as Token))
    }
  }, [isOrigin, token, chain])

  const tokenDecimals = useMemo(() => {
    if (token && chain) {
      const storedAsObject: boolean = isObject(token?.decimals)
      return storedAsObject ? token.decimals[chain?.id] : token.decimals
    }
    return null
  }, [tokenAmount, token, chain])

  return (
    <div
      data-test-id="transaction-payload-detail"
      className="flex flex-col p-1 space-y-1"
    >
      {chain && (
        <TransactionPayloadDetailButton
          data-test-id="transaction-payload-network"
          onClick={handleSelectChainCallback}
        >
          <Image
            src={chain.chainImg}
            className="items-center w-4 h-4 mr-1.5 rounded-full"
            alt={`${chain.name} icon`}
          />
          <div className="whitespace-nowrap">{chain.name}</div>
        </TransactionPayloadDetailButton>
      )}

      {token && tokenAmount && (
        <TransactionPayloadDetailButton
          data-test-id="transaction-payload-token"
          onClick={handleSelectTokenCallback}
        >
          <Image
            src={token?.icon}
            className="items-center w-4 h-4 mr-1.5 rounded-full"
            alt={`${token?.name} icon`}
          />
          <div className="mr-1">
            {typeof tokenAmount === 'string' && tokenDecimals
              ? trimTrailingZeroesAfterDecimal(
                  formatBigIntToString(BigInt(tokenAmount), tokenDecimals, 4)
                )
              : typeof tokenAmount === 'number'
                ? numeral(tokenAmount).format('0,0.000')
                : "..."
            }
          </div>
          <div className="mt-0.5 text-xs md:text-sm">{token?.symbol}</div>
        </TransactionPayloadDetailButton>
      )}
    </div>
  )
}


function TransactionPayloadDetailButton({ children, ...props}) {
  return (
    <div
      {...props}
      className={`
        flex flex-row px-1 items-center cursor-pointer rounded-sm w-fit
        hover:bg-slate-900/50 active:opacity-[67%]
      `}
    >
      {children}
    </div>
  )
}