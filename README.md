# James Robb Ably Takehome Test

## Building/Running

I assume you have `go >= 1.18` installed. Checkout this repository and then `cd` into the directory.

To run the server execute the following. It will start on port `50051`.

    go run ./cmd/server/...

To run the client execute the following. By default it will try to connect to a server located at `localhost:50051`. The parameter `-numMessages` denotes how many numbers to receive from the server.

    go run ./cmd/client/... -numMessages=5

Both the server and the client have command line options that can be viewed by passing `-h` when invoking them. The `-h` output documents the various options (e.g., what port to use).

## Protocol

The protocol is implemented with gRPC and Protobuf. The messages exchanged are located in `protocol/protocol.proto`. A client (identified by `client_id`) issues a `NumbersRequest` message to the server to receive `num_numbers` numbers back (in the form of `NumberResponse` messages). The `seed` field is for testing purposes so that a client can seed the PRNG on the server-side and thus know what numbers to expect. When the seed is `0` the server is expected to generate a seed value at random for the client.

The `NumberResponse` message that contains the last number to be sent to a client will also include a checksum in the `checksum` field (which is an empty string otherwise). The checksum calculation is simple and can be seen in `cmd/client/main.go` in the `calculateChecksum` function. The server calculates the checksum in the same way, but the operations to do so aren't as nicely grouped as in the client's code.

The client is able to verify that it received the correct sequence of numbers by using `calculateChecksum` to produce a checksum with the received numbers and comparing the result with the checksum included in the last `NumberResponse` received.

The protobuf messages and gRPC service are compiled to Golang with `compile_protos.sh`.


## Testing

I chose to implement just a basic test for the project. The client has a test mode which simulates a situation in where the client disconnects and attempts to resume receiving numbers. It does this by connecting to the server, disconnecting after receiving half the numbers, and then connecting again with the same client ID. The test can be seen in `cmd/client/main.go` in the `testOperation` function.

There are command line options that one can specify when running the client to run the test scenario describe above. An example of their usage can be seen in `test.sh`.

## Notes For Reviewers

This is a basic implmentation of the task sent to me. There is room for improvements and optimizations everywhere, but given the purpose of the task I tried to focus on what was most relevant.

All the code for the client is in `cmd/client/main.go`. The server code is a bit more spread out though. The real "meat" of the server is in `cmd/server/number_server.go`.

As asked I created an interface for the client state storage. The interface could easily be satisfied with a different backing (.e.g. Redis) but here I use an in-memory store. The in-memory store is very basic (though safe to be used concurrently) and could use some improvement (specifically with respect to garbage collection).

This project was not stress tested. Any limits on the number of concurrent connections, payload size, and so on will be a function of what gRPC allows by default, how much memory the host machine has, etc. This is and should be treated as a proof of concept.

The client (when not in test mode) can tolerate a disconnect/reconnect because it relies on automatic connection retrying built into the gRPC code. Because the gRPC code will attempt to restablish the connection the state is not lost on the client side. An improvement to the client would be command line options to specify the state so that the binary could be be stopped and started again.

Lastly, one important note is that the server implementation does not handle the corner case where two distinct clients connect with the same client ID concurrently. The assumption is that client IDs are generated randomly from a very large space of values and the likelihood that two clients would generate the same client ID is negligible. In production I would not make such an assumption. However, what still does work is that if a client connects with client ID `X` and loses a connection, said client can still use client ID `X` when trying to resume the stream of numbers (which is simulated in the test).