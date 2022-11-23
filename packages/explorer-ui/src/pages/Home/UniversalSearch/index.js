import _ from "lodash"
import { useState } from 'react'

import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { validateAndParseHash } from '@utils/validateAndParseHash'

import { ChainId } from '@constants/networks'

import { SearchBox } from './SearchBox'

import { SearchResults } from './SearchResults'


export function UniversalSearch() {
  const [searchField, setSearchField] = useState("")

  let isValid
  let error
  let inputType
  if (
    (!searchField)
    || (searchField === "")
  ) {
    error = "Field cannot be empty."
  } else if (validateAndParseAddress(searchField)) {
    isValid = true
    inputType = "ADDRESS"
  } else if (validateAndParseHash(searchField)) {
    isValid = true
    inputType = "TRANSACTION"
  } else if (_.values(ChainId).includes(searchField) ) {
    isValid = true
    inputType = "CHAIN"
  } else {
    error = "Not a valid address or transaction hash"
  }

  return (
    <>
      <SearchBox
        searchField={searchField}
        setSearchField={setSearchField}
        inputType={inputType}
      />
      <SearchResults
        searchField={searchField}
        inputType={inputType}
      />
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
    </>
  )
}
