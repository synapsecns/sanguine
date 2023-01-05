import {gql} from '@apollo/client'

const SINGLE_SIDE_INFO_FRAGMENT = gql`
  fragment SingleSideInfo on PartialInfo {
    chainId
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
    $chainId:         Int
    $page:            Int
    $tokenAddress:    String
    $includePending:  Boolean
    $kappa:           String
  ) {
    bridgeTransactions(
      txnHash:          $txnHash
      address:          $address
      chainId:          $chainId
      page:             $page
      tokenAddress:     $tokenAddress
      includePending:   $includePending
      kappa:            $kappa
    ) {
      ...TransactionInfo
    }
  }
  ${BRIDGE_TRANSACTION_INFO_FRAGMENT}
`
export const COUNT_BY_CHAIN_ID = gql`
  query CountByChainId(
    $chainId:   Int
    $direction: Direction
    $hours:     Int
  ) {
    countByChainId(
      chainId:    $chainId
      direction:  $direction
      hours:      $hours
    ) {
      chainId
      count
    }
  }
`

export const COUNT_BY_TOKEN_ADDRESS = gql`
  query CountByTokenAddress(
    $chainId:   Int
    $direction: Direction
    $hours:     Int
    $address:   String
  ) {
    countByTokenAddress(
      chainId:    $chainId
      direction:  $direction
      hours:      $hours
      address:    $address
    ) {
      tokenAddress
      chainId
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
    $chainId:       Int
    $address:       String
    $tokenAddress:  String
  ) {
    bridgeAmountStatistic(
      type:           $type
      duration:       $duration
      chainId:        $chainId
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
    $chainId: Int
    $type: HistoricalResultType!
    $days: Int
  ) {
    historicalStatistics(
      chainId: $chainId
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
