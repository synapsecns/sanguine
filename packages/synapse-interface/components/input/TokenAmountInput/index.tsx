import { formatBNToString } from '@bignumber/format'
import React, { useMemo } from 'react'
import SwitchButton from '@components/buttons/SwitchButton'
import MiniMaxButton from '@components/buttons/MiniMaxButton'
import Spinner from '@/components/icons/Spinner'
import { BigNumber } from '@ethersproject/bignumber'
import { cleanNumberInput } from '@utils/cleanNumberInput'

import { Token } from '@/utils/types'
import { ChainLabel } from '@components/ChainLabel'
import { DisplayType } from '@/pages/bridge/BridgeCard'
import SelectTokenDropdown from './SelectTokenDropdown'

const BridgeInputContainer = ({
  address,
  isOrigin,
  isSwap,
  chains,
  chainId,
  inputString,
  selectedToken,
  connectedChainId,
  onChangeChain,
  onChangeAmount,
  setDisplayType,
  fromTokenBalance,
  isQuoteLoading = false,
}: {
  address: `0x${string}`
  isOrigin: boolean
  isSwap: boolean
  chains: string[]
  chainId: number
  inputString: string
  selectedToken: Token
  connectedChainId: number
  setDisplayType: (v: DisplayType) => void
  onChangeAmount?: (v: string) => void
  onChangeChain: (chainId: number, flip: boolean, type: 'from' | 'to') => void
  fromTokenBalance?: BigNumber
  isQuoteLoading?: boolean
}) => {
  const formattedBalance = useMemo(() => {
    if (!fromTokenBalance) return '0.0'
    return formatBNToString(
      fromTokenBalance,
      selectedToken?.decimals[chainId as keyof Token['decimals']],
      3
    )
  }, [fromTokenBalance])

  const isConnected = address !== null
  const isMaxDisabled = formattedBalance === '0.0'

  const onClickBalance = () => {
    onChangeAmount(
      formatBNToString(
        fromTokenBalance,
        selectedToken?.decimals[chainId as keyof Token['decimals']]
      )
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
          {!isOrigin && !isSwap && (
            <div className="absolute">
              <div className="-mt-12">
                <SwitchButton
                  onClick={() =>
                    onChangeChain(chainId, true, isOrigin ? 'from' : 'to')
                  }
                />
              </div>
            </div>
          )}
          {!(isSwap && !isOrigin) && (
            <ChainLabel
              isOrigin={isOrigin}
              chainId={chainId}
              setDisplayType={setDisplayType}
              onChangeChain={onChangeChain}
              chains={chains}
              connectedChainId={connectedChainId}
            />
          )}
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
              isOrigin &&
              ' transform-gpu transition-all duration-75 hover:border-opacity-30'
            }
          `}
        >
          <SelectTokenDropdown
            chainId={chainId}
            selectedToken={selectedToken}
            isOrigin={isOrigin}
            onClick={() => {
              setDisplayType(isOrigin ? DisplayType.FROM : DisplayType.TO)
            }}
          />
          <div className={isOrigin && 'md:mt-[10px]'}>
            <input
              pattern="[0-9.]+"
              disabled={!isOrigin} // may cause issues idk goal is to prevent to result from being selectable
              className={`
                ml-4 pr-4 w-5/6
                ${isOrigin && isConnected ? '-mt-0 md:-mt-4' : '-mt-0'}
                focus:outline-none bg-transparent
               placeholder:text-[#88818C]
               text-white text-opacity-80 text-lg font-medium
                md:w-2/3 md:text-2xl lg:text-2xl
              `}
              placeholder="0.0000"
              onChange={
                isOrigin
                  ? (e) => onChangeAmount(cleanNumberInput(e.target.value))
                  : () => null
              }
              value={inputString === '0' ? null : inputString}
              name="inputRow"
              autoComplete="off"
            />
            {isOrigin && isConnected && (
              <label
                htmlFor="inputRow"
                className="hidden ml-4 text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
                onClick={onClickBalance}
              >
                {formattedBalance}
                <span className="text-opacity-50 text-secondaryTextColor">
                  {' '}
                  available
                </span>
              </label>
            )}
          </div>
          {isOrigin && isConnected && (
            <div className="sm:mr-2 sm:inline-block md:hidden">
              <MiniMaxButton
                disabled={isMaxDisabled}
                onClickBalance={onClickBalance}
              />
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default BridgeInputContainer
