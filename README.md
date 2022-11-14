## Go Ronin

Official Golang execution layer implementation of the Ronin protocol. It is a fork of Go Ethereum - 
[https://github.com/ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) and EVM compatible.

Ronin consensus use Proof of Staked Authority, a combination of dPoS and PoA, to increase the level of 
decentralization and allows the token holders to join the network as validators

Check out the [whitepaper]() for more information.

[![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.com/invite/pjgPrrZJyZ)

## Executables

The go-ethereum project comes with several wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`ronin`**   | Our main Ronin CLI client. It is the entry point into the Ronin network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Ronin network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `ronin --help` and the [CLI page](https://geth.ethereum.org/docs/interface/command-line-options) for command line options.          |
|   `clef`    | Stand-alone signing tool, which can be used as a backend signer for `ronin`.  |
|   `devp2p`    | Utilities to interact with nodes on the networking layer, without running a full blockchain. |
|   `abigen`    | Source code generator to convert Ethereum contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Ethereum contract ABIs](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://geth.ethereum.org/docs/dapp/native-bindings) page for details. |
|  `bootnode`   | Stripped down version of our Ethereum client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks.                                                                                                                                                                                                                                                                 |
|     `evm`     | Developer utility version of the EVM (Ethereum Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow isolated, fine-grained debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                                                     |
|   `rlpdump`   | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://eth.wiki/en/fundamentals/rlp)) dumps (data encoding used by the Ethereum protocol both network as well as consensus wise) to user-friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                                                                                 |
|   `puppeth`   | a CLI wizard that aids in creating a new Ethereum network.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

## Running `ronin`

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://geth.ethereum.org/docs/interface/command-line-options)),
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `geth` instance.

### Requirements
The minimum recommended hardware specification for nodes connected to Mainnet is:
- CPU: Equivalent of 8 AWS vCPU
- RAM: 16GiB
- Storage: 1 TiB
- OS: Ubuntu 20.04 or macOS >= 12
- Network: Reliable IPv4 or IPv6 network connection, with an open public port

### Building the source
Building `ronin` requires both a Go (version 1.17 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make ronin
```

or, to build the full suite of utilities:

```shell
make all
```

### Full node on the main Ronin network

```shell
$ ronin --http.api eth,net,web3,consortium --networkid 2020 --bootnodes enode://a166ab6437cf370bc604097529a0fb6a8a4836bb85833fbf588b130cb73fe0517940d10c5d89c0e3e1c2800a774ac1ae2108d62cb4608556e41bc1fc4482241a@35.193.159.26:30303 --datadir /opt/ronin --port 30303 --http --http.corsdomain '*' --http.addr 0.0.0.0 --http.port 8545 --http.vhosts '*' --ws --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins '*' 
```

This command will:
 * Start `ronin` in snap sync mode (default, can be changed with the `--syncmode` flag),
   causing it to download more data in exchange for avoiding processing the entire history
   of the Ronin network, which is very CPU intensive.

### Configuration

As an alternative to passing the numerous flags to the `ronin` binary, you can also pass a
configuration file via:

```shell
$ ronin --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ ronin --your-favourite-flags dumpconfig
```

### Programmatically interfacing `ronin` nodes

As a developer, sooner rather than later you'll want to start interacting with `ronin` and the
Ronin network via your own programs and not manually through the console. To aid
this, `ronin` has built-in support for a JSON-RPC based APIs which are the same as Ethereum that can be found at ([standard APIs](https://eth.wiki/json-rpc/API)
and [`ronin` specific APIs](https://geth.ethereum.org/docs/rpc/server)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `ronin`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `geth` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
Ethereum nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

## How to contribute

### Contribution guidelines
- Quality: Code in the Ronin project should meet the style guidelines, with sufficient test-cases, descriptive commit
  messages, evidence that the contribution does not break any compatibility commitments or cause adverse feature
  interactions, and evidence of high-quality peer-review.
- Size: The Ronin project's culture is one of small pull-requests, regularly submitted. The larger a pull-request,
  the more likely it is that you will be asked to resubmit as a series of self-contained and individually reviewable
  smaller PRs.
- Maintainability: If the feature will require ongoing maintenance (eg support for a particular branch of database),
  we nay ask you to accept responsibility for maintaining this feature
- Commit message: Commit messages of Ronin project follows [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/)

### Submit an issue
- Create a new issue
- Comment on the issue (if you'd like to be assigned to it) - that way our team can assign the issue to you
- If you do not have a specific contribution in mind, you can also browse the issues labelled as `help wanted`
- Issues that additionally have the `good first issue` label are considered ideal for first-timers

### Submit your PR
- After your changes are committed to your GitHub fork, submit a pull request (PR) to the `master` branch of the
  axieinfinity/ronin repo
- In your PR description, reference the issue it resolves (see [linking a pull request to an issue using a keyword](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue#linking-a-pull-request-to-an-issue-using-a-keyword))
    - ex: `[FIXES #123] feat: update out of date content`

### Wait for review
- The team reviews every PR
- Acceptable PRs will be approved & merged into the `master` branch

### Release
- You can [view the history of release](https://github.com/axieinfinity/ronin/releases), which include PR highlights

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
