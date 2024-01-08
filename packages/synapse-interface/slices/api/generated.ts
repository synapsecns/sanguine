import { api } from '@/slices/api/slice'
export type Maybe<T> = T | null
export type InputMaybe<T> = Maybe<T>
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K]
}
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>
}
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>
}
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T
> = { [_ in K]?: never }
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never
    }
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string }
  String: { input: string; output: string }
  Boolean: { input: boolean; output: boolean }
  Int: { input: number; output: number }
  Float: { input: number; output: number }
}

export type AddressChainRanking = {
  __typename?: 'AddressChainRanking'
  chainID?: Maybe<Scalars['Int']['output']>
  rank?: Maybe<Scalars['Int']['output']>
  volumeUsd?: Maybe<Scalars['Float']['output']>
}

export type AddressDailyCount = {
  __typename?: 'AddressDailyCount'
  count?: Maybe<Scalars['Int']['output']>
  date?: Maybe<Scalars['String']['output']>
}

export type AddressData = {
  __typename?: 'AddressData'
  bridgeFees?: Maybe<Scalars['Float']['output']>
  bridgeTxs?: Maybe<Scalars['Int']['output']>
  bridgeVolume?: Maybe<Scalars['Float']['output']>
  chainRanking?: Maybe<Array<Maybe<AddressChainRanking>>>
  dailyData?: Maybe<Array<Maybe<AddressDailyCount>>>
  earliestTx?: Maybe<Scalars['Int']['output']>
  rank?: Maybe<Scalars['Int']['output']>
  swapFees?: Maybe<Scalars['Float']['output']>
  swapTxs?: Maybe<Scalars['Int']['output']>
  swapVolume?: Maybe<Scalars['Float']['output']>
}

/** AddressRanking gives the amount of transactions that occurred for a specific address across all chains. */
export type AddressRanking = {
  __typename?: 'AddressRanking'
  address?: Maybe<Scalars['String']['output']>
  count?: Maybe<Scalars['Int']['output']>
}

export type BlockHeight = {
  __typename?: 'BlockHeight'
  blockNumber?: Maybe<Scalars['Int']['output']>
  chainID?: Maybe<Scalars['Int']['output']>
  type?: Maybe<ContractType>
}

/**
 * BridgeTransaction represents an entire bridge transaction, including both
 * to and from transactions. If a `from` transaction does not have a corresponding
 * `to` transaction, `pending` will be true.
 */
export type BridgeTransaction = {
  __typename?: 'BridgeTransaction'
  fromInfo?: Maybe<PartialInfo>
  kappa?: Maybe<Scalars['String']['output']>
  pending?: Maybe<Scalars['Boolean']['output']>
  swapSuccess?: Maybe<Scalars['Boolean']['output']>
  toInfo?: Maybe<PartialInfo>
}

export enum BridgeTxType {
  Destination = 'DESTINATION',
  Origin = 'ORIGIN',
  Rfq = 'RFQ',
}

export enum BridgeType {
  Bridge = 'BRIDGE',
  Cctp = 'CCTP',
  Rfq = 'RFQ',
}

/** BridgeWatcherTx represents a single sided bridge transaction specifically for the bridge watcher. */
export type BridgeWatcherTx = {
  __typename?: 'BridgeWatcherTx'
  bridgeTx?: Maybe<PartialInfo>
  kappa?: Maybe<Scalars['String']['output']>
  kappaStatus?: Maybe<KappaStatus>
  pending?: Maybe<Scalars['Boolean']['output']>
  type?: Maybe<BridgeTxType>
}

export type ContractQuery = {
  chainID?: InputMaybe<Scalars['Int']['input']>
  type?: InputMaybe<ContractType>
}

export enum ContractType {
  Bridge = 'BRIDGE',
  Cctp = 'CCTP',
  Rfq = 'RFQ',
}

export enum DailyStatisticType {
  Addresses = 'ADDRESSES',
  Fee = 'FEE',
  Transactions = 'TRANSACTIONS',
  Volume = 'VOLUME',
}

/** DateResult is a given statistic for a given date. */
export type DateResult = {
  __typename?: 'DateResult'
  date?: Maybe<Scalars['String']['output']>
  total?: Maybe<Scalars['Float']['output']>
}

