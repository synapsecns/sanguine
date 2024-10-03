import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { TRANSACTIONS_PATH, ACCOUNTS_PATH } from '@urls'
import { ChainInfo } from '@components/misc/ChainInfo'
import { Error } from '@components/Error'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { useRouter } from 'next/router'
//@ts-ignore
import { CHAINS } from '@synapsecns/synapse-constants'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { API_URL } from '@graphql'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { formatDateTimestamp } from '@utils/formatDate'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { addressToSymbol } from '@utils/addressToSymbol'
const CHAINS_BY_ID = CHAINS.CHAINS_BY_ID

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
})

const truncateAddress = (address) => {
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const truncateHash = (hash) => {
  return `${hash.slice(0, 8)}...${hash.slice(-5)}`
}

export const BridgeTransaction = ({ queryResult }) => {
  const router = useRouter()
  const { kappa } = router.query
  const transaction = queryResult.bridgeTransactions[0]
  const { pending, fromInfo, toInfo } = transaction

  // Get time taken to complete tx w/ appropriate units.
  const getTimeElapsedStr = (start, end) => {
    const diff = end - start
    if (diff <= 0) {
      return '1 second'
    }
    return diff === 1 ? '1 second' : `${diff} seconds`
  }
  let content

  if (!!transaction) {
    content = (
      <>
        <div className="flex flex-row pb-12 mt-5 mb-1 text-white">
          <a
            className="text-[#99E6FF] cursor-pointer mr-2"
            href={TRANSACTIONS_PATH}
          >
            ← Explorer
          </a>
          <span className="mr-2 text-gray-500">/</span>
          <span
            className="text-white cursor-pointer hover:text-gray-500"
            onClick={() => {
              try {
                navigator.clipboard
                  .writeText(String(kappa))
                  .then(() => '')
                  .catch((err) => alert('Failed to copy text: ' + err))
              } catch (err) {
                alert('Clipboard not supported on this browser')
              }
            }}
          >
            <span className="hidden sm:inline">{kappa}</span>
            <span className="sm:hidden">{truncateAddress(kappa)}</span>
          </span>
        </div>
        <div className="py-4 border border-[#252537] rounded-xl w-4/5 mx-auto">
          <div className="flex flex-col px-4 gap-y-2">
            <div className="flex flex-col mt-4">
              <div className="flex flex-col">
                <div className="flex items-center justify-between w-full">
                  <div className="flex gap-x-[1.8rem] py-1">
                    <p className="w-24 text-white text-opacity-60">Timestamp</p>
                    <p className="text-white">
                      {formatDateTimestamp(new Date(fromInfo.time * 1000))}
                    </p>
                  </div>
                  <div className="flex gap-x-[1.8rem] py-1">
                    <p className="hidden text-white sm:block">
                      {new Date(fromInfo.time * 1000)
                        .toISOString()
                        .replace('T', ' ')
                        .slice(0, 16) + ' UTC'}
                    </p>
                  </div>
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="w-24 text-white text-opacity-60">Status</p>
                  <p className="text-white ">
                    {pending ? 'Pending' : 'Confirmed'}{' '}
                  </p>
                </div>
                <div className="flex gap-x-[1.8rem] py-1 pb-4">
                  <p className="w-24 text-white text-opacity-60">Elapsed</p>
                  <p className="text-white ">
                    {toInfo
                      ? getTimeElapsedStr(fromInfo.time, toInfo.time)
                      : '--'}{' '}
                  </p>
                </div>
                <HorizontalDivider className="mx-[-1rem] bg-[#252537]" />
                <div className="flex gap-x-[1.8rem] py-1 pt-4">
                  <p className="w-24 text-white text-opacity-60">From</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-sm text-white break-all cursor-pointer hover:text-gray-500"
                    href={ACCOUNTS_PATH + '/' + fromInfo.address}
                  >
                    <span className="hidden sm:inline">{fromInfo.address}</span>
                    <span className="sm:hidden">
                      {truncateAddress(fromInfo.address)}
                    </span>
                  </a>
                </div>
                <div className="flex gap-x-[1.8rem] py-1 pb-4">
                  <p className="w-24 text-white text-opacity-60">To</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-sm text-white break-all cursor-pointer hover:text-gray-500"
                    href={toInfo ? ACCOUNTS_PATH + '/' + toInfo.address : ''}
                  >
                    {toInfo ? (
                      <>
                        <span className="hidden sm:inline">
                          {toInfo.address}
                        </span>
                        <span className="sm:hidden">
                          {truncateAddress(toInfo.address)}
                        </span>
                      </>
                    ) : (
                      '--'
                    )}
                  </a>
                </div>
                <HorizontalDivider className="mx-[-1rem] bg-[#252537]" />

                <div className="flex gap-x-[1.8rem] pt-4">
                  <p className="w-24 text-white text-opacity-60">Sent</p>
                  <div className="flex flex-col items-center sm:flex-row">
                    <IconAndAmount
                      value={fromInfo.value}
                      tokenAddress={fromInfo.tokenAddress}
                      chainId={fromInfo.chainID}
                      tokenSymbol={addressToSymbol({
                        tokenAddress: fromInfo.tokenAddress,
                        chainId: fromInfo.chainID,
                      })}
                      iconSize="w-4 h-4"
                      // textSize="text-sm"
                      // styledCoin={true}
                    />
                    <div className="flex flex-col items-center sm:flex-row">
                      <span className="px-2 text-white text-opacity-60">
                        on
                      </span>
                      <ChainInfo
                        chainId={fromInfo?.chainID}
                        noLink={true}
                        imgClassName="w-4 h-4 rounded-full"
                      />
                    </div>
                  </div>
                </div>
                <HorizontalDivider className="mx-[-1rem] bg-[#252537] block sm:hidden my-4" />
                <div className="flex gap-x-[1.8rem] mt-3 pb-1 items-center">
                  <p className="w-24 text-white text-opacity-60">Received</p>
                  {toInfo ? (
                    <div className="flex flex-col items-center sm:flex-row">
                      <IconAndAmount
                        value={toInfo.value}
                        tokenAddress={toInfo.tokenAddress}
                        chainId={toInfo.chainID}
                        tokenSymbol={addressToSymbol({
                          tokenAddress: toInfo.tokenAddress,
                          chainId: toInfo.chainID,
                        })}
                        iconSize="w-4 h-4"
                        // textSize="text-sm"
                        // styledCoin={true}
                      />
                      <div className="flex flex-col items-center sm:flex-row">
                        <span className="px-2 text-white text-opacity-60">
                          on
                        </span>
                        <ChainInfo
                          chainId={toInfo?.chainID}
                          noLink={true}
                          imgClassName="w-4 h-4 rounded-full"
                        />
                      </div>
                    </div>
                  ) : (
                    <p className="text-sm text-white break-all">--</p>
                  )}
                </div>
              </div>
              <br />

              <HorizontalDivider className="mx-[-1rem] bg-[#252537]" />

              <div className="flex flex-col mt-8">
                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="w-24 text-white text-opacity-60">Txn Hash</p>
                  <p className="text-sm text-white break-all">
                    <span className="hidden sm:inline">{kappa}</span>
                    <span className="sm:hidden">{truncateHash(kappa)}</span>
                  </p>
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="w-24 text-white text-opacity-60">From Txn</p>
                  {fromInfo ? (
                    <a
                      target="_blank"
                      rel="noreferrer"
                      className="text-sm text-white break-all cursor-pointer hover:text-gray-500"
                      href={
                        CHAINS_BY_ID[fromInfo.chainID]?.explorerUrl +
                        '/tx/' +
                        fromInfo.hash
                      }
                    >
                      <span className="hidden sm:inline">{fromInfo.hash}</span>
                      <span className="sm:hidden">
                        {truncateHash(fromInfo.hash)}
                      </span>
                    </a>
                  ) : (
                    <p className="text-sm text-white break-all">--</p>
                  )}
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="w-24 text-white text-opacity-60">Dest Txn</p>
                  {toInfo ? (
                    <a
                      target="_blank"
                      rel="noreferrer"
                      className="text-sm text-white break-all cursor-pointer hover:text-gray-500"
                      href={
                        CHAINS_BY_ID[toInfo.chainID]?.explorerUrl +
                        '/tx/' +
                        toInfo.hash
                      }
                    >
                      <span className="hidden sm:inline">{toInfo.hash}</span>
                      <span className="sm:hidden">
                        {truncateHash(toInfo.hash)}
                      </span>
                    </a>
                  ) : (
                    <p className="text-sm text-white break-all">--</p>
                  )}
                </div>
              </div>
            </div>
          </div>
          <br />
        </div>
      </>
    )
  } else {
    content = (
      <Error
        text="Sorry, there was a problem with that transaction hash."
        param={kappa}
        subtitle="Unknown"
      />
    )
  }

  return <StandardPageContainer>{content}</StandardPageContainer>
}

const TransactionPage = ({ queryResult }) => {
  return <BridgeTransaction queryResult={queryResult} />
}

export default TransactionPage

export const getServerSideProps = async (context) => {
  const { data } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: context.params.chainIdFrom,
      kappa: context.params.kappa,
      useMv: true,
    },
  })

  return {
    props: {
      queryResult: data,
    }, // will be passed to the page component as props
  }
}
