import { ethers, Contract } from 'ethers'

// ERC20 Contract ABI for "Approve" and "Allowance" functions
const erc20_abi = [
  {
    name: 'approve',
    inputs: [
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    outputs: [
      {
        internalType: 'bool',
        name: '',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    name: 'allowance',
    inputs: [
      {
        internalType: 'address',
        name: 'owner',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
    ],
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
]

// Fetch the current allowance and update if needed
export const checkAndSetAllowance = async (
  wallet,
  tokenAddress,
  approvalAddress,
  amount
) => {
  // Transactions with the native token don't need approval
  if (tokenAddress === ethers.constants.AddressZero) {
    return
  }

  const erc20 = new Contract(tokenAddress, erc20_abi, wallet)
  const allowance = await erc20.allowance(
    await wallet.getAddress(),
    approvalAddress
  )
  if (allowance.lt(amount)) {
    const approveTx = await erc20.approve(approvalAddress, amount, {
      gasPrice: await wallet.provider.getGasPrice(),
    })
    try {
      await approveTx.wait()
      console.log(`Transaction mined succesfully: ${approveTx.hash}`)
    } catch (error) {
      console.log(`Transaction failed with error: ${error}`)
    }
  }
}

// export const main = async () => {

//     // setting up a signer
//     const provider = new ethers.providers.JsonRpcProvider("https://polygon-rpc.com", 137);
//     // use provider.getSigner() method to get a signer if you're using this for a UI
//     const wallet = new ethers.Wallet("YOUR_PRIVATE_KEY", provider)

//     await checkAndSetAllowance(
//         wallet,
//         '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174', // fromTokenAddress (USDC on Polygon)
//         '0x6e14f48576265272B6CAA3A7cC500a26050Be64E', // addressToApprove (Router's Reserve Token Handler on Polygon)
//         ethers.constants.MaxUint256 // amount to approve (infinite approval)
//     );
// }
