import TransactionArrow from '../icons/TransactionArrow'
import { getErc20TokenTransfers } from '@/utils/actions/getErc20TokenTransfers'
import { useEffect } from 'react'
import { arbitrum } from 'viem/chains'

export const AirdropRewards = () => {
  useEffect(() => {
    ;(async () => {
      const transfers = await getErc20TokenTransfers(
        '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
        '0xF080B794AbF6BB905F2330d25DF545914e6027F8',
        '0x81EF4608B796265F1e3695cE00FdCfC8aA5933Dd',
        arbitrum,
        173545720n
      )
      console.log('transfers:', transfers)
    })()
  }, [])

  return (
    <div
      id="airdrop-rewards"
      className="flex border rounded-lg text-secondary border-surface bg-background"
    >
      <div className="text-green-500">Rebate</div>
      <TransactionArrow />
      <div>
        <div>Token Img, Token Name</div>
        <div>Token Img, Airdropped Total</div>
      </div>
    </div>
  )
}
