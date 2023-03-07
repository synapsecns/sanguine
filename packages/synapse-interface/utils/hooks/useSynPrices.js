import { formatUnits } from '@ethersproject/units'
import { Zero } from '@ethersproject/constants'

import { WETH, SYN } from '@constants/tokens/basic'

import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/lp'

import { ChainId } from '@constants/networks'

import { useMultipleContractSingleData } from '@hooks/multicall'
import { useEthPrice } from '@hooks/usePrices'
import { useGenericTokenContract } from '@hooks/contracts/useContract'

import { ERC20_INTERFACE } from '@constants/interfaces'



export function useSynPrices() {
  const ethPrice = useEthPrice()
  const synContract = useGenericTokenContract(ChainId.ETH, SYN)
  const wethContract = useGenericTokenContract(ChainId.ETH, WETH)

  const [sushiSynBalanceObj, sushiEthBalanceObj] = useMultipleContractSingleData(
    ChainId.ETH,
    [synContract.address, wethContract.address],
    ERC20_INTERFACE,
    'balanceOf',
    [SYN_ETH_SUSHI_TOKEN.addresses[ChainId.ETH]],
    {resultOnly: true}
  )

  const sushiSynBalance = sushiSynBalanceObj?.balance ?? Zero
  const sushiEthBalance = sushiEthBalanceObj?.balance ?? Zero

  const ethBalanceNumber     = Number(formatUnits(sushiEthBalance, 'ether'))
  const synBalanceNumber     = Number(formatUnits(sushiSynBalance, 'ether'))

  const synPerEth = synBalanceNumber / ethBalanceNumber

  const synPrice = ethPrice / synPerEth

  return {
    synBalanceNumber,
    ethBalanceNumber,
    synPrice,
    ethPrice
  }
}
