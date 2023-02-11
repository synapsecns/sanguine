## Ethereum Transaction Signer

This Go library provides support for signing Ethereum transactions using three different signers: AwsSigner, GcpSigner, and LocalSigner.

## Local Signer

The LocalSigner is an option for signing transactions on a single machine. It requires the user to provide their private key, which is stored unencrypted in memory and used to sign the transactions. While this approach is straightforward to set up and use, it is not suitable for production environments.

In a multi-user environment, it is important to secure private keys. The use of an Hardware Security Module (HSM) is a good idea for storing the private keys of the secp256k1 curve, which is used in the Ethereum blockchain. An HSM provides an extra layer of security by storing the private keys in a physically secure device and limiting access to the keys through strict authentication methods.


## AWS Signer

The AwsSigner leverages Amazon Web Services' (AWS) authentication methods to sign transactions. This signer is ideal for use in an AWS cloud environment and allows for efficient scaling and management of the signing process. However, it is important for the user to check that their AWS Identity and Access Management (IAM) permissions are secure before using this signer.


## GCP Signer

The GcpSigner leverages Google Cloud Platform's (GCP) authentication methods to sign transactions. This signer is ideal for use in a GCP cloud environment and allows for efficient scaling and management of the signing process. However, it is important for the user to check that their Google Cloud Identity and Access Management (IAM) permissions are secure before using this signer.


### Authorization & Authentication

The recommended approach for signing transactions in a production kubernetes cluster is through workload identity federation, which offers a completely keyless solution. Unlike JSON-based service accounts or Hashicorp Vault, there is no need to store or manage private keys with this method.

Workload identity federation leverages the built-in security features of the kubernetes cluster, such as OpenID Connect (OIDC) or Identity and Access Management (IAM), to authenticate the signer without the use of private keys. This eliminates the risk of key theft or unauthorized access, and makes the signing process more secure.

Additionally, workload identity federation is fully integrated with kubernetes, making it easier to manage and maintain in a production environment. It also eliminates the need for additional infrastructure, as opposed to Hashicorp Vault, which requires setup and maintenance.


