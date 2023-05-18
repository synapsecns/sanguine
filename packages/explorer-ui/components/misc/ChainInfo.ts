import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { getChainUrl, getExplorerTxUrl } from '@urls'

import { CHAIN_EXPLORER_URLS, CHAIN_INFO_MAP } from '@constants/networks'
import Image from 'next/image'

export function ChainInfo({
  chainId,
  imgClassName = 'w-4 h-4 rounded-full',
  linkClassName = 'float-right text-white transition ease-out hover:text-[#8FEBFF] px-1.5  rounded-md ease-in-out bg-[#191919]',
  textClassName = "pl-1 whitespace-nowrap text-sm text-white",
  txHash,
  useExplorerLink = false,
  noLink = false,
}) {
  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId] ?? {}
  var link = ""
  if (txHash) {
    // @ts-expect-error TS(2345): Argument of type '{ hash: any; chainId: any; }' is... Remove this comment to see the full error message
    link = getExplorerTxUrl({ hash: txHash, chainId: chainId })
  }
  if (useExplorerLink) {
    link = getChainUrl({ chainId: chainId })
  }
  if (chainName) {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="w-full relative">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex justify-between ">
          // @ts-expect-error TS(2304): Cannot find name 'div'.
          <div className='flex flex-row w-[90%] items-center'>
            // @ts-expect-error TS(2749): 'Image' refers to a value, but is being used as a ... Remove this comment to see the full error message
            <Image
              // @ts-expect-error TS(2304): Cannot find name 'className'.
              className={`inline mr-[.5rem] rounded-full ${imgClassName}`}
              // @ts-expect-error TS(2304): Cannot find name 'src'.
              src={chainImg}
              // @ts-expect-error TS(2304): Cannot find name 'alt'.
              alt={chainImg}
            />
            // @ts-expect-error TS(2304): Cannot find name 'p'.
            <p className={textClassName}>{chainName}</p>
          </div>
          {noLink ? null : (
            // @ts-expect-error TS(2304): Cannot find name 'div'.
            <div className='flex items-center'>
              // @ts-expect-error TS(2304): Cannot find name 'a'.
              <a
                // @ts-expect-error TS(2304): Cannot find name 'type'.
                type="link"
                // @ts-expect-error TS(2304): Cannot find name 'target'.
                target="_blank"
                // @ts-expect-error TS(2304): Cannot find name 'href'.
                href={link}
                // @ts-expect-error TS(2304): Cannot find name 'className'.
                className={linkClassName}
              >
                ↗
              // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
              </a>
            </div>)}
        </div>
      </div>
    )
  } else {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="flex items-center">
        // @ts-expect-error TS(2749): 'QuestionMarkCircleIcon' refers to a value, but is... Remove this comment to see the full error message
        <QuestionMarkCircleIcon
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          // @ts-expect-error TS(2304): Cannot find name 'strokeWidth'.
          strokeWidth={1}
        />
        // @ts-expect-error TS(2304): Cannot find name 'span'.
        <span>--</span>
      </div>
    )
  }
}
