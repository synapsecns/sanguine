# Modules

Modules are implementations that can modify the inputs to or outputs of an rpc call. They are meant to deal w/ specific application level limitations or requirements. For example, a module could be used to add a custom header to all requests, or to modify the response of a call to a specific service. These do not neccesarily emulate the original functionality of omnirpc and are run through seperate commands.
