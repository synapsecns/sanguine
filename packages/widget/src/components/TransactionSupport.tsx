import { TRANSACTION_SUPPORT_URL } from '@/constants/index'

export const TransactionSupport = () => {
  return (
    <div
      id="transaction-support"
      className="flex items-center justify-between w-full"
    >
      <div>What's taking so long?</div>
      <div className="flex items-center">
        <a
          href={TRANSACTION_SUPPORT_URL}
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-[--synapse-select-bg] hover:no-underline
          `}
        >
          F.A.Q
        </a>
        <div>/</div>
        <a
          href="https://discord.gg/4rMzuEnKqe"
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-[--synapse-select-bg] hover:no-underline
            min-[360px]:after:content-['_(Discord)']
          `}
        >
          Support
        </a>
      </div>
    </div>
  )
}
