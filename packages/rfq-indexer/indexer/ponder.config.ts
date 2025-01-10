import { createConfig, loadBalance, rateLimit } from '@ponder/core'
import { http } from 'viem'
import dotenv from 'dotenv';

import { ABI_FastBridgeV1 } from '@/abis/FastBridgeV1'
import { ABI_FastBridgeV2 } from '@/abis/FastBridgeV2';
import { privateEncrypt } from 'crypto';


dotenv.config();

const getContractNetwork = (chainId: number, contractLabel: string) => {
  contractLabel=contractLabel.toUpperCase()
  const startBlockEnvVar = `CONTRACT_${contractLabel}_STARTBLOCK_${chainId}`;
  const addressEnvVar = `CONTRACT_${contractLabel}_ADDRESS_${chainId}`;

  const startBlock = process.env[startBlockEnvVar];
  const contractAddr = process.env[addressEnvVar];

  // if no env vars found, assume this chain+contract is legitimately not applicable & return nothing
  if (!startBlock && !contractAddr) {
    return {};
  }

  if (!startBlock) {
    throw new Error(`Environment variable ${startBlockEnvVar} must be defined`);
  }

  if (!contractAddr) {
    throw new Error(`Environment variable ${addressEnvVar} must be defined`);
  }

  console.log(`Including ${contractLabel} ${chainId.toString().padStart(10)}:${contractAddr.slice(0,6)} blocks ${startBlock.toString().padStart(10)} - <CURRENT>`);

  return {
    [chainId]: {
      contractAddr,
      startBlock: parseInt(startBlock, 10),
    }
  };
};

const getConfigNetwork = (chainId: number) => {
  const transportEnvVarName = `RPC_URL_${chainId}`;
  const url = process.env[transportEnvVarName];

  if (!url) {
    throw new Error(`RPC_URL_${chainId} must be defined`);
  }

  if (!url.startsWith('http')) {
    throw new Error(`RPC_URL_${chainId} must be an HTTP/S URL`);
  }

  // conditionally apply rate limits to RPCs
  // const requestsPerSecond = process.env[`RPC_LIMIT_RPS_${chainId}`] ? parseInt(process.env[`RPC_LIMIT_RPS_${chainId}`]!, 10) : 9999;
  // const transport = rateLimit(http(url), { requestsPerSecond });

  const transport = http(url);

  return {
    chainId,
    transport,
  };
};


// infer chain ID list from RPC_URL_ env vars that are configured
const chainIds = Object.keys(process.env)
  .filter(key => key.startsWith('RPC_URL_'))
  .map(key => parseInt(key.replace('RPC_URL_', ''), 10));

// generate contract network records
export const contractNetworks_FastBridgeV1 = chainIds.reduce((acc, chainId) => {
  return { ...acc, ...getContractNetwork(chainId, "FASTBRIDGEV1") };
}, {});

export const contractNetworks_FastBridgeV2 = chainIds.reduce((acc, chainId) => {
  return { ...acc, ...getContractNetwork(chainId, "FASTBRIDGEV2") };
}, {});

const configNetworks = chainIds.reduce((acc, chainId) => {
  return { 
    ...acc, 
    [chainId]: getConfigNetwork(chainId) 
  };
}, {});


console.log(contractNetworks_FastBridgeV2)

export default createConfig({
  networks: configNetworks,
  contracts: {
    v2: {
      abi: ABI_FastBridgeV2,
      network: contractNetworks_FastBridgeV2,
    },
    v1: {
      abi: ABI_FastBridgeV1,
      network: contractNetworks_FastBridgeV1,
    },
  },
});


