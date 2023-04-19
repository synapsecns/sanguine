import Card from '@/components/ui/tailwind/Card'

export function LandingPageContainer({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div
      data-test-id="landing-page-container"
      className="relative px-4 md:px-24"
    >
      {children}
    </div>
  )
}

export function SectionContainer({
  children,
  styles,
}: {
  children: React.ReactNode
  styles?: string
}) {
  return (
    <div
      className={`
        py-6 md:py-12 space-y-[1rem]
        ${styles}
      `}
    >
      {children}
    </div>
  )
}

export function SupportCard({
  header,
  children,
  image,
}: {
  header: string
  children: React.ReactNode
  image?: any
}) {
  return (
    <Card
      title={header}
      titleClassName="text-[1.69rem] font-medium text-white"
      className="px-0 bg-transparent text-secondaryTextColor"
      divider={false}
      image={image}
    >
      {children}
    </Card>
  )
}
