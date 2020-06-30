# DumpCN

A simple script that reads a list of domains (starting with `https://` or
not) from standard input, grabs the certificate and prints the CN. By default it
runs 32 threads (goroutines actually).


Running it against
[opendns-top-domains.txt](https://raw.githubusercontent.com/opendns/public-domain-lists/master/opendns-top-domains.txt)
(10000 domains) takes approximately 1.5 seconds:
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
