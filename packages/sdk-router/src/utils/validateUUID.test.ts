import { validateUUID } from './validateUUID'

describe('#validateUUID', () => {
  it('returns true if valid', () => {
    expect(validateUUID('0190c847-f0f3-79cd-abbf-0f885153e5dc')).toBe(true)
  })

  it('returns false if not valid', () => {
    expect(validateUUID('test-id')).toBe(false)
    expect(validateUUID('0190c83b-d74a-7b0b-aa6e-57e1cdb2f84bcc')).toBe(false)
  })
})
