# Synapse Node Helm Chart

This helm chart is used to deploy a synapse node to a kubernetes cluster.

## Configuration

| Parameter                     | Description                                                                                                                                                                                                                                                                                                                | Default                           |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |-----------------------------------|
| `image.repository`                 | docker image to use files                                                                                                                             | `ghcr.io/synapsecns/synapse-node` |
| `image.tag` | tag of docker image to pull | `latest`                          |
| `image.pullPolicy` | When to pull images | `Always`                          |
| `nameOverride` | Override service name | `null`                            |
| `ingress` | Configurable [ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) to expose the  service. | `enabled: false`                  |
| `resources` | Allows you to set the [resources](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/) for the statefulset | `{}`                              |
| `nodeSelector` | Configurable [nodeSelector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector) so that you can target specific nodes for btcd cluster | `{}`                              |
| `affinity` | Value for the [node affinity settings](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature) | `{}`                              |
| `secretName` | Value for the secret in the same namespace that contains docker credentials | `{}`                              |
