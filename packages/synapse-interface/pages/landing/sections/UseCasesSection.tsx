import { SectionContainer } from '../../../components/landing/shared'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
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
      <div className="flex-col items-center py-6 mt-0 justify-left md:mt-0 lg:mt-0 md:flex md:justify-center lg:flex lg:justify-center ">
        <div className="flex items-center mb-4">
          <span className="mr-6 text-4xl text-white">Use cases</span>
        </div>
        <div className="mt-2 text-left text-secondaryTextColor md:text-center lg:text-center md:mt-0 lg:mt-0">
          Here’s a preview of what you can do using Synapse.
        </div>
      </div>

      <Grid
        cols={{ xs: 1, sm: 1, md: 1, lg: 3 }}
        gap={4}
        className="py-6 pt-6 pb-24 mx-auto place-items-center 2xl:w-3/4"
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
        bg-opacity-70 bg-[#2F2F2F] px-4 py-4
        md:py-0 w-full md:w-[300px]
      `}
      divider={false}
    >
      <div className="pb-4">
        <div className="flex justify-center mb-2">{image}</div>
        <div className="px-2">
          <div className="text-lg font-medium text-left text-white">
            {title}
          </div>
          <div className="mt-1 text-sm leading-6 text-left text-opacity-75 text-secondaryTextColor">
            {description}
          </div>
        </div>
      </div>
    </Card>
  )
}
