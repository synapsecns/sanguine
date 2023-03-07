export function getWalletStyle(walletId) {
  switch (walletId) {
    case 'metamask':
      return 'hover:!border-orange-500'
    case 'walletconnect':
      return 'hover:!border-sky-500'
    case 'binancewallet':
      return 'hover:!border-yellow-500'
    default:
      return 'hover:!border-blue-500'
  }
}
