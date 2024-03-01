import { TRANSACTION_SUPPORT_URL, DISCORD_URL } from '@/constants/urls'

export const TransactionSupport = ({ isReverted }: { isReverted: boolean }) => {
  return (
    <div
      id="transaction-support"
      className="flex items-center justify-between w-full py-1 pl-3 pr-1 text-sm"
    >
      {isReverted ? (
        <div>Transaction reverted, funds returned</div>
      ) : (
        <div>What's taking so long?</div>
      )}

      <div className="flex items-center">
        <a
          href={TRANSACTION_SUPPORT_URL}
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-zinc-700 hover:no-underline
          `}
        >
          F.A.Q
        </a>
        <div>/</div>
        <a
          href={DISCORD_URL}
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-zinc-700 hover:no-underline
          `}
        >
          Support
        </a>
      </div>
    </div>
  )
}
