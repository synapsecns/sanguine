
-- run this on ponder index to get a summary of a test run. use WHERE params to point it at a particular test period.

select to_timestamp(min(depositTs))                                          period_start,
	   to_timestamp(max(depositTs))                                          period_end,
	   max(depositTs) - min(depositTs)                                       period_seconds,
	   round(count(1) * 1.00 / (max(depositTs) - min(depositTs)), 2)         deposits_per_second,
	   round(count(1) * 1.00 / (max(depositTs) - min(depositTs)) * 86400, 0) deposits_per_day,
	   count(1)                                                              count_deposits,
	   count(case when proofSeconds is not null then 1 else null end)        count_proofs,
	   count(case when claimSeconds is not null then 1 else null end)        count_claims,
	   count(case when disputeId is not null then 1 else null end)           count_disputes,
	   round(avg(relaySeconds), 2)                                           relaySeconds_AVG,
	   max(relaySeconds)                                                     relaySeconds_MAX,
	   round(avg(proofSeconds), 2)                                           proofSeconds_AVG,
	   max(proofSeconds)                                                     proofSeconds_MAX,
	   round(avg(claimSeconds), 2)                                           claimSeconds_AVG,
	   max(claimSeconds)                                                     claimSeconds_MAX
from (
-- this subquery can be executed by itself for detail data
		 select xdeposit."blockTimestamp"                                depositTs,
				xdeposit."transactionId",
				xrelay."blockTimestamp" - xdeposit."blockTimestamp"      relaySeconds,
				xproof."blockTimestamp" - xrelay."blockTimestamp"        proofSeconds,
				xclaim."blockTimestamp" - xproof."blockTimestamp" - 1800 claimSeconds,
				xdispute.id                                              disputeId,
				xrelay.relayer
		 from "BridgeRequestEvents" xdeposit
				  left join "BridgeRelayedEvents" xrelay on
			 xdeposit."transactionId" = xrelay."transactionId"
				  left join "BridgeProofProvidedEvents" xproof on
			 xdeposit."transactionId" = xproof."transactionId"
				  left join "BridgeDepositClaimedEvents" xclaim on
			 xdeposit."transactionId" = xclaim."transactionId"
				  left join "BridgeProofDisputedEvents" xdispute on
			 xdeposit."transactionId" = xdispute."transactionId"
		 where xdeposit."originChainId" = 480
		   and xdeposit."originAmount" <= 71000000000000

		   -- change these to point at a particular test period
		   and xdeposit."blockTimestamp" between 1728666774 and 1728667721
        ) sqData;

