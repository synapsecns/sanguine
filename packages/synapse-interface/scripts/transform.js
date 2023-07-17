let connectedTokensList = require('./results')

const tokenDirections = {}

// Create an inverted index where token symbols are keys and values are arrays of chain-token pairs where they appear
const invertedIndex = {}
for (let tokenName in connectedTokensList) {
  for (let chainName in connectedTokensList[tokenName]) {
    for (let tokenSymbol of connectedTokensList[tokenName][chainName]) {
      if (!invertedIndex[tokenSymbol]) {
        invertedIndex[tokenSymbol] = new Set()
      }
      for (let otherTokenName in connectedTokensList) {
        for (let otherChainName in connectedTokensList[otherTokenName]) {
          if (
            connectedTokensList[otherTokenName][otherChainName].includes(
              tokenSymbol
            )
          ) {
            const chainTokenPair = `${otherTokenName}-${otherChainName.toUpperCase()}`
            invertedIndex[tokenSymbol].add(chainTokenPair)
          }
        }
      }
    }
  }
}

// Use this inverted index to generate the final directions object
for (let tokenName in connectedTokensList) {
  for (let chainName in connectedTokensList[tokenName]) {
    for (let tokenSymbol of connectedTokensList[tokenName][chainName]) {
      const from = `${tokenName}-${chainName.toUpperCase()}`
      if (!tokenDirections[from]) {
        tokenDirections[from] = new Set()
      }
      for (let to of invertedIndex[tokenSymbol]) {
        // Ignore if 'to' is same as 'from'
        if (to !== from) {
          tokenDirections[from].add(to)
        }
      }
    }
  }
}

// Convert sets into arrays
for (let from in tokenDirections) {
  tokenDirections[from] = Array.from(tokenDirections[from])
}

console.log(JSON.stringify(tokenDirections, null, 2))
