# Rollmelette + Cap'n Proto

This repository aims to demonstrate the usage of [Cap'n Proto](https://capnproto.org) as a schema language for [Cartesi Rollups](https://cartesi.io/) applications.
The example application in this repository was built on top of the [Rollmelette]((https://github.com/gligneul/rollmelette)) high-level framework for added convenience.
It uses [Cap'n Proto](https://capnproto.org) to encode/decode advance and inspect requests.

When developing [BugLess](https://github.com/crypto-bug-hunters/bug-less), we decided to describe the application state in Go, since it was the back-end language of choice.
In order for the front-end to be aware of the current state of the application, we made inspect requests return the whole application state in JSON format.
Because we were using TypeScript in the front-end, we decided to mirror the application state definition from Go to TypeScript.
Every time the application state definition in Go would change, we would have to update the TypeScript mirror.
This led to code duplication, and room for inconsistencies between back-end and front-end.

Another issue that arose during development was related to the input encoding.
Since we were using the now-deprecated [EggRoll](https://github.com/gligneul/eggroll) high-level framework, we made use of all tools provided by it.
One of such tools was the JSON Codec, which encoded inputs as the concatenation of a Solidity-like function selector and JSON data.
Inputs were encoded seamlessly by EggRoll, but not-so-seamlessly by the front-end.
The front-end had to define several magic numbers (the function selectors) and serialize data in JSON format without any schema.

Here is where [Cap'n Proto](https://capnproto.org) comes in handy.
Instead of describing data in any specific programming language, you describe it in a schema language.
From this schema definition, you can then generate bindings for any of the 15+ [supported programming languages](https://capnproto.org/otherlang.html).

## Dependencies

- All dependencies for [`rollmelette`](https://github.com/gligneul/rollmelette)
- [`capnp`](https://capnproto.org/install.html)
- [`go-capnp`](https://github.com/capnproto/go-capnp/blob/main/docs/Installation.md)
