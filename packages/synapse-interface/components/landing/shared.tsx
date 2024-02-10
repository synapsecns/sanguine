import Card from '@/components/ui/tailwind/Card'

export const LandingPageContainer = ({
  children,
}: {
  children: React.ReactNode
}) => {
  return (
    <div data-test-id="landing-page" className="p-4">
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
    <section
      className={`
        py-6 md:my-8 space-y-[1rem]
        ${styles}
      `}
      data-test-id={dataTestId}
    >
      {children}
    </section>
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
      title={
        <>
          {header}
          <div className="inline-block float-right">
            {image}
          </div>
        </>

      }
      titleClassName="text-3xl py-2"
      className="bg-transparent text-secondaryTextColor  max-w-md"
      divider={false}
    >
      {children}
    </Card>
  )
}
