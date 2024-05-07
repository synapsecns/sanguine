import { Contract, ethers, AbiCoder, ZeroAddress } from 'ethers'
import { BridgeableToken } from 'types'

import { formatBigIntToString } from '../formatBigIntToString'
import multicallAbi from '../../constants/abis/multicall.json'
import erc20Abi from '../../constants/abis/erc20.json'

const multicallAddress: string = `0xcA11bde05977b3631167028862bE2a173976CA11`

const useMulticallContract = (signerOrProvider: any): Contract => {
  return new Contract(multicallAddress, multicallAbi, signerOrProvider)
}

export interface TokenBalance {
  token: BridgeableToken
  balance: bigint
  parsedBalance: string
}

export const fetchTokenBalances = async ({
  address,
  chainId,
  tokens,
  signerOrProvider,
}: {
  address: string
  chainId: number
  tokens: BridgeableToken[]
  signerOrProvider: any
}): Promise<TokenBalance[]> => {
  const multicall: Contract = useMulticallContract(signerOrProvider)

  if (!signerOrProvider) {
    console.error(
      '[Synapse Widget] Error fetching token balances: Require valid Signer or Provider'
    )
    return
  }
  if (Number(signerOrProvider?._network.chainId.toString()) !== chainId) {
    console.error(
      '[Synapse Widget] Error fetching token balances: Signer or Provider does not match selected chainId'
    )
    return
  }

  const validTokens = tokens.filter((token: BridgeableToken) => {
    const tokenAddress: string = token && token.addresses[chainId]
    return token && tokenAddress !== undefined
  })

  const calls = validTokens.map((token: BridgeableToken) => {
    const tokenAddress: string = token.addresses[chainId]

    if (tokenAddress === ZeroAddress) {
      const tokenContract = new ethers.Contract(
        tokenAddress,
        multicallAbi,
        signerOrProvider
      )

      return {
        target: multicallAddress,
        callData: tokenContract.interface.encodeFunctionData('getEthBalance', [
          address,
        ]),
      }
    } else {
      const tokenContract = new ethers.Contract(
        tokenAddress,
        erc20Abi,
        signerOrProvider
      )

      return {
        target: tokenAddress,
        callData: tokenContract.interface.encodeFunctionData('balanceOf', [
          address,
        ]),
      }
    }
  })

  try {
    const [, response] = await multicall.aggregate(calls)
    const coder = AbiCoder.defaultAbiCoder()

    const data = response.map((encodedBalance, index) => {
      const balance: bigint = coder.decode(['uint256'], encodedBalance)[0]
      const token: BridgeableToken = validTokens[index]
      const decimals: number = token.decimals[chainId]

      return {
        token: validTokens[index],
        balance: String(balance),
        parsedBalance: formatBigIntToString(balance, decimals, 4),
      }
    })

    return data
  } catch (error) {
    console.error('[Synapse Widget] Error fetching token balances: ', error)
    return error
  }
}
