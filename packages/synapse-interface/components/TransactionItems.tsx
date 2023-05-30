import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkLinkTextColor } from '@styles/chains'
import { Chain } from '@types'

export const CheckingConfPlaceholder = ({ chain }: { chain: Chain }) => {
  return (
    <div className="flex items-center p-1 max-w-[80%] rounded-lg ">
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chain?.color)}
                opacity-70 pr-2
              `}
            >
              Confirmations left on {chain?.name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}

export const PendingCreditTransactionItem = ({
  chainId,
}: {
  chainId: number
}) => {
  const chain = CHAINS_BY_ID[chainId]

  return (
    <div className="flex items-center p-1 rounded-lg ">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chain?.chainImg?.src}
        />
      </div>
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chain?.color)}
                opacity-70 pr-2
              `}
            >
              Waiting to be credited on
              <br />
              {chain?.name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}

export const EmptySubTransactionItem = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return (
    <div className="flex items-center py-1 pl-3 rounded-lg ">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chain?.chainImg.src}
        />
      </div>
    </div>
  )
}

export const CreditedTransactionItem = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return (
    <div className="flex items-center p-1 border border-gray-700 rounded-lg">
      <div className="flex-shrink-0">
        <img
          className="inline w-4 h-4 ml-1 mr-2 -mt-1 rounded"
          src={chain?.chainImg.src}
        />
      </div>
      <div>
        <div>
          <div className="w-full text-sm">
            <div
              className={`
                ${getNetworkLinkTextColor(chain?.color)}
                opacity-70 pr-2
              `}
            >
              Bridging Completed on
              <br />
              {chain?.name}
            </div>
          </div>
        </div>
        <div className="w-full"></div>
      </div>
    </div>
  )
}
