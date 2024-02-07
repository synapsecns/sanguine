import {Head, Html} from 'next/document'

export default function Document() {
  return (
    <Html>
      <Head>
        <title>Synapse Explorer</title>
        <meta charset="utf-8" />
        <link rel="icon" href="%PUBLIC_URL%/favicon.ico" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="theme-color" content="#000000" />
        <meta
          name="description"
          content="Bridge Explorer for Synapse Protocol"
        />
        <link rel="manifest" href="%PUBLIC_URL%/manifest.json" />

      </Head>
    </Html>
  )
}
