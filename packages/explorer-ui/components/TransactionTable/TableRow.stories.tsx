import React from 'react'

import { TableRow } from './TableRow'

export default {
  component: TableRow,
  title: 'TableRow',
}

const Template = (args) => <TableRow {...args} />

export const Default = Template.bind({})
Default.args = {
  items: [
    '0x1e79...7ebd',
    '0x1e79...7ebd',
    '1.070 JEWEL',
    '0.070 JEWEL',
    'DFK Chain',
    'Klaytn',
    'about 1 hour ago',
    '1d854d...b2d5',
  ],
  key: '1d854d...b2d5',
}
