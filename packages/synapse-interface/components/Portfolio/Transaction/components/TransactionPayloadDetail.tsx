import { useCallback, useMemo } from 'react'
import numeral from 'numeral'
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
  className,
  showChain = true,
}: {
  chain?: Chain
  token?: Token
  tokenAmount?: string | number
  isOrigin: boolean
  className?: string
  showChain?: boolean
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

  const buttonStyle = `
    flex gap-1.5 pl-1.5 pr-2.5 py-0.5 -my-0.5 items-center cursor-pointer rounded
    border border-transparent hover:bg-bgBase/20 active:bg-bgBase/30
    active:opacity-60
    w-fit
  `

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
        <TransactionPayloadDetailButton
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
              {trimTrailingZeroesAfterDecimal(
                formatBigIntToString(BigInt(tokenAmount), tokenDecimals, 4)
              )}
            </span>
          ) : typeof tokenAmount === 'number' ? (
            <span>
              {trimTrailingZeroesAfterDecimal(
                numeral(tokenAmount).format('0,0.000')
              )}
            </span>
          ) : (
            <span>…</span>
          )}
          <span className="mt-0.5 text-sm">{token?.symbol}</span>
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
        hover:bg-slate-400/20 active:opacity-60 space-x-1
      `}
    >
      {children}
    </div>
  )
}