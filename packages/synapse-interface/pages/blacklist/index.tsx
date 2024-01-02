import { EXCLUDED_ADDRESSES } from '@/constants/blacklist'

export default () => {
  return (
    <div>
      <ul>
        {EXCLUDED_ADDRESSES.map((address, index) => (
          <li key={index}>{address}</li>
        ))}
      </ul>
    </div>
  )
}

export async function getServerSideProps(context) {
  if (context.req.headers['user-agent'].includes('curl')) {
    context.res.setHeader('Content-Type', 'text/plain')
    context.res.write(EXCLUDED_ADDRESSES.join('\n'))
    context.res.end()
    return { props: {} }
  }

  return {
    props: {
      addresses: EXCLUDED_ADDRESSES,
    },
  }
}
