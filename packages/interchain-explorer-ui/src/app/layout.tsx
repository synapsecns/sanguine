import type { Metadata } from 'next'
import './globals.css'
import { GqlProvider } from '@/providers/GqlProvider'

export const metadata: Metadata = {
  title: 'Interchain Explorer',
  description: 'An explorer for the Synapse Interchain Network',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en">
      <body>
        <GqlProvider>{children}</GqlProvider>
      </body>
    </html>
  )
}
