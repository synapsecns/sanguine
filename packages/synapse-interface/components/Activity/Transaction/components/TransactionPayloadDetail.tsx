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
import { formatAmount } from '@/utils/formatAmount'

function isObject(object): boolean {
  return typeof object === 'object' && object !== null
}

export const TransactionPayloadDetail = ({
  chain,
  token,
  tokenAmount,
  isOrigin,
  className,
  showChain = true,
  disabled,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: string | number
  isOrigin: boolean
  className?: string
  showChain?: boolean
  disabled: boolean
}) => {
  const dispatch = useAppDispatch()

  const handleSelectChainCallback = useCallback(() => {
    if (!disabled) {
      if (isOrigin) {
        dispatch(setFromChainId(chain?.id as number))
      } else {
        dispatch(setToChainId(chain?.id as number))
      }
    }
  }, [isOrigin, chain, disabled])

  const handleSelectTokenCallback = useCallback(() => {
    if (!disabled) {
      if (isOrigin && chain && token) {
        dispatch(setFromChainId(chain?.id as number))
        dispatch(setFromToken(token as Token))
      } else {
        dispatch(setToChainId(chain?.id as number))
        dispatch(setToToken(token as Token))
      }
    }
  }, [isOrigin, token, chain, disabled])

  const tokenDecimals = useMemo(() => {
    if (token && chain) {
      const storedAsObject: boolean = isObject(token?.decimals)
      return storedAsObject ? token.decimals[chain?.id] : token.decimals
    }
    return null
  }, [tokenAmount, token, chain])

  const buttonStyle =
    'flex gap-1.5 pl-1.5 pr-2.5 py-0.5 -my-0.5 items-center cursor-pointer rounded border border-transparent hover:border-surface hover:bg-tint active:opacity-70 w-fit'

  return (
    <div data-test-id="transaction-payload-detail" className={className}>
      {chain && showChain && (
        <div
          data-test-id="transaction-payload-network"
          onClick={handleSelectChainCallback}
          className={buttonStyle}
        >
          <Image
            src={chain.chainImg}
            className="w-4 h-4 pt-0.5 ml-0.5"
            alt={`${chain.name} icon`}
          />
          {chain.name}
        </div>
      )}

      {token && tokenAmount && (
        <div
          data-test-id="transaction-payload-token"
          onClick={handleSelectTokenCallback}
          className={buttonStyle}
        >
          <Image
            src={token?.icon}
            className="items-center w-5 h-5 mt-px"
            alt={`${token?.name} icon`}
          />
          {typeof tokenAmount === 'string' && tokenDecimals ? (
            <span>
              {formatAmount(
                formatBigIntToString(BigInt(tokenAmount), tokenDecimals, 18)
              )}
            </span>
          ) : typeof tokenAmount === 'number' ? (
            <span>
              {trimTrailingZeroesAfterDecimal(
                formatAmount(tokenAmount.toString())
              )}
            </span>
          ) : (
            <span>…</span>
          )}
          <span className="mt-0.5 text-sm">{token?.symbol}</span>
        </div>
      )}
    </div>
  )
}
