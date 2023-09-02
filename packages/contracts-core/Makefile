include ../../make/solidity.Makefile

devnet-clean:
	rm -rf deployments/chain_a/*.json
	rm -rf deployments/chain_b/*.json
	rm -rf deployments/chain_c/*.json

	rm -rf broadcast/42/*
	rm -rf broadcast/43/*
	rm -rf broadcast/44/*

	cd ../../docker/devnet && docker-compose down -v

devnet-deploy:
	forge script script/DeployMessaging003SynChain.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployMessaging003LightChain.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast

	forge script script/DeployClients003.s.sol  --ffi -f chain_a --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployClients003.s.sol  --ffi -f chain_b --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
	forge script script/DeployClients003.s.sol  --ffi -f chain_c --private-key 63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9   --broadcast
