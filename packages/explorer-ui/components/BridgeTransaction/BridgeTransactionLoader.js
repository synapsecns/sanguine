import {ChevronDoubleRightIcon} from '@heroicons/react/outline'

import Grid from '@components/tailwind/Grid'

import {ContainerCard} from '@components/ContainerCard'

export function BridgeTransactionLoader() {
  return (
    <div className="flex items-center justify-center mt-10 mb-10">
      <Grid gap={8} cols={{ sm: 1, md: 7, lg: 7 }}>
        <div className="col-span-3 ">
          <FullInfoLoader subtitle="Origin" />
        </div>
        <div className="flex items-center justify-center lg:col-span-1">
          <ChevronDoubleRightIcon className="w-10 h-10 text-purple-500" />
        </div>
        <div className="col-span-3">
          <FullInfoLoader subtitle="Destination" />
        </div>
      </Grid>
    </div>
  )
}

function FullInfoLoader({ subtitle }) {
  return (
    <div className="items-center">
      <div className="flex items-center mb-2 sm:text-xs md:text-lg ">
        <LoadingAnimation className="w-full h-5" />
      </div>
      <ContainerCard
        title={<LoadingAnimation className="w-40 h-5 mr-5" />}
        subtitle={subtitle}
      >
        <div className="mt-10 mb-10 text-4xl text-center">
          <LoadingAnimation className="w-full h-10" />
        </div>
        <LoadingAnimation className="w-full h-5 mb-2" />
        <LoadingAnimation className="w-full h-5 mb-2" />
      </ContainerCard>
    </div>
  )
}

function LoadingAnimation({ className = 'w-10 h-4' }) {
  return <div className={`${className} rounded bg-slate-500 animate-pulse`} />
}
