import type { Metadata } from 'next'
import '@/styles/globals.css'

export const metadata: Metadata = {
  title: 'Synapse Widget Landing Page',
  description: 'Example Synapse Bridge widget',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body id="root">{children}</body>
    </html>
  )
}
