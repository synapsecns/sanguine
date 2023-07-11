# Terraform IAP Proxy Provider

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-helmproxy.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/terraform-provider-helmproxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/terraform-provider-helmproxy)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/terraform-provider-helmproxy)

This provider is a wrapper for the Helm provider that allows for the use of an IAP (Identity-Aware Proxy) when interacting with GCP resources. This is necessary because Terraform resources are short-lived, so spinning up the IAP proxy separately and having it provide access to the resources is not an option.

## Why use an IAP proxy?
IAP (Identity-Aware Proxy) is a feature of GCP that allows you to authenticate and authorize access to resources in a more fine-grained manner than just using a service account. By using IAP, you can ensure that only authorized users and applications can access your resources.

## How does the provider work?
The provider wraps the Helm provider and combines the schemas of the two providers. It also adds some new fields to the schema, such as the project, zone, service_account, instance, and remote_port fields, which are necessary for configuring the IAP proxy.

When the provider is used to create or update resources, it first starts the IAP proxy on the specified instance, using the specified service account and project. It then passes the requests for resources through the proxy, allowing for authenticated and authorized access.

When the resources are destroyed, the provider stops the IAP proxy on the specified instance.

## How to use the provider

To use the provider, you will need to specify the project, zone, service_account, instance, and remote_port fields in your Terraform configuration. You will also need to provide credentials for the service account that will be used to start the IAP proxy. Please see the example folder for an example.


