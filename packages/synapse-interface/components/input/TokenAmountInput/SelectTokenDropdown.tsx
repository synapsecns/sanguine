import { ChevronDownIcon } from '@heroicons/react/outline'
import { displaySymbol } from '@utils/displaySymbol'
import React from 'react'
import Image from 'next/image'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@styles/tokens'
import { Token } from '@/utils/types'
import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID } from '@/constants/chains'
import { Chain } from '@/utils/types'

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
  const currentChain: Chain = CHAINS_BY_ID[chainId]
  const isUnsupportedChain: boolean = currentChain ? false : true
  const symbol = selectedToken ? displaySymbol(chainId, selectedToken) : ''
  const dataId = isOrigin ? 'bridge-origin-token' : 'bridge-destination-token'

  return (
    <button
      data-test-id="select-token-dropdown"
      className="flex-shrink-0 cursor-pointer focus:outline-none"
      onClick={isUnsupportedChain ? () => null : onClick}
    >
      <div
        className={`
          group rounded-md
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
            rounded-md
            py-1.5 pl-2 h-14
          `}
        >
          <div className="self-center flex-shrink-0 block mr-1">
            <div className="relative flex p-1 rounded-full">
              {selectedToken?.icon ? (
                <Image
                  alt="token image"
                  className="rounded-md w-7 h-7"
                  src={selectedToken?.icon}
                />
              ) : (
                <QuestionMarkCircleIcon className="w-6 h-6 mr-3 text-white rounded-md" />
              )}
            </div>
          </div>
          <div className="text-left cursor-pointer">
            <h4 className="w-24 text-lg font-medium text-white">
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
