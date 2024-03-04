import Card from '@tw/Card'
import InfoSection from './InfoSection'

const InfoSectionCard = ({ title, children }) => {
  return (
    <Card
      title={title}
      className="pt-3 rounded-lg bg-bgBase"
      titleClassName="font-medium text-secondaryTextColor mb-3"
      divider={false}
    >
      <InfoSection showDivider={false} showOutline={false}>
        {children}
      </InfoSection>
    </Card>
  )
}
export default InfoSectionCard
