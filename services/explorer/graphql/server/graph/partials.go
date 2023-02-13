package graph

// TODO make more dynamic.
const originToDestCol = `
f.pre_ftoken AS ftoken,
f.pre_famount AS famount,
f.pre_famount_usd AS famount_usd,
f.pre_fevent_type AS fevent_type,
f.pre_ftoken_raw AS ftoken_raw,
f.pre_ftx_hash AS ftx_hash,
f.pre_fchain_id AS fchain_id,
f.pre_fcontract_address AS fcontract_address,
f.pre_ftoken_symbol AS ftoken_symbol,
f.pre_fdestination_kappa AS fdestination_kappa,
f.pre_fsender AS fsender,
f.pre_frecipient AS frecipient,
f.pre_frecipient_bytes AS frecipient_bytes,
f.pre_ffee AS ffee,
f.pre_fkappa AS fkappa,
f.pre_ftoken_index_from AS ftoken_index_from,
f.pre_ftoken_index_to AS ftoken_index_to,
f.pre_fmin_dy AS fmin_dy,
f.pre_fdeadline AS fdeadline,
f.pre_fswap_success AS fswap_success,
f.pre_fswap_token_index AS fswap_token_index,
f.pre_fswap_min_amount AS fswap_min_amount,
f.pre_fswap_deadline AS fswap_deadline,
f.pre_ftoken_id AS ftoken_id,
f.pre_ffee_amount_usd AS ffee_amount_usd,
f.pre_ftoken_decimal AS ftoken_decimal,
f.pre_ftimestamp AS ftimestamp,
f.pre_fdestination_chain_id AS fdestination_chain_id,
f.pre_finsert_time AS finsert_time,
(
  IF(
    ti.token_address = '', be.token, ti.token_address
  )
) AS ttoken,
toUInt256(
  IF(
    se.tokens_sold > 0, se.tokens_sold,
    be.amount
  )
) AS tamount,
(
  IF(
    se.swap_amount_usd[ti.token_index] > 0,
    (
      (
        toFloat64(
          (
            IF(
              se.tokens_sold > 0, se.tokens_sold,
              be.amount
            )
          )
        )/ exp10(be.token_decimal)
      ) * se.swap_amount_usd[ti.token_index]
    ),
    be.amount_usd
  )
) AS tamount_usd,
be.event_type AS tevent_type,
be.token AS ttoken_raw,
be.tx_hash AS ttx_hash,
be.chain_id AS tchain_id,
be.contract_address AS tcontract_address,
be.token_symbol AS ttoken_symbol,
be.destination_kappa AS tdestination_kappa,
be.sender AS tsender,
be.recipient AS trecipient,
be.recipient_bytes AS trecipient_bytes,
toUInt256(be.fee) AS tfee,
be.kappa AS tkappa,
be.token_index_from AS ttoken_index_from,
be.token_index_to AS ttoken_index_to,
be.min_dy AS tmin_dy,
be.deadline AS tdeadline,
be.swap_success AS tswap_success,
be.swap_token_index AS tswap_token_index,
be.swap_min_amount AS tswap_min_amount,
be.swap_deadline AS tswap_deadline,
be.token_id AS ttoken_id,
be.fee_amount_usd AS tfee_amount_usd,
be.token_decimal AS ttoken_decimal,
be.timestamp AS ttimestamp,
be.destination_chain_id AS tdestination_chain_id,
be.insert_time AS tinsert_time
FROM
  (
    SELECT
      (
        IF(
          ti.token_address = '', be.token, ti.token_address
        )
      ) AS pre_ftoken,
      toUInt256(
        IF(
          se.tokens_bought > 0, se.tokens_bought,
          be.amount
        )
      ) AS pre_famount,
      (
        IF(
          se.swap_amount_usd[ti.token_index] > 0,
          (
            (
              toFloat64(
                (
                  IF(
                    se.tokens_bought > 0, se.tokens_bought,
                    be.amount
                  )
                )
              )/ exp10(be.token_decimal)
            ) * se.swap_amount_usd[ti.token_index]
          ),
          be.amount_usd
        )
      ) AS pre_famount_usd,
      be.event_type AS pre_fevent_type,
      be.token AS pre_ftoken_raw,
      be.tx_hash AS pre_ftx_hash,
      be.chain_id AS pre_fchain_id,
      be.contract_address AS pre_fcontract_address,
      be.token_symbol AS pre_ftoken_symbol,
      be.destination_kappa AS pre_fdestination_kappa,
      be.sender AS pre_fsender,
      be.recipient AS pre_frecipient,
      be.recipient_bytes AS pre_frecipient_bytes,
      toUInt256(be.fee) AS pre_ffee,
      be.kappa AS pre_fkappa,
      be.token_index_from AS pre_ftoken_index_from,
      be.token_index_to AS pre_ftoken_index_to,
      be.min_dy AS pre_fmin_dy,
      be.deadline AS pre_fdeadline,
      be.swap_success AS pre_fswap_success,
      be.swap_token_index AS pre_fswap_token_index,
      be.swap_min_amount AS pre_fswap_min_amount,
      be.swap_deadline AS pre_fswap_deadline,
      be.token_id AS pre_ftoken_id,
      be.fee_amount_usd AS pre_ffee_amount_usd,
      be.token_decimal AS pre_ftoken_decimal,
      be.timestamp AS pre_ftimestamp,
      be.destination_chain_id AS pre_fdestination_chain_id,
      be.insert_time AS pre_finsert_time
`
const originToDestJoins = `
be
LEFT JOIN (
  SELECT
    amount_usd AS swap_amount_usd,
    tokens_bought,
    tokens_sold,
    sold_id,
    bought_id,
    contract_address AS swap_address,
    tx_hash AS swap_tx_hash,
    chain_id AS swap_chain_id
  FROM
    swap_events
) se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
  SELECT
    DISTINCT ON (
      chain_id, token_index, contract_address
    ) *
  FROM
    token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = be.sold_id
) AS f
LEFT JOIN (
  SELECT
    *
  from
    bridge_events
  ORDER BY
    block_number DESC,
    event_index DESC,
    insert_time DESC
  LIMIT
    1 BY chain_id,
    contract_address,
    event_type,
    block_number,
    event_index,
    tx_hash
) be ON fdestination_chain_id = be.chain_id
AND fdestination_kappa = be.kappa
LEFT JOIN (
  SELECT
    amount_usd AS swap_amount_usd,
    tokens_bought,
    tokens_sold,
    sold_id,
    bought_id,
    contract_address AS swap_address,
    tx_hash AS swap_tx_hash,
    chain_id AS swap_chain_id
  FROM
    swap_events
) se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
  SELECT
    DISTINCT ON (
      chain_id, token_index, contract_address
    ) *
  FROM
    token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = se.bought_id
`

