export const STATUSES = ['Null', 'Success', 'Fail'] as const
export type StatusType = (typeof STATUSES)[number]

export enum Status {
  Null = 0,
  Success = 1,
  Fail = 2,
}

export const statusToString = (status: number): StatusType => {
  const result = Status[status]
  if (!result) {
    throw new Error(`Unknown status: ${status}`)
  }
  return result as StatusType
}
