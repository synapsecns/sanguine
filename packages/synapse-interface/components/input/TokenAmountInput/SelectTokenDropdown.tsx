import { ChevronDownIcon } from '@heroicons/react/outline'
import { displaySymbol } from '@utils/displaySymbol'
import React from 'react'
import Image from 'next/image'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@styles/tokens'
import { Token } from '@/utils/types'

const SelectTokenDropdown = ({
  chainId,
  selectedToken,
  onClick,
  isOrigin,
}: {
  chainId: number
  selectedToken: Token
  onClick: () => void
  isOrigin: boolean
}) => {
  const symbol = selectedToken ? displaySymbol(chainId, selectedToken) : ''
  const dataId = isOrigin ? 'bridge-origin-token' : 'bridge-destination-token'

  return (
    <button
      className="sm:mt-[-1px] flex-shrink-0 mr-[-1px] w-[35%] cursor-pointer focus:outline-none"
      onClick={onClick}
    >
      <div
        className={`
          group rounded-xl
          -ml-2
          bg-white bg-opacity-10
        `}
      >
        <div
          className={`
            flex justify-center md:justify-start
            bg-[#49444c] bg-opacity-100
            transform-gpu transition-all duration-100
            ${getMenuItemHoverBgForCoin(selectedToken?.color)}
            border border-transparent
            ${getBorderStyleForCoinHover(selectedToken?.color)}
            items-center
            rounded-lg
            py-1.5 pl-2 h-14
          `}
        >
          <div className="self-center flex-shrink-0 hidden mr-1 sm:block">
            <div className="relative flex p-1 rounded-full">
              <Image
                alt="token image"
                className="rounded-md w-7 h-7"
                src={selectedToken?.icon}
              />
            </div>
          </div>
          <div className="text-left cursor-pointer">
            <h4 className="text-lg font-medium text-white">
              <span data-test-id={dataId}>{symbol}</span>
              <ChevronDownIcon
                className={`
                inline w-4 ml-2 -mt-1 transition-all ease-in
                transform-gpu duration-75 focus:rotate-180
                `}
              />
            </h4>
          </div>
        </div>
      </div>
    </button>
  )
}
export default SelectTokenDropdown
