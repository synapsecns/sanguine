# Change Log

All notable changes to this project will be documented in this file.
See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

## [0.11.11](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.10...@synapsecns/sdk-router@0.11.11) (2025-02-25)


### Bug Fixes

* **sdk-router:** fixed fee calculation [SYN-49] ([#3533](https://github.com/synapsecns/sanguine/issues/3533)) ([5b49675](https://github.com/synapsecns/sanguine/commit/5b49675bdee175daa9abedf16804a0e62f6f6fa8))





## [0.11.10](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.9...@synapsecns/sdk-router@0.11.10) (2025-02-21)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.9](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.8...@synapsecns/sdk-router@0.11.9) (2024-12-18)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.8](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.7...@synapsecns/sdk-router@0.11.8) (2024-12-01)


### Reverts

* Revert "reducing dependency bloat (#3411)" (#3421) ([5feb5a0](https://github.com/synapsecns/sanguine/commit/5feb5a0883e297bafa328fbe5c86935ed1ed2fa5)), closes [#3411](https://github.com/synapsecns/sanguine/issues/3411) [#3421](https://github.com/synapsecns/sanguine/issues/3421)





## [0.11.7](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.6...@synapsecns/sdk-router@0.11.7) (2024-11-30)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.6](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.5...@synapsecns/sdk-router@0.11.6) (2024-11-07)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.5](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.4...@synapsecns/sdk-router@0.11.5) (2024-10-29)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.4](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.3...@synapsecns/sdk-router@0.11.4) (2024-10-11)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.11.3](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.2...@synapsecns/sdk-router@0.11.3) (2024-10-03)


### Bug Fixes

* **sdk-router:** quote unmarshall hotfix [SLT-302] ([#3223](https://github.com/synapsecns/sanguine/issues/3223)) ([e88685f](https://github.com/synapsecns/sanguine/commit/e88685f6ca763a3402ddcea9f6c465f67dea617b))





## [0.11.2](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.1...@synapsecns/sdk-router@0.11.2) (2024-09-26)


### Bug Fixes

* **sdk-router:** disable ARB airdrop tests ([#3195](https://github.com/synapsecns/sanguine/issues/3195)) ([fc6ddae](https://github.com/synapsecns/sanguine/commit/fc6ddaedf03f7769dab362f0bcdf81a3dd010516))





## [0.11.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.11.0...@synapsecns/sdk-router@0.11.1) (2024-09-04)

**Note:** Version bump only for package @synapsecns/sdk-router





# [0.11.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.10.0...@synapsecns/sdk-router@0.11.0) (2024-08-26)


### Features

* **sdk:** uuid `BridgeQuote.id` ([#2896](https://github.com/synapsecns/sanguine/issues/2896)) ([85b5f53](https://github.com/synapsecns/sanguine/commit/85b5f538034a47f513d434aac2e55979bdbe390c))





# [0.10.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.9.0...@synapsecns/sdk-router@0.10.0) (2024-08-05)


* feat(sdk-router)!: add support for FastBridgeRouterV2 (#2957) ([175e0cd](https://github.com/synapsecns/sanguine/commit/175e0cd32a7e93e12af5bf458cbad49276f98518)), closes [#2957](https://github.com/synapsecns/sanguine/issues/2957)


### BREAKING CHANGES

* The `bridgeQuote` and `allBridgeQuotes` functions now accept an options object instead of individual optional parameters.

* docs: new options

* test: update

* test: cleanup

* fix: make `createRFQDestQuery` static, expose for tests

* test: add unit tests for `createRFQDestQuery`

* feat: use FastBridgeRouterV2 address

* docs: be more explicit about smart contract integrations

* docs: add address section

* docs: strike through deprecated addresses

* feat: use newest address

* docs: more explicit language

* docs: be more explicit in README as well

* fix: update docs link, slippage wording





# [0.9.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.8.0...@synapsecns/sdk-router@0.9.0) (2024-07-24)


### Features

* **sdk-router:** add Blast, Linea support ([#2903](https://github.com/synapsecns/sanguine/issues/2903)) ([c5cf52a](https://github.com/synapsecns/sanguine/commit/c5cf52a80c6f0bc16c29a0a9496713777d49654d))





# [0.8.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.7.1...@synapsecns/sdk-router@0.8.0) (2024-07-15)


### Features

* **sdk:** Adds BSC to RFQ ([#2830](https://github.com/synapsecns/sanguine/issues/2830)) ([5fb5e8a](https://github.com/synapsecns/sanguine/commit/5fb5e8a429511da6cd271719fcd2a43dad47d1f7))





## [0.7.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.7.0...@synapsecns/sdk-router@0.7.1) (2024-07-03)

**Note:** Version bump only for package @synapsecns/sdk-router





# [0.7.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.6.0...@synapsecns/sdk-router@0.7.0) (2024-06-27)


### Features

* **sdk:** Adds origin/dest chain ids to bridge quote ([#2804](https://github.com/synapsecns/sanguine/issues/2804)) ([b045125](https://github.com/synapsecns/sanguine/commit/b0451251b48d35909bb9d976c3a6bdd03b99d6ab))





# [0.6.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.5.1...@synapsecns/sdk-router@0.6.0) (2024-05-09)


### Features

* **widget:** suppress console errors ([#2594](https://github.com/synapsecns/sanguine/issues/2594)) ([925447b](https://github.com/synapsecns/sanguine/commit/925447badd2065444d5df7cc14a6e6c56e10355f)), closes [#2591](https://github.com/synapsecns/sanguine/issues/2591) [#2593](https://github.com/synapsecns/sanguine/issues/2593)





## [0.5.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.5.0...@synapsecns/sdk-router@0.5.1) (2024-05-08)


### Bug Fixes

* **sdk-rouder:** remove cache hydration ([#2597](https://github.com/synapsecns/sanguine/issues/2597)) ([1c304ee](https://github.com/synapsecns/sanguine/commit/1c304ee60af8c12311fcbb48fa92da6903e1c653))





# [0.5.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.4.3...@synapsecns/sdk-router@0.5.0) (2024-05-03)


### Features

* **sdk-router:** Adds Scroll to SDK ([#2554](https://github.com/synapsecns/sanguine/issues/2554)) ([ed7f6de](https://github.com/synapsecns/sanguine/commit/ed7f6de23b538a5fb378e6167facfe30b352c703))





## [0.4.3](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.4.2...@synapsecns/sdk-router@0.4.3) (2024-04-03)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.4.2](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.4.1...@synapsecns/sdk-router@0.4.2) (2024-03-31)


### Bug Fixes

* **sdk-router:** Update default dest deadline ([#2407](https://github.com/synapsecns/sanguine/issues/2407)) ([8a9b628](https://github.com/synapsecns/sanguine/commit/8a9b62831beffe7e382332093be12d95b42f4bc2))





## [0.4.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.4.0...@synapsecns/sdk-router@0.4.1) (2024-03-29)

**Note:** Version bump only for package @synapsecns/sdk-router





# [0.4.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.29...@synapsecns/sdk-router@0.4.0) (2024-03-29)


### Features

* **sdk-router:** Add Base to the list of supported RFQ chains ([#2398](https://github.com/synapsecns/sanguine/issues/2398)) ([0425595](https://github.com/synapsecns/sanguine/commit/0425595c026bbc698a5f8a9ffd8f500fd2db1ae9))





## [0.3.29](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.28...@synapsecns/sdk-router@0.3.29) (2024-03-21)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.28](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.27...@synapsecns/sdk-router@0.3.28) (2024-03-04)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.27](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.26...@synapsecns/sdk-router@0.3.27) (2024-03-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.26](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.25...@synapsecns/sdk-router@0.3.26) (2024-03-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.25](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.24...@synapsecns/sdk-router@0.3.25) (2024-02-22)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.24](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.23...@synapsecns/sdk-router@0.3.24) (2024-02-13)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.23](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.22...@synapsecns/sdk-router@0.3.23) (2024-02-12)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.22](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.21...@synapsecns/sdk-router@0.3.22) (2024-02-06)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.21](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.20...@synapsecns/sdk-router@0.3.21) (2024-01-26)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.20](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.19...@synapsecns/sdk-router@0.3.20) (2024-01-19)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.19](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.18...@synapsecns/sdk-router@0.3.19) (2024-01-16)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.18](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.17...@synapsecns/sdk-router@0.3.18) (2024-01-12)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.17](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.16...@synapsecns/sdk-router@0.3.17) (2024-01-11)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.16](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.15...@synapsecns/sdk-router@0.3.16) (2024-01-11)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.15](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.14...@synapsecns/sdk-router@0.3.15) (2024-01-10)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.14](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.13...@synapsecns/sdk-router@0.3.14) (2024-01-10)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.13](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.12...@synapsecns/sdk-router@0.3.13) (2024-01-10)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.12](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.11...@synapsecns/sdk-router@0.3.12) (2024-01-10)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.11](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.10...@synapsecns/sdk-router@0.3.11) (2024-01-10)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.10](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.9...@synapsecns/sdk-router@0.3.10) (2024-01-09)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.9](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.8...@synapsecns/sdk-router@0.3.9) (2024-01-08)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.8](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.7...@synapsecns/sdk-router@0.3.8) (2023-12-27)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.7](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.6...@synapsecns/sdk-router@0.3.7) (2023-12-27)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.6](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.5...@synapsecns/sdk-router@0.3.6) (2023-12-22)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.5](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.4...@synapsecns/sdk-router@0.3.5) (2023-12-22)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.4](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.3...@synapsecns/sdk-router@0.3.4) (2023-12-22)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.3](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.2...@synapsecns/sdk-router@0.3.3) (2023-12-21)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.2](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.1...@synapsecns/sdk-router@0.3.2) (2023-12-21)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.3.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.3.0...@synapsecns/sdk-router@0.3.1) (2023-12-18)

**Note:** Version bump only for package @synapsecns/sdk-router





# [0.3.0](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.24...@synapsecns/sdk-router@0.3.0) (2023-12-12)


### Features

* Add Polygon support to New CCTP Router Contracts ([#1617](https://github.com/synapsecns/sanguine/issues/1617)) ([a891148](https://github.com/synapsecns/sanguine/commit/a89114852ef7fd2a0a3aaf3546ee5377f8fa9a93))





## [0.2.24](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.23...@synapsecns/sdk-router@0.2.24) (2023-12-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.23](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.22...@synapsecns/sdk-router@0.2.23) (2023-11-21)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.22](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.21...@synapsecns/sdk-router@0.2.22) (2023-11-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.21](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.20...@synapsecns/sdk-router@0.2.21) (2023-11-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.20](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.19...@synapsecns/sdk-router@0.2.20) (2023-10-26)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.19](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.18...@synapsecns/sdk-router@0.2.19) (2023-10-05)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.18](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.17...@synapsecns/sdk-router@0.2.18) (2023-09-12)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.17](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.16...@synapsecns/sdk-router@0.2.17) (2023-09-11)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.16](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.15...@synapsecns/sdk-router@0.2.16) (2023-09-07)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.15](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.14...@synapsecns/sdk-router@0.2.15) (2023-09-05)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.14](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.13...@synapsecns/sdk-router@0.2.14) (2023-08-28)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.13](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.12...@synapsecns/sdk-router@0.2.13) (2023-08-05)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.12](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.11...@synapsecns/sdk-router@0.2.12) (2023-08-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.11](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.10...@synapsecns/sdk-router@0.2.11) (2023-08-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.10](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.9...@synapsecns/sdk-router@0.2.10) (2023-08-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.9](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.8...@synapsecns/sdk-router@0.2.9) (2023-08-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.8](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.7...@synapsecns/sdk-router@0.2.8) (2023-08-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.7](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.6...@synapsecns/sdk-router@0.2.7) (2023-08-01)


### Reverts

* Revert "Revert "try to fix router install (#1201)"" (#1210) ([84d7cbb](https://github.com/synapsecns/sanguine/commit/84d7cbb0d8913c33ce91d8f380591fb401594c82)), closes [#1201](https://github.com/synapsecns/sanguine/issues/1201) [#1210](https://github.com/synapsecns/sanguine/issues/1210)





## [0.2.6](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.5...@synapsecns/sdk-router@0.2.6) (2023-08-01)


### Reverts

* Revert "try to fix router install (#1201)" ([67d6ef5](https://github.com/synapsecns/sanguine/commit/67d6ef5f1d7c1307b603ebe916cadb7eea7fe0bc)), closes [#1201](https://github.com/synapsecns/sanguine/issues/1201)





## [0.2.5](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.4...@synapsecns/sdk-router@0.2.5) (2023-08-01)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.4](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.3...@synapsecns/sdk-router@0.2.4) (2023-06-29)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.3](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.2...@synapsecns/sdk-router@0.2.3) (2023-06-28)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.2](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.2.1...@synapsecns/sdk-router@0.2.2) (2023-06-28)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.2.1](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.27...@synapsecns/sdk-router@0.2.1) (2023-06-27)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.27](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.26...@synapsecns/sdk-router@0.1.27) (2023-06-13)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.26](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.23...@synapsecns/sdk-router@0.1.26) (2023-06-13)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.23](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.22...@synapsecns/sdk-router@0.1.23) (2023-06-02)


### Reverts

* Revert "Publish" ([8202042](https://github.com/synapsecns/sanguine/commit/8202042d2485f1104f72910183bdb9d17cacb197))





## [0.1.22](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.21...@synapsecns/sdk-router@0.1.22) (2023-05-16)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.21](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.20...@synapsecns/sdk-router@0.1.21) (2023-05-03)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.20](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.19...@synapsecns/sdk-router@0.1.20) (2023-05-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.19](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.18...@synapsecns/sdk-router@0.1.19) (2023-05-02)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.18](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.17...@synapsecns/sdk-router@0.1.18) (2023-04-20)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.17](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.14...@synapsecns/sdk-router@0.1.17) (2023-04-18)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.14](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.13...@synapsecns/sdk-router@0.1.14) (2023-04-03)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.13](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.11...@synapsecns/sdk-router@0.1.13) (2023-03-31)

**Note:** Version bump only for package @synapsecns/sdk-router





## 0.1.11 (2023-03-30)


### Reverts

* Revert "feat/sdk-updates" ([92d842a](https://github.com/synapsecns/sanguine/commit/92d842a32a91ebbda7ad790ce25290e51d690056))





## [0.1.7](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.6...@synapsecns/sdk-router@0.1.7) (2023-02-15)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.6](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.5...@synapsecns/sdk-router@0.1.6) (2023-02-15)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.5](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.4...@synapsecns/sdk-router@0.1.5) (2023-02-15)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.4](https://github.com/synapsecns/sanguine/compare/@synapsecns/sdk-router@0.1.3...@synapsecns/sdk-router@0.1.4) (2023-02-15)

**Note:** Version bump only for package @synapsecns/sdk-router





## 0.1.3 (2023-02-15)

**Note:** Version bump only for package @synapsecns/sdk-router





## [0.1.2](https://github.com/synapsecns/sanguine/compare/sdk-router@0.1.1...sdk-router@0.1.2) (2023-02-15)

**Note:** Version bump only for package sdk-router





## 0.1.1 (2023-02-15)

**Note:** Version bump only for package sdk-router
