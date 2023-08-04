import { gql } from '@apollo/client'

const SINGLE_SIDE_INFO_FRAGMENT = gql`
  fragment SingleSideInfo on PartialInfo {
    chainID
    destinationChainID
    address
    hash: txnHash
    value
    formattedValue
    tokenAddress
    tokenSymbol
    time
    eventType
  }
`

const BRIDGE_TRANSACTION_INFO_FRAGMENT = gql`
  fragment TransactionInfo on BridgeTransaction {
    fromInfo {
      ...SingleSideInfo
    }
    toInfo {
      ...SingleSideInfo
    }

    kappa
    pending
    swapSuccess
  }

  ${SINGLE_SIDE_INFO_FRAGMENT}
`

export const GET_BRIDGE_TRANSACTIONS_QUERY = gql`
  query GetBridgeTransactionsQuery(
    $chainIDFrom: [Int]
    $chainIDTo: [Int]
    $addressFrom: String
    $addressTo: String
    $maxAmount: Int
    $minAmount: Int
    $maxAmountUsd: Int
    $minAmountUsd: Int
    $startTime: Int
    $endTime: Int
    $txnHash: String
    $kappa: String
    $pending: Boolean
    $page: Int
    $tokenAddressFrom: [String]
    $tokenAddressTo: [String]
    $useMv: Boolean
  ) {
    bridgeTransactions(
      chainIDFrom: $chainIDFrom
      chainIDTo: $chainIDTo
      addressFrom: $addressFrom
      addressTo: $addressTo
      maxAmount: $maxAmount
      minAmount: $minAmount
      maxAmountUsd: $maxAmountUsd
      minAmountUsd: $minAmountUsd
      startTime: $startTime
      endTime: $endTime
      txnHash: $txnHash
      kappa: $kappa
      pending: $pending
      page: $page
      useMv: $useMv
      tokenAddressFrom: $tokenAddressFrom
      tokenAddressTo: $tokenAddressTo
    ) {
      ...TransactionInfo
    }
  }
  ${BRIDGE_TRANSACTION_INFO_FRAGMENT}
`
export const COUNT_BY_CHAIN_ID = gql`
  query CountByChainId($chainID: Int, $direction: Direction, $hours: Int) {
    countByChainId(chainID: $chainID, direction: $direction, hours: $hours) {
      chainID
      count
    }
  }
`

export const COUNT_BY_TOKEN_ADDRESS = gql`
  query CountByTokenAddress(
    $chainID: Int
    $direction: Direction
    $hours: Int
    $address: String
  ) {
    countByTokenAddress(
      chainID: $chainID
      direction: $direction
      hours: $hours
      address: $address
    ) {
      tokenAddress
      chainID
      count
    }
  }
`

export const ADDRESS_RANKING = gql`
  query AddressRanking($hours: Int) {
    addressRanking(hours: $hours) {
      address
      count
    }
  }
`

export const GET_CSV = gql`
  query GetCsv($address: String!) {
    getCsv(address: $address) {
      cid
      ipfsGatewayUrl
    }
  }
`

export const GET_DAILY_STATS = gql`
  query DailyStatistics(
    $chainID: Int
    $type: DailyStatisticType!
    $platform: Platform
    $days: Int
  ) {
    dailyStatistics(
      chainID: $chainID
      type: $type
      days: $days
      platform: $platform
    ) {
      total
      dateResults {
        date
        total
      }
    }
  }
`

export const AMOUNT_STATISTIC = gql`
  query AmountStatistic(
    $type: StatisticType!
    $duration: Duration!
    $platform: Platform
    $chainID: Int
    $address: String
    $tokenAddress: String
    $useCache: Boolean
    $useMv: Boolean
  ) {
    amountStatistic(
      type: $type
      duration: $duration
      platform: $platform
      chainID: $chainID
      address: $address
      tokenAddress: $tokenAddress
      useCache: $useCache
      useMv: $useMv
    ) {
      value
    }
  }
`

export const DAILY_STATISTICS_BY_CHAIN = gql`
  query DailyStatisticsByChain(
    $chainID: Int
    $type: DailyStatisticType
    $duration: Duration
    $useCache: Boolean
    $platform: Platform
    $useMv: Boolean
  ) {
    dailyStatisticsByChain(
      chainID: $chainID
      type: $type
      duration: $duration
      useCache: $useCache
      platform: $platform
      useMv: $useMv
    ) {
      date
      ethereum
      optimism
      cronos
      bsc
      polygon
      fantom
      boba
      metis
      moonbeam
      moonriver
      klaytn
      arbitrum
      avalanche
      dfk
      aurora
      harmony
      canto
      dogechain
      base
      total
    }
  }
`

export const RANKED_CHAINIDS_BY_VOLUME = gql`
  query RankedChainIDsByVolume($duration: Duration) {
    rankedChainIDsByVolume(duration: $duration) {
      chainID
      total
    }
  }
`
