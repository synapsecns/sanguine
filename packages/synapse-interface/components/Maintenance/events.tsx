import { AnnouncementBanner } from './AnnouncementBanner'
import { useUpgradeProgressBar } from './UpgradeProgressBar'

const startDate = new Date(Date.UTC(2024, 2, 13, 1, 50, 0))
const endDate = new Date(Date.UTC(2024, 2, 13, 1, 60, 0))

export const EthDencunUpgradeBanner = (
  <AnnouncementBanner
    bannerId="03122024-eth-dencun"
    bannerContents={
      <div>
        The Bridge + RFQ will be globally offline 15mins ahead of the Dencun
        upgrade (March 13, 13:55 UTC, 9:55 EST). Will be back online about 15 -
        30 mins after Dencun.
      </div>
    }
    startDate={startDate}
    endDate={endDate}
  />
)

export const { UpgradeProgressBar: EthDencunUpgradeProgressBar } =
  useUpgradeProgressBar('Dencun upgrade in progress', startDate, endDate)
