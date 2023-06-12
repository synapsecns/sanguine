import _ from 'lodash'

import { WETH } from '@constants/tokens/swapMaster'
import { AVWETH, ETH, WETHE } from '@constants/tokens/master'
import { stringToBigNum } from '@/utils/stringToBigNum'
import { getAddress } from '@ethersproject/address'
import { DepositTokenInput } from '@components/TokenInput'
import PriceImpactDisplay from '../components/PriceImpactDisplay'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { TransactionButton } from '@/components/buttons/TransactionButton'
import { Zero } from '@ethersproject/constants'
import { PoolToken, Token } from '@types'
import { useState, useEffect, useMemo } from 'react'
import { BigNumber } from '@ethersproject/bignumber'
import { calculateExchangeRate } from '@utils/calculateExchangeRate'
import { getTokenAllowance } from '@/utils/actions/getTokenAllowance'
import { approve, deposit } from '@/utils/actions/approveAndDeposit'
import { QUOTE_POLLING_INTERVAL } from '@/constants/bridge' // TODO CHANGE
import { PoolData, PoolUserData } from '@types'
import LoadingTokenInput from '@components/loading/LoadingTokenInput'
import { fetchBalance, fetchToken } from '@wagmi/core'
import { formatBNToString } from '@/utils/bignumber/format'

const DEFAULT_DEPOSIT_QUOTE = {
  priceImpact: undefined,
  allowances: {},
  routerAddress: '',
}

// TODO: Find mistmatch in pool tokens vs. native tokens

