package db

// LEAF is a the prefix of the laf in the db.
const LEAF = "leaf_"

// MESSAGE is the prefix of the message in the db.
const MESSAGE = "message_"

// PROOF is the prefix of the proof message in the db.
const PROOF = "PROOF_"

// HEIGHT is the height of the index.
const HEIGHT = "HEIGHT_"

// LatestLeafIndex is the latest leaf index.
const LatestLeafIndex = "latest_known_leaf_index_"

// MessagesLastBlockEnd is the db key for the latest block for messages.
const MessagesLastBlockEnd = "messages_last_block"

// LatestRoot is the latest root.
const LatestRoot = "update_latest_root_"

// UpdaterProducedUpdate adds a produced update.
const UpdaterProducedUpdate = "updater_produced_update_"
