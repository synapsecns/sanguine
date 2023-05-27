---
sidebar_position: 6
---

# Estimating Total Gas
Each chain in the Synapse network has a special Gas Oracle contract that is used to determine the  current price of Gas on that chain.
<br/>
What this means is that part of the Origin "state" will include the current cost of gas on that chain.
<br/>
This means that when a remote chain hears about the state of another chain, it will get a rough idea of the goind rate of gas.
<br/>
Thus, when a message is sent from chain A to chain B, chain A will estimate the cost of gas on SYN chain and also the cost of gas on chain B, and it will collect a sufficient amount to cover the gas required to execute the message.
<br/>
TODO: More work on this later
