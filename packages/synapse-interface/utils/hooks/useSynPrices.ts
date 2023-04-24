import { formatUnits } from '@ethersproject/units'
import { Zero } from '@ethersproject/constants'
import { SYN } from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
import { useEthPrice } from '@hooks/usePrices'
import { fetchBalance } from '@wagmi/core'
import * as ALL_CHAINS from '@constants/chains/master'
import { BigNumber } from 'ethers'

export const useSynPrices = async () => {
  const ethPrice: BigNumber = await useEthPrice()
  const sushiSynBalance =
    (
      await fetchBalance({
        token: `0x${SYN.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? Zero
  const sushiEthBalance =
    (
      await fetchBalance({
        token: `0x${WETH.addresses[ALL_CHAINS.ETH.id].slice(2)}`,
        chainId: ALL_CHAINS.ETH.id,
        address: `0x${SYN_ETH_SUSHI_TOKEN.addresses[ALL_CHAINS.ETH.id].slice(
          2
        )}`,
      })
    )?.value ?? Zero
  console.log('sushiSynBalance', sushiSynBalance)
  console.log('sushiEthBalance', sushiEthBalance)
  const ethBalanceNumber = Number(formatUnits(sushiEthBalance, 'ether'))
  const synBalanceNumber = Number(formatUnits(sushiSynBalance, 'ether'))

  const synPerEth = synBalanceNumber / ethBalanceNumber

  const synPrice: number = ethPrice.toNumber() / synPerEth

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice: ethPrice.toNumber(),
  }
}
