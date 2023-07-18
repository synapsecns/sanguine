import { CHAINS_BY_ID } from '@/constants/chains'
import {
  clearTokenSelectorState,
  resetTokenSelectorState,
  setSelectFromChainId,
  setSelectFromToken,
  setSelectToChainId,
  setSelectToToken,
} from '@/slices/tokenSelectorSlice'
import { RootState } from '@/store/store'
import { useSelector } from 'react-redux'
import { useDispatch } from 'react-redux'
import Select from 'react-select'

const ExampleTokenSelector = () => {
  const {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    toChainIds,
    fromTokens,
    toTokens,
  } = useSelector((state: RootState) => state.tokenSelector)

  const dispatch = useDispatch()

  const fromChainOptions = fromChainIds.map((option) => ({
    label: (
      <span className="flex items-center space-x-1">
        <img src={CHAINS_BY_ID[option].chainImg.src} className="w-5 h-5" />
        <div>
          {CHAINS_BY_ID[option].name} [{option}]
        </div>
      </span>
    ),
    value: option,
  }))

  const fromTokenOptions = fromTokens.map((option) => ({
    label: option,
    value: option,
  }))

  const toChainOptions = toChainIds.map((option) => ({
    label: (
      <span className="flex items-center space-x-1">
        <img src={CHAINS_BY_ID[option].chainImg.src} className="w-5 h-5" />
        <div>
          {CHAINS_BY_ID[option].name} [{option}]
        </div>
      </span>
    ),
    value: option,
  }))

  const toTokenOptions = toTokens.map((option) => ({
    label: option,
    value: option,
  }))

  const handleFromChainChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectFromChainId(Number(selectedOption.value)))
    } else {
      dispatch(setSelectFromChainId(null))
    }
  }

  const handleFromTokenChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectFromToken(selectedOption.value))
    } else {
      dispatch(setSelectFromToken(null))
    }
  }

  const handleToChainChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectToChainId(Number(selectedOption.value)))
    } else {
      dispatch(setSelectToChainId(null))
    }
  }

  const handleToTokenChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectToToken(selectedOption.value))
    } else {
      dispatch(setSelectToToken(null))
    }
  }

  const customFilterOption = (option, rawInput) => {
    const searchTerm = rawInput.toLowerCase()

    return (
      option.data.label.props.children[1].props.children[0]
        .toLowerCase()
        .includes(searchTerm) ||
      option.value.toString().toLowerCase().includes(searchTerm)
    )
  }

  return (
    <div className="flex flex-col items-center p-2 justify-ceneter">
      <div className="">
        <h1 className="mb-2 text-xl text-bold">
          Example chain & token selector{' '}
        </h1>
        <div>
          <Select
            key={fromChainId}
            options={fromChainOptions}
            onChange={handleFromChainChange}
            isSearchable={true}
            isClearable={true}
            filterOption={customFilterOption}
            placeholder="Search and select origin chain"
            value={fromChainOptions.find(
              (option) => Number(option.value) === fromChainId
            )}
          />
        </div>

        <div>
          <Select
            key={fromToken}
            options={fromTokenOptions}
            onChange={handleFromTokenChange}
            isSearchable={true}
            isClearable={true}
            placeholder="Search and select token you want to bridge"
            value={fromTokenOptions.find(
              (option) => option.value === fromToken
            )}
          />
        </div>

        <div>
          <Select
            key={toChainId}
            options={toChainOptions}
            filterOption={customFilterOption}
            onChange={handleToChainChange}
            isSearchable={true}
            isClearable={true}
            placeholder="Search and select destination chain"
            value={toChainOptions.find(
              (option) => Number(option.value) === toChainId
            )}
          />
        </div>

        <div>
          <Select
            key={toToken}
            options={toTokenOptions}
            onChange={handleToTokenChange}
            isSearchable={true}
            isClearable={true}
            placeholder="Search and select token you want to receive"
            value={toTokenOptions.find((option) => option.value === toToken)}
          />
        </div>

        <button
          className="p-2 bg-blue-200 border-2"
          onClick={() => dispatch(clearTokenSelectorState())}
        >
          Clear everything
        </button>

        <button
          className="p-2 bg-purple-300 border-2"
          onClick={() => dispatch(resetTokenSelectorState())}
        >
          Get Initial Defaults
        </button>

        <BridgeNote
          fromToken={fromToken}
          fromChainId={fromChainId}
          toToken={toToken}
          toChainId={toChainId}
        />
      </div>
    </div>
  )
}

const BridgeNote = ({ fromToken, fromChainId, toToken, toChainId }) => {
  return (
    <div className="mt-2">
      You are bridging [{fromToken !== null ? fromToken : 'SELECT'}] on Chain [
      {fromChainId !== null ? fromChainId : 'SELECT'}] to [
      {toToken !== null ? toToken : 'SELECT'}] on Chain [
      {toChainId !== null ? toChainId : 'SELECT'}]
    </div>
  )
}

export default ExampleTokenSelector
