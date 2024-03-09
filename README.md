# Rollmelette + Cap'n Proto

This repository aims to demonstrate the usage of [Cap'n Proto](https://capnproto.org) as a schema language for [Cartesi Rollups](https://cartesi.io/) applications.
The example application in this repository was built on top of the [Rollmelette](https://github.com/gligneul/rollmelette) high-level framework for added convenience.
It uses [Cap'n Proto](https://capnproto.org) to encode/decode advance requests.

When developing [BugLess](https://github.com/crypto-bug-hunters/bug-less), we decided to describe the application state in Go, since it was the back-end language of choice.
In order for the front-end to be aware of the current state of the application, we made inspect requests return the whole application state in JSON format.
Because we were using TypeScript in the front-end, we decided to mirror the application state definition from Go to TypeScript.
Every time the application state definition in Go would change, we would have to update the TypeScript mirror.
This led to code duplication, and room for inconsistencies between back-end and front-end.

Another issue that arose during development was related to the input encoding.
Since we were using the now-deprecated [EggRoll](https://github.com/gligneul/eggroll) high-level framework, we made use of all tools provided by it.
One of such tools was the JSON Codec, which encoded inputs as the concatenation of a Solidity-like function selector and JSON data.
Inputs were seamlessly decoded by EggRoll, but not-so-seamlessly encoded by the front-end.
The front-end code manually defined the function selectors and serialized data in JSON format without any schema.

Here is where [Cap'n Proto](https://capnproto.org) comes in handy.
Instead of describing data in any specific programming language, you describe it in a schema language.
From this schema definition, you can then generate bindings for any of the 15+ [supported programming languages](https://capnproto.org/otherlang.html).

## Dependencies

- All dependencies for [`rollmelette`](https://github.com/gligneul/rollmelette)
- [`capnp`](https://capnproto.org/install.html)
- [`go-capnp`](https://github.com/capnproto/go-capnp/blob/main/docs/Installation.md)

## Build

You can build the application by running:

```sh
make build
```

## Running

Then, you can run the application with the following command:

```sh
make run
```

## Encoding advance requests

On another terminal, you can encode inputs with the `capnp` CLI tool.
You'll also need `xxd`, `tr` and `sed` to manipulate binary and text data.

```sh
echo '(add=(operand=123))' | \
    capnp convert text:binary request.capnp AdvanceRequest | \
    xxd -ps | tr -d '\n' | sed 's/^/0x/;$a\'
```

You can then copy the output hexstring, and send it to the application it via `sunodo send`.
Be sure to select the "Hex string encoding" option.

## Inspecting

You can inspect the application state as usual.

```sh
curl http://localhost:8080/inspect/
```
