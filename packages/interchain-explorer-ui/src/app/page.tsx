'use client'

import { SearchInput } from '@/components/SearchInput'
import { Stats } from '@/components/Stats'
import { TransactionsTable } from '@/components/TransactionsTable'

const Home = () => {
  return (
    <div className="bg-gray-900 text-white min-h-screen">
      <div className="container mx-auto p-4">
        <div className="flex justify-between items-center mb-4">
          <SearchInput />
        </div>

        <Stats />

        <TransactionsTable />
      </div>
    </div>
  )
}

export default Home
