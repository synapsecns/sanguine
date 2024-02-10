import Link from 'next/link'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'
import { SectionContainer } from '../../../components/landing/shared'
import { GITHUB_URL, DOCS_URL, MEDIUM_URL } from '@/constants/urls'

export default function ResourcesSection() {
  return (
    <SectionContainer dataTestId="landing-resources-section" styles="max-w-5xl m-auto">
      <h2 className="text-4xl font-medium text-center text-white">
        Get started now
      </h2>
      <p className="mb-8 text-center text-secondaryTextColor">
        Find the resources you need to create integrations with Synapse.
      </p>

      <Grid
        cols={{ sm: 1, md: 1, lg: 3 }}
        gap={6}
        className="p-4"
      >
        <ResourceCard
          title="References"
          description="Find the resources you need to create integrations with Synapse."
          buttonText="See references"
          linkUrl={GITHUB_URL}
        />
        <ResourceCard
          title="Documentation"
          description="Read a detailed breakdown of our APIs and smart contracts."
          buttonText="Read the docs"
          linkUrl={DOCS_URL}
        />
        <ResourceCard
          title="Tutorials"
          description="Watch interactive tutorials to learn how Synapse works."
          buttonText="Go to tutorials"
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
