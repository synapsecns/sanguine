# Example Project

   1. Create the abi-bound contract

       The example project contains a simple example using ethergo. This provides a much simpler example then synapse code by testing a simple counter. How is this done?

       First, we create a `counter` folder. This will store our contract. In a more complicated project, we might pull our abigen from a hardhat or foundry repository or even etherscan with the [abigen tool](../../tools/abigen/readme.md). In this instance, we just copy and paste `counter.sol` into `example/counter`. Next, we need to create a [`generate.go`](https://go.dev/blog/generate) file to generate the abi for us. This way, when we run `go build ./...` all `go:build` commands will be run.

       Our [`generate.go`](counter/generate.go) looks like this:

        ```go
        package counter

        //go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ./counter.sol --pkg counter --sol-version 0.8.4 --filename counter

        ```

       As you can see, we've added `0.8.4` as our solidity version, but any compatible version with the [`pragma string`](https://docs.soliditylang.org/en/develop/layout-of-source-files.html) will work. `--pkg` *must* be the same name as the folder.

      1. Now, we need to create a contract type file. Any type that satisfies the `ContractType` interface will work, but we'll add a few tricks here to make these easier to work with.
         1. Create a `contractType` type. This shouldn't be exported:
            ```go
               // contractTypeImpl is the type of the contract being saved/fetched.
              // we use an interface here so the deploy helper here can be abstracted away from the synapse contracts
              //go:generate go run golang.org/x/tools/cmd/stringer -type=contractTypeImpl -linecomment
              type contractTypeImpl int
            ```
            As a result of our `//go:generate`, any new type we define will have a string generated based on a comment we put after it
         2. Define a contract type. You should run go:generate after this step. `contractTypeImpl` should've gained a `String()` method that will return `CounterType`, this will be assumed to be the contract name:
            ```go
              // CounterType is the type of the counter contract
              CounterType contractTypeImpl = 0 // CounterType
            ```
         3. (Optional) add sanity checks. Go provides a variety of compile/runtime sanity checks through both `init()` and interface assertions `_ X = y`. We want to assert that `go build ./...` has been updated. We also want to add an `AllContractTypes` method. These checks make it much harder to forget to rerun `go:generate`/have tests fail because of a config issue
            ```go
            // AllContractTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
            // of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
            // any method called where the index is higher than the stringer array length.
            // TODO: find a compile time way to do this.
            var AllContractTypes []contractTypeImpl

            // set all contact types.
            func init() {
                 for i := 0; i < len(_contractTypeImpl_index); i++ {
                      contractType := contractTypeImpl(i)
                      AllContractTypes = append(AllContractTypes, contractType)
                      // assert type is correct
                      var _ ContractType = contractType
             }

            // verifyStringerUpdated verifies stringer is up to date (this index is included in stringer).
            func verifyStringerUpdated(contractType contractTypeImpl) {
                if int(contractType) > len(_contractTypeImpl_index) {
                    panic("please update stringer before running test again")
                }
            }
            ```
         4. Define methods. Above, we said that `contractTypeImpl` had to satisfy `ContractType` which has three methods. Let's take a look at that interface and see what we need to do here:
           ```go
           // ContractType is a contract type interface that contracts need to comply with.
           type ContractType interface {
               // ID gets the unique identifier for the contracts
               ID() int
               // Name gets a the contracts name
               Name() string
               // ContractInfo gets the contract info from the compiler contract.
               ContractInfo() *compiler.Contract
               // ContractName gets the name fo the deployed contract
               ContractName() string
           }
           ```
           First up is `ID()`. This is a unique identifier represented as an int. If you remember, `contractTypeImpl` is an int! If you followed step 3, we added safety checks to make sure every number had a unique type so we can just cast the type to an int:
           ```go
           // ID get sthe contract type as an id.
           func (c contractTypeImpl) ID() int {
              verifyStringerUpdated(c) // ignore this if you skipped step 3
              return int(c)
           }
           ```
         Next up is Name. We already set this using stringer so we can just call the stringer method.
         ```go
         // Name gets the name of the contract.
         func (c contractTypeImpl) Name() string {
           verifyStringerUpdated(c) // ignore this if you skipped step 3
           return c.String()
         }
         Next, we have to define a contract name. In our case, this is the same as name. But these may vary.
         ```go
         func (c contractTypeImpl) ContractName() string {
             verifyStringerUpdated(c) // ignore this if you skipped step 3
             return c.String()
         }
         ```
         Finally, we define the `ContractInfo`. This is used when we upload data to tenderly. If you try this make sure that `counter.Contracts` is correctly imported to the path do your contract. If you need help figuring out the path to your contract, check out `x.contractinfo.json` where `x` is filename from above
         ```go
         // ContractInfo gets the source code of every contract.
         func (c contractTypeImpl) ContractInfo() *compiler.Contract {
             verifyStringerUpdated(c)
             switch c {
               case CounterType:
                   return counter.Contracts["/solidity/counter.sol:Counter"]
               default:
                   panic("not yet implemented")
              }
            }
         ```
   2. Add a deployer- deployers are responsible for telling the contract registry what their dependencies are and specifying deploy instructions. These are the equivelant of individual migrations in hardhat. Deployers can get quite complicated when you have to deal with complex initializations, but for now we'll juse use a simple delpoyer in [deployer.go](deployer.go)
        ```go
      package example

      import (
        "context"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/common"
        "github.com/ethereum/go-ethereum/core/types"
        "github.com/synapsecns/sanguine/ethergo/deployer"
        "github.com/synapsecns/sanguine/ethergo/example/counter"
        "github.com/synapsecns/sanguine/ethergo/backends"
      )

      // CounterDeployer deploys a counter.
      type CounterDeployer struct {
        *deployer.BaseDeployer
      }

      // NewCounterDeployer creates a deployer for the new counter.
      func NewCounterDeployer(registry deployer.GetOnlyContractRegistry, backend backends.SimulatedTestBackend) deployer.ContractDeployer {
        return &CounterDeployer{
          deployer.NewSimpleDeployer(registry, backend, CounterType),
        }
      }

      // Deploy deploys the contract.
      func (n *CounterDeployer) Deploy(ctx context.Context) (backends.DeployedContract, error) {
        //nolint: wrapcheck
        return n.DeploySimpleContract(ctx, func(transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
          return counter.DeployCounter(transactOps, backend)
        }, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
          // this is kept separate because we often want to add an address handle to this so it's compatible with vm.ContractRef
          return counter.NewCounter(address, backend)
        })
      }

      // compile time assertion.
      var _ deployer.ContractDeployer = &CounterDeployer{}

   3. Add a Test - at this point, we've constructed a rudimentary contract deployment setup that we're ready to try out. Let's write a test to make sure we can increment the country

      ```go
      package example_test

      import (
        "context"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        . "github.com/stretchr/testify/assert"
        "github.com/synapsecns/sanguine/ethergo/example"
        "github.com/synapsecns/sanguine/ethergo/example/counter"
        "github.com/synapsecns/sanguine/ethergo/manager"
        "github.com/synapsecns/sanguine/ethergo/backends/simulated"
        "testing"
        "time"
      )

      func TestCounter(t *testing.T) {
        // register a test timeout
        testContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
        defer cancel()

        // since extra deployers don't neccesarily deploy anything (only when requested in the GetOnlyContractRegistry)
        // adding them here won't slow anyhting down. It's reccomended you have a global slice of these deployers you register every time.
        deployer := manager.NewDeployerManager(t, example.NewCounterDeployer)

        newTestBackend := simulated.NewSimulatedBackend(testContext, t)

        deployedContract := deployer.Get(testContext, newTestBackend, example.CounterType)
        // if you're using these often, it's recommended you extend manager and add type casted getters here, along with the global registry
        counterHandle := deployedContract.ContractHandle().(*counter.Counter)

        // first up, let's make sure we're at 0
        count, err := counterHandle.GetCount(&bind.CallOpts{Context: testContext})
        Nil(t, err)
        True(t, count.Int64() == 0)

        // let's increment the counter
        authOpts := newTestBackend.GetTxContext(testContext, nil)
        tx, err := counterHandle.IncrementCounter(authOpts.TransactOpts)
        Nil(t, err)

        newTestBackend.WaitForConfirmation(testContext, tx)

        // we should be at 1
        count, err = counterHandle.GetCount(&bind.CallOpts{Context: testContext})
        Nil(t, err)
        True(t, count.Int64() == 1)
      }
      ```


That's it! You should be done. As you can see, there's a lot more that can be done here. Passing in a list of all your deployers every time doesn't make sense. You'll want to create a standard testutil and extend it. We also haven't covered that any backend here is interchangable: you can use simulated, ganache, or embedded geth. This tutorial should've covered the basics though
