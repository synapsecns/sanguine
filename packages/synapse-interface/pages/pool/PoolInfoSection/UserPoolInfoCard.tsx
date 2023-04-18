import { commify } from '@ethersproject/units'

import { formatBNToPercentString, formatBNToString } from '@bignumber/format'

import InfoListItem from '@components/InfoListItem'
import InfoSectionCard from './InfoSectionCard'

function UserPoolInfoCard({ data }) {
  if (!data) return null


  const share = formatBNToPercentString(data.share, 18, 4)
  const value = commify(formatBNToString(data.value, 18, 6))


  return (
    <InfoSectionCard title="My Pool LP">
      <InfoListItem
        labelText="Percent of Pool"
        content={share}
      />
      <InfoListItem
        labelText="Total amount:"
        content={value}
      />
    </InfoSectionCard>
  )
}

export default UserPoolInfoCard
