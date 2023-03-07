import Vibrant from 'node-vibrant'

import { useLayoutEffect, useEffect, useState } from 'react'

import { ChainId } from '@constants/networks'

import { uriToHttp } from '@urls/uriToHttp'



async function getColorFromUriPath(uri) {
  const formattedPath = uriToHttp(uri)[0]

  try {
    return Vibrant.from(formattedPath)
      .getPalette()
      .then(palette => {
        console.log(palette)
        if (palette?.Vibrant) {
          return palette.Vibrant.hex
        }
        return null
      })
      .catch(() => null)
  } catch (e) {
    console.error(e)
    return
  }
}

async function getColorsFromUriPath(uri) {
  const formattedPath = uriToHttp(uri)[0]

  try {
    return Vibrant.from(formattedPath)
      .getPalette()
      .then(palette => {
        let aggregatedSwatch = {}
        for (const [vibrantKey, swatch] of Object.entries(palette))  {
          aggregatedSwatch[vibrantKey] = swatch.hex
        }
        return aggregatedSwatch
      })
      .catch( e => null)
  } catch (e) {
    console.error(e)
    return
  }
}

const DEFAULT_COLOR_PALETTE = {}
export function useImgPalette(listImageUri) {
  const [colorPalette, setColorPalette] = useState(DEFAULT_COLOR_PALETTE)

  useLayoutEffect(() => {
    let stale = false

    if (listImageUri) {
      getColorsFromUriPath(listImageUri).then(calculatedPalette => {
        if (!stale && calculatedPalette !== null) {
          setColorPalette(calculatedPalette)
        }
      })
    }

    return () => {
      stale = true
      setColorPalette(DEFAULT_COLOR_PALETTE)
    }
  }, [listImageUri])

  return colorPalette
}


export function useImgColor(listImageUri) {
  const [color, setColor] = useState('#0094ec')

  useLayoutEffect(() => {
    let stale = false

    if (listImageUri) {
      getColorFromUriPath(listImageUri).then(color => {
        if (!stale && color !== null) {
          setColor(color)
        }
      })
    }

    return () => {
      stale = true
      setColor('#0094ec')
    }
  }, [listImageUri])

  return color
}