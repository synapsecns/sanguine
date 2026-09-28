import { act, renderHook } from '@testing-library/react'

import { useStaleQuoteUpdater } from '@/components/StateManagedBridge/hooks/useStaleQuoteUpdater'

const quote = (requestId: number) => ({ requestId } as any)

describe('interface stale quote refresh', () => {
  beforeEach(() => jest.useFakeTimers())
  afterEach(() => jest.useRealTimers())

  it('cleans stale listeners on disable and unmount', async () => {
    const refresh = jest.fn().mockResolvedValue(undefined)
    const currentQuote = quote(1)
    const { rerender, unmount, result } = renderHook(
      ({ enabled }) =>
        useStaleQuoteUpdater(currentQuote, refresh, enabled, 15000, 0),
      { initialProps: { enabled: true } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    expect(result.current).toBe(true)
    rerender({ enabled: false })
    await act(async () => {
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).not.toHaveBeenCalled()
    rerender({ enabled: true })
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    unmount()
    await act(async () => {
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).not.toHaveBeenCalled()
  })

  it('uses the latest callback and refreshes once for repeated mouse events', async () => {
    const previous = jest.fn().mockResolvedValue(undefined)
    const latest = jest.fn().mockResolvedValue(undefined)
    const currentQuote = quote(1)
    const { rerender } = renderHook(
      ({ refresh }) =>
        useStaleQuoteUpdater(currentQuote, refresh, true, 15000, 0),
      { initialProps: { refresh: previous } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    rerender({ refresh: latest })
    await act(async () => {
      document.dispatchEvent(new MouseEvent('mousemove'))
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(previous).not.toHaveBeenCalled()
    expect(latest).toHaveBeenCalledTimes(1)
  })

  it('invalidates callbacks from replaced quotes', async () => {
    const added = jest.spyOn(document, 'addEventListener')
    const refresh = jest.fn().mockResolvedValue(undefined)
    const { rerender } = renderHook(
      ({ currentQuote }) =>
        useStaleQuoteUpdater(currentQuote, refresh, true, 15000, 0),
      { initialProps: { currentQuote: quote(1) } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    const obsolete = added.mock.calls.findLast(
      ([name]) => name === 'mousemove'
    )[1] as () => void
    rerender({ currentQuote: quote(2) })
    await act(async () => {
      obsolete()
    })
    expect(refresh).not.toHaveBeenCalled()
    await act(async () => {
      jest.advanceTimersByTime(15000)
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).toHaveBeenCalledTimes(1)
    added.mockRestore()
  })

  it('lets replacement quotes become stale and refresh while an older refresh never settles', async () => {
    const refresh = jest
      .fn()
      .mockReturnValue(new Promise<void>(() => undefined))
    const { rerender, result } = renderHook(
      ({ currentQuote }) =>
        useStaleQuoteUpdater(currentQuote, refresh, true, 15000, 0),
      { initialProps: { currentQuote: quote(1) } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).toHaveBeenCalledTimes(1)

    rerender({ currentQuote: quote(2) })
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    expect(result.current).toBe(true)
    await act(async () => {
      document.dispatchEvent(new MouseEvent('mousemove'))
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).toHaveBeenCalledTimes(2)
  })

  it('does not let an old completion clear the replacement cycle pending refresh', async () => {
    const added = jest.spyOn(document, 'addEventListener')
    let finishOld: () => void
    let finishCurrent: () => void
    const refresh = jest
      .fn()
      .mockImplementationOnce(
        () =>
          new Promise<void>((done) => {
            finishOld = done
          })
      )
      .mockImplementationOnce(
        () =>
          new Promise<void>((done) => {
            finishCurrent = done
          })
      )
      .mockResolvedValue(undefined)
    const { rerender } = renderHook(
      ({ currentQuote }) =>
        useStaleQuoteUpdater(currentQuote, refresh, true, 15000, 0),
      { initialProps: { currentQuote: quote(1) } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    rerender({ currentQuote: quote(2) })
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    const currentRefresh = added.mock.calls.findLast(
      ([name]) => name === 'mousemove'
    )[1] as () => void
    await act(async () => {
      currentRefresh()
    })
    expect(refresh).toHaveBeenCalledTimes(2)
    await act(async () => {
      finishOld()
    })
    await act(async () => {
      currentRefresh()
    })
    expect(refresh).toHaveBeenCalledTimes(2)
    await act(async () => {
      finishCurrent()
    })
    added.mockRestore()
  })

  it('keeps the existing automatic refresh window before requiring activity', async () => {
    const refresh = jest.fn().mockResolvedValue(undefined)
    const { rerender, result } = renderHook(
      ({ currentQuote }) => useStaleQuoteUpdater(currentQuote, refresh, true),
      { initialProps: { currentQuote: quote(1) } }
    )
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    expect(refresh).toHaveBeenCalledTimes(1)
    rerender({ currentQuote: quote(2) })
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    expect(refresh).toHaveBeenCalledTimes(2)
    rerender({ currentQuote: quote(3) })
    await act(async () => {
      jest.advanceTimersByTime(15000)
    })
    expect(result.current).toBe(true)
    expect(refresh).toHaveBeenCalledTimes(2)
    await act(async () => {
      document.dispatchEvent(new MouseEvent('mousemove'))
    })
    expect(refresh).toHaveBeenCalledTimes(3)
  })

  it.each(['resolve', 'reject'])(
    'rearms when a refresh %s does not replace the quote',
    async (outcome) => {
      const refresh = jest.fn()
      if (outcome === 'reject') refresh.mockRejectedValue(new Error('offline'))
      else refresh.mockResolvedValue(undefined)
      const currentQuote = quote(1)
      const { result } = renderHook(() =>
        useStaleQuoteUpdater(currentQuote, refresh, true)
      )
      await act(async () => {
        jest.advanceTimersByTime(15000)
      })
      expect(refresh).toHaveBeenCalledTimes(1)
      await act(async () => {
        jest.advanceTimersByTime(15000)
      })
      expect(refresh).toHaveBeenCalledTimes(2)
      await act(async () => {
        jest.advanceTimersByTime(15000)
      })
      expect(result.current).toBe(true)
      await act(async () => {
        document.dispatchEvent(new MouseEvent('mousemove'))
      })
      expect(refresh).toHaveBeenCalledTimes(3)
    }
  )
})
