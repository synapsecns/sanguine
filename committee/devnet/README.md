# Verifier Devnet

This module provides everything that's needed to spin up a Verifier development network in order to test functionality of Verifier nodes and Module verification requests.

# Provisioner

Package `provisioner` sets up the development network with the addresses of your Verifier nodes and changes the threshold for an accepted verification in the network.

# Sender

Package `sender` sends verification requests to the `SynapseModule`, which will then be picked up by Verifier nodes in the network for verification in the network.
