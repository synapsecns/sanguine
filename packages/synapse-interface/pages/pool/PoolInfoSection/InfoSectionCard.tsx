import Card from '@tw/Card'
import InfoSection from './InfoSection'

const InfoSectionCard = ({ title, children }) => {
  return (
    <Card
      title={title}
      className="p-3 rounded-md bg-bgBase"
      titleClassName="font-thin text-secondaryTextColor mb-3"
      divider={false}
    >
      <InfoSection showDivider={false} showOutline={false}>
        {children}
      </InfoSection>
    </Card>
  )
}
export default InfoSectionCard
