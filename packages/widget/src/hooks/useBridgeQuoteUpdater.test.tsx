import { useEffect, useLayoutEffect } from 'react'
import { act, renderHook } from '@testing-library/react'

import { useBridgeQuoteUpdater } from '@/hooks/useBridgeQuoteUpdater'
import {
  EMPTY_BRIDGE_QUOTE,
  type BridgeQuote,
} from '@/state/slices/bridgeQuote/reducer'

const STALE_TIMEOUT = 1000

const createQuote = (overrides: Partial<BridgeQuote> = {}): BridgeQuote => ({
  ...EMPTY_BRIDGE_QUOTE,
  outputAmount: 1n,
  outputAmountString: '1.00',
  delta: 1n,
  requestId: 1,
  timestamp: 1,
  ...overrides,
})

const dispatchMouseMove = () => {
  act(() => {
    document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
  })
}

type EventListenerSpy = {
  mock: {
    calls: Array<[string, EventListenerOrEventListenerObject, ...unknown[]]>
  }
}

const getMouseMoveListenerCalls = (spy: EventListenerSpy) =>
  spy.mock.calls.filter(([eventName]) => eventName === 'mousemove')

const getLatestMouseMoveListener = (spy: EventListenerSpy) =>
  getMouseMoveListenerCalls(spy).at(-1)?.[1]

