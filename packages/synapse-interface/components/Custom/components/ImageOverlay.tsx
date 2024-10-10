import React from 'react'

export const ImageOverlayComponent = ({
  bigImageSrc,
  smallImageSrc,
  altTextBig,
  altTextSmall,
}) => {
  return (
    <div className="relative">
      <img
        src={bigImageSrc}
        alt={altTextBig}
        className="object-cover w-5 h-5"
      />
      <img
        src={smallImageSrc}
        alt={altTextSmall}
        className="absolute bottom-0 object-cover w-3 h-3 -right-1"
      />
    </div>
  )
}
