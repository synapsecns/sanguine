package hardhat

import (
	"time"
)

// Contract is a contract representation in the hardhat parser.
type Contract struct {
	Name              string                     `json:"contractName"`
	Abi               interface{}                `json:"abi"`
	Bytecode          string                     `json:"bytecode"`
	DeployedBytecode  string                     `json:"deployedBytecode"`
	SourceMap         string                     `json:"sourceMap"`
	DeployedSourceMap string                     `json:"deployedSourceMap"`
	Source            string                     `json:"source"`
	SourcePath        string                     `json:"sourcePath"`
	Ast               ContractAst                `json:"legacyAST"`
	Compiler          ContractCompiler           `json:"compiler"`
	Networks          map[string]ContractNetwork `json:"networks"`

	SchemaVersion string    `json:"schemaVersion"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// Node is a node in the ast.
type Node struct {
	NodeType     string `json:"nodeType"`
	AbsolutePath string `json:"absolutePath"`
	File         string `json:"file"`
}

// ContractAst is the ast of a contract.
type ContractAst struct {
	AbsolutePath    string           `json:"absolutePath"`
	ExportedSymbols map[string][]int `json:"exportedSymbols"`
	ID              int              `json:"id"`
	NodeType        string           `json:"nodeType"`
	Nodes           []Node           `json:"nodes"`
	Src             string           `json:"src"`
}

// ContractCompiler is the compiler metadata used for contracts.
type ContractCompiler struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ContractNetwork is the network metadata used for contracts.
type ContractNetwork struct {
	Events          interface{} `json:"events"`
	Links           interface{} `json:"links"`
	Address         string      `json:"address"`
	TransactionHash string      `json:"transactionHash"`
}

// HardhatContract is a contract representation in the hardhat parser.
// nolint: golint, revive
type HardhatContract struct {
	Contract
	Address  string         `json:"address"`
	Receipt  hardhatReceipt `json:"receipt"`
	Metadata string         `json:"metadata"`
}

type hardhatReceipt struct {
	TransactionHash string `json:"transactionHash"`
}

// ContractSources is the source of a contract.
type ContractSources struct {
	Content string `json:"content"`
}
