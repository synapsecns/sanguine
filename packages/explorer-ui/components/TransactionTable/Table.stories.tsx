import React from 'react'

import { Table } from './Table'
import { Default as TableBody } from './TableBody.stories'
import { Default as TableHeader } from './TableHeader.stories'

export default {
  component: Table,
  title: 'Table',
}

const Template = (args) => <Table {...args} />

export const Default = Template.bind({})
Default.args = {
  header: <TableHeader {...TableHeader.args} />,
  body: <TableBody {...TableBody.args} />,
}
