// function NetworkPausedButton({ networkName }) {
//   return (
//     <Button disabled={true} type="button" className={ACTION_BTN_CLASSNAME}>
//       {networkName} Undergoing Chain Downtime
//     </Button>
//   )
// }
// // Undergoing Network Upgrades

// const PAUSED_BASE_PROPERTIES = `
//     w-full rounded-lg my-2 px-4 py-3
//     text-white text-opacity-100 transition-all
//     hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
//     disabled:from-bgLight disabled:to-bgLight
//     bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
//   `

// function PausedButton({ networkName }) {
//   return (
//     <Button
//       disabled={true}
//       fancy={true}
//       type="button"
//       className={`${PAUSED_BASE_PROPERTIES}`}
//     >
//      Temporarily paused due to chain connectivity issues
//     </Button>
//   )
// }

// function HeavyLoadButton() {
//   return (
//     <Button
//       disabled={true}
//       fancy={true}
//       type="button"
//       className={ACTION_BTN_CLASSNAME}
//     >
//       Synapse is experiencing heavy load
//     </Button>
//   )
// }

// function AdvancedOptionsButton({ className, onClick }) {
//   return (
//     <div
//       className={`
//         group rounded-lg hover:bg-gray-900 ${className} p-1`}
//       onClick={onClick}
//     >
//       <CogIcon className="w-6 h-6 text-gray-500 group-hover:text-gray-300" />
//     </div>
//   )
// }