/** DateResult is a given statistic for a given date. */
export type DateResultByChain = {
  __typename?: 'DateResultByChain'
  arbitrum?: Maybe<Scalars['Float']['output']>
  aurora?: Maybe<Scalars['Float']['output']>
  avalanche?: Maybe<Scalars['Float']['output']>
  base?: Maybe<Scalars['Float']['output']>
  boba?: Maybe<Scalars['Float']['output']>
  bsc?: Maybe<Scalars['Float']['output']>
  canto?: Maybe<Scalars['Float']['output']>
  cronos?: Maybe<Scalars['Float']['output']>
  date?: Maybe<Scalars['String']['output']>
  dfk?: Maybe<Scalars['Float']['output']>
  dogechain?: Maybe<Scalars['Float']['output']>
  ethereum?: Maybe<Scalars['Float']['output']>
  fantom?: Maybe<Scalars['Float']['output']>
  harmony?: Maybe<Scalars['Float']['output']>
  klaytn?: Maybe<Scalars['Float']['output']>
  metis?: Maybe<Scalars['Float']['output']>
  moonbeam?: Maybe<Scalars['Float']['output']>
  moonriver?: Maybe<Scalars['Float']['output']>
  optimism?: Maybe<Scalars['Float']['output']>
  polygon?: Maybe<Scalars['Float']['output']>
  total?: Maybe<Scalars['Float']['output']>
}

export enum Direction {
  In = 'IN',
  Out = 'OUT',
}

export enum Duration {
  AllTime = 'ALL_TIME',
  Past_3Months = 'PAST_3_MONTHS',
  Past_6Months = 'PAST_6_MONTHS',
  PastDay = 'PAST_DAY',
  PastMonth = 'PAST_MONTH',
  PastYear = 'PAST_YEAR',
}

export type HeroType = {
  __typename?: 'HeroType'
  heroID: Scalars['String']['output']
  recipient: Scalars['String']['output']
}

/** HistoricalResult is a given statistic for dates. */
export type HistoricalResult = {
  __typename?: 'HistoricalResult'
  dateResults?: Maybe<Array<Maybe<DateResult>>>
  total?: Maybe<Scalars['Float']['output']>
  type?: Maybe<HistoricalResultType>
}

export enum HistoricalResultType {
  Addresses = 'ADDRESSES',
  Bridgevolume = 'BRIDGEVOLUME',
  Transactions = 'TRANSACTIONS',
}

export enum KappaStatus {
  Exists = 'EXISTS',
  Pending = 'PENDING',
  Unknown = 'UNKNOWN',
}

export type Leaderboard = {
  __typename?: 'Leaderboard'
  address?: Maybe<Scalars['String']['output']>
  avgVolumeUSD?: Maybe<Scalars['Float']['output']>
  fees?: Maybe<Scalars['Float']['output']>
  rank?: Maybe<Scalars['Int']['output']>
  txs?: Maybe<Scalars['Int']['output']>
  volumeUSD?: Maybe<Scalars['Float']['output']>
}

export type MessageBusTransaction = {
  __typename?: 'MessageBusTransaction'
  fromInfo?: Maybe<PartialMessageBusInfo>
  messageID?: Maybe<Scalars['String']['output']>
  pending?: Maybe<Scalars['Boolean']['output']>
  toInfo?: Maybe<PartialMessageBusInfo>
}

export type MessageType = HeroType | PetType | TearType | UnknownType

/** PartialInfo is a transaction that occurred on one chain. */
export type PartialInfo = {
  __typename?: 'PartialInfo'
  USDValue?: Maybe<Scalars['Float']['output']>
  address?: Maybe<Scalars['String']['output']>
  blockNumber?: Maybe<Scalars['Int']['output']>
  chainID?: Maybe<Scalars['Int']['output']>
  destinationChainID?: Maybe<Scalars['Int']['output']>
  eventType?: Maybe<Scalars['Int']['output']>
  formattedEventType?: Maybe<Scalars['String']['output']>
  formattedTime?: Maybe<Scalars['String']['output']>
  formattedValue?: Maybe<Scalars['Float']['output']>
  time?: Maybe<Scalars['Int']['output']>
  tokenAddress?: Maybe<Scalars['String']['output']>
  tokenSymbol?: Maybe<Scalars['String']['output']>
  txnHash?: Maybe<Scalars['String']['output']>
  value?: Maybe<Scalars['String']['output']>
}

