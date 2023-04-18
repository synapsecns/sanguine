import _ from 'lodash'

import { AVWETH, ETH, WETH, WETHE } from '@constants/tokens/basic'

import { useTokenBalance } from '@hooks/tokens/useTokenBalances'
import { useSwapPoolDeposit } from '@hooks/pools/useSwapPoolDeposit'
import { useApproveAndDeposit } from '@hooks/actions/useApproveAndDeposit'

import Button from '@tw/Button'
import TokenInput from '@components/TokenInput'
import PoolStakingButton from '@components/buttons/PoolStakingButton'

import PriceImpactDisplay from './PriceImpactDisplay'

import { TransactionButton } from '@components/buttons/SubmitTxButton'

export default function PoolManagementDeposit({
  poolName,
  poolStakingLink,
  poolStakingLinkText,
}) {
  const {
    onChangeTokenInputValue,
    clearInputs,
    priceImpact,
    poolTokens,
    inputState,
    tokenInputSum,
    depositAmount,
  } = useSwapPoolDeposit(poolName)

  const approveAndDeposit = useApproveAndDeposit(poolName)

  return (
    <div className="flex-col">
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {poolTokens.map((token, i) => (
          <TokenInputWithBalance
            token={token}
            inputValue={inputState[token.symbol]}
            onChangeTokenInputValue={onChangeTokenInputValue}
            key={i}
          />
        ))}
      </div>
      <TransactionButton
        label="Add Liquidity"
        pendingLabel="Adding Liquidity"
        disabled={tokenInputSum.eq(0)}
        className="items-center w-full px-6 py-3 mt-6 text-md rounded-xl"
        onClick={async () => {
          const appAndDeposit = await approveAndDeposit({
            slippageCustom: null,
            slippageSelected: 'ONE_TENTH',
            infiniteApproval: true,
            inputState,
            depositAmount,
          })
          // Clear input after deposit
          clearInputs()
        }}
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

function TokenInputWithBalance({ token, inputValue, onChangeTokenInputValue }) {
  let balanceToken
  if (token.symbol == WETH.symbol) {
    balanceToken = ETH
  } else if (token.symbol == AVWETH.symbol) {
    // token = WETHE
    balanceToken = WETHE
  } else {
    balanceToken = token
  }
  const balance = useTokenBalance(balanceToken)

  return (
    <TokenInput
      token={balanceToken}
      key={balanceToken.symbol}
      max={balance}
      inputValue={inputValue}
      onChange={(value) => onChangeTokenInputValue(token.symbol, value)}
      disabled={false}
    />
  )
}
