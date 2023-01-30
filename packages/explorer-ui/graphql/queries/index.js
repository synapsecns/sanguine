import {gql} from '@apollo/client'

const SINGLE_SIDE_INFO_FRAGMENT = gql`
  fragment SingleSideInfo on PartialInfo {
    chainID
    address
    hash: txnHash
    value
    formattedValue
    tokenAddress
    tokenSymbol
    time
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
    $txnHash:         String
    $address:         String
    $chainID:         [Int]
    $page:            Int
    $tokenAddress:    [String]
    $pending:         Boolean
    $kappa:           String
  ) {
    bridgeTransactions(
      txnHash:          $txnHash
      address:          $address
      chainID:          $chainID
      page:             $page
      tokenAddress:     $tokenAddress
      pending:          $pending
      kappa:            $kappa
    ) {
      ...TransactionInfo
    }
  }
  ${BRIDGE_TRANSACTION_INFO_FRAGMENT}
`
export const COUNT_BY_CHAIN_ID = gql`
  query CountByChainId(
    $chainID:   Int
    $direction: Direction
    $hours:     Int
  ) {
    countByChainId(
      chainID:    $chainID
      direction:  $direction
      hours:      $hours
    ) {
      chainID
      count
    }
  }
`

export const COUNT_BY_TOKEN_ADDRESS = gql`
  query CountByTokenAddress(
    $chainID:   Int
    $direction: Direction
    $hours:     Int
    $address:   String
  ) {
    countByTokenAddress(
      chainID:    $chainID
      direction:  $direction
      hours:      $hours
      address:    $address
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

export const BRIDGE_AMOUNT_STATISTIC = gql`
  query BridgeAmountStatistic(
    $type:          StatisticType!
    $duration:      Duration!
    $chainID:       Int
    $address:       String
    $tokenAddress:  String
  ) {
    bridgeAmountStatistic(
      type:           $type
      duration:       $duration
      chainID:        $chainID
      address:        $address
      tokenAddress:   $tokenAddress
    ) {
      value
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

export const GET_HISTORICAL_STATS = gql`
  query HistoricalStatistics(
    $chainID: Int
    $type: HistoricalResultType!
    $days: Int
  ) {
    historicalStatistics(
      chainID: $chainID
      type: $type
      days: $days
    ) {
      total
      dateResults {
        date
        total
      }
    }
  }
`


export const GET_DAILY_STATS = gql`
  query DailyStatistics(
    $chainID: Int
    $type: DailyStatisticType!,
    $platform: Platform,
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
    $type:          StatisticType!
    $duration:      Duration!
    $platform:      Platform
    $chainID:       Int
    $address:       String
    $tokenAddress:  String
  ) {
    amountStatistic(
      type: $type
      duration: $duration
      platform: $platform
      chainID: $chainID
      address: $address
      tokenAddress: $tokenAddress
    ) {
      value
    }
  }
`
