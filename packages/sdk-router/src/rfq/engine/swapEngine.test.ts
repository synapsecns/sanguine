import { EngineID, validateEngineID } from './swapEngine'

describe('EngineID', () => {
  it('validates values within enum', () => {
    expect(validateEngineID(EngineID.Null)).toBe(true)
    expect(validateEngineID(EngineID.NoOp)).toBe(true)
    expect(validateEngineID(EngineID.Default)).toBe(true)
  })

  it('does not validate values outside enum', () => {
    expect(validateEngineID(EngineID.Default + 1)).toBe(false)
    expect(validateEngineID(10)).toBe(false)
    expect(validateEngineID(1000)).toBe(false)
  })
})
