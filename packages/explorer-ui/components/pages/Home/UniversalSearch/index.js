import _ from 'lodash'
import {useState} from 'react'

import {validateAndParseAddress} from '@utils/validateAndParseAddress'
import {validateAndParseHash} from '@utils/validateAndParseHash'

import {ChainId} from '@constants/networks'

import {SearchBox} from './SearchBox'

export function UniversalSearch({ placeholder }) {
  const [searchField, setSearchField] = useState('')
  const [showText, setShowText] = useState(false)

  let isValid
  let error
  let inputType
  let searchLink
  if (!searchField || searchField === '') {
    error = 'Field cannot be empty.'
  } else if (validateAndParseAddress(searchField)) {
    isValid = true
    inputType = 'ADDRESS'
    searchLink = '/txs?account=' + searchField
  } else if (validateAndParseHash(searchField)) {
    isValid = true
    inputType = 'TRANSACTION'
    // searchLink = '/tx/' + searchField
  } else if (_.values(ChainId).includes(searchField)) {
    isValid = true
    inputType = 'CHAIN'
    searchLink = '/txs?chainId=' + searchField
  } else {
    error = 'Not a valid address or transaction hash'
  }

  return (
    <>
      <div className="border-y border-white border-opacity-10 ">
        <div className="flex justify-center items-center p-2 gap-x-4 py-6">
          <h3
            className="text-white flex items-center mr-4"
            onClick={() => setShowText(!showText)}
          >
            {!showText ? (
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="currentColor"
                className="w-4 h-4 mr-2"
              >
                <path
                  fillRule="evenodd"
                  d="M4.72 3.97a.75.75 0 011.06 0l7.5 7.5a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 01-1.06-1.06L11.69 12 4.72 5.03a.75.75 0 010-1.06zm6 0a.75.75 0 011.06 0l7.5 7.5a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 11-1.06-1.06L17.69 12l-6.97-6.97a.75.75 0 010-1.06z"
                  clipRule="evenodd"
                />
              </svg>
            ) : (
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="currentColor"
                className="w-4 h-4 mr-2"
              >
                <path
                  fillRule="evenodd"
                  d="M20.03 4.72a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 01-1.06 0l-7.5-7.5a.75.75 0 011.06-1.06L12 11.69l6.97-6.97a.75.75 0 011.06 0zm0 6a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 01-1.06 0l-7.5-7.5a.75.75 0 111.06-1.06L12 17.69l6.97-6.97a.75.75 0 011.06 0z"
                  clipRule="evenodd"
                />
              </svg>
            )}
            Filters
          </h3>
          <div className="grow">
            <SearchBox
              searchField={searchField}
              setSearchField={setSearchField}
              inputType={inputType}
              placeholder={placeholder}
            />
          </div>
          <button className="font-medium rounded-r-md border border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 px-4 py-2">
            <a href={searchLink}>
              Search
            </a>
          </button>
          <div className="">
            <button className="font-medium rounded-l-md text-white border  border-[#BE78FF] bg-synapse-radial px-4 py-2">
              Confirmed
            </button>
            <button className="font-medium rounded-r-md border border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 px-4 py-2">
              Pending
            </button>
          </div>
          {/* {!isValid && error ? (
          <div
          className="absolute  font-medium p-4 mt-1 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800"
            role="alert"
          >
            {error}
          </div>
        ) : (
          ""
        )} */}
        </div>
        {showText ? (
          <div>
            {/* THIS IS WALLET ADDRESS */}
            <div className="flex justify-center items-center p-2 gap-x-4 py-4">
              <h3
                className="text-white flex items-center mr-10"
                onClick={() => setShowText(!showText)}
              >
                Wallet
              </h3>
              <div className="grow">
                <form className="flex items-center">
                  <div className="relative w-full group">
                    <input
                      type="text"
                      id="simple-search"
                      className={`
                        bg-white bg-opacity-5
                        rounded-md
                        border border-white border-opacity-20
                        focus:outline-none focus-within:border-gray-500
                        block w-full  px-4 py-2
                        text-white
                        placeholder:text-white placeholder:text-opacity-60
                      `}
                      placeholder="Wallet Address"
                      onChange={(e) => {
                        setSearchField(e.target.value)
                      }}
                      value={searchField}
                    />
                  </div>
                </form>
              </div>
              <input
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                checked
              />
              <h3 className="text-white font-semibold">To</h3>
              <input
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                checked
              />
              <h3 className="text-white font-semibold">From</h3>
            </div>
            {/* THIS IS MIN/MAX SIZE */}
            <div className="flex justify-center items-center p-2 gap-x-4 py-4">
              <h3
                className="text-white flex items-center mr-6"
                onClick={() => setShowText(!showText)}
              >
                Min Size
              </h3>
              <div className="grow mr-12">
                <form className="flex items-center">
                  <div className="relative w-full group">
                    <input
                      type="text"
                      id="simple-search"
                      className={`
                        bg-white bg-opacity-5
                        rounded-md
                        border border-white border-opacity-20
                        focus:outline-none focus-within:border-gray-500
                        block w-full  px-4 py-2
                        text-white
                        placeholder:text-white placeholder:text-opacity-60
                      `}
                      placeholder="Min Size"
                      onChange={(e) => {
                        setSearchField(e.target.value)
                      }}
                      value={searchField}
                    />
                  </div>
                </form>
              </div>
              <h3
                className="text-white flex items-center"
                onClick={() => setShowText(!showText)}
              >
                Max Size
              </h3>
              <div className="grow">
                <form className="flex items-center">
                  <div className="relative w-full group ">
                    <input
                      type="text"
                      id="simple-search"
                      className={`
                        bg-white bg-opacity-5
                        rounded-md
                        border border-white border-opacity-20
                        focus:outline-none focus-within:border-gray-500
                        block w-full  px-4 py-2
                        text-white
                        placeholder:text-white placeholder:text-opacity-60
                      `}
                      placeholder="Max Size"
                      onChange={(e) => {
                        setSearchField(e.target.value)
                      }}
                      value={searchField}
                    />
                  </div>
                </form>
              </div>
            </div>
            {/* THIS IS START/DATE */}
            <div className="flex justify-center items-center p-2 gap-x-4 py-4">
              <h3
                className="text-white flex items-center mr-2"
                onClick={() => setShowText(!showText)}
              >
                Start date
              </h3>
              <div className="grow mr-12">
                <form className="flex items-center">
                  <div className="relative w-full group">
                    <input
                      type="text"
                      id="simple-search"
                      className={`
                        bg-white bg-opacity-5
                        rounded-md
                        border border-white border-opacity-20
                        focus:outline-none focus-within:border-gray-500
                        block w-full  px-4 py-2
                        text-white
                        placeholder:text-white placeholder:text-opacity-60
                      `}
                      placeholder="yyyy-mm-dd"
                      onChange={(e) => {
                        setSearchField(e.target.value)
                      }}
                      value={searchField}
                    />
                  </div>
                </form>
              </div>
              <h3
                className="text-white flex items-center"
                onClick={() => setShowText(!showText)}
              >
                End date
              </h3>
              <div className="grow">
                <form className="flex items-center">
                  <div className="relative w-full group ">
                    <input
                      type="text"
                      id="simple-search"
                      className={`
                        bg-white bg-opacity-5
                        rounded-md
                        border border-white border-opacity-20
                        focus:outline-none focus-within:border-gray-500
                        block w-full  px-4 py-2
                        text-white
                        placeholder:text-white placeholder:text-opacity-60
                      `}
                      placeholder="yyyy-mm-dd"
                      onChange={(e) => {
                        setSearchField(e.target.value)
                      }}
                      value={searchField}
                    />
                  </div>
                </form>
              </div>
            </div>
            {/* THIS IS BUTTONS */}
            <div className="flex items-center p-2 gap-x-4 mb-3">
              <button className="font-medium rounded-[4px] text-white bg-[#333333] px-4 py-2">
                Apply
              </button>
              <button className="font-medium rounded-[4px] text-[#333333]  bg-opacity-40 bg-[#333333] px-4 py-2">
                Reset
              </button>
            </div>
          </div>
        ) : null}
      </div>
    </>
  )
}