export type PartialMessageBusInfo = {
  __typename?: 'PartialMessageBusInfo'
  blockNumber?: Maybe<Scalars['Int']['output']>
  chainID?: Maybe<Scalars['Int']['output']>
  chainName?: Maybe<Scalars['String']['output']>
  contractAddress?: Maybe<Scalars['String']['output']>
  destinationChainID?: Maybe<Scalars['Int']['output']>
  destinationChainName?: Maybe<Scalars['String']['output']>
  formattedTime?: Maybe<Scalars['String']['output']>
  message?: Maybe<Scalars['String']['output']>
  messageType?: Maybe<MessageType>
  revertedReason?: Maybe<Scalars['String']['output']>
  time?: Maybe<Scalars['Int']['output']>
  txnHash?: Maybe<Scalars['String']['output']>
}

export type PetType = {
  __typename?: 'PetType'
  name: Scalars['String']['output']
  petID: Scalars['String']['output']
  recipient: Scalars['String']['output']
}

export enum Platform {
  All = 'ALL',
  Bridge = 'BRIDGE',
  MessageBus = 'MESSAGE_BUS',
  Swap = 'SWAP',
}

export type Query = {
  __typename?: 'Query'
  /** Get wallet information */
  addressData?: Maybe<AddressData>
  /**
   * Returns addresses and transaction count (origin) over time.
   * Specifying no parameters defaults to 24 hours.
   */
  addressRanking?: Maybe<Array<Maybe<AddressRanking>>>
  /**
   * Returns mean/median/total/count of transactions transacted for a given duration, chain and address.
   * Specifying no duration defaults to ALL_TIME, and no chain or address searches across all.
   */
  amountStatistic?: Maybe<ValueResult>
  /** Returns bridged transactions filterable by chain, to/from address, to/from txn hash, token address, and keccak hash. */
  bridgeTransactions?: Maybe<Array<Maybe<BridgeTransaction>>>
  /**
   * Returns the COUNT of bridged transactions for a given chain. If direction of bridge transactions
   * is not specified, it defaults to IN.
   * Specifying no duration defaults to the last 24 hours.
   */
  countByChainId?: Maybe<Array<Maybe<TransactionCountResult>>>
  /**
   * Returns counts of token addresses source and time.
   * Specifying no parameters defaults to origin and 24 hours.
   */
  countByTokenAddress?: Maybe<Array<Maybe<TokenCountResult>>>
  /** Daily statistic data */
  dailyStatisticsByChain?: Maybe<Array<Maybe<DateResultByChain>>>
  /** GetBlockHeight gets block heights from the current bridge. Returns results in an array of increased block heights. */
  getBlockHeight?: Maybe<Array<Maybe<BlockHeight>>>
  /** GetDestinationBridgeTx is the bridge watcher endpoint for getting an destination bridge tx (BETA). */
  getDestinationBridgeTx?: Maybe<BridgeWatcherTx>
  /** GetOriginBridgeTx is the bridge watcher endpoint for getting an origin bridge tx (BETA). */
  getOriginBridgeTx?: Maybe<BridgeWatcherTx>
  /** Get LeaderBoard */
  leaderboard?: Maybe<Array<Maybe<Leaderboard>>>
  /** Message bus transactions */
  messageBusTransactions?: Maybe<Array<Maybe<MessageBusTransaction>>>
  /** Ranked chainIDs by volume */
  rankedChainIDsByVolume?: Maybe<Array<Maybe<VolumeByChainId>>>
}

export type QueryAddressDataArgs = {
  address: Scalars['String']['input']
}

export type QueryAddressRankingArgs = {
  hours?: InputMaybe<Scalars['Int']['input']>
}

export type QueryAmountStatisticArgs = {
  address?: InputMaybe<Scalars['String']['input']>
  chainID?: InputMaybe<Scalars['Int']['input']>
  duration?: InputMaybe<Duration>
  platform?: InputMaybe<Platform>
  tokenAddress?: InputMaybe<Scalars['String']['input']>
  type: StatisticType
  useCache?: InputMaybe<Scalars['Boolean']['input']>
  useMv?: InputMaybe<Scalars['Boolean']['input']>
}

