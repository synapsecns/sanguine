// @ts-expect-error TS(2556): A spread argument must either have a tuple type or... Remove this comment to see the full error message
export const fetcher = (...args) => fetch(...args).then((res) => res.json())
