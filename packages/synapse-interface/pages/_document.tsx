import { Head, Html, Main, NextScript } from 'next/document'
import HotJar from '@/components/HotJar/HotJar'

const Document = () => {
  return (
    <Html lang="en">
      <Head>
        <HotJar />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
export default Document
