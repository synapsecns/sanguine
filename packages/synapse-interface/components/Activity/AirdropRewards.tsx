import TransactionArrow from '../icons/TransactionArrow'

export const AirdropRewards = () => {
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
