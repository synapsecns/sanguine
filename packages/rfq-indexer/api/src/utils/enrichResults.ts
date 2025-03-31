import axios from 'axios';

import { chainConfig } from '../constants/config';

export function scannerLink(chainId: number, value: string): string {
  try {
    const chainInfo = chainConfig[chainId];
    if (!chainInfo) {
      throw new Error('Chain ID not found');
    }

    const baseUrl = chainInfo.scannerUrl;
    let link = '';
    if (value.length === 42) {
      // Assume it's an address
      link = `${baseUrl}/address/${value}`;
    } else if (value.length === 66) {
      // Assume it's a transaction
      link = `${baseUrl}/tx/${value}`;
    }

    if (link) {
      return `${value} <a href="${link}" target="_blank" style="margin-left: 5px; color: white;"><i class="fas fa-external-link-alt"></i></a>`;
    }
  } catch (error: any) {
    console.error('Error in scannerLink:', error.message);
  }

  return value; // Return as-is if any error occurs
}


export async function addSenderStatus(results: any): Promise<void> {
  if (!results.every((result:any) => 'sender' in result)) {
    throw new Error('err addSenderStatus: Missing sender in one or more results');
  }

  const axiosRequests = results.map((result:any) => {
    return axios.get(`https://screener.omnirpc.io/fe/address/${result.sender}`, { timeout: 2500 })
      .then(response => {
        const { risk } = response.data;
        result.senderStatus = risk ? 'SCREENED' : 'OK';
      })
      .catch(error => {
        result.senderStatus = 'LOOKUP_FAILED';
        console.log('Error calling screener:', error.message);
      });
  });

  await Promise.all(axiosRequests);
}

export async function addTokenSymbols(results: any): Promise<void> {
  results.forEach((result: any) => {
    if ('originChainId' in result && 'originToken' in result) {
      const originChain = chainConfig[result.originChainId];
      const originTokenInfo = originChain?.tokens?.[result.originToken];
      result.originTokenSymbol = originTokenInfo?.symbol ?? 'unknown';
    }
    if ('destChainId' in result && 'destToken' in result) {
      const destChain = chainConfig[result.destChainId];
      const destTokenInfo = destChain?.tokens?.[result.destToken];
      result.destTokenSymbol = destTokenInfo?.symbol ?? 'unknown';
    }
  });
}

// TODO: add cg api key. free rate limit seems to be OK for now
export async function addUsdPricesCurrent(results: any): Promise<void> {
  const coingeckoBaseUrl = 'https://api.coingecko.com/api/v3/simple/price';
  const coingeckoIdsToFetch = new Set<string>();

  // Collect all unique coingeckoIds that need price lookup
  results.forEach((result: any) => {
    if ('originChainId' in result && 'originToken' in result) {
      const originChain = chainConfig[result.originChainId];
      const originTokenInfo = originChain?.tokens?.[result.originToken];
      if (originTokenInfo?.coingeckoId) {
        coingeckoIdsToFetch.add(originTokenInfo.coingeckoId);
      }
    }
    if ('destChainId' in result && 'destToken' in result) {
      const destChain = chainConfig[result.destChainId];
      const destTokenInfo = destChain?.tokens?.[result.destToken];
      if (destTokenInfo?.coingeckoId) {
        coingeckoIdsToFetch.add(destTokenInfo.coingeckoId);
      }
    }
  });

  // Fetch prices from CoinGecko
  const coingeckoIdsArray = Array.from(coingeckoIdsToFetch);
  const coingeckoIdsQuery = coingeckoIdsArray.join(',');
  const url = `${coingeckoBaseUrl}?ids=${coingeckoIdsQuery}&vs_currencies=usd`;

  try {
    const response = await axios.get(url);
    const prices = response.data;

    // Add USD prices and calculate UsdValue
    results.forEach((result: any) => {
      if ('originChainId' in result && 'originToken' in result) {
        const originChain = chainConfig[result.originChainId];
        const originTokenInfo = originChain?.tokens?.[result.originToken];
        const originCoingeckoId = originTokenInfo?.coingeckoId;
        const originPrice = originCoingeckoId ? prices[originCoingeckoId]?.usd : undefined;
        if (originPrice !== undefined) {
          result.originTokenPriceUsd = originPrice;
          if (result.originAmountFormatted) {
            result.originUsdValue = result.originAmountFormatted * originPrice;
          }
        } else {
          result.originTokenPriceUsd = null;
          result.originUsdValue = null;
        }
      }
      if ('destChainId' in result && 'destToken' in result) {
        const destChain = chainConfig[result.destChainId];
        const destTokenInfo = destChain?.tokens?.[result.destToken];
        const destCoingeckoId = destTokenInfo?.coingeckoId;
        const destPrice = destCoingeckoId ? prices[destCoingeckoId]?.usd : undefined;
        if (destPrice !== undefined) {
          result.destTokenPriceUsd = destPrice;
          if (result.destAmountFormatted) {
            result.destUsdValue = result.destAmountFormatted * destPrice;
          }
        } else {
          result.destTokenPriceUsd = null;
          result.destUsdValue = null;
        }
      }
    });
  } catch (error: any) {
    console.log('Error fetching prices from CoinGecko:', error.message);
    results.forEach((result: any) => {
      result.originTokenPriceUsd = null;
      result.originUsdValue = null;
      result.destTokenPriceUsd = null;
      result.destUsdValue = null;
    });
  }
}
