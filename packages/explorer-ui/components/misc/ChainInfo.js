import {QuestionMarkCircleIcon} from '@heroicons/react/outline'

import {CHAIN_EXPLORER_URLS, CHAIN_INFO_MAP} from '@constants/networks'
import {getNetworkTextColor} from '@styles/networks'
import Image from 'next/image'

export function ChainInfo({
  chainId,
  imgClassName = 'w-4 h-4',
  textClassName = getNetworkTextColor(chainId),
  txHash,
}) {
  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId] ?? {}
  const explorer = CHAIN_EXPLORER_URLS[chainId] ?? ''

  if (chainName) {
    return (
      <div className="flex items-center">
        <Image
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          src={chainImg}
          alt={chainImg}
        />
        <span className={textClassName}>{chainName}</span>
        <a
          target="_blank"
          href={explorer + '/tx/' + txHash}
          className="bg-gray-700 p-1 rounded-md ml-1"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="#8FEBFF"
            className="w-4 h-4"
          >
            <path
              fillRule="evenodd"
              d="M15.75 2.25H21a.75.75 0 01.75.75v5.25a.75.75 0 01-1.5 0V4.81L8.03 17.03a.75.75 0 01-1.06-1.06L19.19 3.75h-3.44a.75.75 0 010-1.5zm-10.5 4.5a1.5 1.5 0 00-1.5 1.5v10.5a1.5 1.5 0 001.5 1.5h10.5a1.5 1.5 0 001.5-1.5V10.5a.75.75 0 011.5 0v8.25a3 3 0 01-3 3H5.25a3 3 0 01-3-3V8.25a3 3 0 013-3h8.25a.75.75 0 010 1.5H5.25z"
              clipRule="evenodd"
            />
          </svg>
        </a>
      </div>
    )
  } else {
    return (
      <div className="flex items-center">
        <QuestionMarkCircleIcon
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          strokeWidth={1}
        />
        <span>--</span>
      </div>
    )
  }
}
