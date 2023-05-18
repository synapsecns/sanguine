import _ from 'lodash'
import { useState } from 'react'
import TextField from '@mui/material/TextField'
import {
  inputStyle,
  dateInputStyle,
  comboSelectStyle,
  comboSelectStyleSmall,
  inputStyleRounded,
} from '@utils/styles/muiStyles'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { validateAndParseHash } from '@utils/validateAndParseHash'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'
import Box from '@mui/material/Box'
import InputLabel from '@mui/material/InputLabel'
import MenuItem from '@mui/material/MenuItem'
import FormControl from '@mui/material/FormControl'
import Select from '@mui/material/Select'
import { ChainId } from '@constants/networks'
import dayjs from 'dayjs'
import { CHAIN_ENUM_BY_ID } from '@constants/networks'

import { SearchBox } from './SearchBox'

export function UniversalSearch({
  placeholder,
  setPending,
  pending,
  loading,
  setWallet,
  wallet,
  setMinSize,
  minSize,
  setMaxSize,
  maxSize,

  setStartDate,
  startDate,
  setEndDate,
  endDate,
  setToTx,
  toTx,
  setFromTx,
  fromTx,
  setKappa,
  kappa,
  executeSearch,
  chains,
  setChains,
  tokens,
  setTokens,
  chainsLocale,
  setChainsLocale,
  walletLocale,
  setWalletLocale,
}) {
  // const search = useSearchParams()

  const [searchField, setSearchField] = useState('')
  const [showText, setShowText] = useState(false)
  // const [startDate, setStartDate] = useState("s");
  const handleChains = (event) => {
    const {
      target: { value },
    } = event
    setChains(
      // On autofill we get a stringified value.
      typeof value === 'string' ? value.split(',') : value
    )
  }

  const handleTokens = (event) => {
    const {
      target: { value },
    } = event
    setTokens(
      // On autofill we get a stringified value.
      typeof value === 'string' ? value.split(',') : value
    )
  }
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 hover:text-white'
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'
  const resetFields = () => {
    setWallet('')
    setMinSize({ type: 'USD', value: '' })
    setMaxSize({ type: 'USD', value: '' })
    setStartDate(null)
    setEndDate(null)
    setToTx(true)
    setFromTx(true)
    setChains([])
    setKappa('')
  }
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
    // @ts-ignore
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
            <TextField
              size="small"
              value={kappa}
              onChange={(e) => {
                setKappa(e.target.value)
              }}
              id="outlined-basic"
              label="Search by TXID / TXHash"
              variant="outlined"
              sx={inputStyle}
            />
          </div>
          <button
            onClick={() => executeSearch()}
            className={
              'font-medium rounded-md border border-l-0 border-gray-700 text-white bg-gray-700  px-4 py-1 hover:bg-opacity-70 ease-in-out duration-200 ml-[-105px] pointer-cursor z-10' +
              (loading ? ' pointer-events-none opacity-[0.4]' : '')
            }
          >
            Search
          </button>
          {/* <button onClick={() => executeSearch()} className="font-medium rounded-md border border-l-0 border-gray-700 text-white bg-gray-700  px-4 py-2 hover:bg-opacity-70 ease-in-out duration-200">
            <a href={searchLink}>Search</a>
          </button> */}
          <div className="">
            <button
              disabled={loading}
              onClick={() => setPending(false)}
              className={
                'font-medium rounded-l-md px-4 py-2 border ' +
                (pending ? unSelectStyle : selectStyle) +
                (loading ? ' pointer-events-none' : '')
              }
            >
              Confirmed
            </button>
            <button
              disabled={loading}
              onClick={() => setPending(true)}
              className={
                'font-medium rounded-r-md px-4 py-2 border ' +
                (pending ? selectStyle : unSelectStyle) +
                (loading ? ' pointer-events-none' : '')
              }
            >
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
                <TextField
                  size="small"
                  value={wallet}
                  onChange={(e) => {
                    setWallet(e.target.value)
                  }}
                  id="outlined-basic"
                  label="Wallet Address"
                  variant="outlined"
                  sx={inputStyle}
                />
              </div>
              <div className="flex justify-center rounded-md border-l-0 border-gray-700 border-opacity-70 bg-[#333333]  bg-opacity-30  py-[9px] px-3">
                <div
                  className="form-check form-check-inline mx-1"
                  onClick={() => setWalletLocale(!walletLocale)}
                >
                  <input
                    checked={walletLocale}
                    className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-purple-300 checked:bg-opacity-90 checked:border-purple-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer"
                    type="checkbox"
                    id="walletFrom"
                    value="option1"
                  />
                  <label
                    className="form-check-label inline-block text-gray-500 "
                    htmlFor="walletFrom"
                  >
                    From
                  </label>
                </div>
                <div
                  className="form-check form-check-inline mx-1"
                  onClick={() => setWalletLocale(!walletLocale)}
                >
                  <input
                    checked={!walletLocale}
                    className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white  checked:bg-purple-300 checked:bg-opacity-90 checked:border-purple-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer"
                    type="checkbox"
                    id="walletTo"
                    value="option2"
                  />
                  <label
                    className="form-check-label inline-block text-gray-500  "
                    htmlFor="walletTo"
                  >
                    To
                  </label>
                </div>
              </div>

              {/* <input
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                checked={toTx}
                onClick={() => { setToTx(!toTx) }}

              />
              <h3 className="text-white font-semibold">To</h3>
              <input
                type="checkbox"
                className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                checked={fromTx}
                onClick={() => { setFromTx(!fromTx) }}
              />
              <h3 className="text-white font-semibold">From</h3> */}
            </div>

            {/* THIS IS MIN/MAX SIZE */}
            <div className="flex justify-center items-center p-2 gap-x-14 py-4">
              <h3
                className="text-white flex items-center mr-1"
                onClick={() => setShowText(!showText)}
              >
                Chain
              </h3>
              <div className="grow">
                <div className="flex flex-row items-center ">
                  <TextField
                    select
                    id="Chains"
                    name="Chains"
                    variant="outlined"
                    label="Chains"
                    size="small"
                    value={chains}
                    sx={comboSelectStyle}
                    SelectProps={{
                      multiple: true,
                      onChange: (e) => handleChains(e),
                    }}
                  >
                    {Object.values(CHAIN_ENUM_BY_ID).map((chain) => (
                      <MenuItem key={chain} value={chain}>
                        {chain.charAt(0).toUpperCase() + chain.slice(1)}
                      </MenuItem>
                    ))}
                  </TextField>

                  <div className="ml-4 w-fit flex justify-center rounded-md border-l-0 border-gray-700 border-opacity-70 bg-[#333333]  bg-opacity-30 py-[9px] px-3">
                    <div
                      className="form-check form-check-inline mx-2"
                      onClick={() => setChainsLocale(!chainsLocale)}
                    >
                      <input
                        checked={chainsLocale}
                        className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-purple-300 checked:bg-opacity-90 checked:border-purple-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer"
                        type="checkbox"
                        id="walletFrom"
                        value="option1"
                      />
                      <label
                        className="form-check-label inline-block text-gray-500 "
                        htmlFor="walletFrom"
                      >
                        From
                      </label>
                    </div>
                    <div
                      className="form-check form-check-inline mx-2"
                      onClick={() => setChainsLocale(!chainsLocale)}
                    >
                      <input
                        checked={!chainsLocale}
                        className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-purple-300 checked:bg-opacity-90 checked:border-purple-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer"
                        type="checkbox"
                        id="walletTo"
                        value="option2"
                      />
                      <label
                        className="form-check-label inline-block text-gray-500  "
                        htmlFor="walletTo"
                      >
                        To
                      </label>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            {/* <div className="flex justify-center items-center p-2 gap-x-14 py-4">
              <h3
                className="text-white flex items-center mr-1"
                onClick={() => setShowText(!showText)}
              >
                Token
              </h3>
              <div className="grow">
                <div className="flex flex-row items-center ">
                <TextField size="small" value={token} onChange={(e) => {
                  setToken(e.target.value)
                }} id="outlined-basic" label="Token Address" variant="outlined" sx={inputStyle} />


                </div>
              </div>
            </div> */}
            {/* THIS IS MIN/MAX SIZE */}
            <div className="flex justify-center items-center p-2 gap-x-6 py-4">
              <h3
                className="text-white flex items-center mr-6"
                onClick={() => setShowText(!showText)}
              >
                Volume
              </h3>
              <div className="flex flex-row w-full justify-between">
                <div className="w-[49%] flex flex-row ">
                  <div className="w-[80%]">
                    <TextField
                      type="number"
                      size="small"
                      value={minSize.value}
                      onChange={(e) => {
                        setMinSize({ ...minSize, value: e.target.value })
                      }}
                      id="outlined-basic"
                      label="Min Size"
                      variant="outlined"
                      sx={inputStyleRounded}
                    />
                  </div>
                  <div className="w-[20%]">
                    <TextField
                      select
                      id="UnitsMin"
                      name="UnitsMin"
                      variant="outlined"
                      label="Units"
                      size="small"
                      sx={comboSelectStyleSmall}
                      SelectProps={{
                        value: minSize.type,
                        onChange: (e) =>
                          setMinSize({ ...minSize, type: e.target.value }),
                      }}
                    >
                      <MenuItem key="USD" value="USD" selected>
                        USD
                      </MenuItem>
                      {/* <MenuItem
                        key="Amount"
                        value="Amount"
                      >
                        Amount
                      </MenuItem> */}
                    </TextField>
                  </div>
                </div>
                <div className="w-[49%] flex flex-row">
                  <div className="w-[80%]">
                    <TextField
                      size="small"
                      type="number"
                      onChange={(e) => {
                        setMaxSize({ ...maxSize, value: e.target.value })
                      }}
                      value={maxSize.value}
                      id="outlined-basic"
                      label="Max Size"
                      variant="outlined"
                      sx={inputStyleRounded}
                    />
                  </div>
                  <div className="w-[20%]">
                    <TextField
                      select
                      id="UnitsMax"
                      name="UnitsMax"
                      variant="outlined"
                      label="Units"
                      size="small"
                      sx={comboSelectStyleSmall}
                      SelectProps={{
                        value: maxSize.type,
                        onChange: (e) =>
                          setMaxSize({ ...maxSize, type: e.target.value }),
                      }}
                    >
                      <MenuItem key="USD" value="USD">
                        USD
                      </MenuItem>
                      {/* <MenuItem
                        key="Amount"
                        value="Amount"
                      >
                        Amount
                      </MenuItem> */}
                    </TextField>
                  </div>
                </div>
              </div>
            </div>
            {/* THIS IS START/DATE */}
            <div className="flex justify-center items-center p-2 gap-x-14 py-4">
              <h3
                className="text-white flex items-center mr-2"
                onClick={() => setShowText(!showText)}
              >
                Time
              </h3>
              <div className="flex flex-row w-full justify-between">
                <div className="w-[49%]">
                  <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DatePicker
                      label="Start Date"
                      value={startDate}
                      onChange={(newValue) => {
                        setStartDate(newValue)
                      }}
                      renderInput={(params) => (
                        <TextField
                          size="small"
                          sx={dateInputStyle}
                          {...params}
                        />
                      )}
                    />
                  </LocalizationProvider>
                </div>
                <div className="w-[49%]">
                  <LocalizationProvider
                    sx={{ width: '100%' }}
                    dateAdapter={AdapterDayjs}
                  >
                    <DatePicker
                      label="End Date"
                      value={endDate}
                      onChange={(newValue) => {
                        setEndDate(newValue)
                      }}
                      renderInput={(params) => (
                        <TextField
                          size="small"
                          sx={dateInputStyle}
                          {...params}
                        />
                      )}
                    />
                  </LocalizationProvider>
                </div>
              </div>{' '}
            </div>
            {/* THIS IS BUTTONS */}
            <div className="flex items-center p-2 gap-x-4 mb-3">
              <button
                className="font-medium rounded-md border border-l-0 border-gray-700 text-white bg-gray-700  px-4 py-2 hover:bg-opacity-70 ease-in-out duration-200"
                onClick={() => resetFields()}
              >
                Reset
              </button>
            </div>
          </div>
        ) : null}
      </div>
    </>
  )
}
