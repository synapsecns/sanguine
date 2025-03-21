export enum EngineID {
  Null,
  NoOp,
  DefaultPools,
  KyberSwap,
  ParaSwap,
  LiFi,
}

export const validateEngineID = (engineID: number): engineID is EngineID => {
  return Object.values(EngineID).includes(engineID)
}
