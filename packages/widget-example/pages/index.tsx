import dynamic from 'next/dynamic'

const WidgetExampleClient = dynamic(
  () => import('../components/WidgetExampleClient'),
  {
    loading: () => (
      <div className="widget-status-panel">
        Loading the browser wallet integration...
      </div>
    ),
    ssr: false,
  }
)

const Home = () => {
  return (
    <main className="page-shell">
      <section className="content-shell">
        <div className="hero-copy">
          <span className="eyebrow">Synapse Widget</span>
          <h1>Widget Example</h1>
          <p className="lede">
            Minimal Next.js consumer package for rendering the bridge widget
            from this monorepo.
          </p>
        </div>
        <WidgetExampleClient />
      </section>
    </main>
  )
}

export default Home
