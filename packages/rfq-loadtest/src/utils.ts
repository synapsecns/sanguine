export function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export function tStamp(startTimeStamp = 0) {
  const timeCur = new Date().toISOString().replace('T', ' ').replace('Z', '')

  const timeDiff =
    startTimeStamp > 0
      ? ` +${(Date.now() - startTimeStamp).toString().padStart(5)}ms`
      : ''

  return `${timeCur}${timeDiff} - `
}

export function print(...outputs: any[]) {
  outputs = outputs.map((output: any) => {
    if (typeof output == 'string') {
      // Replace %ts with formatted timestamp
      output = output.replaceAll('%ts', tStamp())
    }
    return output
  })

  console.log(...outputs)
}

export function getRandomInt(min: number, max: number) {
  if (min > max) {
    // fix mistake inputs
    ;[min, max] = [max, min]
  }
  return Math.floor(Math.random() * (max - min + 1)) + min
}
