import { useLocation } from 'react-router-dom'

export function useUrlQuery() {
  const location = useLocation()
  return new URLSearchParams(location.search);
}