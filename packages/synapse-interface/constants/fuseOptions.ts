export const DEFAULT_FUSE_OPTIONS = {
  includeScore: true,
  threshold: 0.1
}


export const CHAIN_FUSE_OPTIONS = {
  ...DEFAULT_FUSE_OPTIONS,
  keys: [
    {
      name: 'name',
      weight: 2,
    },
    'id',
    'nativeCurrency.symbol',
  ],
}

export const TOKEN_FUSE_OPTIONS = {
    ...DEFAULT_FUSE_OPTIONS,
    ignoreLocation: true,
}

export function getTokenFuseOptions(chainId: number) {
  return {
    ...TOKEN_FUSE_OPTIONS,
    keys: [
      {
        name: 'symbol',
        weight: 2,
      },
      'routeSymbol',
      `addresses.${chainId}`,
      'name',
    ],
  }
}