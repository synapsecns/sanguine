import {Head, Html} from 'next/document'

export default function Document() {
  return (
    // @ts-expect-error TS(2749): 'Html' refers to a value, but is being used as a t... Remove this comment to see the full error message
    <Html>
      <Head>
        // @ts-expect-error TS(2304): Cannot find name 'meta'.
        <meta charset="utf-8" />
        // @ts-expect-error TS(2304): Cannot find name 'link'.
        <link rel="icon" href="%PUBLIC_URL%/favicon.ico" />
        // @ts-expect-error TS(2304): Cannot find name 'meta'.
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        // @ts-expect-error TS(2304): Cannot find name 'meta'.
        <meta name="theme-color" content="#000000" />
        // @ts-expect-error TS(2304): Cannot find name 'meta'.
        <meta
          name="description"
          // @ts-expect-error TS(2304): Cannot find name 'content'.
          content="Bridge Explorer for Synapse Protocol"
        />
        // @ts-expect-error TS(2304): Cannot find name 'link'.
        <link rel="manifest" href="%PUBLIC_URL%/manifest.json" />
        // @ts-expect-error TS(2304): Cannot find name 'title'.
        <title>Synapse Explorer</title>
      </Head>
    </Html>
  )
}
