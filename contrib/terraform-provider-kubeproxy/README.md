# Terraform Kubernetes IAP Proxy Provider

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/terraform-provider-kubeproxy)

This provider is a wrapper for the Kubernetes provider that allows for the use of an IAP (Identity-Aware Proxy) when interacting with GCP Kubernetes clusters. This is necessary because Terraform resources are short-lived, so spinning up the IAP proxy separately and having it provide access to the resources is not an option.

## Why use an IAP proxy?
IAP (Identity-Aware Proxy) is a feature of GCP that allows you to authenticate and authorize access to resources in a more fine-grained manner than just using a service account. By using IAP, you can ensure that only authorized users and applications can access your resources.

## How does the provider work?
The provider wraps the Kubernetes provider and adds some new fields to the schema, such as the project, zone, service_account, instance, and remote_port fields, which are necessary for configuring the IAP proxy.

When the provider is used to create or update resources, it first starts the IAP proxy on the specified instance, using the specified service account and project. It then sets the KUBECONFIG environment variable to use the proxy to access the Kubernetes cluster.

When the resources are destroyed, the provider stops the IAP proxy on the specified instance.

## How to use the provider
To use the provider, you will need to specify the project, zone, service_account, instance, and remote_port fields in your Terraform configuration. You will also need to provide credentials for the service account that will be used to start the IAP proxy.

After configuring the provider, you can use it in your Terraform resources just like you would use the Kubernetes provider. The provider will automatically handle starting and stopping the IAP proxy as needed.

## Conclusion
The Terraform Kubernetes IAP Proxy Provider allows you to use an IAP proxy when interacting with GCP Kubernetes clusters in Terraform, allowing for more fine-grained authentication and authorization of access to your resources. It is easy to use and seamlessly integrates with the Kubernetes provider, making it a great choice for securing your GCP Kubernetes clusters in Terraform.

It's good to note that this is a conceptual explanation and the implementation may differ and depend on the specific details of the kubernetes provider and how it interacts with the gcp iap.

## Incompatibilities

This provider does not support kubernetes_manifest resources. The kubectl provider should be used for this
