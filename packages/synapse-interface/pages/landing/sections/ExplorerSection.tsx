import Link from 'next/link'
import Grid from '@/components/ui/tailwind/Grid'
import Button from '@/components/ui/tailwind/Button'
import Card from '@/components/ui/tailwind/Card'
import { SectionContainer } from '../../../components/landing/shared'
import { EXPLORER_PATH } from '@/constants/urls'
import {
  getTotalBridgeVolume,
  getTotalTxCount,
  getTotalValueLocked,
  ExplorerQueryStatsResponse,
} from '@/utils/hooks/useExplorerStats'

export default function ExplorerSection() {
  const totalBridgeVolume = getTotalBridgeVolume()
  const totalTxCount = getTotalTxCount()
  const totalValueLocked = getTotalValueLocked()

  return (
    <SectionContainer dataTestId="landing-explorer-section">
      <Grid
        cols={{ sm: 1, md: 2 }}
        gap={4}
        className={`
          flex items-center
          px-8 pt-6 mx-auto md:px-12
          absolute-lightened-bg z-10
        `}
      >
        <div className="max-w-sm mx-auto mt-12 text-left">
          <div className="mb-3 text-3xl font-medium text-white">
            Battle-tested infrastructure
          </div>
          <div className="text-secondaryTextColor ">
            Synapse has processed millions of transactions and tens of billions
            in bridged assets.
          </div>
        </div>
        <div className="hidden col-span-1 text-center md:block">
          <Link href={EXPLORER_PATH} target="_blank">
            <Button
              className={`
                  border-[#AC8FFF] border text-sm
                  px-4 py-3 hover:opacity-75
                `}
              style={{
                background:
                  'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
                borderRadius: '10px',
              }}
            >
              Go to Explorer
            </Button>
          </Link>
        </div>
      </Grid>

      <Grid
        cols={{ sm: 1, md: 2, lg: 3 }}
        gap={4}
        className={`
            max-w-4xl pb-12
            mx-auto space-x-0
            absolute-lightened-bg z-10 no-mt
            justify-center
        `}
      >
        <StatisticsCard title="Total Value Locked" value={totalValueLocked} />
        <StatisticsCard title="Total Bridge Volume" value={totalBridgeVolume} />
        <StatisticsCard title="Total TX Count" value={totalTxCount} />
      </Grid>
    </SectionContainer>
  )
}

function StatisticsCard({
  title,
  value,
}: {
  title: string
  value: ExplorerQueryStatsResponse
}) {
  return (
    <Card
      title={title}
      titleClassName="text-white opacity-75"
      className="justify-center px-2 pt-16 pb-8 bg-transparent"
      divider={false}
    >
      {value ? (
        <div className={`flex text-3xl font-medium text-white justify-left`}>
          {value}
        </div>
      ) : (
        <div
          className="w-full h-8 bg-slate-700 animate-pulse"
          style={{ maxWidth: '200px' }}
        />
      )}
    </Card>
  )
}
