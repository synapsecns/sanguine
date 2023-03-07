export function reduceTokensIntoDict(arr) {
  return (
    arr.reduce(
      (acc, token) => ({ ...acc, [token.symbol]: token }),
      {}
    )
  )

}


export function reduceNestedTokensIntoDict(nestedTokensByChainId) {
  let obj = {}
  for (const [chainId, arr] of Object.entries(nestedTokensByChainId)) {
    obj[chainId] = reduceTokensIntoDict(arr)
  }
  return obj
}