const Deposit = ({
  pool,
  chainId,
  address,
  poolData,
  poolUserData,
  refetchCallback,
}: {
  pool: Token
  chainId: number
  address: string
  poolData: PoolData
  poolUserData: PoolUserData
  refetchCallback: () => void
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
  }>(DEFAULT_DEPOSIT_QUOTE)
  const [time, setTime] = useState(Date.now())
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
      } else {
        setDepositQuote(DEFAULT_DEPOSIT_QUOTE)
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
  }, [inputValue, time, pool, chainId, address])

  const onChangeInputValue = (token: Token, value: string) => {
    // console.log(`[changeInputValue] token`, token)
    // const ga = getAddress(token.addresses[chainId])
    // console.log(`ga`, ga)
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
    setDepositQuote(DEFAULT_DEPOSIT_QUOTE)
  }

  let isFromBalanceEnough = true
  let isAllowanceEnough = true

  const getButtonProperties = () => {
    let properties = {
      label: 'Deposit',
      pendingLabel: 'Depositing funds...',
      className: '',
      disabled: false,
      buttonAction: () => {
        console.log(`[buttonAction] inputValue.bn`, inputValue.bn)
        console.log(`[buttonAction] pool`, pool)
        console.log(`[buttonAction] poolData`, poolData)
        console.log(`[buttonAction] poolUserData`, poolUserData)

        const filteredInputValues = filterInputValues(inputValue, pool, chainId)

        console.log(`[buttonAction] filteredInputValues`, filteredInputValues)

        // return deposit(pool, 'ONE_TENTH', null, inputValue.bn, chainId)
        return deposit(pool, 'ONE_TENTH', null, filteredInputValues, chainId)
      },

      postButtonAction: () => {
        console.log('Post Button Action')
        refetchCallback()
        resetInputs()
      },
    }

    // if (sumBigNumbersFromState().eq(0)) {
    //   console.log(`'am hi here `)
    //   properties.disabled = true
    // }

    if (!isFromBalanceEnough) {
      console.log(`m i here `)
      properties.label = `Insufficient Balance`
      properties.disabled = true
      return properties
    }

    if (!isAllowanceEnough) {
      properties.label = `Approve Token(s)`
      properties.pendingLabel = `Approving Token(s)`
      properties.className = 'from-[#feba06] to-[#FEC737]'
      properties.disabled = false
      properties.buttonAction = () =>
        approve(pool, depositQuote, inputValue.bn, chainId).then(() =>
          calculateMaxDeposits()
        )
      properties.postButtonAction = () => setTime(0)
      return properties
    }

    return properties
  }

  for (const [tokenAddr, amount] of Object.entries(inputValue.bn)) {
    if (
      typeof amount !== 'undefined' &&
      Object.keys(depositQuote.allowances).length > 0 &&
      !amount.isZero() &&
      typeof depositQuote.allowances[tokenAddr] !== 'undefined' &&
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

  const {
    label: btnLabel,
    pendingLabel,
    className: btnClassName,
    buttonAction,
    postButtonAction,
    disabled,
  } = useMemo(getButtonProperties, [
    isFromBalanceEnough,
    isAllowanceEnough,
    address,
    inputValue,
    depositQuote,
  ])

  const actionBtn = useMemo(
    () => (
      <TransactionButton
        className={btnClassName}
        // disabled={sumBigNumbersFromState().eq(0) || disabled}
        disabled={disabled}
        onClick={() => buttonAction()}
        onSuccess={() => postButtonAction()}
        label={btnLabel}
        pendingLabel={pendingLabel}
      />
    ),
    [
      buttonAction,
      postButtonAction,
      btnLabel,
      pendingLabel,
      btnClassName,
      isFromBalanceEnough,
      isAllowanceEnough,
      inputValue,
    ]
  )

  // console.log(`pooluserData`, poolUserData)

  return (
    <div className="flex-col">
      <div className="px-2 pt-1 pb-4 bg-bgLight rounded-xl">
        {pool && poolUserData && poolData ? (
          poolUserData.tokens.map((tokenObj, i) => {
            return (
              <SerializedDepositInput
                key={i}
                tokenObj={tokenObj}
                address={address}
                chainId={chainId}
                inputValue={inputValue}
                onChangeInputValue={onChangeInputValue}
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

const SerializedDepositInput = ({
  tokenObj,
  address,
  chainId,
  inputValue,
  onChangeInputValue,
}) => {
  const [serializedToken, setSerializedToken] = useState(undefined)
  const balanceToken = correctToken(tokenObj.token)

  useEffect(() => {
    const fetchSerializedData = async () => {
      try {
        const token = await serializeToken(
          address,
          chainId,
          balanceToken,
          tokenObj
        )
        setSerializedToken(token)
      } catch (error) {
        console.log(`error`, error)
      }
    }

    fetchSerializedData()
  }, [])

  // console.log(`serializedToken`, serializedToken)
  // console.log(`inputValue`, inputValue)
  return (
    serializedToken && (
      <DepositTokenInput
        token={serializedToken}
        key={serializedToken.symbol}
        rawBalance={serializedToken.rawBalance}
        balanceStr={String(serializedToken.balanceStr)}
        inputValueStr={inputValue.str[serializedToken.addresses[chainId]]}
        onChange={(value) => onChangeInputValue(serializedToken, value)}
        chainId={chainId}
        address={address}
      />
    )
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

const filterInputValues = (inputValues, pool, chainId) => {
  // take the inputvalues and only keep the ones that have pool native tokens

  // inputvalues.bn
  // 0x0000000000000000000000000000000000000000: BigNumber {_hex: '0x11c37937e08000', _isBigNumber: true}
  // 0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e: BigNumber {_hex: '0x00', _isBigNumber: true}
  // 0x82aF49447D8a07e3bd95BD0d56f35241523fBab1 : BigNumber {_hex: '0x00',

  const poolTokens = pool.nativeTokens ?? pool.poolTokens

  const poolTokenAddresses = []

  poolTokens.map((nativeToken) => {
    poolTokenAddresses.push(nativeToken.addresses[chainId])
  })

  let filteredInputValues = Object.keys(inputValues.bn)
    .filter((key) => poolTokenAddresses.includes(key))
    .reduce((obj, key) => {
      obj[key] = inputValues.bn[key]
      return obj
    }, {})

  return filteredInputValues
}

const serializeToken = async (
  address: string,
  chainId: number,
  balanceToken: Token,
  tokenObj: any
  // tokenObj: PoolToken
) => {
  let fetchedBalance

  if (balanceToken === ETH) {
    fetchedBalance = await fetchBalance({
      address: address as `0x${string}`,
      chainId,
    })

    return {
      ...balanceToken,
      rawBalance: fetchedBalance.value,
      balanceStr: formatBNToString(
        fetchedBalance.value,
        balanceToken.decimals[chainId],
        4
      ),
    }
  } else if (balanceToken === WETHE) {
    fetchedBalance = await fetchBalance({
      address: address as `0x${string}`,
      chainId,
      token: balanceToken.addresses[chainId] as `0x${string}`,
    })

    return {
      ...balanceToken,
      rawBalance: fetchedBalance.value,
      balanceStr: fetchedBalance.formatted,
    }
  } else {
    return {
      ...balanceToken,
      rawBalance: tokenObj.rawBalance,
      balanceStr: tokenObj.balanceStr,
    }
  }
}

export default Deposit
