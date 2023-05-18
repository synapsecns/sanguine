export function mode(arr) {
  const mode = {}
  let max = 0
  let count = 0

  for (let i = 0; i < arr.length; i++) {
    const item = arr[i]

    if (mode[item]) {
      mode[item]++
    } else {
      mode[item] = 1
    }

    if (count < mode[item]) {
      max = item
      count = mode[item]
    }
  }

  return max
}
