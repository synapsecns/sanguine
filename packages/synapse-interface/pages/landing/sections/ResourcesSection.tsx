import Link from 'next/link'
import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'
import { SectionContainer } from '../../../components/landing/shared'
import { GITHUB_URL, DOCS_URL, MEDIUM_URL } from '@/constants/urls'

export default function ResourcesSection() {
  return (
    <SectionContainer dataTestId="landing-resources-section" styles="max-w-5xl m-auto">
      <h2 className="text-4xl font-medium text-center">
        Get started now
      </h2>
      <p className="mb-8 text-center text-zinc-700 dark:text-zinc-400">
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
        bg-zinc-200 dark:bg-zinc-800 p-4 max-w-xs m-auto
      `}
      divider={false}
    >
      <div className="text-xl font-medium text-left text-zinc-900 dark:text-zinc-100">
        {title}
      </div>
      <div
        className={`
          mt-2 mb-4 text-sm text-left
          text-opacity-75 text-zinc-700 dark:text-zinc-400
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
