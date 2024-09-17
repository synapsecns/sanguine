import express from 'express'

import { tokenListController } from '../controllers/tokenListController'

const router = express.Router()

router.get('/', tokenListController)

export default router
