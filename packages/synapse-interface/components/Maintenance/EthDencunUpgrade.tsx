import { AnnouncementBanner } from './AnnouncementBanner'

/**
 * Start: 15 min prior to Eth Dencun Upgrade Time @ 3/13/24 13:55 UTC
 * End: 30 min after start of Eth Decun Upgrade Time
 */
export const ETH_DENCUN_BANNER_START = new Date(
  Date.UTC(2024, 2, 13, 13, 20, 0)
)
export const ETH_DENCUN_START_DATE = new Date(Date.UTC(2024, 2, 13, 13, 40, 0))
export const ETH_DENCUN_END_DATE = new Date(Date.UTC(2024, 2, 13, 14, 25, 0))

export const EthDencunUpgradeBanner = () => {
  return (
    <AnnouncementBanner
      bannerId="03132024-eth-dencun"
      bannerContents={
        <div className="flex flex-col justify-center space-y-1 text-center">
          <div>
            Synapse Bridge is upgrading ahead of the Ethereum Dencun upgrade
            (March 13, 13:55 UTC, 9:55 EST).
          </div>
          <div>Will be back online shortly ahead of the network upgrade.</div>
        </div>
      }
      startDate={ETH_DENCUN_BANNER_START}
      endDate={ETH_DENCUN_END_DATE}
    />
  )
}
