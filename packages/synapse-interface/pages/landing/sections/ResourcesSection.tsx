import Link from 'next/link'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'
import { SectionContainer } from '../../../components/landing/shared'
import { GITHUB_URL, DOCS_URL, MEDIUM_URL } from '@/constants/urls'

export default function ResourcesSection() {
  return (
    <SectionContainer dataTestId="landing-resources-section">
      <div
        className={`
          mt-8 mb-4 text-4xl font-medium text-left
          text-white lg:text-center md:text-center
        `}
      >
        Get started now
      </div>
      <div
        className={`
          mb-8 text-left text-secondaryTextColor
          md:text-center lg:text-center
        `}
      >
        Find the resources you need to create integrations with Synapse.
      </div>

      <Grid
        cols={{ sm: 1, md: 1, lg: 3 }}
        gap={6}
        className="py-4 mx-auto lg:px-12 2xl:w-3/4"
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
        border-white border-opacity-10
        bg-[#2F2F2F] bg-opacity-70 py-6 px-6
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
