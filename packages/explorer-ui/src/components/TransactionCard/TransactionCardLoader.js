import Card from '@tw/Card'

export function TransactionCardLoader({ ordinal }) {
  const backgroundColor =
    ordinal % 2 === 0 ? 'bg-transparent' : 'bg-[#D9D9D9] bg-opacity-5'

  return (
    <Card
      className={`pb-5 pt-5 flex-wrap md:flex lg:flex mt-2 mb-2 rounded-none justify-center ${backgroundColor}`}
    >
      <div className="w-3/4 h-10 rounded bg-slate-400 animate-pulse"></div>
    </Card>
  )
}

export function TransactionsLoader({ number }) {
  return [...Array(number).keys()].map((i) => <TransactionCardLoader key={i} />)
}
