// chunks array into chunks of maximum size
// evenly distributes items among the chunks
export function chunkArray(items, maxChunkSize) {
  if (maxChunkSize < 1) {
    throw new Error('maxChunkSize must be gte 1')
  }

  if (items.length <= maxChunkSize) {
    return [items]
  }

  const numChunks = Math.ceil(items.length / maxChunkSize)
  const chunkSize = Math.ceil(items.length / numChunks)

  return [...Array(numChunks).keys()].map((ix) => items.slice(ix * chunkSize, ix * chunkSize + chunkSize))
}
