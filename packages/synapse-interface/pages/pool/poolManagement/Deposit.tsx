import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { getAddress } from '@ethersproject/address'
import TokenInput from '@components/TokenInput'
import PriceImpactDisplay from '../components/PriceImpactDisplay'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { Zero } from '@ethersproject/constants'
import { Token } from '@types'
import { useState, useEffect } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { approve, deposit } from '@/utils/actions/approveAndDeposit'
import { QUOTE_POLLING_INTERVAL } from '@/constants/bridge' // TODO CHANGE
import { PoolData, PoolUserData } from '@types'
import LoadingTokenInput from '@components/loading/LoadingTokenInput'

const Deposit = ({
  pool,
  chainId,
  address,
  poolData,
  poolUserData,
}: {
  pool: Token
  chainId: number
  address: string
  poolData: PoolData
  poolUserData: PoolUserData
}) => {
  // todo store sum in here?
  const [inputValue, setInputValue] = useState<{
    bn: Record<string, BigNumber>
    str: Record<string, string>
  }>({ bn: {}, str: {} })
  const [depositQuote, setDepositQuote] = useState<{
    priceImpact: BigNumber
    allowances: Record<string, BigNumber>
    routerAddress: string
  }>({ priceImpact: undefined, allowances: {}, routerAddress: '' })
  const [time, setTime] = useState(Date.now())
  const [successTx, setSuccessTx] = useState<string>(null)
  const { synapseSDK } = useSynapseContext()

  // TODO move this to utils
  const sumBigNumbersFromState = () => {
    let sum = Zero
    pool?.poolTokens &&
      pool.poolTokens.map((token) => {
        if (!token.addresses[chainId]) return
        const tokenAddress = getAddress(token.addresses[chainId])
        if (inputValue.bn[tokenAddress]) {
          sum = sum.add(
            inputValue.bn[getAddress(token.addresses[chainId])].mul(
              BigNumber.from(10).pow(18 - token.decimals[chainId])
            )
          )
        }
      })
    return sum
  }

  const calculateMaxDeposits = async () => {
    try {
      if (poolUserData == null || address == null) {
        return
      }
      let inputSum = sumBigNumbersFromState()
      if (poolData.totalLocked.gt(0) && inputSum.gt(0)) {
        const { amount } = await synapseSDK.calculateAddLiquidity(
          chainId,
          pool.swapAddresses[chainId],
          inputValue.bn
        )

        let allowances: Record<string, BigNumber> = {}
        for (const [key, value] of Object.entries(inputValue.bn)) {
          allowances[key] = await getTokenAllowance(
            pool.swapAddresses[chainId],
            key,
            address,
            chainId
          )
        }

        const priceImpact = calculateExchangeRate(
          inputSum,
          18,
          inputSum.sub(amount),
          18
        )
        // TODO: DOUBLE CHECK THIS
        setDepositQuote({
          priceImpact,
          allowances,
          routerAddress: pool.swapAddresses[chainId],
        })
      }
    } catch (e) {
      console.log(e)
    }
  }
  useEffect(() => {
    const interval = setInterval(
      () => setTime(Date.now()),
      QUOTE_POLLING_INTERVAL
    )
    return () => {
      clearInterval(interval)
    }
  }, [])

  useEffect(() => {
    calculateMaxDeposits()
  }, [inputValue, time, pool, chainId, address, successTx])

  const onChangeInputValue = (token: Token, value: string) => {
    const bigNum = stringToBigNum(value, token.decimals[chainId]) ?? Zero
    if (chainId && token) {
      setInputValue({
        bn: {
          ...inputValue.bn,
          [getAddress(token.addresses[chainId])]: bigNum,
        },
        str: {
          ...inputValue.str,
          [getAddress(token.addresses[chainId])]: value,
        },
      })
    }
  }

  useEffect(() => {
    if (poolData && poolUserData && pool && chainId && address) {
      resetInputs()
    }
  }, [poolUserData])

  const resetInputs = () => {
    let initInputValue: {
      bn: Record<string, BigNumber>
      str: Record<string, string>
    } = { bn: {}, str: {} }
    poolUserData.tokens.map((tokenObj, i) => {
      initInputValue.bn[tokenObj.token.addresses[chainId]] = Zero
      initInputValue.str[tokenObj.token.addresses[chainId]] = ''
    })
    setInputValue(initInputValue)
  }

  // some messy button gen stuff (will re-write)
  let isFromBalanceEnough = true
  let isAllowanceEnough = true
  let btnLabel = 'Deposit'
  let pendingLabel = 'Depositing funds...'
  let btnClassName = ''
  let buttonAction = () =>
    deposit(pool, 'ONE_TENTH', null, inputValue.bn, chainId).then((successTx) =>
      setSuccessTx(successTx)
    )
  let postButtonAction = () => {
    resetInputs()
  }

  for (const [tokenAddr, amount] of Object.entries(inputValue.bn)) {
    if (
      Object.keys(depositQuote.allowances).length > 0 &&
      !amount.isZero() &&
      amount.gt(depositQuote.allowances[tokenAddr])
    ) {
      isAllowanceEnough = false
    }
    poolUserData.tokens.map((tokenObj, i) => {
      if (
        tokenObj.token.addresses[chainId] === tokenAddr &&
        amount.gt(tokenObj.balance)
      ) {
        isFromBalanceEnough = false
      }
    })
  }

  if (!isFromBalanceEnough) {
    btnLabel = `Insufficient Balance`
  } else if (!isAllowanceEnough) {
    buttonAction = () => approve(pool, depositQuote, inputValue.bn, chainId)
    btnLabel = `Approve Token(s)`
    pendingLabel = `Approving Token(s)`
    btnClassName = 'from-[#feba06] to-[#FEC737]'
    postButtonAction = () => setTime(0)
  }
  const actionBtn = (
    <TransactionButton
      className={btnClassName}
      disabled={sumBigNumbersFromState().eq(0) || !isFromBalanceEnough}
      onClick={() => buttonAction()}
      onSuccess={() => postButtonAction()}
      label={btnLabel}
      pendingLabel={pendingLabel}
    />
  )

  return (
    <div className="flex-col">
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {pool && poolUserData && poolData ? (
          poolUserData.tokens.map((tokenObj, i) => {
            const balanceToken = correctToken(tokenObj.token)
            return (
              <TokenInput
                token={balanceToken}
                key={balanceToken.symbol}
                balanceStr={String(tokenObj.balanceStr)}
                inputValueStr={inputValue.str[balanceToken.addresses[chainId]]}
                onChange={(value) => onChangeInputValue(balanceToken, value)}
                chainId={chainId}
                address={address}
              />
            )
          })
        ) : (
          <>
            <LoadingTokenInput />
            <LoadingTokenInput />
          </>
        )}
      </div>
      {actionBtn}
      {depositQuote.priceImpact && depositQuote.priceImpact?.gt(Zero) && (
        <PriceImpactDisplay priceImpact={depositQuote.priceImpact} />
      )}
    </div>
  )
}
const correctToken = (token: Token) => {
  let balanceToken: Token | undefined
  if (token.symbol == WETH.symbol) {
    balanceToken = ETH
  } else if (token.symbol == AVWETH.symbol) {
    // token = WETHE
    balanceToken = WETHE
  } else {
    balanceToken = token
  }
  return balanceToken
}

export default Deposit
