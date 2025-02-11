/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from 'ethers'
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from '@ethersproject/abi'
import type { Listener, Provider } from '@ethersproject/providers'
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
} from './common'

export type PoolTokenStruct = { isWeth: boolean; token: string }

export type PoolTokenStructOutput = [boolean, string] & {
  isWeth: boolean
  token: string
}

export type PoolStruct = {
  pool: string
  lpToken: string
  tokens: PoolTokenStruct[]
}

export type PoolStructOutput = [string, string, PoolTokenStructOutput[]] & {
  pool: string
  lpToken: string
  tokens: PoolTokenStructOutput[]
}

export type SwapQueryStruct = {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumberish
  deadline: BigNumberish
  rawParams: BytesLike
}

export type SwapQueryStructOutput = [
  string,
  string,
  BigNumber,
  BigNumber,
  string
] & {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export type BridgeTokenStruct = { symbol: string; token: string }

export type BridgeTokenStructOutput = [string, string] & {
  symbol: string
  token: string
}

export type DestRequestStruct = { symbol: string; amountIn: BigNumberish }

export type DestRequestStructOutput = [string, BigNumber] & {
  symbol: string
  amountIn: BigNumber
}

export declare namespace LocalBridgeConfig {
  export type BridgeTokenConfigStruct = {
    id: string
    token: string
    decimals: BigNumberish
    tokenType: BigNumberish
    bridgeToken: string
    bridgeFee: BigNumberish
    minFee: BigNumberish
    maxFee: BigNumberish
  }

  export type BridgeTokenConfigStructOutput = [
    string,
    string,
    BigNumber,
    number,
    string,
    BigNumber,
    BigNumber,
    BigNumber
  ] & {
    id: string
    token: string
    decimals: BigNumber
    tokenType: number
    bridgeToken: string
    bridgeFee: BigNumber
    minFee: BigNumber
    maxFee: BigNumber
  }
}

export declare namespace MulticallView {
  export type ResultStruct = { success: boolean; returnData: BytesLike }

  export type ResultStructOutput = [boolean, string] & {
    success: boolean
    returnData: string
  }
}

export interface SynapseRouterInterface extends utils.Interface {
  functions: {
    'adapterSwap(address,address,uint256,address,bytes)': FunctionFragment
    'addToken(string,address,uint8,address,uint256,uint256,uint256)': FunctionFragment
    'addTokens((string,address,uint256,uint8,address,uint256,uint256,uint256)[])': FunctionFragment
    'allPools()': FunctionFragment
    'bridge(address,uint256,address,uint256,(address,address,uint256,uint256,bytes),(address,address,uint256,uint256,bytes))': FunctionFragment
    'bridgeTokens()': FunctionFragment
    'bridgeTokensAmount()': FunctionFragment
    'calculateAddLiquidity(address,uint256[])': FunctionFragment
    'calculateBridgeFee(address,uint256)': FunctionFragment
    'calculateRemoveLiquidity(address,uint256)': FunctionFragment
    'calculateSwap(address,uint8,uint8,uint256)': FunctionFragment
    'calculateWithdrawOneToken(address,uint256,uint8)': FunctionFragment
    'config(address)': FunctionFragment
    'fee(address)': FunctionFragment
    'getAmountOut(address,address,uint256)': FunctionFragment
    'getConnectedBridgeTokens(address)': FunctionFragment
    'getDestinationAmountOut((string,uint256)[],address)': FunctionFragment
    'getOriginAmountOut(address,string[],uint256)': FunctionFragment
    'multicallView(bytes[])': FunctionFragment
    'owner()': FunctionFragment
    'poolInfo(address)': FunctionFragment
    'poolTokens(address)': FunctionFragment
    'poolsAmount()': FunctionFragment
    'removeToken(address)': FunctionFragment
    'removeTokens(address[])': FunctionFragment
    'renounceOwnership()': FunctionFragment
    'setAllowance(address,address,uint256)': FunctionFragment
    'setSwapQuoter(address)': FunctionFragment
    'setTokenConfig(address,uint8,address)': FunctionFragment
    'setTokenFee(address,uint256,uint256,uint256)': FunctionFragment
    'swap(address,address,uint256,(address,address,uint256,uint256,bytes))': FunctionFragment
    'swapQuoter()': FunctionFragment
    'symbolToToken(string)': FunctionFragment
    'synapseBridge()': FunctionFragment
    'tokenToSymbol(address)': FunctionFragment
    'transferOwnership(address)': FunctionFragment
  }

  getFunction(
    nameOrSignatureOrTopic:
      | 'adapterSwap'
      | 'addToken'
      | 'addTokens'
      | 'allPools'
      | 'bridge'
      | 'bridgeTokens'
      | 'bridgeTokensAmount'
      | 'calculateAddLiquidity'
      | 'calculateBridgeFee'
      | 'calculateRemoveLiquidity'
      | 'calculateSwap'
      | 'calculateWithdrawOneToken'
      | 'config'
      | 'fee'
      | 'getAmountOut'
      | 'getConnectedBridgeTokens'
      | 'getDestinationAmountOut'
      | 'getOriginAmountOut'
      | 'multicallView'
      | 'owner'
      | 'poolInfo'
      | 'poolTokens'
      | 'poolsAmount'
      | 'removeToken'
      | 'removeTokens'
      | 'renounceOwnership'
      | 'setAllowance'
      | 'setSwapQuoter'
      | 'setTokenConfig'
      | 'setTokenFee'
      | 'swap'
      | 'swapQuoter'
      | 'symbolToToken'
      | 'synapseBridge'
      | 'tokenToSymbol'
      | 'transferOwnership'
  ): FunctionFragment

  encodeFunctionData(
    functionFragment: 'adapterSwap',
    values: [string, string, BigNumberish, string, BytesLike]
  ): string
  encodeFunctionData(
    functionFragment: 'addToken',
    values: [
      string,
      string,
      BigNumberish,
      string,
      BigNumberish,
      BigNumberish,
      BigNumberish
    ]
  ): string
  encodeFunctionData(
    functionFragment: 'addTokens',
    values: [LocalBridgeConfig.BridgeTokenConfigStruct[]]
  ): string
  encodeFunctionData(functionFragment: 'allPools', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'bridge',
    values: [
      string,
      BigNumberish,
      string,
      BigNumberish,
      SwapQueryStruct,
      SwapQueryStruct
    ]
  ): string
  encodeFunctionData(
    functionFragment: 'bridgeTokens',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'bridgeTokensAmount',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'calculateAddLiquidity',
    values: [string, BigNumberish[]]
  ): string
  encodeFunctionData(
    functionFragment: 'calculateBridgeFee',
    values: [string, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'calculateRemoveLiquidity',
    values: [string, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'calculateSwap',
    values: [string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'calculateWithdrawOneToken',
    values: [string, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(functionFragment: 'config', values: [string]): string
  encodeFunctionData(functionFragment: 'fee', values: [string]): string
  encodeFunctionData(
    functionFragment: 'getAmountOut',
    values: [string, string, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'getConnectedBridgeTokens',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'getDestinationAmountOut',
    values: [DestRequestStruct[], string]
  ): string
  encodeFunctionData(
    functionFragment: 'getOriginAmountOut',
    values: [string, string[], BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'multicallView',
    values: [BytesLike[]]
  ): string
  encodeFunctionData(functionFragment: 'owner', values?: undefined): string
  encodeFunctionData(functionFragment: 'poolInfo', values: [string]): string
  encodeFunctionData(functionFragment: 'poolTokens', values: [string]): string
  encodeFunctionData(
    functionFragment: 'poolsAmount',
    values?: undefined
  ): string
  encodeFunctionData(functionFragment: 'removeToken', values: [string]): string
  encodeFunctionData(
    functionFragment: 'removeTokens',
    values: [string[]]
  ): string
  encodeFunctionData(
    functionFragment: 'renounceOwnership',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'setAllowance',
    values: [string, string, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'setSwapQuoter',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'setTokenConfig',
    values: [string, BigNumberish, string]
  ): string
  encodeFunctionData(
    functionFragment: 'setTokenFee',
    values: [string, BigNumberish, BigNumberish, BigNumberish]
  ): string
  encodeFunctionData(
    functionFragment: 'swap',
    values: [string, string, BigNumberish, SwapQueryStruct]
  ): string
  encodeFunctionData(functionFragment: 'swapQuoter', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'symbolToToken',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'synapseBridge',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'tokenToSymbol',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'transferOwnership',
    values: [string]
  ): string

  decodeFunctionResult(functionFragment: 'adapterSwap', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'addToken', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'addTokens', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'allPools', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'bridge', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'bridgeTokens',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'bridgeTokensAmount',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'calculateAddLiquidity',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'calculateBridgeFee',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'calculateRemoveLiquidity',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'calculateSwap',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'calculateWithdrawOneToken',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'config', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'fee', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'getAmountOut',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'getConnectedBridgeTokens',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'getDestinationAmountOut',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'getOriginAmountOut',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'multicallView',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'owner', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'poolInfo', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'poolTokens', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'poolsAmount', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'removeToken', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'removeTokens',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'renounceOwnership',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setAllowance',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setSwapQuoter',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setTokenConfig',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'setTokenFee', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'swap', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'swapQuoter', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'symbolToToken',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'synapseBridge',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'tokenToSymbol',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'transferOwnership',
    data: BytesLike
  ): Result

  events: {
    'OwnershipTransferred(address,address)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'OwnershipTransferred'): EventFragment
}

export interface OwnershipTransferredEventObject {
  previousOwner: string
  newOwner: string
}
export type OwnershipTransferredEvent = TypedEvent<
  [string, string],
  OwnershipTransferredEventObject
>

export type OwnershipTransferredEventFilter =
  TypedEventFilter<OwnershipTransferredEvent>

export interface SynapseRouter extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  interface: SynapseRouterInterface

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>
  listeners(eventName?: string): Array<Listener>
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this
  removeAllListeners(eventName?: string): this
  off: OnEvent<this>
  on: OnEvent<this>
  once: OnEvent<this>
  removeListener: OnEvent<this>

  functions: {
    adapterSwap(
      to: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<ContractTransaction>

    addToken(
      symbol: string,
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    addTokens(
      tokens: LocalBridgeConfig.BridgeTokenConfigStruct[],
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    allPools(
      overrides?: CallOverrides
    ): Promise<[PoolStructOutput[]] & { pools: PoolStructOutput[] }>

    bridge(
      to: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<ContractTransaction>

    bridgeTokens(
      overrides?: CallOverrides
    ): Promise<[string[]] & { tokens: string[] }>

    bridgeTokensAmount(
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { amount: BigNumber }>

    calculateAddLiquidity(
      pool: string,
      amounts: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<[BigNumber]>

    calculateBridgeFee(
      token: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { feeAmount: BigNumber }>

    calculateRemoveLiquidity(
      pool: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber[]] & { amountsOut: BigNumber[] }>

    calculateSwap(
      pool: string,
      tokenIndexFrom: BigNumberish,
      tokenIndexTo: BigNumberish,
      dx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { amountOut: BigNumber }>

    calculateWithdrawOneToken(
      pool: string,
      tokenAmount: BigNumberish,
      tokenIndex: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { amountOut: BigNumber }>

    config(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<[number, string] & { tokenType: number; bridgeToken: string }>

    fee(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<
      [number, BigNumber, BigNumber] & {
        bridgeFee: number
        minFee: BigNumber
        maxFee: BigNumber
      }
    >

    getAmountOut(
      tokenIn: string,
      tokenOut: string,
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[SwapQueryStructOutput]>

    getConnectedBridgeTokens(
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<
      [BridgeTokenStructOutput[]] & { tokens: BridgeTokenStructOutput[] }
    >

    getDestinationAmountOut(
      requests: DestRequestStruct[],
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<
      [SwapQueryStructOutput[]] & { destQueries: SwapQueryStructOutput[] }
    >

    getOriginAmountOut(
      tokenIn: string,
      tokenSymbols: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [SwapQueryStructOutput[]] & { originQueries: SwapQueryStructOutput[] }
    >

    multicallView(
      data: BytesLike[],
      overrides?: CallOverrides
    ): Promise<
      [MulticallView.ResultStructOutput[]] & {
        callResults: MulticallView.ResultStructOutput[]
      }
    >

    owner(overrides?: CallOverrides): Promise<[string]>

    poolInfo(
      pool: string,
      overrides?: CallOverrides
    ): Promise<[BigNumber, string]>

    poolTokens(
      pool: string,
      overrides?: CallOverrides
    ): Promise<[PoolTokenStructOutput[]] & { tokens: PoolTokenStructOutput[] }>

    poolsAmount(
      overrides?: CallOverrides
    ): Promise<[BigNumber] & { amount: BigNumber }>

    removeToken(
      token: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    removeTokens(
      tokens: string[],
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setAllowance(
      token: string,
      spender: string,
      amount: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setSwapQuoter(
      _swapQuoter: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setTokenConfig(
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setTokenFee(
      token: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    swap(
      to: string,
      token: string,
      amount: BigNumberish,
      query: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<ContractTransaction>

    swapQuoter(overrides?: CallOverrides): Promise<[string]>

    symbolToToken(arg0: string, overrides?: CallOverrides): Promise<[string]>

    synapseBridge(overrides?: CallOverrides): Promise<[string]>

    tokenToSymbol(arg0: string, overrides?: CallOverrides): Promise<[string]>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>
  }

  adapterSwap(
    to: string,
    tokenIn: string,
    amountIn: BigNumberish,
    tokenOut: string,
    rawParams: BytesLike,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractTransaction>

  addToken(
    symbol: string,
    token: string,
    tokenType: BigNumberish,
    bridgeToken: string,
    bridgeFee: BigNumberish,
    minFee: BigNumberish,
    maxFee: BigNumberish,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  addTokens(
    tokens: LocalBridgeConfig.BridgeTokenConfigStruct[],
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  allPools(overrides?: CallOverrides): Promise<PoolStructOutput[]>

  bridge(
    to: string,
    chainId: BigNumberish,
    token: string,
    amount: BigNumberish,
    originQuery: SwapQueryStruct,
    destQuery: SwapQueryStruct,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractTransaction>

  bridgeTokens(overrides?: CallOverrides): Promise<string[]>

  bridgeTokensAmount(overrides?: CallOverrides): Promise<BigNumber>

  calculateAddLiquidity(
    pool: string,
    amounts: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<BigNumber>

  calculateBridgeFee(
    token: string,
    amount: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  calculateRemoveLiquidity(
    pool: string,
    amount: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber[]>

  calculateSwap(
    pool: string,
    tokenIndexFrom: BigNumberish,
    tokenIndexTo: BigNumberish,
    dx: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  calculateWithdrawOneToken(
    pool: string,
    tokenAmount: BigNumberish,
    tokenIndex: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>

  config(
    arg0: string,
    overrides?: CallOverrides
  ): Promise<[number, string] & { tokenType: number; bridgeToken: string }>

  fee(
    arg0: string,
    overrides?: CallOverrides
  ): Promise<
    [number, BigNumber, BigNumber] & {
      bridgeFee: number
      minFee: BigNumber
      maxFee: BigNumber
    }
  >

  getAmountOut(
    tokenIn: string,
    tokenOut: string,
    amountIn: BigNumberish,
    overrides?: CallOverrides
  ): Promise<SwapQueryStructOutput>

  getConnectedBridgeTokens(
    tokenOut: string,
    overrides?: CallOverrides
  ): Promise<BridgeTokenStructOutput[]>

  getDestinationAmountOut(
    requests: DestRequestStruct[],
    tokenOut: string,
    overrides?: CallOverrides
  ): Promise<SwapQueryStructOutput[]>

  getOriginAmountOut(
    tokenIn: string,
    tokenSymbols: string[],
    amountIn: BigNumberish,
    overrides?: CallOverrides
  ): Promise<SwapQueryStructOutput[]>

  multicallView(
    data: BytesLike[],
    overrides?: CallOverrides
  ): Promise<MulticallView.ResultStructOutput[]>

  owner(overrides?: CallOverrides): Promise<string>

  poolInfo(
    pool: string,
    overrides?: CallOverrides
  ): Promise<[BigNumber, string]>

  poolTokens(
    pool: string,
    overrides?: CallOverrides
  ): Promise<PoolTokenStructOutput[]>

  poolsAmount(overrides?: CallOverrides): Promise<BigNumber>

  removeToken(
    token: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  removeTokens(
    tokens: string[],
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  renounceOwnership(
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setAllowance(
    token: string,
    spender: string,
    amount: BigNumberish,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setSwapQuoter(
    _swapQuoter: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setTokenConfig(
    token: string,
    tokenType: BigNumberish,
    bridgeToken: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setTokenFee(
    token: string,
    bridgeFee: BigNumberish,
    minFee: BigNumberish,
    maxFee: BigNumberish,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  swap(
    to: string,
    token: string,
    amount: BigNumberish,
    query: SwapQueryStruct,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractTransaction>

  swapQuoter(overrides?: CallOverrides): Promise<string>

  symbolToToken(arg0: string, overrides?: CallOverrides): Promise<string>

  synapseBridge(overrides?: CallOverrides): Promise<string>

  tokenToSymbol(arg0: string, overrides?: CallOverrides): Promise<string>

  transferOwnership(
    newOwner: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  callStatic: {
    adapterSwap(
      to: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    addToken(
      symbol: string,
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: CallOverrides
    ): Promise<boolean>

    addTokens(
      tokens: LocalBridgeConfig.BridgeTokenConfigStruct[],
      overrides?: CallOverrides
    ): Promise<void>

    allPools(overrides?: CallOverrides): Promise<PoolStructOutput[]>

    bridge(
      to: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: CallOverrides
    ): Promise<void>

    bridgeTokens(overrides?: CallOverrides): Promise<string[]>

    bridgeTokensAmount(overrides?: CallOverrides): Promise<BigNumber>

    calculateAddLiquidity(
      pool: string,
      amounts: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateBridgeFee(
      token: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateRemoveLiquidity(
      pool: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber[]>

    calculateSwap(
      pool: string,
      tokenIndexFrom: BigNumberish,
      tokenIndexTo: BigNumberish,
      dx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateWithdrawOneToken(
      pool: string,
      tokenAmount: BigNumberish,
      tokenIndex: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    config(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<[number, string] & { tokenType: number; bridgeToken: string }>

    fee(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<
      [number, BigNumber, BigNumber] & {
        bridgeFee: number
        minFee: BigNumber
        maxFee: BigNumber
      }
    >

    getAmountOut(
      tokenIn: string,
      tokenOut: string,
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<SwapQueryStructOutput>

    getConnectedBridgeTokens(
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<BridgeTokenStructOutput[]>

    getDestinationAmountOut(
      requests: DestRequestStruct[],
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<SwapQueryStructOutput[]>

    getOriginAmountOut(
      tokenIn: string,
      tokenSymbols: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<SwapQueryStructOutput[]>

    multicallView(
      data: BytesLike[],
      overrides?: CallOverrides
    ): Promise<MulticallView.ResultStructOutput[]>

    owner(overrides?: CallOverrides): Promise<string>

    poolInfo(
      pool: string,
      overrides?: CallOverrides
    ): Promise<[BigNumber, string]>

    poolTokens(
      pool: string,
      overrides?: CallOverrides
    ): Promise<PoolTokenStructOutput[]>

    poolsAmount(overrides?: CallOverrides): Promise<BigNumber>

    removeToken(token: string, overrides?: CallOverrides): Promise<boolean>

    removeTokens(tokens: string[], overrides?: CallOverrides): Promise<void>

    renounceOwnership(overrides?: CallOverrides): Promise<void>

    setAllowance(
      token: string,
      spender: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    setSwapQuoter(_swapQuoter: string, overrides?: CallOverrides): Promise<void>

    setTokenConfig(
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      overrides?: CallOverrides
    ): Promise<void>

    setTokenFee(
      token: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: CallOverrides
    ): Promise<void>

    swap(
      to: string,
      token: string,
      amount: BigNumberish,
      query: SwapQueryStruct,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    swapQuoter(overrides?: CallOverrides): Promise<string>

    symbolToToken(arg0: string, overrides?: CallOverrides): Promise<string>

    synapseBridge(overrides?: CallOverrides): Promise<string>

    tokenToSymbol(arg0: string, overrides?: CallOverrides): Promise<string>

    transferOwnership(
      newOwner: string,
      overrides?: CallOverrides
    ): Promise<void>
  }

  filters: {
    'OwnershipTransferred(address,address)'(
      previousOwner?: string | null,
      newOwner?: string | null
    ): OwnershipTransferredEventFilter
    OwnershipTransferred(
      previousOwner?: string | null,
      newOwner?: string | null
    ): OwnershipTransferredEventFilter
  }

  estimateGas: {
    adapterSwap(
      to: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<BigNumber>

    addToken(
      symbol: string,
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    addTokens(
      tokens: LocalBridgeConfig.BridgeTokenConfigStruct[],
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    allPools(overrides?: CallOverrides): Promise<BigNumber>

    bridge(
      to: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<BigNumber>

    bridgeTokens(overrides?: CallOverrides): Promise<BigNumber>

    bridgeTokensAmount(overrides?: CallOverrides): Promise<BigNumber>

    calculateAddLiquidity(
      pool: string,
      amounts: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateBridgeFee(
      token: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateRemoveLiquidity(
      pool: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateSwap(
      pool: string,
      tokenIndexFrom: BigNumberish,
      tokenIndexTo: BigNumberish,
      dx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    calculateWithdrawOneToken(
      pool: string,
      tokenAmount: BigNumberish,
      tokenIndex: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    config(arg0: string, overrides?: CallOverrides): Promise<BigNumber>

    fee(arg0: string, overrides?: CallOverrides): Promise<BigNumber>

    getAmountOut(
      tokenIn: string,
      tokenOut: string,
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getConnectedBridgeTokens(
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getDestinationAmountOut(
      requests: DestRequestStruct[],
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    getOriginAmountOut(
      tokenIn: string,
      tokenSymbols: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    multicallView(
      data: BytesLike[],
      overrides?: CallOverrides
    ): Promise<BigNumber>

    owner(overrides?: CallOverrides): Promise<BigNumber>

    poolInfo(pool: string, overrides?: CallOverrides): Promise<BigNumber>

    poolTokens(pool: string, overrides?: CallOverrides): Promise<BigNumber>

    poolsAmount(overrides?: CallOverrides): Promise<BigNumber>

    removeToken(
      token: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    removeTokens(
      tokens: string[],
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setAllowance(
      token: string,
      spender: string,
      amount: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setSwapQuoter(
      _swapQuoter: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setTokenConfig(
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setTokenFee(
      token: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    swap(
      to: string,
      token: string,
      amount: BigNumberish,
      query: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<BigNumber>

    swapQuoter(overrides?: CallOverrides): Promise<BigNumber>

    symbolToToken(arg0: string, overrides?: CallOverrides): Promise<BigNumber>

    synapseBridge(overrides?: CallOverrides): Promise<BigNumber>

    tokenToSymbol(arg0: string, overrides?: CallOverrides): Promise<BigNumber>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>
  }

  populateTransaction: {
    adapterSwap(
      to: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<PopulatedTransaction>

    addToken(
      symbol: string,
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    addTokens(
      tokens: LocalBridgeConfig.BridgeTokenConfigStruct[],
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    allPools(overrides?: CallOverrides): Promise<PopulatedTransaction>

    bridge(
      to: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<PopulatedTransaction>

    bridgeTokens(overrides?: CallOverrides): Promise<PopulatedTransaction>

    bridgeTokensAmount(overrides?: CallOverrides): Promise<PopulatedTransaction>

    calculateAddLiquidity(
      pool: string,
      amounts: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    calculateBridgeFee(
      token: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    calculateRemoveLiquidity(
      pool: string,
      amount: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    calculateSwap(
      pool: string,
      tokenIndexFrom: BigNumberish,
      tokenIndexTo: BigNumberish,
      dx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    calculateWithdrawOneToken(
      pool: string,
      tokenAmount: BigNumberish,
      tokenIndex: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    config(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    fee(arg0: string, overrides?: CallOverrides): Promise<PopulatedTransaction>

    getAmountOut(
      tokenIn: string,
      tokenOut: string,
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getConnectedBridgeTokens(
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getDestinationAmountOut(
      requests: DestRequestStruct[],
      tokenOut: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    getOriginAmountOut(
      tokenIn: string,
      tokenSymbols: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    multicallView(
      data: BytesLike[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>

    poolInfo(
      pool: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    poolTokens(
      pool: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    poolsAmount(overrides?: CallOverrides): Promise<PopulatedTransaction>

    removeToken(
      token: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    removeTokens(
      tokens: string[],
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setAllowance(
      token: string,
      spender: string,
      amount: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setSwapQuoter(
      _swapQuoter: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setTokenConfig(
      token: string,
      tokenType: BigNumberish,
      bridgeToken: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setTokenFee(
      token: string,
      bridgeFee: BigNumberish,
      minFee: BigNumberish,
      maxFee: BigNumberish,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    swap(
      to: string,
      token: string,
      amount: BigNumberish,
      query: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<PopulatedTransaction>

    swapQuoter(overrides?: CallOverrides): Promise<PopulatedTransaction>

    symbolToToken(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    synapseBridge(overrides?: CallOverrides): Promise<PopulatedTransaction>

    tokenToSymbol(
      arg0: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>
  }
}
