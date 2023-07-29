# DFK

This package contains a monitoring interface for the DFK community api as a secondary check to the explorer api in cases of lag. The structure of this folder is as follows:

```
├── <a href="./queries">queries</a>: graphql queries used to query the api
├── <a href="./.gqlgenc.yaml">.gqlgenc.yaml</a>: <a href="https://github.com/Yamashou/gqlgenc"> gqlgenc config </a>
```


The diagram below explains the generation workflow and where errors could occur:

```mermaid
graph LR;

subgraph Step1_Introspect_Graphql_API[Fetch Graphql Schema]
  A[Introspect Via API]
  A -->|Success| B[Generate Models from GraphQL]
  A -->|Error| C[Exit Status 1]
end

subgraph Step2_Generate_Models[Generate Models for Entities]
  B --> D[Generate Queries using Models]
end

subgraph Step3_Generate_Queries[Use graphql file to generate schema]
  D -->|Models| E[Generate queries from ./queries.graphql]
  E -->|Success| F[Exit Status 0 and generate files]
  E -->|Failure| G[Exit Status 1]
end

```
