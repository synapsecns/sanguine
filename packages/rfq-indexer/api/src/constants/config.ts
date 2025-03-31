type ChainTokenInfo = {
  symbol: string;
  decimals: number;
  coingeckoId: string;
};

type ChainInfo = {
  label: string;
  scannerUrl: string;
  tokens: {
    [tokenAddress: string]: ChainTokenInfo;
  };
};

export const chainConfig: { [chainId: number]: ChainInfo } = {
    1: {
      label: 'ethereum',
      scannerUrl: 'https://etherscan.io',
      tokens: {
        "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    56: {
      label: 'bnb',
      scannerUrl: 'https://bscscan.com',
      tokens: {
        "0x2170Ed0880ac9A755fd29B2688956BD959F933F8": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    10: {
      label: 'optimism',
      scannerUrl: 'https://optimistic.etherscan.io',
      tokens: {
        "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    42161: {
      label: 'arbitrum',
      scannerUrl: 'https://arbiscan.io',
      tokens: {
        "0xaf88d065e77c8cC2239327C5EDb3A432268e5831": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    8453: {
      label: 'base',
      scannerUrl: 'https://basescan.org',
      tokens: {
        "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    534352: {
      label: 'scroll',
      scannerUrl: 'https://scrollscan.com',
      tokens: {
        "0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    81457: {
      label: 'blast',
      scannerUrl: 'https://blastscan.io',
      tokens: {
        "0x4300000000000000000000000000000000000003": {
          symbol: "USDC",
          decimals: 18,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    59144: {
      label: 'linea',
      scannerUrl: 'https://lineascan.build',
      tokens: {
        "0x176211869cA2b568f2A7D4EE941E073a821EE1ff": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    480: {
      label: 'worldchain',
      scannerUrl: 'https://worldscan.org',
      tokens: {
        "0x79A02482A880bCE3F13e09Da970dC34db4CD24d1": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    130: {
      label: 'unichain',
      scannerUrl: 'https://uniscan.xyz',
      tokens: {
        "0x078D782b760474a361dDA0AF3839290b0EF57AD6": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
      },
    },
    80094: {
      label: 'berachain',
      scannerUrl: 'https://berascan.com',
      tokens: {
        "0x549943e04f40284185054145c6E4e9568C1D3241": {
          symbol: "USDC",
          decimals: 6,
          coingeckoId: "usd-coin",
        },
        "0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590": {
          symbol: "ETH",
          decimals: 18,
          coingeckoId: "ethereum",
        },
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "BERA",
          decimals: 18,
          coingeckoId: "berachain",
        },
      },
    },
    999: {
      label: 'hyperevm',
      scannerUrl: 'https://purrsec.com',
      tokens: {
        "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {
          symbol: "HYPE",
          decimals: 18,
          coingeckoId: "hyperliquid",
        },
      },
    },
  };
