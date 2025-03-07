import { EngineID, validateEngineID } from '../engineID'

describe('EngineID', () => {
  it('validates values within enum', () => {
    expect(validateEngineID(EngineID.Null)).toBe(true)
    expect(validateEngineID(EngineID.NoOp)).toBe(true)
    expect(validateEngineID(EngineID.Default)).toBe(true)
    expect(validateEngineID(EngineID.KyberSwap)).toBe(true)
    expect(validateEngineID(EngineID.ParaSwap)).toBe(true)
    expect(validateEngineID(EngineID.LiFi)).toBe(true)
  })

  it('does not validate values outside enum', () => {
    expect(validateEngineID(EngineID.LiFi + 1)).toBe(false)
    expect(validateEngineID(10)).toBe(false)
    expect(validateEngineID(1000)).toBe(false)
  })
})
