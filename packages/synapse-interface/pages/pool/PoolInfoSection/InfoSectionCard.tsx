import Card from '@tw/Card'
import InfoSection from './InfoSection'

const InfoSectionCard = ({ title, children }) => {
  return (
    <Card
      title={title}
      className="p-3 rounded-md bg-zinc-100 border border-zinc-200 dark:bg-zinc-800 dark:border-transparent"
      divider={false}
    >
      <InfoSection showDivider={false} showOutline={false}>
        {children}
      </InfoSection>
    </Card>
  )
}
export default InfoSectionCard
