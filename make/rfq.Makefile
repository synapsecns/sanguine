devnet-clean: ## Delete all devnet data and docker containers
	rm -rf contracts/deployments/chain_a/*.json
	rm -rf contracts/deployments/chain_b/*.json

	# TODO: this should also delete broadcast/**/42, broadcast/**/43, broadcast/**/44
	rm -rf contracts/broadcast/42/*
	rm -rf contracts/broadcast/43/*

	docker-compose down --volumes

devnet-up: ## This should be run to start the devnet docker-containers. This does not deploy anything.
	docker-compose up -d --build

devnet-deploy: ## This should be run exactly once to deploy the contracts to the devnet.
	cd contracts/ && \
	forge create FastBridge --private-key 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9 --unlocked --from 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947 --constructor-args 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947 --rpc-url http://localhost:8042 \
	&&  \
	forge create FastBridge --private-key 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9 --unlocked --from 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947 --constructor-args 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947 --rpc-url http://localhost:8043 \
	&&  \
	cast send 0x6438CB36cb18520774EfC7A172410D8BBBe9a428 --rpc-url http://localhost:8042 --private-key 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9 "addRelayer(address)()" 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559 \
	&& \
	cast send 0x6438CB36cb18520774EfC7A172410D8BBBe9a428 --rpc-url http://localhost:8043 --private-key 0x63e21d10fd50155dbba0e7d3f7431a400b84b4c2ac1ee38872f82448fe3ecfb9 "addRelayer(address)()" 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
