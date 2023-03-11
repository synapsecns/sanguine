// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BasicClientTools.t.sol";
import { ClientHarness } from "../../harnesses/client/ClientHarness.t.sol";

abstract contract ClientTools is BasicClientTools {
    function deployBasicClient(uint32 domain, uint32 optimisticSeconds)
        public
        virtual
        override
        returns (BasicClient)
    {
        return
            new ClientHarness(
                address(suiteOrigin(domain)),
                address(suiteDestination(domain)),
                optimisticSeconds
            );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectLogClientMessage() public {
        emit LogClientMessage(messageOrigin, messageNonce, messageBody);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function clientSendMessage(bytes memory revertMessage) public {
        vm.expectRevert(revertMessage);
        suiteClient(messageOrigin).sendMessage{ value: tipsTotal }(
            messageDestination,
            tipsRaw,
            messageBody
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function clientSendMessage() public {
        suiteClient(messageOrigin).sendMessage{ value: tipsTotal }(
            messageDestination,
            tipsRaw,
            messageBody
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            CHAIN GETTERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function suiteClient(uint32 domain) public view returns (ClientHarness) {
        return ClientHarness(address(basicClients[domain]));
    }
}
