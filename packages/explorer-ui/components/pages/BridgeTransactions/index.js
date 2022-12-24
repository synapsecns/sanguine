import { useState, useEffect, useRef } from 'react'
import _ from 'lodash'
import { useLazyQuery } from '@apollo/client'
import { SearchIcon, XCircleIcon } from '@heroicons/react/outline'
import Image from 'next/image'
import { useSearchParams } from 'next/navigation'

import Button from '@components/tailwind/Button'

import { GET_LATEST_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'

import { Error } from '@components/Error'
import { CHAIN_INFO_MAP } from '@constants/networks'
import { Pagination } from '@components/Pagination'
import {
  AllTransactions,
  TransactionsLoader,
} from '@components/TransactionCard'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import {
  getNetworkButtonBgClassName,
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderHover,
  getNetworkButtonBorderActive,
} from '@styles/networks'
import { getChainUrl } from '@urls'
import { nameToChainIds, suggestions } from '@utils/chainAutocomplete'
import { UniversalSearch } from '@components/pages/Home/UniversalSearch'

export function BridgeTransactions() {
  const search = useSearchParams()
  const p = Number(search.get('page')) || 1

  const [page, setPage] = useState(p)
  const [transactions, setTransactions] = useState([])

  const [bridgeTransactions, { loading, error, data }] = useLazyQuery(
    GET_LATEST_BRIDGE_TRANSACTIONS_QUERY
  )

  useEffect(() => {
    if (data) {
      setTransactions(data.latestBridgeTransactions)
    }

    const num = Number(search.get('page'))

    if (num === 0) {
      setPage(1)
      bridgeTransactions({
        variables: { includePending: false, page: 1 },
      })
    } else {
      setPage(num)
      bridgeTransactions({
        variables: { includePending: false, page: num },
      })
    }
  }, [data, search])

  let content

  const nextPage = () => {
    let newPage = page + 1
    setPage(newPage)
    // setSearch({ page: newPage })

    bridgeTransactions({
      variables: { includePending: false, page: newPage },
    })
  }

  const prevPage = () => {
    if (page > 1) {
      let newPage = page - 1
      setPage(newPage)
      // setSearch({ page: newPage })
      bridgeTransactions({
        variables: { includePending: false, page: newPage },
      })
    }
  }

  const resetPage = () => {
    setPage(1)
    // setSearch({ page: 1 })
    bridgeTransactions({
      variables: { includePending: false, page: 1 },
    })
  }

  if (loading) {
    content = (
      <>
        <AutoCompleteSearch suggestions={suggestions} />
        <TransactionsLoader number={50} />
      </>
    )
  } else if (error) {
    content = <Error text="Sorry, there was a problem." subtitle="Unknown" />
  } else {
    let latestBridgeTransactions = transactions

    latestBridgeTransactions = _.orderBy(
      latestBridgeTransactions,
      'fromInfo.time',
      ['desc']
    )

    content = (
      <>
        <div className="mt-6">
          <UniversalSearch placeholder="Search all transactions by address or chain" />
          <div className="px-4 sm:px-6 lg:px-8">
            <div className="mt-8 flex flex-col">
              <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div className="inline-block min-w-full py-2 align-middle">
                  <div className="overflow-hidden shadow-sm ring-1 ring-black ring-opacity-5">
                    <table className="min-w-full">
                      <thead className="">
                        <tr>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            From
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            To
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Initial
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Final
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Origin
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Destination
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Date
                          </th>
                          <th
                            scope="col"
                            className="px-2 py-2 text-left text-md font-bold text-white"
                          >
                            Tx ID
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        <AllTransactions txns={latestBridgeTransactions} />
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <Pagination
            page={page}
            resetPage={resetPage}
            prevPage={prevPage}
            nextPage={nextPage}
          />
        </div>
      </>
    )
  }

  return (
    <StandardPageContainer title="Bridge Transactions">
      {content}
    </StandardPageContainer>
  )
}

/* TODO: future optimization, replace all this with headless ui combobox */

function AutoCompleteSearch({ suggestions }) {
  const [filteredSuggestions, setFilteredSuggestions] = useState([])
  const [activeSuggestionIndex, setActiveSuggestionIndex] = useState(0)
  const [showSuggestions, setShowSuggestions] = useState(false)
  const [input, setInput] = useState('')
  const [isValid, setValid] = useState(false)

  // const navigate = useNavigate()

  const onChange = (e) => {
    const userInput = e.target.value

    const unLinked = suggestions.filter(
      (suggestion) =>
        suggestion.toLowerCase().indexOf(userInput.toLowerCase()) > -1
    )

    if (unLinked.length === 0) {
      setValid(false)
    }

    if (suggestions.includes(e.target.value)) {
      setValid(true)
    } else {
      setValid(false)
    }

    setInput(e.target.value)
    setFilteredSuggestions(unLinked)
    setActiveSuggestionIndex(0)
    setShowSuggestions(true)
  }

  const onClick = (e) => {
    setFilteredSuggestions([])
    setInput(e.target.innerText)
    setActiveSuggestionIndex(0)
    setShowSuggestions(false)
    setValid(true)
  }

  const onKeyDown = (key) => {
    if (key.keyCode === 13 || key.keyCode === 9) {
      setInput(filteredSuggestions[activeSuggestionIndex])
      setFilteredSuggestions([])
    }
  }

  const SuggestionsListComponent = () => {
    let chainIds = filteredSuggestions.map(
      (suggestion) => nameToChainIds[suggestion]
    )

    return chainIds.length ? (
      <ul className="absolute z-10 mt-1 bg-white border border-purple-500 rounded-lg shadow-lg dark:bg-gray-700">
        {chainIds.map((suggestion, index) => {
          const itemChainId = parseInt(suggestion)

          return (
            <SelectSpecificNetworkButton
              itemChainId={itemChainId}
              onClick={onClick}
              key={index}
            />
          )
        })}
      </ul>
    ) : (
      ''
    )
  }

  return (
    <>
      <div className="flex items-center mt-5">
        <div className="relative w-full">
          <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
            <SearchIcon className="w-5 h-5 text-gray-500 dark:text-gray-400" />
          </div>
          <input
            type="text"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:outline-none focus:ring-purple-700 focus:border-purple-500 block w-full pl-10 p-2.5  dark:bg-transparent dark:border-gray-600 dark:placeholder-gray-400 dark:text-white "
            placeholder="Search by chain"
            onChange={onChange}
            onKeyDown={onKeyDown}
            value={input}
          />
          <button
            type="button"
            className="absolute inset-y-0 right-0 flex items-center pr-3"
            onClick={() => window.location.reload(false)}
          >
            <XCircleIcon
              className="w-4 h-4 text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white"
              strokeWidth={2}
            />
          </button>
          {showSuggestions && input && <SuggestionsListComponent />}
        </div>
        <button
          type="button"
          className={`p-2.5 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 ${
            !isValid ? 'pointer-events-none opacity-50' : ''
          }`}
          // onClick={() =>
          //   navigate(getChainUrl({ chainId: nameToChainIds[input] }))
          // }
        >
          <SearchIcon className="w-5 h-5" strokeWidth={1} />
        </button>
      </div>
    </>
  )
}

export function SelectSpecificNetworkButton({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
}) {
  const { chainImg, chainName } = CHAIN_INFO_MAP[itemChainId]

  const ref = useRef(null)

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  let activeClassName
  let activeTextClassName
  if (isCurrentChain) {
    activeClassName = getNetworkButtonBgClassName(itemChainId)
    activeTextClassName = ''
  } else {
    activeClassName = getNetworkButtonBgClassNameActive(itemChainId)
    activeTextClassName = `
      dark:text-gray-400
      dark:group-hover:text-gray-300
    `
  }

  return (
    <Button
      innerRef={ref}
      tabIndex={active ? '1' : '0'}
      outline={!isCurrentChain}
      className={`
        flex items-center w-full rounded-md
        !p-4
        cursor-pointer
        border border-transparent
        ${getNetworkButtonBorderHover(itemChainId)}
        ${getNetworkButtonBorderActive(itemChainId)}
        ${activeClassName}
        focus:outline-none
      `}
      onClick={onClick}
    >
      <Image
        src={chainImg}
        alt="Switch Network"
        className="w-5 h-5 mr-2 rounded-md"
      />
      <div
        className={`
          text-primary font-medium ${activeTextClassName}
        `}
      >
        {chainName}
      </div>
    </Button>
  )
}
