export function PageLink({ text, url, external = false }) {
  if (external) {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="mt-2 mb-14">
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="text-white text-opacity-100 hover:text-opacity-90 hover:underline"
          // @ts-expect-error TS(2304): Cannot find name 'href'.
          href={url}
          // @ts-expect-error TS(2304): Cannot find name 'target'.
          target="_blank"
          // @ts-expect-error TS(2304): Cannot find name 'rel'.
          rel="noreferrer"
        >
          {text}
        </a>
      </div>
    )
  } else {
    return (
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="mt-2 mb-14 ">
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a href={url} className="text-white text-opacity-100 hover:text-opacity-90 hover:underline">
          {text}
        </a>
      </div>
    )
  }
}
