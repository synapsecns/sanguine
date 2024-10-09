import { useTranslations } from 'next-intl'
import { TRANSACTION_SUPPORT_URL, DISCORD_URL } from '@/constants/urls'

export const TransactionSupport = ({
  status,
}: {
  status: 'pending' | 'completed' | 'reverted' | 'refunded'
}) => {
  const t = useTranslations('Time')

  return (
    <div
      id="transaction-support"
      className="flex items-center justify-between w-full py-1 pl-3 pr-1 text-sm"
    >
      {status === 'reverted' && (
        <div>{t('Transaction reverted, funds returned')}</div>
      )}

      {status === 'refunded' && (
        <div>{t('Transaction refunded, funds returned')}</div>
      )}

      {status === 'pending' && <div>{t("What's taking so long?")}</div>}

      <div className="flex items-center">
        <a
          href={TRANSACTION_SUPPORT_URL}
          target="_blank"
          className={`
            px-2 py-1 underline 
            hover:rounded hover:bg-zinc-700 hover:no-underline
          `}
        >
          {t('FAQ')}
        </a>
        <div>/</div>
        <a
          href={DISCORD_URL}
          target="_blank"
          className={`
            px-2 py-1 underline
            hover:rounded hover:bg-zinc-700 hover:no-underline
          `}
        >
          {t('Support')}
        </a>
      </div>
    </div>
  )
}
