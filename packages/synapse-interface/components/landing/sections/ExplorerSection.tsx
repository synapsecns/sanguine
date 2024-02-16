import Link from 'next/link'
import Grid from '@tw/Grid'
import Button from '@tw/Button'
import Card from '@tw/Card'
import { SectionContainer } from '@/components/landing/shared'
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
    <SectionContainer dataTestId="landing-explorer-section" styles="-mx-4">
      <div className="w-full bg-bgBase/10 pb-6">
        <Grid
          cols={{ sm: 1, md: 2 }}
          className="flex items-center p-6 m-auto max-w-4xl"
        >
          <div>
            <h2 className="mb-3 text-3xl font-medium text-white text-center md:text-left">
              Battle-tested infrastructure
            </h2>
            <p className="text-secondaryTextColor text-center md:text-left">
              Synapse has processed millions of transactions and tens of billions
              in bridged assets.
            </p>
          </div>
          <div className="hidden col-span-1 text-right md:block">
            <Link href={EXPLORER_PATH} target="_blank">
              <Button
                className={`
                    border border-[#AC8FFF] text-sm
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
          cols={{ sm: 1, md: 3 }}
          gap={4}
          className="max-w-4xl justify-center m-auto"
        >
          <StatisticsCard title="Total Value Locked" value={totalValueLocked} />
          <StatisticsCard title="Total Bridge Volume" value={totalBridgeVolume} />
          <StatisticsCard title="Total TX Count" value={totalTxCount} />
        </Grid>
      </div>
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
      className="justify-center p-4 bg-transparent"
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
