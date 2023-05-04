import Card from '@/components/ui/tailwind/Card'

export const LandingPageContainer = ({
  children,
}: {
  children: React.ReactNode
}) => {
  return (
    <div data-test-id="landing-page" className="relative px-4 md:px-24">
      {children}
    </div>
  )
}

export const SectionContainer = ({
  children,
  styles,
  dataTestId,
}: {
  children: React.ReactNode
  styles?: string
  dataTestId?: string
}) => {
  return (
    <div
      className={`
        py-6 md:my-8 space-y-[1rem]
        ${styles}
      `}
      data-test-id={dataTestId}
    >
      {children}
    </div>
  )
}

export const SupportCard = ({
  header,
  children,
  image,
}: {
  header: string
  children: React.ReactNode
  image?: any
}) => {
  return (
    <Card
      title={header}
      titleClassName="text-[1.69rem] font-medium text-white"
      className="px-0 bg-transparent text-secondaryTextColor sm:pb-0"
      divider={false}
      image={image}
    >
      {children}
    </Card>
  )
}
