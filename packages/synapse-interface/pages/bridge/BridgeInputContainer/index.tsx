import _ from 'lodash'
import { Zero } from '@ethersproject/constants'
import { formatBNToString } from '@bignumber/format'

import SelectTokenDropdown from './SelectTokenDropdown'
import { ChainLabel } from './ChainLabel'
import { Token } from '@utils/classes/Token'
import { useAccount, useBalance, useNetwork } from 'wagmi'

import { ChainId } from '@constants/networks'

import SwitchButton from '@components/buttons/SwitchButton'
import MiniMaxButton from '@components/buttons/MiniMaxButton'
import { BigNumber } from '@ethersproject/bignumber'
// import { useTokenBalance } from '@hooks/tokens/useTokenBalances'
// import { useNetworkController } from '@hooks/wallet/useNetworkController'

import { cleanNumberInput } from '@utils/cleanNumberInput'

export default function BridgeInputContainer({
  selected,
  inputValue,
  isSwapFrom,
  onChangeAmount,
  swapFromToChains,
  setDisplayType,
  inputRef,
  chainId,
  onChangeChain,
}: {
  isSwapFrom: boolean
  selected: Token
  onChangeSelected: (v: Token) => void
  onChangeAmount: (value: string) => void
  inputValue: string
  inputRef: any
  swapFromToChains?: () => void
  tokens: any[]
  chainId: number
  setDisplayType: (value: string) => void
  onChangeChain: () => void
}) {
  const { address } = useAccount()
  const { chain } = useNetwork()
  const tokenAddr = selected.addresses[chain?.id as keyof Token['addresses']]

  const {
    data: evmFromTokenBalance,
    isError: balanceError,
    isLoading: balanceLoading,
  } = useBalance({
    address: `0x${tokenAddr.slice(2)}`,
  })

  let tokenBalance: BigNumber = evmFromTokenBalance?.value ?? Zero
  const formattedBalance = formatBNToString(
    tokenBalance,
    selected.decimals[chainId as keyof Token['decimals']],
    4
  )

  let isConnected = address !== null

  function onChange(e: any) {
    let val = e.target.value

    let cleanValue = cleanNumberInput(val)

    onChangeAmount(cleanValue)
  }

  function onClickBalance() {
    onChangeAmount(
      formatBNToString(
        tokenBalance,
        4,
        selected.decimals[chainId as keyof Token['decimals']]
      )
    )
  }

  let balanceLabel
  if (isSwapFrom) {
    balanceLabel = (
      <a onClick={onClickBalance} className="hover:underline group">
        <small className="text-xs text-gray-500 cursor-pointer group-hover:underline">
          Max:{' '}
          <span className="font-medium text-gray-400 ">{formattedBalance}</span>{' '}
          {selected.symbol}
        </small>
      </a>
    )
  }
  return (
    <div
      className={`
        text-left px-2 sm:px-4 pt-2 pb-4 rounded-xl
        bg-bgLight
      `}
    >
      <div>
        <div className="pt-1 pb-3">
          {!isSwapFrom && (
            <div className="absolute">
              <div className="-mt-12">
                <SwitchButton onClick={swapFromToChains ?? (() => null)} />
              </div>
            </div>
          )}
          <ChainLabel
            isSwapFrom={isSwapFrom}
            chainId={chainId}
            setDisplayType={setDisplayType}
            onChangeChain={onChangeChain}
          />
        </div>
      </div>
      <div className="flex h-16 mb-4 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-20
            rounded-xl
            border border-white border-opacity-20
            ${
              isSwapFrom &&
              ' transform-gpu transition-all duration-75 hover:border-opacity-30'
            }
          `}
        >
          <SelectTokenDropdown
            chainId={chainId}
            selected={selected}
            onClick={() => {
              setDisplayType(isSwapFrom ? 'from' : 'to')
            }}
          />

          <input
            ref={inputRef}
            pattern="[0-9.]+"
            disabled={!isSwapFrom} // may cause issues idk goal is to prevent to result from being selectable
            className={`
              ml-4
              ${isSwapFrom && isConnected ? '-mt-0 md:-mt-4' : '-mt-0'}
              focus:outline-none
              bg-transparent
              pr-4
              w-2/3
             placeholder:text-[#88818C]
             text-white text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
            `}
            placeholder="0.0000"
            onChange={isSwapFrom ? onChange : () => {}}
            value={inputValue}
            name="inputRow"
          />
          {isSwapFrom && isConnected && (
            <label
              htmlFor="inputRow"
              className="absolute hidden pt-1 mt-8 ml-40 text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
              onClick={onClickBalance}
            >
              {formattedBalance}
              <span className="text-opacity-50 text-secondaryTextColor">
                {' '}
                available
              </span>
            </label>
          )}
          {isSwapFrom && isConnected && (
            <div className="hidden mr-2 sm:inline-block">
              <MiniMaxButton onClickBalance={onClickBalance} />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