const destToOriginCol = `
t.pre_ttoken AS ttoken,
t.pre_tamount AS tamount,
t.pre_tamount_usd AS tamount_usd,
t.pre_tevent_type AS tevent_type,
t.pre_ttoken_raw AS ttoken_raw,
t.pre_ttx_hash AS ttx_hash,
t.pre_tchain_id AS tchain_id,
t.pre_tcontract_address AS tcontract_address,
t.pre_ttoken_symbol AS ttoken_symbol,
t.pre_tdestination_kappa AS tdestination_kappa,
t.pre_tsender AS tsender,
t.pre_trecipient AS trecipient,
t.pre_trecipient_bytes AS trecipient_bytes,
t.pre_tfee AS tfee,
t.pre_tkappa AS tkappa,
t.pre_ttoken_index_from AS ttoken_index_from,
t.pre_ttoken_index_to AS ttoken_index_to,
t.pre_tmin_dy AS tmin_dy,
t.pre_tdeadline AS tdeadline,
t.pre_tswap_success AS tswap_success,
t.pre_tswap_token_index AS tswap_token_index,
t.pre_tswap_min_amount AS tswap_min_amount,
t.pre_tswap_deadline AS tswap_deadline,
t.pre_ttoken_id AS ttoken_id,
t.pre_tfee_amount_usd AS tfee_amount_usd,
t.pre_ttoken_decimal AS ttoken_decimal,
t.pre_ttimestamp AS ttimestamp,
t.pre_tdestination_chain_id AS tdestination_chain_id,
t.pre_tinsert_time AS tinsert_time,
(
  if(
    ti.token_address = '', be.token, ti.token_address
  )
) AS ftoken,
toUInt256(
  if(
    se.tokens_sold > 0, se.tokens_sold,
    be.amount
  )
) AS famount,
(
  if(
    se.swap_amount_usd[ti.token_index] > 0,
    (
      (
        toFloat64(
          (
            if(
              se.tokens_sold > 0, se.tokens_sold,
              be.amount
            )
          )
        )/ exp10(be.token_decimal)
      ) * se.swap_amount_usd[ti.token_index]
    ),
    be.amount_usd
  )
) AS famount_usd,
be.event_type AS fevent_type,
be.token AS ftoken_raw,
be.tx_hash AS ftx_hash,
be.chain_id AS fchain_id,
be.contract_address AS fcontract_address,
be.token_symbol AS ftoken_symbol,
be.destination_kappa AS fdestination_kappa,
be.sender AS fsender,
be.recipient AS frecipient,
be.recipient_bytes AS frecipient_bytes,
toUInt256(be.fee) AS ffee,
be.kappa AS fkappa,
be.token_index_from AS ftoken_index_from,
be.token_index_to AS ftoken_index_to,
be.min_dy AS fmin_dy,
be.deadline AS fdeadline,
be.swap_success AS fswap_success,
be.swap_token_index AS fswap_token_index,
be.swap_min_amount AS fswap_min_amount,
be.swap_deadline AS fswap_deadline,
be.token_id AS ftoken_id,
be.fee_amount_usd AS ffee_amount_usd,
be.token_decimal AS ftoken_decimal,
be.timestamp AS ftimestamp,
be.destination_chain_id AS fdestination_chain_id,
be.insert_time AS finsert_time
FROM
  (
    SELECT
      (
        if(
          ti.token_address = '', be.token, ti.token_address
        )
      ) AS pre_ttoken,
      toUInt256(
        if(
          se.tokens_bought > 0, se.tokens_bought,
          be.amount
        )
      ) AS pre_tamount,
      (
        if(
          se.swap_amount_usd[ti.token_index] > 0,
          (
            (
              toFloat64(
                (
                  if(
                    se.tokens_bought > 0, se.tokens_bought,
                    be.amount
                  )
                )
              )/ exp10(be.token_decimal)
            ) * se.swap_amount_usd[ti.token_index]
          ),
          be.amount_usd
        )
      ) AS pre_tamount_usd,
      be.event_type AS pre_tevent_type,
      be.token AS pre_ttoken_raw,
      be.tx_hash AS pre_ttx_hash,
      be.chain_id AS pre_tchain_id,
      be.contract_address AS pre_tcontract_address,
      be.token_symbol AS pre_ttoken_symbol,
      be.destination_kappa AS pre_tdestination_kappa,
      be.sender AS pre_tsender,
      be.recipient AS pre_trecipient,
      be.recipient_bytes AS pre_trecipient_bytes,
      toUInt256(be.fee) AS pre_tfee,
      be.kappa AS pre_tkappa,
      be.token_index_from AS pre_ttoken_index_from,
      be.token_index_to AS pre_ttoken_index_to,
      be.min_dy AS pre_tmin_dy,
      be.deadline AS pre_tdeadline,
      be.swap_success AS pre_tswap_success,
      be.swap_token_index AS pre_tswap_token_index,
      be.swap_min_amount AS pre_tswap_min_amount,
      be.swap_deadline AS pre_tswap_deadline,
      be.token_id AS pre_ttoken_id,
      be.fee_amount_usd AS pre_tfee_amount_usd,
      be.token_decimal AS pre_ttoken_decimal,
      be.timestamp AS pre_ttimestamp,
      be.destination_chain_id AS pre_tdestination_chain_id,
      be.insert_time AS pre_tinsert_time
`
const destToOriginJoins = `
be
LEFT JOIN (
  SELECT
    amount_usd AS swap_amount_usd,
    tokens_bought,
    tokens_sold,
    sold_id,
    bought_id,
    contract_address AS swap_address,
    tx_hash AS swap_tx_hash,
    chain_id AS swap_chain_id
  FROM
    swap_events
) se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
  SELECT
    DISTINCT ON (
      chain_id, token_index, contract_address
    ) *
  FROM
    token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = be.bought_id
) AS t
LEFT JOIN (
  SELECT
    *
  from
    bridge_events
  ORDER BY
    block_number DESC,
    event_index DESC,
    insert_time DESC
  LIMIT
    1 BY chain_id,
    contract_address,
    event_type,
    block_number,
    event_index,
    tx_hash
) be ON pre_tchain_id = be.destination_chain_id
AND pre_tkappa = be.destination_kappa
LEFT JOIN (
  SELECT
    amount_usd AS swap_amount_usd,
    tokens_bought,
    tokens_sold,
    sold_id,
    bought_id,
    contract_address AS swap_address,
    tx_hash AS swap_tx_hash,
    chain_id AS swap_chain_id
  FROM
    swap_events
) se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
  SELECT
    DISTINCT ON (
      chain_id, token_index, contract_address
    ) *
  FROM
    token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = se.sold_id
`

