import { readContract, fetchBalance, Address, multicall } from '@wagmi/core'
import { SYN, WETH } from '@constants/tokens/bridgeable'
import * as ALL_CHAINS from '@constants/chains/master'
import {
  CHAINLINK_ETH_PRICE_ADDRESSES,
  CHAINLINK_AVAX_PRICE_ADDRESSES,
  CHAINLINK_METIS_PRICE_ADDRESSES,
  CHAINLINK_ARB_PRICE_ADDRESSES,
  CHAINLINK_GMX_PRICE_ADDRESSES,
  CHAINLINK_FRAX_PRICE_ADDRESSES,
  CHAINLINK_USDT_PRICE_ADDRESSES,
  CHAINLINK_USDC_PRICE_ADDRESSES,
  CHAINLINK_CRVUSD_PRICE_ADDRESSES,
  CHAINLINK_DAI_PRICE_ADDRESSES,
} from '@constants/chainlink'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import CHAINLINK_AGGREGATOR_ABI from '@abis/chainlinkAggregator.json'

export const getEthPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const ethPriceResult: bigint = (await readContract({
    address: CHAINLINK_ETH_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })) as bigint

  const ethPriceBigInt = ethPriceResult ?? 0n
  if (ethPriceBigInt === 0n) {
    return 0
  } else {
    // Note: BigInt to Number conversion happens here
    return Number(ethPriceBigInt) / Math.pow(10, 8)
  }
}

export const getAvaxPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const avaxPriceResult: bigint = (await readContract({
    address: CHAINLINK_AVAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: 1,
  })) as bigint

  const avaxPriceBigInt = avaxPriceResult ?? 0n

  if (avaxPriceBigInt === 0n) {
    return 0
  } else {
    return Number(avaxPriceBigInt) / Math.pow(10, 8)
  }
}

export const getMetisPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const metisPriceResult: bigint = (await readContract({
    address: CHAINLINK_METIS_PRICE_ADDRESSES[ALL_CHAINS.METIS.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: ALL_CHAINS.METIS.id,
  })) as bigint

  const metisPriceBigInt = metisPriceResult ?? 0n

  if (metisPriceBigInt === 0n) {
    return 0
  } else {
    return Number(metisPriceBigInt) / Math.pow(10, 8)
  }
}

export const getSynPrices = async () => {
  const ethPrice: number = await getEthPrice()
  const sushiSynBalance =
    (
      await fetchBalance({
        token: SYN.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id,
        address: SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id] as Address,
      })
    )?.value ?? 0n

  const sushiEthBalance =
    (
      await fetchBalance({
        token: WETH.addresses[ALL_CHAINS.ETH.id] as Address,
        chainId: ALL_CHAINS.ETH.id,
        address: SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id] as Address,
      })
    )?.value ?? 0n

  const synBalanceNumber = Number(sushiSynBalance) / Math.pow(10, 18)
  const ethBalanceNumber = Number(sushiEthBalance) / Math.pow(10, 18)
  const synPerEth = synBalanceNumber / ethBalanceNumber
  const synPrice: number = ethPrice * (1 / synPerEth)

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice,
  }
}

export const getArbPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const arbPriceResult: bigint = (await readContract({
    address: CHAINLINK_ARB_PRICE_ADDRESSES[ALL_CHAINS.ARBITRUM.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: ALL_CHAINS.ARBITRUM.id,
  })) as bigint

  const arbPriceBigInt = arbPriceResult ?? 0n

  if (arbPriceBigInt === 0n) {
    return 0
  } else {
    return Number(arbPriceBigInt) / Math.pow(10, 8)
  }
}

export const getGmxPrice = async (): Promise<number> => {
  // the price result returned by latestAnswer is 8 decimals
  const gmxPriceResult: bigint = (await readContract({
    address: CHAINLINK_GMX_PRICE_ADDRESSES[ALL_CHAINS.ARBITRUM.id] as Address,
    abi: CHAINLINK_AGGREGATOR_ABI,
    functionName: 'latestAnswer',
    chainId: ALL_CHAINS.ARBITRUM.id,
  })) as bigint

  const gmxPriceBigInt = gmxPriceResult ?? 0n

  if (gmxPriceBigInt === 0n) {
    return 0
  } else {
    return Number(gmxPriceBigInt) / Math.pow(10, 8)
  }
}

interface EthStablecoinPrices {
  usdcPrice: number
  usdtPrice: number
  daiPrice: number
  crvUsdPrice: number
  fraxPrice: number
}

