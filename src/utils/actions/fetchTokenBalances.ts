import { Contract, ethers, AbiCoder } from 'ethers'
import { BridgeableToken } from 'types'
import multicallAbi from '../../constants/abis/multicall.json'
import erc20Abi from '../../constants/abis/erc20.json'
import { formatBigIntToString } from '../formatBigIntToString'

const multicallAddress: string = `0xcA11bde05977b3631167028862bE2a173976CA11`

function useMulticallContract(signerOrProvider: any) {
  return new Contract(multicallAddress, multicallAbi, signerOrProvider)
}

export interface TokenBalance {
  token: BridgeableToken
  balance: bigint
  parsedBalance: string
}

export async function fetchTokenBalances({
  address,
  chainId,
  tokens,
  signer,
}: {
  address: string
  chainId: number
  tokens: any[]
  signer: any
}): Promise<TokenBalance[]> {
  const multicall = useMulticallContract(signer)

  const calls = tokens.map((token: BridgeableToken) => {
    const tokenAddress: string = token.addresses[chainId]
    const tokenContract = new ethers.Contract(tokenAddress, erc20Abi, signer)
    return {
      target: tokenAddress,
      callData: tokenContract.interface.encodeFunctionData('balanceOf', [
        address,
      ]),
    }
  })

  try {
    const [, response] = await multicall.aggregate(calls)
    const coder = AbiCoder.defaultAbiCoder()

    const data = response.map((encodedBalance, index) => {
      const balance: bigint = coder.decode(['uint256'], encodedBalance)[0]
      const token: BridgeableToken = tokens[index]
      const decimals: number =
        typeof token.decimals === 'number'
          ? token.decimals
          : token.decimals[chainId]
      return {
        token: tokens[index],
        balance: String(balance),
        parsedBalance: formatBigIntToString(balance, decimals, 4),
      }
    })

    console.log('multicall data:', data)
    return data
  } catch (error) {
    console.error('Error fetching token balances:', error)
    return error
  }
}
