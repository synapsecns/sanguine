import { Head, Html, Main, NextScript } from 'next/document'
export const runtime = 'experimental-edge'

const Document = () => {
  return (
    <Html lang="en">
      <Head />
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
export default Document
