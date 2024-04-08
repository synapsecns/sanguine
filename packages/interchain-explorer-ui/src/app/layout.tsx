import type { Metadata } from 'next'

import { GqlProvider } from '@/providers/GqlProvider'
import './globals.css'

export const metadata: Metadata = {
  title: 'Interchain Explorer',
  description: 'An explorer for the Synapse Interchain Network',
}

const RootLayout = ({
  children,
}: Readonly<{
  children: React.ReactNode
}>) => {
  return (
    <html lang="en">
      <body>
        <GqlProvider>{children}</GqlProvider>
      </body>
    </html>
  )
}

export default RootLayout
