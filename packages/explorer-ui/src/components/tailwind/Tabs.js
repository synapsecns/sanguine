import Grid from '@tw/Grid'

export default function Tabs({ children, numTabs = 2, ...props }) {
  return (
    <Grid gap={2} aria-label="Tabs" cols={{ xs: 2, sm: numTabs }} {...props}>
      {children}
    </Grid>
  )
}
