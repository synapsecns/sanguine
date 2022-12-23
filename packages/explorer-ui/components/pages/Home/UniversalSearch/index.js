import _ from 'lodash'
import { useState } from 'react'

import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { validateAndParseHash } from '@utils/validateAndParseHash'

import { ChainId } from '@constants/networks'

import { SearchBox } from './SearchBox'

import { SearchResults } from './SearchResults'

export function UniversalSearch() {
  const [searchField, setSearchField] = useState('')

  let isValid
  let error
  let inputType
  if (!searchField || searchField === '') {
    error = 'Field cannot be empty.'
  } else if (validateAndParseAddress(searchField)) {
    isValid = true
    inputType = 'ADDRESS'
  } else if (validateAndParseHash(searchField)) {
    isValid = true
    inputType = 'TRANSACTION'
  } else if (_.values(ChainId).includes(searchField)) {
    isValid = true
    inputType = 'CHAIN'
  } else {
    error = 'Not a valid address or transaction hash'
  }

  return (
    <>
      <div className="flex justify-center items-center p-2 gap-x-4 border-y border-white border-opacity-10 py-6">
        <h3 className="text-white flex items-center mr-6">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            className="w-4 h-4 mr-1"
          >
            <path
              fillRule="evenodd"
              d="M4.72 3.97a.75.75 0 011.06 0l7.5 7.5a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 01-1.06-1.06L11.69 12 4.72 5.03a.75.75 0 010-1.06zm6 0a.75.75 0 011.06 0l7.5 7.5a.75.75 0 010 1.06l-7.5 7.5a.75.75 0 11-1.06-1.06L17.69 12l-6.97-6.97a.75.75 0 010-1.06z"
              clipRule="evenodd"
            />
          </svg>
          Filters
        </h3>
        <div className="grow">
          <SearchBox
            searchField={searchField}
            setSearchField={setSearchField}
            inputType={inputType}
          />
          <SearchResults searchField={searchField} inputType={inputType} />
        </div>
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
    </>
  )
}
