export function ellipsizeString({ string, limiter = 6 }) {
  if (limiter === 0) {
    return string
  } else {
    return (
      string &&
      `${string.toLowerCase().slice(0, limiter)}...${string
        .toLowerCase()
        .slice(-limiter + 2, string.length)}`
    )
  }
}
