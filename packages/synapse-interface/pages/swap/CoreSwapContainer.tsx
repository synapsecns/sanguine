import _ from 'lodash'

import { Zero } from '@ethersproject/constants'

import { formatBNToString } from '@bignumber/format'

import { ChevronDownIcon } from '@heroicons/react/outline'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'

import { useTokenBalance } from '@hooks/tokens/useTokenBalances'

import {
  getMenuItemHoverBgForCoin,
  getBorderStyleForCoinHover,
} from '@styles/coins'

import MiniMaxButton from '@components/buttons/MiniMaxButton'
import SwitchButton from '@components/buttons/SwitchButton'
import { ETH, WETH } from '@constants/tokens/basic'
import { displaySymbol } from '@utils/displaySymbol'

import { useNetworkController } from '@hooks/wallet/useNetworkController'

import { cleanNumberInput } from '@utils/cleanNumberInput'

export default function CoreSwapContainer({
  selected,
  inputValue,
  isSwapFrom,
  onChangeAmount,
  swapFromToCoins,
  setDisplayType,
  inputRef,
}) {
  const { chainId } = useActiveWeb3React()
  const { terraAddress, account } = useNetworkController()
  let isConnected = account || terraAddress

  let balanceCoin
  if (selected.symbol == WETH.symbol) {
    balanceCoin = ETH
  } else {
    balanceCoin = selected
  }
  const tokenBalance = useTokenBalance(balanceCoin) ?? Zero
  const formattedBalance = formatBNToString(
    tokenBalance,
    selected.decimals[chainId],
    4
  )

  function onChange(e) {
    let val = e.target.value

    let cleanValue = cleanNumberInput(val)

    onChangeAmount(cleanValue)
  }

  function onClickBalance() {
    onChangeAmount(formatBNToString(tokenBalance, selected.decimals[chainId]))
  }

  const symbol = displaySymbol(chainId, balanceCoin)

  return (
    <div className="pt-2">
      <div>
        {!isSwapFrom && (
          <div className="absolute mt-1 ml-2">
            <div className="-mt-8">
              <SwitchButton selected={selected} onClick={swapFromToCoins} />
            </div>
          </div>
        )}
      </div>
      <div className="p-3 border-none bg-bgLight rounded-xl">
        <div className="flex space-x-2">
          <div
            className={`
            flex flex-grow items-center
            pl-4 md:pl-2
            w-full h-20
            rounded-xl
            border border-white border-opacity-20
            hover:border-opacity-30
          `}
          >
            <button
              className="sm:mt-[-1px] flex-shrink-0 mr-[-1px] w-[35%]"
              onClick={() => {
                setDisplayType(isSwapFrom ? 'from' : 'to')
              }}
            >
              <div
                className={`
                  group rounded-xl 
                  border border-transparent
                  transform-gpu transition-all duration-125
                  ${getMenuItemHoverBgForCoin(selected)} 
                  ${getBorderStyleForCoinHover(selected)}
                `}
              >
                <div className="flex justify-center md:justify-start bg-white bg-opacity-10 items-center rounded-lg py-1.5 pl-2 cursor-pointer h-14">
                  <div className="self-center flex-shrink-0 hidden mr-1 sm:block">
                    <div className="relative flex p-1 rounded-full">
                      <img className="w-7 h-7" src={balanceCoin.icon} />
                    </div>
                  </div>
                  <div className="text-left cursor-pointer">
                    <h4 className="text-lg font-medium text-gray-300 ">
                      <span>{symbol}</span>
                      <ChevronDownIcon className="inline w-4 ml-2 -mt-1 transition-all transform focus:rotate-180" />
                    </h4>
                  </div>
                </div>
              </div>
            </button>
            <div
              className={`
                flex flex-grow items-center
                w-full h-16 border-none 
              `}
            >
              <input
                ref={inputRef}
                pattern="[0-9.]+"
                className={`
                  ml-4
                  ${isSwapFrom && isConnected ? '-mt-0 md:-mt-4' : '-mt-0'}
                  focus:outline-none
                  bg-transparent
                  pr-4
                  w-5/6
                placeholder:text-[#88818C] 
                  text-white text-opacity-80 text-lg md:text-2xl lg:text-2xl font-medium
                `}
                placeholder="0.0000"
                onChange={onChange}
                value={inputValue}
                label="inputRow"
              />
              {isSwapFrom && isConnected && (
                <label
                  htmlFor="inputRow"
                  className="absolute hidden pt-1 mt-8 ml-4 text-xs text-white transition-all duration-150 md:block transform-gpu hover:text-opacity-70 hover:cursor-pointer"
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
                  <MiniMaxButton
                    tokenBalance={tokenBalance}
                    formattedBalance={formattedBalance}
                    inputValue={inputValue}
                    onClickBalance={onClickBalance}
                    selected={selected}
                  />
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
