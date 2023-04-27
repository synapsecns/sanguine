import Card from '@tw/Card'
import InfoSection from './InfoSection'

const InfoSectionCard = ({ title, children }) => {
  return (
    <Card
      title={title}
      className="p-6 rounded-3xl bg-bgBase"
      titleClassName="text-base font-base text-secondaryTextColor text-opacity-50 mb-5"
      divider={false}
    >
      <InfoSection showDivider={true} showOutline={false}>
        {children}
      </InfoSection>
    </Card>
  )
}
export default InfoSectionCard
