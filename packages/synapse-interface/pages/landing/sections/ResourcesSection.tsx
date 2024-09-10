import Link from 'next/link'
import { useTranslations } from 'next-intl'

import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'
import { SectionContainer } from '@/components/landing/shared'
import { GITHUB_URL, SYNAPSE_DOCS_URL, MEDIUM_URL } from '@/constants/urls'

export default function ResourcesSection() {
  const t = useTranslations('Landing.ResourcesSection')
  return (
    <SectionContainer
      dataTestId="landing-resources-section"
      styles="max-w-5xl m-auto"
    >
      <h2 className="text-4xl font-medium text-center text-white">
        {t('Get started now')}
      </h2>
      <p className="mb-8 text-center text-secondaryTextColor">
        {t('Find the resources')}
      </p>

      <Grid cols={{ sm: 1, md: 1, lg: 3 }} gap={6} className="p-4">
        <ResourceCard
          title={t('References')}
          description={t('Find the resources')}
          buttonText={t('See references')}
          linkUrl={GITHUB_URL}
        />
        <ResourceCard
          title={t('Documentation')}
          description={t('Read detailed')}
          buttonText={t('Read the docs')}
          linkUrl={SYNAPSE_DOCS_URL}
        />
        <ResourceCard
          title={t('Tutorials')}
          description={t('Watch tutorials')}
          buttonText={t('Go to tutorials')}
          linkUrl={MEDIUM_URL}
        />
      </Grid>
    </SectionContainer>
  )
}

function ResourceCard({ title, description, buttonText, linkUrl }) {
  return (
    <Card
      className={`
        text-center rounded-md border
        border-white/10
        bg-[#2F2F2F] bg-opacity-70 p-4 max-w-xs m-auto
      `}
      divider={false}
    >
      <div className="text-lg font-medium text-left text-white">{title}</div>
      <div
        className={`
          mt-1 mb-4 text-sm text-left
          text-opacity-75 text-secondaryTextColor
        `}
      >
        {description}
      </div>
      <div className="float-left">
        <Link href={linkUrl} target="_blank">
          <Button
            className={`
            bg-white hover:opacity-75
            text-sm text-[#18171B] font-medium
            px-4 py-3 border rounded-md
            `}
            onClick={() => null}
          >
            {buttonText}
          </Button>
        </Link>
      </div>
    </Card>
  )
}
