jest.mock('child_process')
import path from 'path'
import { fork } from 'child_process'

import { execute, COLLECT_FILES } from './aggregate-coverage'
describe('scripts', () => {
  it('should execute the collectFiles', () => {
    execute({
      command: COLLECT_FILES,
      commandArgs: ['mock', 'args'],
    })
    expect(fork).toHaveBeenCalledWith(
      path.resolve(__dirname, '../lib/coverage-aggregator.js'),
      ['mock', 'args']
    )
  })
})
