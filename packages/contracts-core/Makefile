include ../../make/solidity.Makefile

devnet-clean: ## Delete all devnet data and docker containers
	rm -rf deployments/chain_a/*.json
	rm -rf deployments/chain_b/*.json
	rm -rf deployments/chain_c/*.json

	# TODO: this should also delete broadcast/**/42, broadcast/**/43, broadcast/**/44
	rm -rf broadcast/42/*
	rm -rf broadcast/43/*
	rm -rf broadcast/44/*

	cd ../../docker/devnet && docker-compose down --volumes

devnet-up: ## This should be run to start the devnet docker-containers. This does not deploy anything.
	cd ../../docker/devnet && docker-compose up -d --build

devnet-deploy: ## This should be run exactly once to deploy the contracts to the devnet.
	# backup the .env in case there's real data in there
	cp .env /tmp/.env.bak || true
	cp devnet.env .env

	forge script script/DeployMessaging003SynChain.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast

	forge script script/DeployClients003.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployClients003.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployClients003.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast

devnet-logs:
	cd ../../docker/devnet && docker-compose logs -f
