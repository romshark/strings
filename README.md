<a href="https://pkg.go.dev/github.com/romshark/strings">
    <img src="https://godoc.org/github.com/romshark/strings?status.svg" alt="GoDoc">
</a>
<a href="https://goreportcard.com/report/github.com/romshark/strings">
    <img src="https://goreportcard.com/badge/github.com/romshark/strings" alt="GoReportCard">
</a>
<a href='https://coveralls.io/github/romshark/strings?branch=main'>
    <img src='https://coveralls.io/repos/github/romshark/strings/badge.svg?branch=main' alt='Coverage Status' />
</a>

# strings

An attempt at improving the performance of Go standard library functions
of package https://pkg.go.dev/strings (mostly just for the fun of it).

## Benchmark `ToLower`

To run all benchmarks, use:

```sh
./to_lower.sh . 10 old.txt new.txt
```

### Apple M1 Max

```
goos: darwin
goarch: arm64
pkg: github.com/romshark/strings
cpu: Apple M1 Max
                                 │   old.txt    │               new.txt               │
                                 │    sec/op    │   sec/op     vs base                │
ToLower/empty________________-10    2.018n ± 0%   2.173n ± 0%   +7.65% (p=0.000 n=10)
ToLower/ascii-1______________-10   13.530n ± 0%   3.725n ± 0%  -72.47% (p=0.000 n=10)
ToLower/ascii-1-low__________-10    2.484n ± 0%   2.794n ± 0%  +12.50% (p=0.000 n=10)
ToLower/ascii-2______________-10    16.90n ± 0%   11.96n ± 0%  -29.23% (p=0.000 n=10)
ToLower/ascii-2-low__________-10    3.725n ± 0%   4.035n ± 0%   +8.32% (p=0.000 n=10)
ToLower/ascii-3______________-10    19.61n ± 0%   13.71n ± 1%  -30.11% (p=0.000 n=10)
ToLower/ascii-3-low__________-10    4.346n ± 0%   4.657n ± 0%   +7.16% (p=0.000 n=10)
ToLower/ascii-7______________-10    30.20n ± 0%   15.93n ± 0%  -47.24% (p=0.000 n=10)
ToLower/ascii-7-low__________-10    6.829n ± 0%   7.141n ± 0%   +4.58% (p=0.000 n=10)
ToLower/ascii-8______________-10    33.31n ± 0%   16.55n ± 0%  -50.31% (p=0.000 n=10)
ToLower/ascii-8-low__________-10    8.381n ± 0%   8.381n ± 0%        ~ (p=0.625 n=10)
ToLower/ascii-9______________-10    39.83n ± 0%   20.34n ± 0%  -48.95% (p=0.000 n=10)
ToLower/ascii-9-low__________-10    9.002n ± 0%   9.001n ± 0%        ~ (p=0.409 n=10)
ToLower/ascii-33-capital_____-10    71.24n ± 0%   51.40n ± 0%  -27.85% (p=0.000 n=10)
ToLower/ascii-33-most-up_____-10    98.93n ± 0%   41.55n ± 0%  -58.00% (p=0.000 n=10)
ToLower/ascii-33-up__________-10    93.38n ± 0%   42.46n ± 0%  -54.54% (p=0.000 n=10)
ToLower/ascii-33-low_________-10    27.64n ± 0%   27.33n ± 0%   -1.12% (p=0.000 n=10)
ToLower/ascii-49-capital_____-10    99.60n ± 0%   76.35n ± 0%  -23.34% (p=0.000 n=10)
ToLower/ascii-49-up---_______-10   134.85n ± 0%   80.43n ± 0%  -40.36% (p=0.000 n=10)
ToLower/ascii-49-low----------10    42.22n ± 0%   36.02n ± 0%  -14.68% (p=0.000 n=10)
ToLower/ascii-loremipsum_____-10    743.4n ± 0%   475.2n ± 0%  -36.08% (p=0.000 n=10)
ToLower/ascii-loremipsum-low_-10    539.8n ± 0%   341.4n ± 0%  -36.75% (p=0.000 n=10)
ToLower/loremipsum_u8end_____-10   1889.5n ± 0%   560.0n ± 0%  -70.36% (p=0.000 n=10)
ToLower/loremipsum-low_u8end_-10   1889.0n ± 0%   559.5n ± 0%  -70.38% (p=0.000 n=10)
ToLower/romeo-juliet_________-10    201.4µ ± 0%   154.8µ ± 0%  -23.13% (p=0.000 n=10)
ToLower/romeo-juliet-low_____-10   329.52µ ± 1%   71.32µ ± 0%  -78.36% (p=0.000 n=10)
ToLower/utf8_4-1_____________-10    18.72n ± 0%   19.71n ± 0%   +5.32% (p=0.000 n=10)
ToLower/utf8_2-1_____________-10    28.71n ± 0%   29.04n ± 0%   +1.15% (p=0.000 n=10)
ToLower/utf8_2-1-low_________-10    18.01n ± 0%   18.43n ± 0%   +2.33% (p=0.000 n=10)
ToLower/utf8_3-1_____________-10    18.94n ± 0%   19.41n ± 0%   +2.48% (p=0.000 n=10)
ToLower/utf8_3-3_____________-10    50.02n ± 0%   50.34n ± 0%   +0.65% (p=0.000 n=10)
ToLower/utf8-japanese________-10    670.3n ± 0%   670.7n ± 0%   +0.05% (p=0.002 n=10)
ToLower/hallo________________-10    2.789µ ± 0%   2.795µ ± 0%   +0.22% (p=0.001 n=10)
ToLower/mixed-764b___________-10    2.790µ ± 0%   2.794µ ± 0%        ~ (p=0.082 n=10)
ToLower/32-ascii_1-utf8______-10   122.70n ± 0%   36.64n ± 0%  -70.14% (p=0.000 n=10)
ToLower/mixed-764b-low_______-10    2.449µ ± 0%   2.447µ ± 0%        ~ (p=0.381 n=10)
ToLower/wiki-japan-en-html___-10    2.864m ± 0%   1.191m ± 0%  -58.42% (p=0.000 n=10)
ToLower/wiki-japan-jp-html___-10    9.214m ± 0%   6.075m ± 0%  -34.06% (p=0.000 n=10)
geomean                             166.8n        115.2n       -30.94%
```
