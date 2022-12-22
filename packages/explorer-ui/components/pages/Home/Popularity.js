import { useState } from 'react'

import { CHAINS_PATH, TOKEN_ADDRESSES_PATH } from '@urls'

import { getChainActivity } from '@components/SummaryDisplay/ChainActivityCard'
import { popularTokens } from '@components/SummaryDisplay/TokenActivityCard'
import { ContainerCard } from '@components/ContainerCard'
import { addressToSymbol } from '@utils/addressToSymbol'
import { AssetImage } from '@components/misc/AssetImage'
import { ChainImage } from '@components/misc/ChainImage'
import { CHAIN_INFO_MAP } from '@constants/networks'

import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'
import Image from 'next/image'

import Link from 'next/link'

import { getNetworkButtonBorderHover } from '@utils/styles/networks'

function Direction({ direction, setDirection }) {
  const activeClass =
    'text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400 hover:underline hover:cursor-pointer'
  const inactiveClass =
    'text-white text-opacity-30 hover:underline hover:cursor-pointer'

  return (
    <div className="flex space-x-2 text-lg">
      <div
        className={direction === 'IN' ? activeClass : inactiveClass}
        onClick={() => setDirection('IN')}
      >
        Sent
      </div>
      <div
        className={direction === 'OUT' ? activeClass : inactiveClass}
        onClick={() => setDirection('OUT')}
      >
        Received
      </div>
    </div>
  )
}

export function PopularTokens({ counts }) {
  const [direction, setDirection] = useState('IN')

  // const counts = popularTokens({ direction })
  const topFive = (counts && counts.slice(0, 5)) || []

  return (
    <ContainerCard
      title="Popular Tokens"
      titleClassName="text-white text-2xl mb-8"
      subtitle={<Direction setDirection={setDirection} direction={direction} />}
      className="bg-transparent border-none"
    >
      <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 5 }} gap={2}>
        {topFive.map(({ tokenAddress, chainId }, i) => {
          const { chainName } = CHAIN_INFO_MAP[chainId]
          const displaySymbol = addressToSymbol({ tokenAddress, chainId })
          return (
            <a
              href={`${TOKEN_ADDRESSES_PATH}/${tokenAddress}?chainId=${chainId}`}
              key={i}
            >
              <Card
                className={`
                p-6
                flex flex-col items-center justify-center
                bg-transparent
                border border-white border-opacity-10
              hover:bg-white hover:bg-opacity-10 hover:cursor-pointer ${getNetworkButtonBorderHover(
                chainId
              )}`}
              >
                <ActiveTokenChainIcon
                  tokenAddress={tokenAddress}
                  chainId={chainId}
                />
                <div className="mt-2 font-light text-white">
                  {displaySymbol}
                </div>
                <div className="text-sm text-white text-opacity-60">
                  on {chainName}
                </div>
              </Card>
            </a>
          )
        })}
      </Grid>
    </ContainerCard>
  )
}

export function PopularChains({ counts }) {
  const [direction, setDirection] = useState('IN')

  // const counts = getChainActivity({ direction })
  const topFive = (counts && counts.slice(0, 5)) || []

  return (
    <ContainerCard
      title="Popular Chains"
      titleClassName="text-white text-2xl mb-8"
      subtitle={<Direction setDirection={setDirection} direction={direction} />}
      className="bg-transparent border-none"
    >
      <Grid cols={{ xs: 1, sm: 1, md: 1, lg: 5 }} gap={2}>
        {topFive.map(({ chainId }, i) => {
          const { chainName, chainImg, layer } = CHAIN_INFO_MAP[chainId]
          return (
            <a href={`${CHAINS_PATH}/${chainId}`} key={i}>
              <Card
                className={`
                p-6
                flex flex-col items-center justify-center
                bg-transparent
                border border-white border-opacity-10
              hover:bg-white hover:bg-opacity-10 hover:cursor-pointer ${getNetworkButtonBorderHover(
                chainId
              )}`}
              >
                <Image
                  className="w-12 h-12 rounded-full"
                  src={chainImg}
                  alt={chainImg}
                />
                <div className="mt-3 text-white">{chainName}</div>
                <div className="text-sm text-white text-opacity-60">
                  Layer {layer}
                </div>
              </Card>
            </a>
          )
        })}
      </Grid>
    </ContainerCard>
  )
}

export function ActiveTokenChainIcon({ tokenAddress, chainId }) {
  return (
    <div className="flex mr-3">
      <AssetImage
        tokenAddress={tokenAddress}
        chainId={chainId}
        className="w-12 h-12"
      />
      <div className="z-10 -ml-6 mt-7">
        <ChainImage chainId={chainId} className="mr-0" imgSize="w-5 h-5" />
      </div>
    </div>
  )
}
