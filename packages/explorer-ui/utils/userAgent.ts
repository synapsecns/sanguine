import { UAParser } from 'ua-parser-js'
import { useContext } from 'react'


const parser = new UAParser(window.navigator.userAgent)
const { type } = parser.getDevice()

export const isMobile = type === 'mobile' || type === 'tablet'
const platform = parser.getOS().name
export const isIOS = platform === 'iOS'
export const isNonIOSPhone = !isIOS && type === 'mobile'
