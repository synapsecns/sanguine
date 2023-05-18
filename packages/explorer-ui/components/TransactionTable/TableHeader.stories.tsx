import React from 'react'

import { TableHeader } from './TableHeader'

export default {
  component: TableHeader,
  title: 'TableHeader',
}

const Template = (args) => <TableHeader {...args} />

export const Default = Template.bind({})
Default.args = {
  headers: [
    'From',
    'To',
    'Initial',
    'Final',
    'Origin',
    'Destination',
    'Date',
    'Tx',
    'ID',
  ],
}
