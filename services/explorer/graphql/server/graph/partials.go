package graph

// TODO make more dynamic.
const swapDeDup = `
swapDeDup AS (
SELECT
        if(amount_id >= 0 and tokens_sold = 0, toUInt256(amount_id), sold_id)     AS sold_id,
        if(amount_id >= 0 and tokens_sold = 0, toUInt256(amount[toUInt256(amount_id)]), tokens_sold)     AS tokens_sold,
        if(amount_id >= 0 and tokens_bought = 0, toUInt256(amount_id), bought_id) AS bought_id,
        if(amount_id >= 0 and tokens_bought = 0, toUInt256(amount[toUInt256(amount_id)]), tokens_bought)     AS tokens_bought,
       *
FROM (
         SELECT tokens_bought,
                tokens_sold,
                sold_id,
                bought_id,
                token_decimal,
                amount,
                amount_usd,
                contract_address                              AS swap_address,
                tx_hash                                       AS swap_tx_hash,
                chain_id                                      AS swap_chain_id,
                if(notEmpty(amount.keys), amount.keys[1], -1) AS amount_id

         FROM swap_events WHERE rowCount > 0 AND timestamp >= minTimestamp
         LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash)
)
`

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
f.pre_fblock_number AS fblock_number,
f.pre_fswap_success AS fswap_success,
f.pre_fswap_token_index AS fswap_token_index,
f.pre_fswap_min_amount AS fswap_min_amount,
f.pre_fswap_deadline AS fswap_deadline,
f.pre_ffee_amount_usd AS ffee_amount_usd,
f.pre_ftoken_decimal AS ftoken_decimal,
f.pre_ftimestamp AS ftimestamp,
f.pre_fdestination_chain_id AS fdestination_chain_id,
f.pre_finsert_time AS finsert_time,
 IF(
			 ti.token_address = '', be.token, ti.token_address
	 )                       AS ttoken,
 IF(
			 se.amount[se.bought_id] != '', toUInt256(se.amount[se.bought_id]),
			 be.amount
	 )                       AS tamount,
 IF(
			 se.amount[se.bought_id] != '', se.token_decimal[se.bought_id],
			 be.token_decimal
	 )                       AS ttoken_decimal,
 IF(
	 se.amount_usd[se.bought_id]  > 0,
	 se.amount_usd[se.bought_id],
	 be.amount_usd
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
be.block_number AS tblock_number,
be.fee_usd AS tfee_amount_usd,
be.timestamp AS ttimestamp,
be.destination_chain_id AS tdestination_chain_id,
be.insert_time AS tinsert_time
FROM
  (
    SELECT
IF(
		   ti.token_address = '', be.token, ti.token_address
   )                   AS pre_ftoken,
IF(
		   se.amount[se.sold_id] != '', toUInt256(se.amount[se.sold_id]),
		   be.amount
   )                   AS pre_famount,
IF(
		   se.amount[se.sold_id] != '', se.token_decimal[se.sold_id],
		   be.token_decimal
   )                   AS pre_ftoken_decimal,
IF(
			    se.amount_usd[se.sold_id] > 0,
			    se.amount_usd[se.sold_id],
			    be.amount_usd
   )                   AS pre_famount_usd,
      be.event_type AS pre_fevent_type,
      be.token AS pre_ftoken_raw,
      be.tx_hash AS pre_ftx_hash,
      be.chain_id AS pre_fchain_id,
      be.block_number AS pre_fblock_number,
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
      be.fee_usd AS pre_ffee_amount_usd,
      be.timestamp AS pre_ftimestamp,
      be.destination_chain_id AS pre_fdestination_chain_id,
      be.insert_time AS pre_finsert_time
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
t.pre_tblock_number AS tblock_number,
t.pre_tdeadline AS tdeadline,
t.pre_tswap_success AS tswap_success,
t.pre_tswap_token_index AS tswap_token_index,
t.pre_tswap_min_amount AS tswap_min_amount,
t.pre_tswap_deadline AS tswap_deadline,
t.pre_tfee_amount_usd AS tfee_amount_usd,
t.pre_ttoken_decimal AS ttoken_decimal,
t.pre_ttimestamp AS ttimestamp,
t.pre_tdestination_chain_id AS tdestination_chain_id,
t.pre_tinsert_time AS tinsert_time,
IF(
		 ti.token_address = '', be.token, ti.token_address
 )                       AS ftoken,
IF(
		 se.amount[se.sold_id] != '', toUInt256(se.amount[se.sold_id]),
		 be.amount
 )                       AS famount,
IF(
		 se.amount[se.sold_id] != '', se.token_decimal[se.sold_id],
		 be.token_decimal
 )                       AS ftoken_decimal,

IF(
			    se.amount_usd[se.sold_id] > 0,
			    se.amount_usd[se.sold_id],
			    be.amount_usd
   )                   AS famount_usd,
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
be.block_number AS fblock_number,
be.fee_usd AS ffee_amount_usd,
be.timestamp AS ftimestamp,
be.destination_chain_id AS fdestination_chain_id,
be.insert_time AS finsert_time
FROM
  (
    SELECT
IF(
		   ti.token_address = '', be.token, ti.token_address
   )                   AS pre_ttoken,
IF(
		   se.amount[se.bought_id] != '', toUInt256(se.amount[se.bought_id]),
		   be.amount
   )                   AS pre_tamount,
IF(
		   se.amount[se.bought_id] != '', se.token_decimal[se.bought_id],
		   be.token_decimal
   )                   AS pre_ttoken_decimal,
IF(
			    se.amount_usd[se.bought_id] > 0,
			    se.amount_usd[se.bought_id],
			    be.amount_usd
   )                   AS pre_tamount_usd,
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
      be.block_number AS pre_tblock_number,
      be.swap_success AS pre_tswap_success,
      be.swap_token_index AS pre_tswap_token_index,
      be.swap_min_amount AS pre_tswap_min_amount,
      be.swap_deadline AS pre_tswap_deadline,
      be.fee_usd AS pre_tfee_amount_usd,
      be.timestamp AS pre_ttimestamp,
      be.destination_chain_id AS pre_tdestination_chain_id,
      be.insert_time AS pre_tinsert_time
`

const singleSideCol = `
IF(
   ti.token_address = '', be.token, ti.token_address
)                   AS token,
IF(
   se.amount[se.sold_id] != '', toUInt256(se.amount[se.sold_id]),
   be.amount
)                   AS amount,
IF(
   se.amount[se.sold_id] != '', se.token_decimal[se.sold_id],
   be.token_decimal
)                   AS token_decimal,
IF(
	   se.amount_usd[se.sold_id] > 0,
	   se.amount_usd[se.sold_id],
	   be.amount_usd
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
be.fee_usd AS fee_usd,
be.timestamp AS timestamp,
be.destination_chain_id AS destination_chain_id,
be.insert_time AS insert_time
`

const singleSideJoinsCTE = `
 be
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
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

const baseSwap = `
SELECT * FROM swap_events LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash
`

const baseSwapWithTokenPt1 = `
SELECT
        if(amount_id >= 0 and tokens_sold = 0, toUInt256(amount_id), sold_id)     AS sold_id,
        if(amount_id >= 0 and tokens_sold = 0, toUInt256(amount[toUInt256(amount_id)]), tokens_sold)     AS tokens_sold,
        if(amount_id >= 0 and tokens_bought = 0, toUInt256(amount_id), bought_id) AS bought_id,
        if(amount_id >= 0 and tokens_bought = 0, toUInt256(amount[toUInt256(amount_id)]), tokens_bought)     AS tokens_bought,
       *
FROM (
         SELECT tokens_bought,
                tokens_sold,
                sold_id,
                event_type,
                chain_id,
                bought_id,
                token_decimal,
                amount,
                amount_usd,
                fee_usd,
                sender,
                tx_hash,
                contract_address                              AS swap_address,
                chain_id                                      AS swap_chain_id,
                if(notEmpty(amount.keys), amount.keys[1], -1) AS amount_id

         FROM swap_events
`
const baseSwapWithTokenPt2 = `
         LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash)
) se
         LEFT JOIN (
    SELECT DISTINCT ON(
                            chain_id, token_index, contract_address
                        ) token_address AS token, *
                    FROM
                        token_indices
    ) ti ON se.chain_id = ti.chain_id
    AND se.swap_address = ti.contract_address
    AND ti.token_index = se.bought_id
`

const baseMessageBus = `
SELECT * FROM message_bus_events LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash
`

const swapVolumeSelect = `
multiIf(event_type = 0, amount_usd[sold_id],event_type = 1, arraySum(mapValues(amount_usd)),event_type = 9, arraySum(mapValues(amount_usd)),event_type = 10,amount_usd[sold_id],0)
`
const orString = " OR "
const whereString = " WHERE ("

// TODO MAKE MORE DYNAMIC

const dailyVolumeBridgeMvPt1 = `
         SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
         FROM (SELECT date, maxMap(map(chain_id, total)) AS results
      FROM (SELECT coalesce(toString(b.date), toString(s.date))   AS date,
                   toInt64(coalesce(b.chain_id, s.chain_id, 0)) as chain_id,
                   (toFloat64(coalesce(b.usdTotal, 0)) + toFloat64(coalesce(s.usdTotal, 0)) )                      as total
                     FROM (
                                SELECT toDate(FROM_UNIXTIME(ftimestamp, '%Y/%m/%d')) as date,
                                     fchain_id AS chain_id,
                                     sumKahan(famount_usd)                         as usdTotal
                              FROM (SELECT *
                        FROM mv_bridge_events

`

const dailyVolumeBridgeMvPt2 = `
ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC
LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)
GROUP BY date, chain_id  order by date, chain_id) b
`

const dailyVolumeBridge = `
         SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
         FROM (SELECT date, maxMap(map(chain_id, total)) AS results
      FROM (SELECT coalesce(toString(b.date), toString(s.date))   AS date,
                   toInt64(coalesce(pre_fchain_id, s.chain_id, 0)) as chain_id,
                   (toFloat64(coalesce(b.usdTotal, 0)) + toFloat64(coalesce(s.usdTotal, 0)) )                      as total
                     FROM (
                              SELECT toDate(FROM_UNIXTIME(pre_ftimestamp, '%Y/%m/%d')) as date,
                                     pre_fchain_id,
                                     sumKahan(pre_famount_usd)                         as usdTotal


                              FROM (
                                       SELECT IF(
                                                          ti.token_address = '', be.token, ti.token_address
                                                  )                   AS pre_ftoken,
                                              IF(
                                                          se.amount[se.sold_id] != '', toUInt256(se.amount[se.sold_id]),
                                                          be.amount
                                                  )                   AS pre_famount,
                                              IF(
                                                          se.amount[se.sold_id] != '', se.token_decimal[se.sold_id],
                                                          be.token_decimal
                                                  )                   AS pre_ftoken_decimal,
                                              IF(
                                                          se.amount_usd[se.sold_id] > 0,
                                                          se.amount_usd[se.sold_id],
                                                          be.amount_usd
                                                  )                   AS pre_famount_usd,
                                              be.event_type           AS pre_fevent_type,
                                              be.token                AS pre_ftoken_raw,
                                              be.tx_hash              AS pre_ftx_hash,
                                              be.chain_id             AS pre_fchain_id,
                                              be.block_number         AS pre_fblock_number,
                                              be.contract_address     AS pre_fcontract_address,
                                              be.token_symbol         AS pre_ftoken_symbol,
                                              be.destination_kappa    AS pre_fdestination_kappa,
                                              be.sender               AS pre_fsender,
                                              be.recipient            AS pre_frecipient,
                                              be.recipient_bytes      AS pre_frecipient_bytes,
                                              toUInt256(be.fee)       AS pre_ffee,
                                              be.kappa                AS pre_fkappa,
                                              be.token_index_from     AS pre_ftoken_index_from,
                                              be.token_index_to       AS pre_ftoken_index_to,
                                              be.min_dy               AS pre_fmin_dy,
                                              be.deadline             AS pre_fdeadline,
                                              be.swap_success         AS pre_fswap_success,
                                              be.swap_token_index     AS pre_fswap_token_index,
                                              be.swap_min_amount      AS pre_fswap_min_amount,
                                              be.swap_deadline        AS pre_fswap_deadline,
                                              be.fee_usd              AS pre_ffee_amount_usd,
                                              be.timestamp            AS pre_ftimestamp,
                                              be.destination_chain_id AS pre_fdestination_chain_id,
                                              be.insert_time          AS pre_finsert_time
                                       FROM baseQuery be
                                                LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
                                           AND be.chain_id = se.swap_chain_id
                                                LEFT JOIN (
                                           SELECT DISTINCT ON(
                                                                   chain_id, token_index, contract_address
                                                               ) *
                                                           FROM
                                               token_indices
                                           ) ti ON be.chain_id = ti.chain_id
                                           AND se.swap_address = ti.contract_address
                                           AND ti.token_index = se.sold_id)
                              group by date, pre_fchain_id
                              order by date, pre_fchain_id) b
`
const toDateSelect = `toDate(FROM_UNIXTIME(timestamp, '%Y/%m/%d')) as date`
const toDateSelectMv = `toDate(FROM_UNIXTIME(ftimestamp, '%Y/%m/%d')) as date`

// TODO MAKE MORE DYNAMIC.
const dailyStatisticGenericSelect = `
SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
FROM (SELECT date, maxMap(map(chain_id, total)) AS results
      FROM (SELECT coalesce(toString(b.date), toString(s.date), toString(m.date)) AS date,
                   toInt64(coalesce(b.chain_id, s.chain_id, m.chain_id, 0))        as chain_id,
                   toFloat64(coalesce(b.sumTotal, 0)) +toFloat64(coalesce(s.sumTotal, 0))+toFloat64(coalesce(m.sumTotal, 0))                       as total`

const rankedChainsBridgeVolume = `
SELECT toInt64(coalesce(pre_fchain_id, s.chain_id, 0)) as chain_id,
       (toFloat64(coalesce(b.usdTotal, 0)) +
        toFloat64(coalesce(s.usdTotal, 0)))           as total
FROM (
         SELECT pre_fchain_id,
                sumKahan(pre_famount_usd) as usdTotal


         FROM (
                  SELECT IF(
                                     ti.token_address = '', be.token,
                                     ti.token_address
                             )                   AS pre_ftoken,
                         IF(
                                     se.amount[se.sold_id] != '',
                                     toUInt256(se.amount[se.sold_id]),
                                     be.amount
                             )                   AS pre_famount,
                         IF(
                                     se.amount[se.sold_id] != '',
                                     se.token_decimal[se.sold_id],
                                     be.token_decimal
                             )                   AS pre_ftoken_decimal,
                         IF(
                                     se.amount_usd[se.sold_id] > 0,
                                     se.amount_usd[se.sold_id],
                                     be.amount_usd
                             )                   AS pre_famount_usd,
                         be.event_type           AS pre_fevent_type,
                         be.token                AS pre_ftoken_raw,
                         be.tx_hash              AS pre_ftx_hash,
                         be.chain_id             AS pre_fchain_id,
                         be.block_number         AS pre_fblock_number,
                         be.contract_address     AS pre_fcontract_address,
                         be.token_symbol         AS pre_ftoken_symbol,
                         be.destination_kappa    AS pre_fdestination_kappa,
                         be.sender               AS pre_fsender,
                         be.recipient            AS pre_frecipient,
                         be.recipient_bytes      AS pre_frecipient_bytes,
                         toUInt256(be.fee)       AS pre_ffee,
                         be.kappa                AS pre_fkappa,
                         be.token_index_from     AS pre_ftoken_index_from,
                         be.token_index_to       AS pre_ftoken_index_to,
                         be.min_dy               AS pre_fmin_dy,
                         be.deadline             AS pre_fdeadline,
                         be.swap_success         AS pre_fswap_success,
                         be.swap_token_index     AS pre_fswap_token_index,
                         be.swap_min_amount      AS pre_fswap_min_amount,
                         be.swap_deadline        AS pre_fswap_deadline,
                         be.fee_usd              AS pre_ffee_amount_usd,
                         be.timestamp            AS pre_ftimestamp,
                         be.destination_chain_id AS pre_fdestination_chain_id,
                         be.insert_time          AS pre_finsert_time
                  FROM baseQuery be
                           LEFT JOIN swapDeDup se
                                     ON be.tx_hash = se.swap_tx_hash
                                         AND be.chain_id = se.swap_chain_id
                           LEFT JOIN (
                      SELECT DISTINCT ON(
                                              chain_id, token_index,
                                              contract_address
                                          ) *
                                      FROM
                          token_indices
                      ) ti ON be.chain_id = ti.chain_id
                      AND se.swap_address = ti.contract_address
                      AND ti.token_index = se.sold_id)
         group by pre_fchain_id
         order by pre_fchain_id) b
         `

const dailyStatisticGenericSinglePlatform = `
SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
FROM (
         SELECT date,
                maxMap(map(chain_id, sumTotal)) AS results
            FROM (SELECT toString(toDate(FROM_UNIXTIME(timestamp, '%Y/%m/%d')))                                                as date,
                         chain_id,
`
const dailyStatisticGenericSinglePlatformMv = `
SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
FROM (
         SELECT date,
                maxMap(map(chain_id, sumTotal)) AS results
            FROM (SELECT toString(toDate(FROM_UNIXTIME(ftimestamp, '%Y/%m/%d')))                                                as date,
                         fchain_id AS chain_id,
`

const dailyStatisticGenericSinglePlatformMvFee = `
SELECT date,
       results[1]                   AS ethereum,
       results[10]                  AS optimism,
       results[25]                  AS cronos,
       results[56]                  AS bsc,
       results[137]                 AS polygon,
       results[250]                 AS fantom,
       results[288]                 AS boba,
       results[1088]                AS metis,
       results[1284]                AS moonbeam,
       results[1285]                AS moonriver,
       results[8217]                AS klaytn,
       results[42161]               AS arbitrum,
       results[43114]               AS avalanche,
       results[53935]               AS dfk,
       results[1313161554]          AS aurora,
       results[1666600000]          AS harmony,
       results[7700]                AS canto,
       results[2000]                AS dogechain,
       results[8453]                AS base,
       arraySum(mapValues(results)) AS total
FROM (
         SELECT date,
                maxMap(map(chain_id, sumTotal)) AS results
            FROM (SELECT toString(toDate(FROM_UNIXTIME(ttimestamp, '%Y/%m/%d')))                                                as date,
                         tchain_id AS chain_id,
`

const dailyStatisticBridge = `
FROM (
SELECT IF(
se.amount_usd[se.sold_id] > 0,
se.amount_usd[se.sold_id],
be.amount_usd
)                   AS amount_usd,
be.tx_hash              AS tx_hash,
be.sender               AS sender,
be.fee_usd              AS fee_usd,
be.chain_id				AS chain_id,
be.timestamp			AS timestamp
FROM baseQuery be
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
AND be.chain_id = se.swap_chain_id
LEFT JOIN (
SELECT DISTINCT ON(
chain_id, token_index, contract_address
) *
FROM
token_indices
) ti ON be.chain_id = ti.chain_id
AND se.swap_address = ti.contract_address
AND ti.token_index = se.sold_id)
`

const destToOriginJoinsPt1 = `
be
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
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
  FROM
    bridge_events WHERE timestamp >= minTimestamp AND destination_chain_id > 0
`

const destToOriginJoinsPt2 = `
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
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
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

const originToDestJoinsPt1 = `
be
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
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
) AS f
LEFT JOIN (
  SELECT
    *
  from
    bridge_events WHERE timestamp >= minTimestamp  AND destination_chain_id = 0
`

const originToDestJoinsPt2 = `
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
LEFT JOIN swapDeDup se ON be.tx_hash = se.swap_tx_hash
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
