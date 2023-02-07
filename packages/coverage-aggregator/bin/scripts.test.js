jest.mock('child_process')
import path from 'path'
import { fork } from 'child_process'

import { execute, COLLECT_FILES } from './aggregate-coverage'
describe('pedalboard-scripts', () => {
  it('should execute the collectFiles', () => {
    execute({
      command: COLLECT_FILES,
      commandArgs: ['mock', 'args'],
    })
    expect(fork).toHaveBeenCalledWith(
      path.resolve(__dirname, '../src/collect-files.js'),
      ['mock', 'args']
    )
  })
})
