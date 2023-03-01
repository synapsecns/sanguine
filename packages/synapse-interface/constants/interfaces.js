import { Interface } from '@ethersproject/abi'

import ERC20_ABI from '@abis/erc20.json'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'




export const ERC20_INTERFACE = new Interface(ERC20_ABI)

export const SYNAPSE_BRIDGE_INTERFACE = new Interface(SYNAPSE_BRIDGE_ABI)