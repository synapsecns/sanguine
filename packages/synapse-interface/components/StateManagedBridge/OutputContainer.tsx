import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'
import { ChainSelector } from '@/components/ui/ChainSelector'
import { TokenSelector } from '@/components/ui/TokenSelector'
import { useToChainListArray } from './hooks/useToChainListArray'
import { shortenAddress } from '@/utils/shortenAddress'
// import { ToTokenSelector } from './ToTokenSelector'
import { useDispatch } from 'react-redux'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeSectionContainer } from '@/components/ui/BridgeSectionContainer'
import { BridgeAmountContainer } from '@/components/ui/BridgeAmountContainer'
import { AmountInput } from '@/components/ui/AmountInput'
import { useToTokenListArray } from '@/components/StateManagedBridge/hooks/useToTokenListArray'

export const OutputContainer = ({}) => {
  const { bridgeQuote, isLoading, toChainId, toToken } = useBridgeState()

  const { address: isConnectedAddress } = useAccount()
  const [address, setAddress] = useState<Address>()

  const dispatch = useDispatch()

  const showValue =
    bridgeQuote?.outputAmountString === '0'
      ? ''
      : bridgeQuote?.outputAmountString

  useEffect(() => {
    setAddress(isConnectedAddress)
  }, [isConnectedAddress])

  // update address for destination address if we have a destination address

  return (
    <BridgeSectionContainer>
      <div className="flex items-center justify-between">
        <ToChainSelector />
        {/* {address && (
          <div className="h-5">
            <DisplayAddress address={address} />
          </div>
        )} */}
      </div>

      <BridgeAmountContainer>
        <ToTokenSelector />
        <AmountInput
          disabled={true}
          showValue={showValue}
          isLoading={isLoading}
        />
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const DisplayAddress = ({ address }) => {
  return (
    <div className="border-[0.5px] border-secondaryTextColor rounded-md pt-1 pb-1 pl-3 pr-3 text-secondaryTextColor text-xxs">
      {shortenAddress(address)}
    </div>
  )
}

const ToChainSelector = () => (
  <ChainSelector
    dataTestId="bridge-destination-chain"
    isOrigin={false}
    selectedItem={CHAINS_BY_ID[useBridgeState().toChainId]}
    label="To"
    itemListFunction={useToChainListArray}
    setFunction={setToChainId}
    action="Bridge"
  />
)

const ToTokenSelector = () => (
  <TokenSelector
    dataTestId="bridge-destination-token"
    isOrigin={false}
    selectedItem={useBridgeState().toToken}
    placeholder="In"
    itemListFunction={useToTokenListArray}
    setFunction={setToToken}
    action="Bridge"
  />
)
