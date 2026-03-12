import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "category",
      label: "ack",
      items: [
        {
          type: "doc",
          id: "04-Routers/RFQ/API/relay-ack",
          label: "Relay ack",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "quotes",
      items: [
        {
          type: "doc",
          id: "04-Routers/RFQ/API/upsert-quotes",
          label: "Upsert quotes",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/get-contract-addresses",
          label: "Get contract addresses",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/get-open-quote-requests",
          label: "Get open quote requests",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/get-quotes",
          label: "Get quotes",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/upsert-quote",
          label: "Upsert quote",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/handle-user-quote-request",
          label: "Initiate an Active RFQ",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "04-Routers/RFQ/API/handle-web-socket-connection-for-active-quote-requests",
          label: "Active RFQ Listener",
          className: "api-method get",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
