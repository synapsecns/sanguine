import type { SidebarsConfig } from '@docusaurus/plugin-content-docs'

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: 'category',
      label: 'quotes',
      items: [
        {
          type: 'doc',
          id: 'rfq/API/get-quotes',
          label: 'Get quotes',
          className: 'api-method get',
        },
        {
          type: 'doc',
          id: 'rfq/API/upsert-quote',
          label: 'Upsert quote',
          className: 'api-method put',
        },
      ],
    },
  ],
}

export default sidebar.apisidebar
