import { useState } from 'react'
import { useRouter } from 'next/navigation'

export default function SearchInput() {
  const [searchTerm, setSearchTerm] = useState('')
  const router = useRouter()

  const handleSearch = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    router.push(`/tx/${searchTerm}`)
  }

  return (
    <form onSubmit={handleSearch} className="w-full">
      <input
        className="px-4 py-2 rounded bg-gray-800 border border-gray-700 w-full"
        type="search"
        placeholder="Search by transaction id"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
    </form>
  )
}