describe('useBridgeQuoteUpdater', () => {
  beforeEach(() => {
    jest.useFakeTimers()
  })

  afterEach(() => {
    jest.clearAllTimers()
    jest.useRealTimers()
    jest.restoreAllMocks()
    jest.clearAllMocks()
  })

  it('keeps a fresh quote fresh for a full timeout from arrival', () => {
    const refreshQuote = jest.fn(async () => undefined)

    renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: createQuote() },
      }
    )

    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 1)
    })
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(1)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('gives the first valid quote a full fresh window after a long idle period', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const { rerender } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: EMPTY_BRIDGE_QUOTE },
      }
    )

    act(() => {
      jest.advanceTimersByTime(5000)
    })

    rerender({
      quote: createQuote({ requestId: 2, timestamp: 2 }),
    })

    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('fires exactly one refresh for each stale cycle', () => {
    const refreshQuote = jest.fn(async () => undefined)

    renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: createQuote() },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    dispatchMouseMove()
    dispatchMouseMove()

    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('does not leave a refreshed quote stale immediately after replacement', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const quoteA = createQuote({ requestId: 1, timestamp: 1 })
    const quoteB = createQuote({ requestId: 2, timestamp: 2 })
    const { rerender } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: quoteA },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)

    rerender({ quote: quoteB })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(2)
  })

  it('clears stale listeners and timers when quote loading starts', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const removeEventListenerSpy = jest.spyOn(document, 'removeEventListener')
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, isQuoteLoading }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          isQuoteLoading,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, isQuoteLoading: false },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    const staleListener = getLatestMouseMoveListener(addEventListenerSpy)

    rerender({ quote, isQuoteLoading: true })
    dispatchMouseMove()

    expect(staleListener).toBeDefined()
    expect(getMouseMoveListenerCalls(removeEventListenerSpy)).toContainEqual([
      'mousemove',
      staleListener,
    ])
    expect(jest.getTimerCount()).toBe(0)
    expect(refreshQuote).not.toHaveBeenCalled()

    rerender({ quote, isQuoteLoading: false })
    dispatchMouseMove()

    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()

    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('clears stale listeners and timers when wallet pending starts', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const removeEventListenerSpy = jest.spyOn(document, 'removeEventListener')
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, isWalletPending }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          isWalletPending,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, isWalletPending: false },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    const staleListener = getLatestMouseMoveListener(addEventListenerSpy)

    rerender({ quote, isWalletPending: true })
    dispatchMouseMove()

    expect(staleListener).toBeDefined()
    expect(getMouseMoveListenerCalls(removeEventListenerSpy)).toContainEqual([
      'mousemove',
      staleListener,
    ])
    expect(jest.getTimerCount()).toBe(0)
    expect(refreshQuote).not.toHaveBeenCalled()

    rerender({ quote, isWalletPending: false })
    dispatchMouseMove()

    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()

    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('gives wallet pending interruptions a new full stale window on resume', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, isWalletPending }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          isWalletPending,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, isWalletPending: false },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 100)
    })

    rerender({ quote, isWalletPending: true })

    act(() => {
      jest.advanceTimersByTime(200)
    })

    expect(getMouseMoveListenerCalls(addEventListenerSpy)).toHaveLength(0)
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    rerender({ quote, isWalletPending: false })

    dispatchMouseMove()
    expect(getMouseMoveListenerCalls(addEventListenerSpy)).toHaveLength(0)
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 1)
    })

    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(1)
    })

    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('gives quote loading interruptions a new full stale window on resume', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, isQuoteLoading }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          isQuoteLoading,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, isQuoteLoading: false },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 100)
    })

    rerender({ quote, isQuoteLoading: true })

    act(() => {
      jest.advanceTimersByTime(200)
    })

    expect(getMouseMoveListenerCalls(addEventListenerSpy)).toHaveLength(0)
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    rerender({ quote, isQuoteLoading: false })

    dispatchMouseMove()
    expect(getMouseMoveListenerCalls(addEventListenerSpy)).toHaveLength(0)
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 1)
    })

    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(1)
    })

    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('removes the stale listener when the quote becomes invalid', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const removeEventListenerSpy = jest.spyOn(document, 'removeEventListener')
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    const staleListener = getLatestMouseMoveListener(addEventListenerSpy)

    rerender({ quote: EMPTY_BRIDGE_QUOTE })
    dispatchMouseMove()

    expect(staleListener).toBeDefined()
    expect(getMouseMoveListenerCalls(removeEventListenerSpy)).toContainEqual([
      'mousemove',
      staleListener,
    ])
    expect(jest.getTimerCount()).toBe(0)
    expect(refreshQuote).not.toHaveBeenCalled()
  })

  it('closes the old stale cycle when the active quote is replaced', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const removeEventListenerSpy = jest.spyOn(document, 'removeEventListener')
    const quoteA = createQuote({ requestId: 1, timestamp: 1 })
    const quoteB = createQuote({ requestId: 2, timestamp: 2 })
    const { rerender } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: quoteA },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    const staleListener = getLatestMouseMoveListener(addEventListenerSpy)

    rerender({ quote: quoteB })
    dispatchMouseMove()

    expect(staleListener).toBeDefined()
    expect(getMouseMoveListenerCalls(removeEventListenerSpy)).toContainEqual([
      'mousemove',
      staleListener,
    ])
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('does not let quote A arm stale work after quote B replaces it before A becomes stale', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const addEventListenerSpy = jest.spyOn(document, 'addEventListener')
    const quoteA = createQuote({ requestId: 1, timestamp: 1 })
    const quoteB = createQuote({ requestId: 2, timestamp: 2 })
    const { rerender } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: quoteA },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 100)
    })

    rerender({ quote: quoteB })

    act(() => {
      jest.advanceTimersByTime(100)
    })

    expect(getMouseMoveListenerCalls(addEventListenerSpy)).toHaveLength(0)

    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 100)
    })

    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('does not re-arm the fresh quote timer on background rerenders', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { rerender } = renderHook(
      (props) => {
        useBridgeQuoteUpdater(
          props.quote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, rerenderTick: 0 },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT - 100)
    })

    rerender({ quote, rerenderTick: 1 })
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    act(() => {
      jest.advanceTimersByTime(100)
    })
    dispatchMouseMove()
    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('cleans up stale timers and listeners on unmount', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { unmount } = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote },
      }
    )

    expect(jest.getTimerCount()).toBe(1)
    unmount()

    expect(jest.getTimerCount()).toBe(0)
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()

    const staleHook = renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    staleHook.unmount()
    dispatchMouseMove()
    expect(refreshQuote).not.toHaveBeenCalled()
  })

  it('does not duplicate refreshes when rerenders happen after the quote is stale', () => {
    const refreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { rerender } = renderHook(
      (props) => {
        useBridgeQuoteUpdater(
          props.quote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, rerenderTick: 0 },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    rerender({ quote, rerenderTick: 1 })
    rerender({ quote, rerenderTick: 2 })

    dispatchMouseMove()
    dispatchMouseMove()

    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })

  it('uses the latest refresh callback for an already stale quote', () => {
    const initialRefreshQuote = jest.fn(async () => undefined)
    const updatedRefreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, refreshQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote, refreshQuote: initialRefreshQuote },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    rerender({ quote, refreshQuote: updatedRefreshQuote })
    dispatchMouseMove()

    expect(initialRefreshQuote).not.toHaveBeenCalled()
    expect(updatedRefreshQuote).toHaveBeenCalledTimes(1)
  })

  it('uses the latest refresh callback for layout-phase activity after a callback rerender', () => {
    const initialRefreshQuote = jest.fn(async () => undefined)
    const updatedRefreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, refreshQuote, triggerLayoutActivity }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )

        useLayoutEffect(() => {
          if (triggerLayoutActivity) {
            document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
          }
        }, [triggerLayoutActivity])
      },
      {
        initialProps: {
          quote,
          refreshQuote: initialRefreshQuote,
          triggerLayoutActivity: false,
        },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    rerender({
      quote,
      refreshQuote: updatedRefreshQuote,
      triggerLayoutActivity: true,
    })

    expect(initialRefreshQuote).not.toHaveBeenCalled()
    expect(updatedRefreshQuote).toHaveBeenCalledTimes(1)
  })

  it('uses the latest refresh callback for passive-effect activity after a callback rerender', () => {
    const initialRefreshQuote = jest.fn(async () => undefined)
    const updatedRefreshQuote = jest.fn(async () => undefined)
    const quote = createQuote()
    const { rerender } = renderHook(
      ({ quote: currentQuote, refreshQuote, triggerPassiveActivity }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )

        useEffect(() => {
          if (triggerPassiveActivity) {
            document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
          }
        }, [triggerPassiveActivity])
      },
      {
        initialProps: {
          quote,
          refreshQuote: initialRefreshQuote,
          triggerPassiveActivity: false,
        },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    rerender({
      quote,
      refreshQuote: updatedRefreshQuote,
      triggerPassiveActivity: true,
    })

    expect(initialRefreshQuote).not.toHaveBeenCalled()
    expect(updatedRefreshQuote).toHaveBeenCalledTimes(1)
  })

  it('keeps a rejected stale refresh one-shot until a new quote arrives', () => {
    const refreshError = new Error('refresh failed')
    const refreshQuote = jest.fn(() => {
      const rejectedPromise = Promise.reject(refreshError)
      rejectedPromise.catch(() => undefined)
      return rejectedPromise
    })

    renderHook(
      ({ quote: currentQuote }) => {
        useBridgeQuoteUpdater(
          currentQuote,
          refreshQuote,
          false,
          false,
          STALE_TIMEOUT
        )
      },
      {
        initialProps: { quote: createQuote() },
      }
    )

    act(() => {
      jest.advanceTimersByTime(STALE_TIMEOUT)
    })

    dispatchMouseMove()
    dispatchMouseMove()

    expect(refreshQuote).toHaveBeenCalledTimes(1)
  })
})
