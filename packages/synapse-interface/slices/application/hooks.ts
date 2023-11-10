import { RootState } from '@/store/store'
import { useAppSelector } from '@/store/hooks'

export const useApplicationState = (): RootState['application'] => {
  return useAppSelector((state) => state.application)
}
