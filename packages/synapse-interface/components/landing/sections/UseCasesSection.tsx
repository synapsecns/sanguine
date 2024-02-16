import { SectionContainer } from '@/components/landing/shared'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import { CrossChainExchangeIcon } from '@/components/icons/LandingIcons/CrossChainExchangeIcon'
import { UniversalMoneyMarketsIcon } from '@/components/icons/LandingIcons/UniversalMoneyMarketsIcon'
import { MultiChainGamingIcon } from '@/components/icons/LandingIcons/MultiChainGamingIcon'

interface useCaseProp {
  title: string
  image: JSX.Element
  description: string
}

const useCases: useCaseProp[] = [
  {
    title: 'Cross-chain exchange',
    image: <CrossChainExchangeIcon />,
    description: 'Swap any asset on any blockchain using Synapse’s token swaps',
  },
  {
    title: 'Universal money markets',
    image: <UniversalMoneyMarketsIcon />,
    description:
      'Borrow and lend assets across any blockchain using Synapse’s pools',
  },
  {
    title: 'Multi-chain gaming',
    image: <MultiChainGamingIcon />,
    description:
      'Create uniqe gaming experiences that access multiple blockchains',
  },
]

export default function UseCasesSection() {
  return (
    <SectionContainer dataTestId="landing-use-cases-section">
      <div className="flex flex-col items-center">
        <h2 className="mr-6 text-4xl text-white mb-4">Use cases</h2>
        <p className="text-secondaryTextColor md:text-center">
          Here’s a preview of what you can do using Synapse.
        </p>
      </div>

      <Grid
        cols={{ xs: 1, sm: 1, md: 1, lg: 3 }}
        gap={4}
        className="p-4 mx-auto place-items-center max-w-5xl"
      >
        {useCases.map((useCase: useCaseProp, index: number) => (
          <UseCaseCard
            key={index}
            image={useCase.image}
            title={useCase.title}
            description={useCase.description}
          />
        ))}
      </Grid>
    </SectionContainer>
  )
}

function UseCaseCard({ image, title, description }) {
  return (
    <Card
      className={`
        border border-white border-opacity-10
        bg-opacity-70 bg-[#2F2F2F]
        w-full md:w-[300px] flex flex-col gap-2 max-w-xs
      `}
      divider={false}
    >
        <div className="flex -my-4 md:-my-5">{image}</div>
        <div className="text-lg font-medium text-white mt-2">
          {title}
        </div>
        <div className="text-sm text-secondaryTextColor">
          {description}
        </div>
    </Card>
  )
}
