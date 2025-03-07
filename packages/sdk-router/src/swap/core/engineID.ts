export enum EngineID {
  Null,
  NoOp,
  Default,
  KyberSwap,
  ParaSwap,
  LiFi,
}

export const validateEngineID = (engineID: number): engineID is EngineID => {
  return Object.values(EngineID).includes(engineID)
}
