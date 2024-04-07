export function shortenHash(hash: string): string | null {
  if (!hash) {
    return null
  }

  return `${hash.substring(0, 6)}...${hash.substring(hash.length - 5)}`
}