export type QueryBridgeTransactionsArgs = {
  addressFrom?: InputMaybe<Scalars['String']['input']>
  addressTo?: InputMaybe<Scalars['String']['input']>
  chainIDFrom?: InputMaybe<Array<InputMaybe<Scalars['Int']['input']>>>
  chainIDTo?: InputMaybe<Array<InputMaybe<Scalars['Int']['input']>>>
  endTime?: InputMaybe<Scalars['Int']['input']>
  kappa?: InputMaybe<Scalars['String']['input']>
  maxAmount?: InputMaybe<Scalars['Int']['input']>
  maxAmountUsd?: InputMaybe<Scalars['Int']['input']>
  minAmount?: InputMaybe<Scalars['Int']['input']>
  minAmountUsd?: InputMaybe<Scalars['Int']['input']>
  onlyCCTP?: InputMaybe<Scalars['Boolean']['input']>
  page?: InputMaybe<Scalars['Int']['input']>
  pending?: InputMaybe<Scalars['Boolean']['input']>
  startTime?: InputMaybe<Scalars['Int']['input']>
  tokenAddressFrom?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>
  tokenAddressTo?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>
  txnHash?: InputMaybe<Scalars['String']['input']>
  useMv?: InputMaybe<Scalars['Boolean']['input']>
}

export type QueryCountByChainIdArgs = {
  address?: InputMaybe<Scalars['String']['input']>
  chainID?: InputMaybe<Scalars['Int']['input']>
  direction?: InputMaybe<Direction>
  hours?: InputMaybe<Scalars['Int']['input']>
}

export type QueryCountByTokenAddressArgs = {
  address?: InputMaybe<Scalars['String']['input']>
  chainID?: InputMaybe<Scalars['Int']['input']>
  direction?: InputMaybe<Direction>
  hours?: InputMaybe<Scalars['Int']['input']>
}

export type QueryDailyStatisticsByChainArgs = {
  chainID?: InputMaybe<Scalars['Int']['input']>
  duration?: InputMaybe<Duration>
  platform?: InputMaybe<Platform>
  type?: InputMaybe<DailyStatisticType>
  useCache?: InputMaybe<Scalars['Boolean']['input']>
  useMv?: InputMaybe<Scalars['Boolean']['input']>
}

export type QueryGetBlockHeightArgs = {
  contracts?: InputMaybe<Array<InputMaybe<ContractQuery>>>
}

export type QueryGetDestinationBridgeTxArgs = {
  address: Scalars['String']['input']
  bridgeType: BridgeType
  chainID: Scalars['Int']['input']
  historical?: InputMaybe<Scalars['Boolean']['input']>
  kappa: Scalars['String']['input']
  timestamp: Scalars['Int']['input']
}

export type QueryGetOriginBridgeTxArgs = {
  bridgeType: BridgeType
  chainID: Scalars['Int']['input']
  txnHash: Scalars['String']['input']
}

export type QueryLeaderboardArgs = {
  chainID?: InputMaybe<Scalars['Int']['input']>
  duration?: InputMaybe<Duration>
  page?: InputMaybe<Scalars['Int']['input']>
  useMv?: InputMaybe<Scalars['Boolean']['input']>
}

export type QueryMessageBusTransactionsArgs = {
  chainID?: InputMaybe<Array<InputMaybe<Scalars['Int']['input']>>>
  contractAddress?: InputMaybe<Scalars['String']['input']>
  endTime?: InputMaybe<Scalars['Int']['input']>
  messageID?: InputMaybe<Scalars['String']['input']>
  page?: InputMaybe<Scalars['Int']['input']>
  pending?: InputMaybe<Scalars['Boolean']['input']>
  reverted?: InputMaybe<Scalars['Boolean']['input']>
  startTime?: InputMaybe<Scalars['Int']['input']>
  txnHash?: InputMaybe<Scalars['String']['input']>
}

export type QueryRankedChainIDsByVolumeArgs = {
  duration?: InputMaybe<Duration>
  useCache?: InputMaybe<Scalars['Boolean']['input']>
}

