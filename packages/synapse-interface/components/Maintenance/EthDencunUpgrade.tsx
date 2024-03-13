import { AnnouncementBanner } from './AnnouncementBanner'
import { useUpgradeProgressBar } from './UpgradeProgressBar'
import { getTimeMinutesBeforeNow } from '@/utils/time'

/** Test Values */
export const ETH_DENCUN_BANNER_START = new Date(Date.UTC(2024, 2, 12, 0, 0, 0))
export const ETH_DENCUN_START_DATE = new Date(Date.UTC(2024, 2, 13, 3, 0, 0))
export const ETH_DENCUN_END_DATE = new Date(Date.UTC(2024, 2, 13, 3, 45, 0))
/** Test Values */

/**
 * Start: 15 min prior to Eth Dencun Upgrade Time @ 3/13/24 13:55 UTC
 * End: 30 min after start of Eth Decun Upgrade Time
 */
// export const ETH_DENCUN_BANNER_START = new Date(Date.UTC(2024, 2, 12, 0, 0, 0))
// export const ETH_DENCUN_START_DATE = new Date(Date.UTC(2024, 2, 13, 13, 40, 0))
// export const ETH_DENCUN_END_DATE = new Date(Date.UTC(2024, 2, 12, 14, 25, 0))

export const EthDencunUpgradeBanner = () => {
  return (
    <AnnouncementBanner
      bannerId="03122024-eth-dencun"
      bannerContents={
        <div>
          The Bridge + RFQ will be globally offline 15mins ahead of the Dencun
          upgrade (March 13, 13:55 UTC, 9:55 EST). Will be back online about 15
          - 30 mins after Dencun.
        </div>
      }
      startDate={ETH_DENCUN_BANNER_START}
      endDate={ETH_DENCUN_END_DATE}
    />
  )
}
