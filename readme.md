# Work in progress
Batagar is not ready yet. It has no integration with mainstream fuzzers.


# batagar 

Batagar is distributed fuzzing platform. I've designed it for easy fuzzing open source projects on my VPSes with minimal effort.

Features:
- Worker fail safe - Worker failure does not affect coordinator nor other workers
- Validates inputs from workers - allows accepting fuzzing from untrusted providers, so anyone can allow crowd-sourced fuzzing for more secure software
- Provides basic user interface with statistics(exposed also as API)
- Provides support for A/B testing: may selectively restart fuzzers, run N group of fuzzers with different options (NIY)
- Automatic crash encryption - in case attacker breaks in on a master server only for a short while, he won't be able to see all previous crashers, found in the past.
- Can work as simple CTF platform: worker has to provide payload, which causes program X to behave in a specific way(NIY)

Batagar already implements support for some popular fuzzers:
- [go-fuzz](https://github.com/dvyukov/go-fuzz)
- [libfuzzer](https://llvm.org/docs/LibFuzzer.html)

### Joining Batagar cluster with docker:
`docker run teawithsand/batagar:master -e SERVER_ADDRESS=/coordinator server address goes here/`

### Creating batagar cluster:
TODO(teawithsand): manual here

### Main components
Batagar is split into two parts:
- Worker: Fuzzer extension, which integrates fuzzers like [go-fuzz](https://github.com/dvyukov/go-fuzz) with Batagar
- Coordinator: component, which accepts connection from workers and collects crashes

### Batagar protocol:
Batagar uses GRPC to communicate clients and servers, in fact API is quite simple so anyone can implement their own client/server.

It even has it's own documentation HERE(TODO: LINK+DOCS HERE)

### Notes
In fact, Batagar can be used for any distributed computation, which works like fuzzing: It's hard to find value X, but once it's found it's clear that X is either right or wrong, which can be easily checked.

Also please note that Batagar was created for purpose of online platform for easy fuzzing of open source software, so any company or person which uses library X, but is not sure about it's security can easily run fuzzer, which automatically reports bug to the creator of X.