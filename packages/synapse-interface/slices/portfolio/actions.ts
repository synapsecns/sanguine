import { createAction } from '@reduxjs/toolkit'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const setActiveTab = createAction<PortfolioTabs>(
  'portfolio/setActiveTab'
)
