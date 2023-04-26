import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'

import { approveAndDeposit } from '@utils/actions/approveAndDeposit'
import TokenInput from '@components/TokenInput'
import PriceImpactDisplay from './PriceImpactDisplay'
import { TransactionResponse } from '@ethersproject/providers'

import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { Zero } from '@ethersproject/constants'
const PoolManagementDeposit = ({ poolName, chainId, address }) => {
  // const {
  //   onChangeTokenInputValue,
  //   clearInputs,
  //   priceImpact,
  //   poolTokens,
  //   inputState,
  //   tokenInputSum,
  //   depositAmount,
  // } = useSwapPoolDeposit(poolName)
  const onChangeTokenInputValue = () => console.log('onChangeTokenInputValue')
  const clearInputs = ''
  const priceImpact = ''
  const poolTokens = []
  const inputState = {}
  const tokenInputSum = Zero
  const depositAmount = ''

  const placeholder = async (): Promise<TransactionResponse> => {
    console.log('placeholder')
    return
  }
  return (
    <div className="flex-col">
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {poolTokens.map((token, i) => (
          <TokenInputWithBalance
            token={token}
            inputValue={inputState[token.symbol]}
            onChangeTokenInputValue={onChangeTokenInputValue}
            key={i}
            address={address}
            chainId={chainId}
          />
        ))}
      </div>
      <TransactionButton
        label="Add Liquidity"
        pendingLabel="Adding Liquidity"
        disabled={tokenInputSum.eq(0)}
        className="items-center w-full px-6 py-3 mt-6 text-md rounded-xl"
        // FIX
        // onClick={async () => {
        //   const appAndDeposit = await approveAndDeposit({
        //     slippageCustom: null,
        //     slippageSelected: 'ONE_TENTH',
        //     infiniteApproval: true,
        //     inputState,
        //     depositAmount,
        //   })
        //   // Clear input after deposit
        //   clearInputs()
        // }}
        onClick={placeholder}
      />
      <PriceImpactDisplay priceImpact={priceImpact} />
      {/* {poolStakingLink && (
        <div className="pb-4">
          <PoolStakingButton
            poolName={poolName}
            poolStakingLink={poolStakingLink}
            poolStakingLinkText={poolStakingLinkText}
          />
        </div>
      )} */}
    </div>
  )
}

function TokenInputWithBalance({
  token,
  inputValue,
  onChangeTokenInputValue,
  address,
  chainId,
}) {
  let balanceToken
  if (token.symbol == WETH.symbol) {
    balanceToken = ETH
  } else if (token.symbol == AVWETH.symbol) {
    // token = WETHE
    balanceToken = WETHE
  } else {
    balanceToken = token
  }
  // ADDDM<E
  // const balance = useTokenBalance(balanceToken)
  const balance = 0
  return (
    <TokenInput
      token={balanceToken}
      key={balanceToken.symbol}
      max={String(balance)}
      inputValue={inputValue}
      onChange={(value) => onChangeTokenInputValue(token.symbol, value)}
      // disabled={false}
      chainId={chainId}
      address={address}
    />
  )
}
export default PoolManagementDeposit
