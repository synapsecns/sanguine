import { useTranslations } from 'next-intl'
import { SectionContainer } from '@/components/landing/shared'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import { CrossChainExchangeIcon } from '@/components/icons/LandingIcons/CrossChainExchangeIcon'
import { UniversalMoneyMarketsIcon } from '@/components/icons/LandingIcons/UniversalMoneyMarketsIcon'
import { MultiChainGamingIcon } from '@/components/icons/LandingIcons/MultiChainGamingIcon'

interface UseCaseProp {
  title: string
  image: JSX.Element
  description: string
}

export default function UseCasesSection() {
  const t = useTranslations('Landing.UseCasesSection')

  const useCases: UseCaseProp[] = [
    {
      title: t('useCases.0.title'),
      image: <CrossChainExchangeIcon />,
      description: t('useCases.0.description'),
    },
    {
      title: t('useCases.1.title'),
      image: <UniversalMoneyMarketsIcon />,
      description: t('useCases.1.description'),
    },
    {
      title: t('useCases.2.title'),
      image: <MultiChainGamingIcon />,
      description: t('useCases.2.description'),
    },
  ]

  return (
    <SectionContainer dataTestId="landing-use-cases-section">
      <div className="flex flex-col items-center">
        <h2 className="mb-4 mr-6 text-4xl text-white">{t('title')}</h2>
        <p className="text-secondaryTextColor md:text-center">
          {t('subtitle')}
        </p>
      </div>
      <Grid
        cols={{ xs: 1, sm: 1, md: 1, lg: 3 }}
        gap={4}
        className="max-w-5xl p-4 mx-auto place-items-center"
      >
        {useCases.map((useCase: UseCaseProp, index: number) => (
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

function UseCaseCard({ image, title, description }: UseCaseProp) {
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
      <div className="mt-2 text-lg font-medium text-white">{title}</div>
      <div className="text-sm text-secondaryTextColor">{description}</div>
    </Card>
  )
}
