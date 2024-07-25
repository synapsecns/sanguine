import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { getChainUrl, getExplorerTxUrl } from '@urls'
import Image from 'next/image'
import { CHAINS } from 'synapse-constants'
import Link from 'next/link'

const CHAINS_BY_ID = CHAINS.CHAINS_BY_ID

interface ChainInfoProps {
  chainId?: any
  imgClassName?: string
  linkClassName?: string
  textClassName?: string
  txHash?: string
  useExplorerLink?: boolean
  noLink?: boolean
  className?: string
}

export function ChainInfo({
  chainId,
  imgClassName = 'w-4 h-4 rounded-full',
  linkClassName = 'float-right text-white transition ease-out hover:text-[#8FEBFF] px-1.5  rounded-md ease-in-out bg-[#191919]',
  textClassName = 'pl-1 whitespace-nowrap text-sm text-white',
  txHash,
  useExplorerLink = false,
  noLink = false,
  className = '',
}: ChainInfoProps) {
  const chain = CHAINS_BY_ID[chainId]
  let link = ''

  if (txHash) {
    link = getExplorerTxUrl({ hash: txHash, chainId })
  }

  if (useExplorerLink) {
    link = getChainUrl({ chainId })
  }

  if (chain) {
    const content = (
      <>
        <Image
          className={`inline rounded-full ${imgClassName}`}
          src={chain?.chainImg}
          alt={chain?.name}
        />
        <p className={textClassName}>{chain.name}</p>
      </>
    )
    return (
      <div className="relative w-full">
        <Link href={link} passHref legacyBehavior>
          <div className="flex items-center justify-start cursor-pointer group">
            <Image
              className={`inline rounded-full ${imgClassName}`}
              src={chain?.chainImg}
              alt={chain?.name}
            />
            <p
              className={`${textClassName} group-hover:text-[#8FEBFF] transition-colors duration-200`}
            >
              {chain.name}
            </p>
          </div>
        </Link>
      </div>
    )
  } else {
    return (
      <div className={`flex items-center ${className}`}>
        <QuestionMarkCircleIcon
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          strokeWidth={1}
        />
        <span>--</span>
      </div>
    )
  }
}
