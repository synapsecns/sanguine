import Grid from '@components/tailwind/Grid'

export default function Tabs({ children, numTabs = 2, ...props }) {
  return (
    // @ts-expect-error TS(2749): 'Grid' refers to a value, but is being used as a t... Remove this comment to see the full error message
    <Grid gap={2} aria-label="Tabs" cols={{ xs: 2, sm: numTabs }} {...props}>
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      {children}
    </Grid>
  )
}
