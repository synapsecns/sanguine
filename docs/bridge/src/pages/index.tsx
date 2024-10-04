// import clsx from 'clsx';
// import Link from '@docusaurus/Link';
// import useDocusaurusContext from '@docusaurus/useDocusaurusContext'
// import Layout from '@theme/Layout'
// import HomepageFeatures from '@site/src/components/HomepageFeatures'
// import Heading from '@theme/Heading'
import { Redirect } from '@docusaurus/router'

// import styles from './index.module.css'

// const HomepageHeader = () => {
//   const { siteConfig } = useDocusaurusContext()
//   return (
//     <header className={clsx('hero hero--primary', styles.heroBanner)}>
//       <div className="container">
//         <Heading as="h1" className="hero__title">
//           {siteConfig.title}
//         </Heading>
//         <p className="hero__subtitle">{siteConfig.tagline}</p>
//         <div className={styles.buttons}>
//           <Link className="button button--secondary button--lg" to="#">
//             Docusaurus Tutorial - 5min ⏱️
//           </Link>
//         </div>
//       </div>
//     </header>
//   )
// }

// export default () => {
//   const { siteConfig } = useDocusaurusContext()
//   // TODO: a homepage
//   // for now, just disable entirely: https://v1.docusaurus.io/docs/en/site-creation#docs-landing-page
//   return (
//     <Layout
//       title={`Hello from ${siteConfig.title}`}
//       description="Description will go into a meta tag in <head />"
//     >
//       <Redirect to={'/docs'} />
//       {/* <HomepageHeader />
//       <main>
//         <HomepageFeatures />
//       </main> */}
//     </Layout>
//   )
// }

export default () => <Redirect to='/docs/About/' />
