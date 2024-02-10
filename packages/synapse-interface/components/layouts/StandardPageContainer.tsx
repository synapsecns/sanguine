import { useEffect, useState } from 'react'
import toast from 'react-hot-toast'
import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkTextColor } from '@styles/chains'
import { Chain } from '@/utils/types/index'
import { AcceptedChainId } from '@constants/chains'

const StandardPageContainer = ({
  title,
  subtitle,
  children,
  rightContent,
  address,
  connectedChainId,
}: {
  title?: string
  subtitle?: string
  children: any
  rightContent?: any
  address: string
  connectedChainId: number
}) => {

  useEffect(() => {
    if (!connectedChainId) return
    const chain = CHAINS_BY_ID[connectedChainId]
    const unsupported = AcceptedChainId[connectedChainId] ? false : true
    let unsupportedToaster

    if (unsupported) {
      unsupportedToaster = toast.error(
        `Connected to an unsupported network. Please switch networks.`,
        { id: 'unsupported-popup', duration: 5000 }
      )
    } else {
      toast(
        <>
          Connected to the{' '}
          <p className={getNetworkTextColor(chain?.color) + ' mx-1'}>
            {' '}
            {chain?.altName ?? chain?.name}{' '}
          </p>
          chain
        </>,
        { id: 'standard-popup', duration: 5000 }
      )
    }
  }, [connectedChainId])

  return (
    <main className="relative z-0 flex-1 h-full overflow-y-auto focus:outline-none">
      <div className="items-center px-4 py-8 mx-auto mt-4 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
        <span
          className={`
            text-3xl font-medium text-default
            bg-clip-text text-transparent bg-gradient-to-r
            from-purple-600 to-blue-600
          `}
        >
          {title}
        </span>
        {rightContent}
        <div className="mt-1 text-sm font-medium text-gray-600">
          {subtitle ?? ''}
        </div>
        {children}
      </div>
    </main>
  )
}
export default StandardPageContainer