export const getAllEthStablecoinPrices =
  async (): Promise<EthStablecoinPrices> => {
    try {
      const multicallInputs = [
        {
          address: CHAINLINK_USDT_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
          abi: CHAINLINK_AGGREGATOR_ABI,
          functionName: 'latestAnswer',
        },
        {
          address: CHAINLINK_USDC_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
          abi: CHAINLINK_AGGREGATOR_ABI,
          functionName: 'latestAnswer',
        },
        {
          address: CHAINLINK_DAI_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
          abi: CHAINLINK_AGGREGATOR_ABI,
          functionName: 'latestAnswer',
        },
        {
          address: CHAINLINK_CRVUSD_PRICE_ADDRESSES[
            ALL_CHAINS.ETH.id
          ] as Address,
          abi: CHAINLINK_AGGREGATOR_ABI,
          functionName: 'latestAnswer',
        },
        {
          address: CHAINLINK_FRAX_PRICE_ADDRESSES[ALL_CHAINS.ETH.id] as Address,
          abi: CHAINLINK_AGGREGATOR_ABI,
          functionName: 'latestAnswer',
        },
      ]

      const response = await multicall({
        contracts: multicallInputs as any,
        chainId: ALL_CHAINS.ETH.id,
      })

      const prices = {
        usdcPrice: convertPrice(response[0].result as bigint),
        usdtPrice: convertPrice(response[1].result as bigint),
        daiPrice: convertPrice(response[2].result as bigint),
        crvUsdPrice: convertPrice(response[3].result as bigint),
        fraxPrice: convertPrice(response[4].result as bigint),
      }

      return prices
    } catch (error) {
      console.error('Failed to fetch Ethereum stablecoin price:', error)
      return {
        usdcPrice: 1,
        usdtPrice: 1,
        daiPrice: 1,
        crvUsdPrice: 1,
        fraxPrice: 1,
      }
    }
  }

interface CoingeckoPrices {
  notePrice: number
  susdPrice: number
  lusdPrice: number
  usdbcPrice: number
  usdcePrice: number
  usdtePrice: number
}

export const getCoingeckoPrices = async (): Promise<CoingeckoPrices> => {
  try {
    const noteId = 'note'
    const susdId = 'nusd' //coingecko id for susd is nusd
    const lusdId = 'liquity-usd'
    const usdbcId = 'bridged-usd-coin-base'
    const usdceId = 'usd-coin-avalanche-bridged-usdc-e'
    const usdteId = 'tether-avalanche-bridged-usdt-e'

    const url = `https://api.coingecko.com/api/v3/simple/price?ids=${noteId},${susdId},${lusdId},${usdbcId},${usdceId},${usdteId}&vs_currencies=usd`

    const response = await fetch(url)

    if (!response.ok) {
      throw new Error(`HTTP error status: ${response.status}`)
    }

    const json = await response.json()

    const prices = {
      notePrice: json[noteId].usd,
      susdPrice: json[susdId].usd,
      lusdPrice: json[lusdId].usd,
      usdbcPrice: json[usdbcId].usd,
      usdcePrice: json[usdceId].usd,
      usdtePrice: json[usdteId].usd,
    }

    return prices
  } catch (error) {
    console.error('Failed to fetch Coingecko prices:', error)
    return {
      lusdPrice: 1,
      notePrice: 1,
      susdPrice: 1,
      usdbcPrice: 1,
      usdcePrice: 1,
      usdtePrice: 1,
    }
  }
}

export const getMusdcPrice = async (): Promise<number> => {
  try {
    const network = 'metis'
    const tokenAddress = '0xEA32A96608495e54156Ae48931A7c20f0dcc1a21'

    const url = geckoTerminalApiUrl(network, tokenAddress)

    const response = await fetch(url)

    if (!response.ok) {
      throw new Error(`HTTP error status: ${response.status}`)
    }

    const json = await response.json()
    const price = Object.values(
      json['data']['attributes']['token_prices']
    )[0] as number

    return Number(price)
  } catch (error) {
    console.error('Failed to fetch musdc price:', error)
    return 1
  }
}

export const getDaiePrice = async (): Promise<number> => {
  try {
    const network = 'avax'
    const tokenAddress = '0xd586E7F844cEa2F87f50152665BCbc2C279D8d70'

    const url = geckoTerminalApiUrl(network, tokenAddress)

    const response = await fetch(url)

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const json = await response.json()

    const price = Object.values(
      json['data']['attributes']['token_prices']
    )[0] as number

    return Number(price)
  } catch (error) {
    console.error('Failed to fetch daie price:', error)
    return 1
  }
}

const convertPrice = (price: bigint) => {
  const priceBigInt = price ?? 0n

  if (priceBigInt === 0n) {
    return 0
  } else {
    return Number(priceBigInt) / Math.pow(10, 8)
  }
}

const geckoTerminalApiUrl = (network: string, address: string) =>
  `https://api.geckoterminal.com/api/v2/simple/networks/${network}/token_price/${address}`
