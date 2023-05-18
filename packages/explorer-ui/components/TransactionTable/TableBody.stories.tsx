import React from 'react'

import { TableBody } from './TableBody'

export default {
  component: TableBody,
  title: 'TableBody',
}

const Template = (args) => <TableBody {...args} />

export const Default = Template.bind({})
Default.args = {
  rows: [
    {
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
    },
    {
      items: [
        '0x2023...9163',
        '0x2023...9163',
        '0.068 nETH',
        '0.065 nETH',
        'Arbitrum',
        'Optimism',
        'about 1 hour ago',
        '83bcbd...0d44',
      ],
      key: '83bcbd...0d44',
    },
  ],
}
