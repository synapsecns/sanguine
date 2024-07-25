import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { TRANSACTIONS_PATH, ACCOUNTS_PATH } from '@urls'
import { ChainInfo } from '@components/misc/ChainInfo'
import { Error } from '@components/Error'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { useRouter } from 'next/router'
import { useSearchParams } from 'next/navigation'
import { CHAINS } from 'synapse-constants'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { API_URL } from '@graphql'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { formatDateTimestamp } from '@utils/formatDate'
import { IconAndAmount } from '@components/misc/IconAndAmount'

const CHAINS_BY_ID = CHAINS.CHAINS_BY_ID
const CCTP_CONTRACTS = CHAINS.CCTP_CONTRACTS
const BRIDGE_CONTRACTS = CHAINS.BRIDGE_CONTRACTS
const FASTBRIDGE_CONTRACTS = CHAINS.FASTBRIDGE_CONTRACTS

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link,
  cache: new InMemoryCache(),
})

function truncateAddress(address) {
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

function truncateHash(hash) {
  return `${hash.slice(0, 8)}...${hash.slice(-5)}`
}

export default function BridgeTransaction({ queryResult }) {
  const router = useRouter()
  const search = useSearchParams()
  const { kappa } = router.query
  const chainId = Number(search.get('chainIdFrom'))
  const handlePending = (date) => {
    const now = new Date().getTime()
    const timeDiff = now - date * 1000
    if (timeDiff > 86400000) {
      return 'Indexing'
    } else {
      return 'Pending'
    }
  }
  const transaction = queryResult.bridgeTransactions[0]
  const { pending, fromInfo, toInfo } = transaction

  const getTimeDifference = (start, end) => {
    const diff = end - start
    if (0 >= diff) {
      return '1'
    }
    return diff.toString()
  }
  let content
  const pendingContent = handlePending(fromInfo?.time)

  if (!!transaction) {
    content = (
      <>
        <div className="flex flex-row mt-5 mb-1 text-white pb-12">
          <a
            className="text-[#99E6FF] cursor-pointer mr-2"
            href={TRANSACTIONS_PATH}
          >
            ← Explorer
          </a>
          <span className="text-gray-500 mr-2">/</span>
          <span
            className="text-white hover:text-gray-500 cursor-pointer"
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
          <div className="flex px-4 gap-y-2 flex-col">
            <div className="flex mt-4 flex-col">
              <div className="flex flex-col">
                <div className="flex justify-between items-center w-full">
                  <div className="flex gap-x-[1.8rem] py-1">
                    <p className="text-white text-opacity-60 w-24">Timestamp</p>
                    <p className="text-white">
                      {formatDateTimestamp(new Date(fromInfo.time * 1000))}
                    </p>
                  </div>
                  <div className="flex gap-x-[1.8rem] py-1">
                    <p className="text-white hidden sm:block">
                      {new Date(fromInfo.time * 1000)
                        .toISOString()
                        .replace('T', ' ')
                        .slice(0, 16) + ' UTC'}
                    </p>
                  </div>
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="text-white text-opacity-60 w-24">Status</p>
                  <p className="text-white ">
                    {pending ? 'Pending' : 'Confirmed'}{' '}
                  </p>
                </div>
                <div className="flex gap-x-[1.8rem] py-1 pb-4">
                  <p className="text-white text-opacity-60 w-24">Elapsed</p>
                  <p className="text-white ">
                    {toInfo
                      ? getTimeDifference(fromInfo.time, toInfo.time) +
                        ' seconds'
                      : '--'}{' '}
                  </p>
                </div>
                <HorizontalDivider className="mx-[-1rem] bg-[#252537]" />
                <div className="flex gap-x-[1.8rem] py-1 pt-4">
                  <p className="text-white text-opacity-60 w-24">From</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm hover:text-gray-500 cursor-pointer"
                    href={ACCOUNTS_PATH + '/' + fromInfo.address}
                  >
                    <span className="hidden sm:inline">{fromInfo.address}</span>
                    <span className="sm:hidden">
                      {truncateAddress(fromInfo.address)}
                    </span>
                  </a>
                </div>
                <div className="flex gap-x-[1.8rem] py-1 pb-4">
                  <p className="text-white text-opacity-60 w-24">To</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm hover:text-gray-500 cursor-pointer"
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
                  <p className="text-white text-opacity-60 w-24">Sent</p>
                  <div className="flex flex-col sm:flex-row items-center">
                    <IconAndAmount
                      formattedValue={fromInfo.formattedValue}
                      tokenAddress={fromInfo.tokenAddress}
                      chainId={fromInfo.chainID}
                      tokenSymbol={fromInfo.tokenSymbol}
                      iconSize="w-4 h-4"
                      textSize="text-sm"
                      styledCoin={true}
                    />
                    <div className="flex flex-col sm:flex-row items-center">
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
                  <p className="text-white text-opacity-60 w-24">Received</p>
                  {toInfo ? (
                    <div className="flex flex-col sm:flex-row items-center">
                      <IconAndAmount
                        formattedValue={toInfo.formattedValue}
                        tokenAddress={toInfo.tokenAddress}
                        chainId={toInfo.chainID}
                        tokenSymbol={toInfo.tokenSymbol}
                        iconSize="w-4 h-4"
                        textSize="text-sm"
                        styledCoin={true}
                      />
                      <div className="flex flex-col sm:flex-row items-center">
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
                    <p className="text-white break-all text-sm">--</p>
                  )}
                </div>
              </div>
              <br />

              <HorizontalDivider className="mx-[-1rem] bg-[#252537]" />

              <div className="flex mt-8 flex-col">
                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="text-white text-opacity-60 w-24">Txn Hash</p>
                  <p className="text-white break-all text-sm">
                    <span className="hidden sm:inline">{kappa}</span>
                    <span className="sm:hidden">{truncateHash(kappa)}</span>
                  </p>
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="text-white text-opacity-60 w-24">From Txn</p>
                  {fromInfo ? (
                    <a
                      target="_blank"
                      rel="noreferrer"
                      className="text-white break-all text-sm hover:text-gray-500 cursor-pointer"
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
                    <p className="text-white break-all text-sm">--</p>
                  )}
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="text-white text-opacity-60 w-24">Dest Txn</p>
                  {toInfo ? (
                    <a
                      target="_blank"
                      rel="noreferrer"
                      className="text-white break-all text-sm hover:text-gray-500 cursor-pointer"
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
                    <p className="text-white break-all text-sm">--</p>
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
export async function getServerSideProps(context) {
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