const singleSideJoins = `
 be
LEFT JOIN (
  SELECT
    amount_usd AS swap_amount_usd,
    tokens_bought,
    tokens_sold,
    sold_id,
    bought_id,
    contract_address AS swap_address,
    tx_hash AS swap_tx_hash,
    chain_id AS swap_chain_id
  FROM
    swap_events
) se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
  SELECT
    DISTINCT ON (
      chain_id, token_index, contract_address
    ) *
  FROM
    token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = be.sold_id
`
const singleSideCol = `
(
  if(
    ti.token_address = '', be.token, ti.token_address
  )
) AS token,
(
  if(
    se.tokens_bought > 0, se.tokens_bought,
    be.amount
  )
) AS amount,
(
  if(
    se.swap_amount_usd[ti.token_index] > 0,
    (
      (
        toFloat64(amount)/ exp10(be.token_decimal)
      ) * se.swap_amount_usd[ti.token_index]
    ),
    be.amount_usd
  )
) AS amount_usd,
be.event_type AS event_type,
be.token AS token_raw,
be.tx_hash AS tx_hash,
be.chain_id AS chain_id,
be.contract_address AS contract_address,
be.token_symbol AS token_symbol,
be.destination_kappa AS destination_kappa,
be.sender AS sender,
be.recipient AS recipient,
be.recipient_bytes AS recipient_bytes,
be.fee AS fee,
be.kappa AS kappa,
be.token_index_from AS token_index_from,
be.token_index_to AS token_index_to,
be.min_dy AS min_dy,
be.deadline AS deadline,
be.swap_success AS swap_success,
be.swap_token_index AS swap_token_index,
be.swap_min_amount AS swap_min_amount,
be.swap_deadline AS swap_deadline,
be.token_id AS token_id,
be.fee_amount_usd AS fee_amount_usd,
be.token_decimal AS token_decimal,
be.timestamp AS timestamp,
be.destination_chain_id AS destination_chain_id,
be.insert_time AS insert_time
`