export enum StatisticType {
  CountAddresses = 'COUNT_ADDRESSES',
  CountTransactions = 'COUNT_TRANSACTIONS',
  MeanFeeUsd = 'MEAN_FEE_USD',
  MeanVolumeUsd = 'MEAN_VOLUME_USD',
  MedianFeeUsd = 'MEDIAN_FEE_USD',
  MedianVolumeUsd = 'MEDIAN_VOLUME_USD',
  TotalFeeUsd = 'TOTAL_FEE_USD',
  TotalVolumeUsd = 'TOTAL_VOLUME_USD',
}

export type TearType = {
  __typename?: 'TearType'
  amount: Scalars['String']['output']
  recipient: Scalars['String']['output']
}

/** TokenCountResult gives the amount of transactions that occurred for a specific token, separated by chain ID. */
export type TokenCountResult = {
  __typename?: 'TokenCountResult'
  chainID?: Maybe<Scalars['Int']['output']>
  count?: Maybe<Scalars['Int']['output']>
  tokenAddress?: Maybe<Scalars['String']['output']>
}

/** TransactionCountResult gives the amount of transactions that occurred for a specific chain ID. */
export type TransactionCountResult = {
  __typename?: 'TransactionCountResult'
  chainID?: Maybe<Scalars['Int']['output']>
  count?: Maybe<Scalars['Int']['output']>
}

export type UnknownType = {
  __typename?: 'UnknownType'
  known: Scalars['Boolean']['output']
}

/** ValueResult is a value result of either USD or numeric value. */
export type ValueResult = {
  __typename?: 'ValueResult'
  value?: Maybe<Scalars['String']['output']>
}

export type VolumeByChainId = {
  __typename?: 'VolumeByChainID'
  chainID?: Maybe<Scalars['Int']['output']>
  total?: Maybe<Scalars['Float']['output']>
}

export type GetBlockHeightQueryVariables = Exact<{
  contracts: Array<InputMaybe<ContractQuery>> | InputMaybe<ContractQuery>
}>

export type GetBlockHeightQuery = {
  __typename?: 'Query'
  getBlockHeight?: Array<{
    __typename?: 'BlockHeight'
    chainID?: number | null
    type?: ContractType | null
    blockNumber?: number | null
  } | null> | null
}

export type GetDestinationBridgeTxFallbackQueryVariables = Exact<{
  chainId: Scalars['Int']['input']
  kappa: Scalars['String']['input']
  address: Scalars['String']['input']
  timestamp: Scalars['Int']['input']
  bridgeType: BridgeType
}>

export type GetDestinationBridgeTxFallbackQuery = {
  __typename?: 'Query'
  getDestinationBridgeTx?: {
    __typename?: 'BridgeWatcherTx'
    kappa?: string | null
    pending?: boolean | null
    bridgeTx?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
  } | null
}

export type GetOriginBridgeTxFallbackQueryVariables = Exact<{
  chainId: Scalars['Int']['input']
  txnHash: Scalars['String']['input']
  bridgeType: BridgeType
}>

export type GetOriginBridgeTxFallbackQuery = {
  __typename?: 'Query'
  getOriginBridgeTx?: {
    __typename?: 'BridgeWatcherTx'
    kappa?: string | null
    pending?: boolean | null
    bridgeTx?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
  } | null
}

export type GetUserHistoricalActivityQueryVariables = Exact<{
  address: Scalars['String']['input']
  startTime: Scalars['Int']['input']
}>

export type GetUserHistoricalActivityQuery = {
  __typename?: 'Query'
  bridgeTransactions?: Array<{
    __typename?: 'BridgeTransaction'
    kappa?: string | null
    fromInfo?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
    toInfo?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
  } | null> | null
}

export type GetUserPendingTransactionsQueryVariables = Exact<{
  address: Scalars['String']['input']
  startTime: Scalars['Int']['input']
}>

export type GetUserPendingTransactionsQuery = {
  __typename?: 'Query'
  bridgeTransactions?: Array<{
    __typename?: 'BridgeTransaction'
    kappa?: string | null
    fromInfo?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
    toInfo?: {
      __typename?: 'PartialInfo'
      chainID?: number | null
      destinationChainID?: number | null
      address?: string | null
      txnHash?: string | null
      value?: string | null
      formattedValue?: number | null
      USDValue?: number | null
      tokenAddress?: string | null
      tokenSymbol?: string | null
      blockNumber?: number | null
      time?: number | null
      formattedTime?: string | null
      formattedEventType?: string | null
      eventType?: number | null
    } | null
  } | null> | null
}

