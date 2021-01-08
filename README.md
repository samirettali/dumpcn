# dumpcn

A tool that takes a list of domains from standard input (prefixed with `https://` or
not), grabs the certificate and prints the CN.

It runs on 32 threads (goroutines) by default.


Running it against
[opendns-top-domains.txt](https://raw.githubusercontent.com/opendns/public-domain-lists/master/opendns-top-domains.txt)
(10000 domains) takes approximately 7 seconds:
```
$ time ./dumpcn -t=100 < opendns-top-domains.txt > /dev/null
./dumpcn -t=100 < opendns-top-domains.txt > /dev/null  1,66s user 1,24s system 39% cpu 7,311 total
```

## Installation

```
$ go get -u github.com/samirettali/dumpcn
```

## Usage

```
$ cat domains.txt | ./dumpcn
```

Change number of threads:
```
$ cat domains.txt | ./dumpcn -t=100
```
