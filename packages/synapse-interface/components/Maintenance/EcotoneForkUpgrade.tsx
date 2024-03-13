import { AnnouncementBanner } from './AnnouncementBanner'

/**
 * Start: 10 min prior to Ecotone Fork Upgrade Time @ (March 14, 00:00 UTC)
 * End: 10 min after start of Ecotone Fork Upgrade Time
 */
// export const ECOTONE_FORK_BANNERS_START = new Date(
//   Date.UTC(2024, 2, 13, 23, 20, 0)
// )
// export const ECOTONE_FORK_START_DATE = new Date(
//   Date.UTC(2024, 2, 13, 23, 50, 0)
// )
// export const ECOTONE_FORK_END_DATE = new Date(Date.UTC(2024, 2, 13, 24, 10, 0))

/** TEST VALUES, REMOVE AFTER TESTING */
export const ECOTONE_FORK_BANNERS_START = new Date(
  Date.UTC(2024, 2, 13, 13, 0, 0)
)
export const ECOTONE_FORK_START_DATE = new Date(
  Date.UTC(2024, 2, 13, 14, 30, 0)
)
export const ECOTONE_FORK_END_DATE = new Date(Date.UTC(2024, 2, 13, 15, 30, 0))
/** TEST VALUES, REMOVE AFTER TESTING */

export const EcotoneForkUpgradeBanner = () => {
  return (
    <AnnouncementBanner
      bannerId="03142024-ecotone-fork"
      bannerContents={
        <div className="flex flex-col justify-center space-y-1 text-center">
          <div>
            Optimism + Base Bridging and RFQ will be paused 10 minutes ahead of
            Ecotone (March 14, 00:00 UTC).
          </div>
          <div>Will be back online shortly following the network upgrade.</div>
        </div>
      }
      startDate={ECOTONE_FORK_BANNERS_START}
      endDate={ECOTONE_FORK_END_DATE}
    />
  )
}
