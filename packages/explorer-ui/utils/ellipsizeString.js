export function ellipsizeString({ string, limiter = 4, isZeroX = false }) {
  if (limiter === 0) {
    return string
  } else {
    if (isZeroX) {
      return (
        string &&
        `${string.toLowerCase().slice(0, limiter + 2)}...${string
          .toLowerCase()
          .slice(-limiter, string.length)}`
      )
    } else {
      return (
        string &&
        `${string.toLowerCase().slice(0, limiter)}...${string
          .toLowerCase()
          .slice(-limiter, string.length)}`
      )
    }
  }
}
