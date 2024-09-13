import express from 'express'

import { tokenListController } from '../controllers/tokensListController'

const router = express.Router()

router.get('/', tokenListController)

export default router