export const GetBlockHeightDocument = `
    query GetBlockHeight($contracts: [ContractQuery]!) {
  getBlockHeight(contracts: $contracts) {
    chainID
    type
    blockNumber
  }
}
    `
export const GetDestinationBridgeTxFallbackDocument = `
    query GetDestinationBridgeTxFallback($chainId: Int!, $kappa: String!, $address: String!, $timestamp: Int!, $bridgeType: BridgeType!) {
  getDestinationBridgeTx(
    chainID: $chainId
    kappa: $kappa
    address: $address
    timestamp: $timestamp
    bridgeType: $bridgeType
  ) {
    bridgeTx {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    kappa
    pending
  }
}
    `
export const GetOriginBridgeTxFallbackDocument = `
    query GetOriginBridgeTxFallback($chainId: Int!, $txnHash: String!, $bridgeType: BridgeType!) {
  getOriginBridgeTx(chainID: $chainId, txnHash: $txnHash, bridgeType: $bridgeType) {
    bridgeTx {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    kappa
    pending
  }
}
    `
export const GetUserHistoricalActivityDocument = `
    query GetUserHistoricalActivity($address: String!, $startTime: Int!) {
  bridgeTransactions(
    pending: false
    addressFrom: $address
    startTime: $startTime
    page: 1
    useMv: true
  ) {
    fromInfo {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    toInfo {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    kappa
  }
}
    `
export const GetUserPendingTransactionsDocument = `
    query GetUserPendingTransactions($address: String!, $startTime: Int!) {
  bridgeTransactions(
    pending: true
    addressFrom: $address
    startTime: $startTime
    page: 1
    useMv: true
  ) {
    fromInfo {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    toInfo {
      chainID
      destinationChainID
      address
      txnHash
      value
      formattedValue
      USDValue
      tokenAddress
      tokenSymbol
      blockNumber
      time
      formattedTime
      formattedEventType
      eventType
    }
    kappa
  }
}
    `

const injectedRtkApi = api.injectEndpoints({
  endpoints: (build) => ({
    GetBlockHeight: build.query<
      GetBlockHeightQuery,
      GetBlockHeightQueryVariables
    >({
      query: (variables) => ({ document: GetBlockHeightDocument, variables }),
    }),
    GetDestinationBridgeTxFallback: build.query<
      GetDestinationBridgeTxFallbackQuery,
      GetDestinationBridgeTxFallbackQueryVariables
    >({
      query: (variables) => ({
        document: GetDestinationBridgeTxFallbackDocument,
        variables,
      }),
    }),
    GetOriginBridgeTxFallback: build.query<
      GetOriginBridgeTxFallbackQuery,
      GetOriginBridgeTxFallbackQueryVariables
    >({
      query: (variables) => ({
        document: GetOriginBridgeTxFallbackDocument,
        variables,
      }),
    }),
    GetUserHistoricalActivity: build.query<
      GetUserHistoricalActivityQuery,
      GetUserHistoricalActivityQueryVariables
    >({
      query: (variables) => ({
        document: GetUserHistoricalActivityDocument,
        variables,
      }),
    }),
    GetUserPendingTransactions: build.query<
      GetUserPendingTransactionsQuery,
      GetUserPendingTransactionsQueryVariables
    >({
      query: (variables) => ({
        document: GetUserPendingTransactionsDocument,
        variables,
      }),
    }),
  }),
})

export { injectedRtkApi as api }
export const {
  useGetBlockHeightQuery,
  useLazyGetBlockHeightQuery,
  useGetDestinationBridgeTxFallbackQuery,
  useLazyGetDestinationBridgeTxFallbackQuery,
  useGetOriginBridgeTxFallbackQuery,
  useLazyGetOriginBridgeTxFallbackQuery,
  useGetUserHistoricalActivityQuery,
  useLazyGetUserHistoricalActivityQuery,
  useGetUserPendingTransactionsQuery,
  useLazyGetUserPendingTransactionsQuery,
} = injectedRtkApi
