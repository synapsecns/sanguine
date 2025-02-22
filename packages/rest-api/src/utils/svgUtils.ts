import { Buffer } from 'buffer'

import { JSDOM } from 'jsdom'

const XML_HEADER = '<?xml version="1.0" encoding="UTF-8" standalone="no"?>'
const DOCTYPE =
  '<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">'

export const addSvgHeaderIfMissing = (buffer: ArrayBuffer): Buffer => {
  const content = Buffer.from(buffer).toString('utf-8')

  // Create a DOM with the SVG content
  const dom = new JSDOM(content, { contentType: 'image/svg+xml' })
  const document = dom.window.document

  // Get the serialized content to check if we need to add headers
  const serialized = dom.serialize()
  const needsXmlHeader = !serialized.includes('<?xml')
  const needsDoctype = !serialized.includes('<!DOCTYPE')

  if (!needsXmlHeader && !needsDoctype) {
    return Buffer.from(content)
  }

  // Build the final SVG with required headers
  let finalSvg = ''
  if (needsXmlHeader) {
    finalSvg += XML_HEADER + '\n'
  }
  if (needsDoctype) {
    finalSvg += DOCTYPE + '\n'
  }

  // Add the SVG content
  const svgElement = document.querySelector('svg')
  if (svgElement) {
    finalSvg += svgElement.outerHTML
  } else {
    // If no SVG element found, return original content
    return Buffer.from(content)
  }

  return Buffer.from(finalSvg)
}
