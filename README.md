# TravelgateX GraphQL Benchmark

This is a repository to compare different types of presenters in means of performance. We are measuring the serialization of Options to an HTTP Response Body, where Options are always the same mock.

It is a fork of [travelgateX/presenters-benchmark](https://github.com/travelgateX/presenters-benchmark) with Artemis candidate added.

## Implementations

	- GraphQL
		- artemis
			- from service models: resolve directly from service models
		- gqlgen
			- mapping: from application models to presentation models then serialize
			- from service models: resolve directly from service models
		- gophers
			- mapping (the only way)
	- REST
		- JSON
			- mapping
			- from service models


To add another implementation create its own package under `pkg/presenter` and implement the interface:

```go
type CandidateHandlerFunc interface {
	HandlerFunc(options []*Option) (http.HandlerFunc, error)
}
```

## Run the tests

A valid candidate must pass tests in [candidate.go](pkg/presenter/candidate.go).

Usage example of the [gophers](pkg/presenter/gophers/gophers_test.go) implementation:

```go
func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
```

## Run the benchmarks

To run the benchmarks, run `BenchmarkSequential`  or `BenchmarkParallel` in [benchmark_test.go](benchmark/benchmark_test.go)

Use [this excel](benchmark/graphics.xlsx) to build graphics

### Environment

- **Model**: MacBook Pro (15-inch, 2016)

- **OS**: macOS Mojave 10.14.2

- **CPU**: Intel Core i7-6820HQ @ 2.7GHz × 4

- **RAM**: 16 GiB 2133MHz LPDDR3

### Sample

Request: [high resolve scale](benchmark/resolveScale_high.txt)

Response: N times [option](benchmark/option.json)


- Sequential benchmarks

```
$ go test -test.benchmem -test.bench BenchmarkSequential ./benchmark
goos: darwin
goarch: amd64
pkg: github.com/travelgateX/presenters-benchmark/benchmark
BenchmarkSequential/artemis/1-8         	   10000	    162258 ns/op	   70190 B/op	     877 allocs/op
BenchmarkSequential/artemis/2-8         	   10000	    201839 ns/op	   87691 B/op	    1066 allocs/op
BenchmarkSequential/artemis/4-8         	    5000	    275541 ns/op	  103744 B/op	    1441 allocs/op
BenchmarkSequential/artemis/8-8         	    3000	    421828 ns/op	  146623 B/op	    2188 allocs/op
BenchmarkSequential/artemis/16-8        	    2000	    711826 ns/op	  236470 B/op	    3678 allocs/op
BenchmarkSequential/artemis/32-8        	    1000	   1297181 ns/op	  416141 B/op	    6657 allocs/op
BenchmarkSequential/artemis/64-8        	     500	   2455579 ns/op	  771580 B/op	   12614 allocs/op
BenchmarkSequential/artemis/128-8       	     300	   4835664 ns/op	 1474365 B/op	   24525 allocs/op
BenchmarkSequential/artemis/256-8       	     200	   9609064 ns/op	 2879662 B/op	   48342 allocs/op
BenchmarkSequential/artemis/512-8       	     100	  19241385 ns/op	 5677035 B/op	   95966 allocs/op
BenchmarkSequential/artemis/1024-8      	      30	  37496044 ns/op	11304069 B/op	  191205 allocs/op
BenchmarkSequential/artemis/2048-8      	      20	  71958972 ns/op	22592159 B/op	  381671 allocs/op
BenchmarkSequential/artemis/4096-8      	      10	 145537484 ns/op	45415664 B/op	  762610 allocs/op
BenchmarkSequential/artemis/8192-8      	       5	 293275125 ns/op	91094294 B/op	 1524475 allocs/op
BenchmarkSequential/artemis/16384-8     	       2	 636053429 ns/op	182411420 B/op	 3048197 allocs/op
BenchmarkSequential/artemis/32768-8     	       1	1370964110 ns/op	365000752 B/op	 6096548 allocs/op
BenchmarkSequential/artemis/65536-8     	       1	2802190032 ns/op	729855872 B/op	12191400 allocs/op
BenchmarkSequential/gophers/1-8         	    2000	    638981 ns/op	  351465 B/op	    1919 allocs/op
BenchmarkSequential/gophers/2-8         	    2000	    736646 ns/op	  388746 B/op	    2609 allocs/op
BenchmarkSequential/gophers/4-8         	    2000	    945624 ns/op	  458172 B/op	    3989 allocs/op
BenchmarkSequential/gophers/8-8         	    1000	   1335921 ns/op	  601703 B/op	    6747 allocs/op
BenchmarkSequential/gophers/16-8        	    1000	   2053868 ns/op	  892068 B/op	   12263 allocs/op
BenchmarkSequential/gophers/32-8        	     500	   3599284 ns/op	 1472416 B/op	   23295 allocs/op
BenchmarkSequential/gophers/64-8        	     200	   6652235 ns/op	 2626624 B/op	   45356 allocs/op
BenchmarkSequential/gophers/128-8       	     100	  12283590 ns/op	 4930106 B/op	   89479 allocs/op
BenchmarkSequential/gophers/256-8       	      50	  25376391 ns/op	 9516811 B/op	  177720 allocs/op
BenchmarkSequential/gophers/512-8       	      30	  48344129 ns/op	18799162 B/op	  354204 allocs/op
BenchmarkSequential/gophers/1024-8      	      20	  95571034 ns/op	37675183 B/op	  707173 allocs/op
BenchmarkSequential/gophers/2048-8      	      10	 190319524 ns/op	76625963 B/op	 1413104 allocs/op
BenchmarkSequential/gophers/4096-8      	       3	 375987624 ns/op	154558637 B/op	 2824984 allocs/op
BenchmarkSequential/gophers/8192-8      	       2	 760864596 ns/op	308727416 B/op	 5648709 allocs/op
BenchmarkSequential/gophers/16384-8     	       1	1486118474 ns/op	617052504 B/op	11296120 allocs/op
BenchmarkSequential/gophers/32768-8     	       1	2929256832 ns/op	1233708360 B/op	22590945 allocs/op
BenchmarkSequential/gophers/65536-8     	       1	6120948099 ns/op	2467058288 B/op	45180819 allocs/op
BenchmarkSequential/gqlgen_mapping/1-8  	    5000	    272282 ns/op	   84402 B/op	    1288 allocs/op
BenchmarkSequential/gqlgen_mapping/2-8  	    3000	    341130 ns/op	  128288 B/op	    2110 allocs/op
BenchmarkSequential/gqlgen_mapping/4-8  	    3000	    516051 ns/op	  215518 B/op	    3753 allocs/op
BenchmarkSequential/gqlgen_mapping/8-8  	    2000	    751735 ns/op	  390012 B/op	    7038 allocs/op
BenchmarkSequential/gqlgen_mapping/16-8 	    1000	   1264483 ns/op	  739852 B/op	   13607 allocs/op
BenchmarkSequential/gqlgen_mapping/32-8 	    1000	   2289591 ns/op	 1442200 B/op	   26745 allocs/op
BenchmarkSequential/gqlgen_mapping/64-8 	     300	   4279979 ns/op	 2859428 B/op	   53022 allocs/op
BenchmarkSequential/gqlgen_mapping/128-8         	     200	   8443408 ns/op	 5692348 B/op	  105575 allocs/op
BenchmarkSequential/gqlgen_mapping/256-8         	     100	  16174498 ns/op	11412084 B/op	  210714 allocs/op
BenchmarkSequential/gqlgen_mapping/512-8         	      50	  33910326 ns/op	23009105 B/op	  420952 allocs/op
BenchmarkSequential/gqlgen_mapping/1024-8        	      20	  69487699 ns/op	46148732 B/op	  842007 allocs/op
BenchmarkSequential/gqlgen_mapping/2048-8        	      10	 137803210 ns/op	92583156 B/op	 1685349 allocs/op
BenchmarkSequential/gqlgen_mapping/4096-8        	       5	 292010364 ns/op	186001816 B/op	 3371292 allocs/op
BenchmarkSequential/gqlgen_mapping/8192-8        	       2	 601478412 ns/op	372237128 B/op	 6746861 allocs/op
BenchmarkSequential/gqlgen_mapping/16384-8       	       1	1381731542 ns/op	749784432 B/op	13510926 allocs/op
BenchmarkSequential/gqlgen_mapping/32768-8       	       1	2752979944 ns/op	1493361296 B/op	26990424 allocs/op
BenchmarkSequential/gqlgen_mapping/65536-8       	       1	5928432639 ns/op	2983088352 B/op	53940233 allocs/op
BenchmarkSequential/gqlgen_service_models/1-8    	    5000	    270317 ns/op	   83373 B/op	    1280 allocs/op
BenchmarkSequential/gqlgen_service_models/2-8    	    3000	    349772 ns/op	  126187 B/op	    2096 allocs/op
BenchmarkSequential/gqlgen_service_models/4-8    	    3000	    514239 ns/op	  211318 B/op	    3727 allocs/op
BenchmarkSequential/gqlgen_service_models/8-8    	    2000	    774149 ns/op	  383010 B/op	    6988 allocs/op
BenchmarkSequential/gqlgen_service_models/16-8   	    1000	   1274907 ns/op	  732515 B/op	   13509 allocs/op
BenchmarkSequential/gqlgen_service_models/32-8   	    1000	   2287811 ns/op	 1425746 B/op	   26551 allocs/op
BenchmarkSequential/gqlgen_service_models/64-8   	     300	   4271489 ns/op	 2812598 B/op	   52634 allocs/op
BenchmarkSequential/gqlgen_service_models/128-8  	     200	   8108841 ns/op	 5605322 B/op	  104806 allocs/op
BenchmarkSequential/gqlgen_service_models/256-8  	     100	  16308579 ns/op	11226694 B/op	  209156 allocs/op
BenchmarkSequential/gqlgen_service_models/512-8  	      50	  34515531 ns/op	22589543 B/op	  417846 allocs/op
BenchmarkSequential/gqlgen_service_models/1024-8 	      20	  71762912 ns/op	45598268 B/op	  835325 allocs/op
BenchmarkSequential/gqlgen_service_models/2048-8 	      10	 142727920 ns/op	91444271 B/op	 1671966 allocs/op
BenchmarkSequential/gqlgen_service_models/4096-8 	       5	 304702011 ns/op	182863731 B/op	 3344023 allocs/op
BenchmarkSequential/gqlgen_service_models/8192-8 	       2	 604461060 ns/op	366567600 B/op	 6700299 allocs/op
BenchmarkSequential/gqlgen_service_models/16384-8         	       1	1235292447 ns/op	733403376 B/op	13404997 allocs/op
BenchmarkSequential/gqlgen_service_models/32768-8         	       1	2645419084 ns/op	1466009264 B/op	26799843 allocs/op
BenchmarkSequential/gqlgen_service_models/65536-8         	       1	5367543581 ns/op	2931217192 B/op	53545082 allocs/op
BenchmarkSequential/rest_json_service_models/1-8          	   20000	     93227 ns/op	   40531 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/2-8          	   10000	    104216 ns/op	   43094 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/4-8          	   10000	    118613 ns/op	   47710 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/8-8          	   10000	    152368 ns/op	   56687 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/16-8         	    5000	    226645 ns/op	   79248 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/32-8         	    3000	    356961 ns/op	  112060 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/64-8         	    2000	    620936 ns/op	  177806 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/128-8        	    1000	   1151344 ns/op	  309183 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/256-8        	    1000	   2232227 ns/op	  574388 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/512-8        	     300	   4322850 ns/op	 1119391 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/1024-8       	     200	   8905152 ns/op	 2240863 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/2048-8       	     100	  17664340 ns/op	 4540108 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/4096-8       	      30	  36191238 ns/op	 9105974 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/8192-8       	      20	  70312710 ns/op	18818422 B/op	     400 allocs/op
BenchmarkSequential/rest_json_service_models/16384-8      	      10	 137443512 ns/op	41461148 B/op	     401 allocs/op
BenchmarkSequential/rest_json_service_models/32768-8      	       5	 292944279 ns/op	98358760 B/op	     404 allocs/op
BenchmarkSequential/rest_json_service_models/65536-8      	       2	 572322823 ns/op	289567944 B/op	     412 allocs/op
BenchmarkSequential/rest_json_mapping/1-8                 	   20000	     94826 ns/op	   41890 B/op	     408 allocs/op
BenchmarkSequential/rest_json_mapping/2-8                 	   10000	    104929 ns/op	   45830 B/op	     416 allocs/op
BenchmarkSequential/rest_json_mapping/4-8                 	   10000	    127969 ns/op	   53135 B/op	     432 allocs/op
BenchmarkSequential/rest_json_mapping/8-8                 	   10000	    163767 ns/op	   67373 B/op	     464 allocs/op
BenchmarkSequential/rest_json_mapping/16-8                	    5000	    239211 ns/op	  100190 B/op	     528 allocs/op
BenchmarkSequential/rest_json_mapping/32-8                	    3000	    382262 ns/op	  154455 B/op	     656 allocs/op
BenchmarkSequential/rest_json_mapping/64-8                	    2000	    682878 ns/op	  263252 B/op	     912 allocs/op
BenchmarkSequential/rest_json_mapping/128-8               	    1000	   1332254 ns/op	  483120 B/op	    1424 allocs/op
BenchmarkSequential/rest_json_mapping/256-8               	     500	   2492459 ns/op	  916368 B/op	    2448 allocs/op
BenchmarkSequential/rest_json_mapping/512-8               	     300	   4813386 ns/op	 1791080 B/op	    4496 allocs/op
BenchmarkSequential/rest_json_mapping/1024-8              	     100	  10072998 ns/op	 3735174 B/op	    8592 allocs/op
BenchmarkSequential/rest_json_mapping/2048-8              	      50	  20562650 ns/op	 7518072 B/op	   16785 allocs/op
BenchmarkSequential/rest_json_mapping/4096-8              	      30	  40147208 ns/op	14478329 B/op	   33168 allocs/op
BenchmarkSequential/rest_json_mapping/8192-8              	      20	  82193279 ns/op	39078949 B/op	   65942 allocs/op
BenchmarkSequential/rest_json_mapping/16384-8             	      10	 162707270 ns/op	78097652 B/op	  131479 allocs/op
BenchmarkSequential/rest_json_mapping/32768-8             	       5	 327206981 ns/op	232337212 B/op	  262563 allocs/op
BenchmarkSequential/rest_json_mapping/65536-8             	       2	 658297356 ns/op	525528492 B/op	  524713 allocs/op
BenchmarkSequential/protobuf_mapping/1-8                  	   20000	     87373 ns/op	   42392 B/op	     422 allocs/op
BenchmarkSequential/protobuf_mapping/2-8                  	   20000	     94866 ns/op	   46528 B/op	     443 allocs/op
BenchmarkSequential/protobuf_mapping/4-8                  	   10000	    102271 ns/op	   55312 B/op	     485 allocs/op
BenchmarkSequential/protobuf_mapping/8-8                  	   10000	    120561 ns/op	   71856 B/op	     569 allocs/op
BenchmarkSequential/protobuf_mapping/16-8                 	   10000	    153252 ns/op	  104432 B/op	     737 allocs/op
BenchmarkSequential/protobuf_mapping/32-8                 	    5000	    217260 ns/op	  178800 B/op	    1073 allocs/op
BenchmarkSequential/protobuf_mapping/64-8                 	    5000	    351166 ns/op	  302960 B/op	    1745 allocs/op
BenchmarkSequential/protobuf_mapping/128-8                	    2000	    635968 ns/op	  567664 B/op	    3089 allocs/op
BenchmarkSequential/protobuf_mapping/256-8                	    2000	   1134163 ns/op	 1097072 B/op	    5777 allocs/op
BenchmarkSequential/protobuf_mapping/512-8                	    1000	   2239386 ns/op	 2139504 B/op	   11153 allocs/op
BenchmarkSequential/protobuf_mapping/1024-8               	     300	   4363707 ns/op	 4224368 B/op	   21905 allocs/op
BenchmarkSequential/protobuf_mapping/2048-8               	     200	   8991229 ns/op	 8394096 B/op	   43409 allocs/op
BenchmarkSequential/protobuf_mapping/4096-8               	     100	  17607544 ns/op	16749936 B/op	   86417 allocs/op
BenchmarkSequential/protobuf_mapping/8192-8               	      50	  35201875 ns/op	33461616 B/op	  172433 allocs/op
BenchmarkSequential/protobuf_mapping/16384-8              	      20	  69439466 ns/op	66884976 B/op	  344465 allocs/op
BenchmarkSequential/protobuf_mapping/32768-8              	      10	 136154325 ns/op	133731696 B/op	  688529 allocs/op
BenchmarkSequential/protobuf_mapping/65536-8              	       5	 260902220 ns/op	267425136 B/op	 1376657 allocs/op
```

![Time sequential](benchmark/time_seq.jpg?raw=true "Title")

![Bytes sequential](benchmark/bytes_seq.jpg?raw=true "Title")

![Allocs sequential](benchmark/allocs_seq.jpg?raw=true "Title")

- Parallel benchmarks

```
$ go test -test.benchmem -test.bench BenchmarkParallel ./benchmark
goos: darwin
goarch: amd64
pkg: github.com/travelgateX/presenters-benchmark/benchmark
BenchmarkParallel/artemis/1-8         	   20000	     70386 ns/op	   70277 B/op	     878 allocs/op
BenchmarkParallel/artemis/2-8         	   20000	     89397 ns/op	   87829 B/op	    1068 allocs/op
BenchmarkParallel/artemis/4-8         	   10000	    130900 ns/op	  103888 B/op	    1443 allocs/op
BenchmarkParallel/artemis/8-8         	   10000	    207386 ns/op	  146909 B/op	    2192 allocs/op
BenchmarkParallel/artemis/16-8        	    5000	    378244 ns/op	  236744 B/op	    3682 allocs/op
BenchmarkParallel/artemis/32-8        	    2000	    682818 ns/op	  416587 B/op	    6664 allocs/op
BenchmarkParallel/artemis/64-8        	    2000	   1122818 ns/op	  772087 B/op	   12623 allocs/op
BenchmarkParallel/artemis/128-8       	    1000	   1968907 ns/op	 1474307 B/op	   24530 allocs/op
BenchmarkParallel/artemis/256-8       	     500	   3675891 ns/op	 2878940 B/op	   48342 allocs/op
BenchmarkParallel/artemis/512-8       	     200	   7101953 ns/op	 5675854 B/op	   95962 allocs/op
BenchmarkParallel/artemis/1024-8      	     100	  13390984 ns/op	11302469 B/op	  191200 allocs/op
BenchmarkParallel/artemis/2048-8      	      50	  27038036 ns/op	22591273 B/op	  381672 allocs/op
BenchmarkParallel/artemis/4096-8      	      20	  54589934 ns/op	45414362 B/op	  762609 allocs/op
BenchmarkParallel/artemis/8192-8      	      10	 126453955 ns/op	91093676 B/op	 1524482 allocs/op
BenchmarkParallel/artemis/16384-8     	       5	 246640112 ns/op	182411390 B/op	 3048221 allocs/op
BenchmarkParallel/artemis/32768-8     	       1	1199253451 ns/op	365001024 B/op	 6096560 allocs/op
BenchmarkParallel/artemis/65536-8     	       1	2416222160 ns/op	729856192 B/op	12191414 allocs/op
BenchmarkParallel/gophers/1-8         	   10000	    201373 ns/op	  351536 B/op	    1919 allocs/op
BenchmarkParallel/gophers/2-8         	   10000	    240680 ns/op	  388871 B/op	    2609 allocs/op
BenchmarkParallel/gophers/4-8         	   10000	    293233 ns/op	  458405 B/op	    3989 allocs/op
BenchmarkParallel/gophers/8-8         	    5000	    413430 ns/op	  602259 B/op	    6747 allocs/op
BenchmarkParallel/gophers/16-8        	    2000	    698227 ns/op	  893713 B/op	   12264 allocs/op
BenchmarkParallel/gophers/32-8        	    1000	   1172611 ns/op	 1477512 B/op	   23295 allocs/op
BenchmarkParallel/gophers/64-8        	    1000	   2198791 ns/op	 2643641 B/op	   45357 allocs/op
BenchmarkParallel/gophers/128-8       	     300	   4196826 ns/op	 4985954 B/op	   89478 allocs/op
BenchmarkParallel/gophers/256-8       	     200	   8123606 ns/op	 9716952 B/op	  177724 allocs/op
BenchmarkParallel/gophers/512-8       	     100	  16434707 ns/op	19223887 B/op	  354211 allocs/op
BenchmarkParallel/gophers/1024-8      	      50	  33172405 ns/op	38752302 B/op	  707185 allocs/op
BenchmarkParallel/gophers/2048-8      	      20	  64198678 ns/op	77249337 B/op	 1413121 allocs/op
BenchmarkParallel/gophers/4096-8      	      10	 140865014 ns/op	154567304 B/op	 2825012 allocs/op
BenchmarkParallel/gophers/8192-8      	       5	 248721156 ns/op	308750718 B/op	 5648797 allocs/op
BenchmarkParallel/gophers/16384-8     	       1	1536392255 ns/op	617066840 B/op	11296182 allocs/op
BenchmarkParallel/gophers/32768-8     	       1	2989488798 ns/op	1233761736 B/op	22591217 allocs/op
BenchmarkParallel/gophers/65536-8     	       1	6246716704 ns/op	2467064400 B/op	45180862 allocs/op
BenchmarkParallel/gqlgen_mapping/1-8  	   20000	     68244 ns/op	   84391 B/op	    1288 allocs/op
BenchmarkParallel/gqlgen_mapping/2-8  	   10000	    107772 ns/op	  128258 B/op	    2110 allocs/op
BenchmarkParallel/gqlgen_mapping/4-8  	   10000	    185105 ns/op	  215446 B/op	    3753 allocs/op
BenchmarkParallel/gqlgen_mapping/8-8  	    5000	    327822 ns/op	  389847 B/op	    7038 allocs/op
BenchmarkParallel/gqlgen_mapping/16-8 	    3000	    641152 ns/op	  739116 B/op	   13608 allocs/op
BenchmarkParallel/gqlgen_mapping/32-8 	    1000	   1232870 ns/op	 1440100 B/op	   26749 allocs/op
BenchmarkParallel/gqlgen_mapping/64-8 	     500	   2459267 ns/op	 2851247 B/op	   53036 allocs/op
BenchmarkParallel/gqlgen_mapping/128-8         	     300	   5003659 ns/op	 5677580 B/op	  105616 allocs/op
BenchmarkParallel/gqlgen_mapping/256-8         	     100	  10459008 ns/op	11420846 B/op	  210914 allocs/op
BenchmarkParallel/gqlgen_mapping/512-8         	     100	  20829629 ns/op	23079177 B/op	  421357 allocs/op
BenchmarkParallel/gqlgen_mapping/1024-8        	      30	  44071865 ns/op	46203592 B/op	  842353 allocs/op
BenchmarkParallel/gqlgen_mapping/2048-8        	      10	 105367711 ns/op	92497748 B/op	 1687005 allocs/op
BenchmarkParallel/gqlgen_mapping/4096-8        	       5	 235787722 ns/op	187244280 B/op	 3378280 allocs/op
BenchmarkParallel/gqlgen_mapping/8192-8        	       3	 427560821 ns/op	374085285 B/op	 6754006 allocs/op
BenchmarkParallel/gqlgen_mapping/16384-8       	       1	1280604932 ns/op	744807480 B/op	13498937 allocs/op
BenchmarkParallel/gqlgen_mapping/32768-8       	       1	2627030182 ns/op	1490526272 B/op	26991444 allocs/op
BenchmarkParallel/gqlgen_mapping/65536-8       	       1	6048744523 ns/op	2978727680 B/op	53957665 allocs/op
BenchmarkParallel/gqlgen_service_models/1-8    	   20000	     70850 ns/op	   83372 B/op	    1280 allocs/op
BenchmarkParallel/gqlgen_service_models/2-8    	   10000	    106039 ns/op	  126168 B/op	    2096 allocs/op
BenchmarkParallel/gqlgen_service_models/4-8    	   10000	    182029 ns/op	  211264 B/op	    3727 allocs/op
BenchmarkParallel/gqlgen_service_models/8-8    	    5000	    330339 ns/op	  382841 B/op	    6988 allocs/op
BenchmarkParallel/gqlgen_service_models/16-8   	    3000	    645176 ns/op	  732255 B/op	   13509 allocs/op
BenchmarkParallel/gqlgen_service_models/32-8   	    1000	   1196134 ns/op	 1424130 B/op	   26553 allocs/op
BenchmarkParallel/gqlgen_service_models/64-8   	     500	   2358728 ns/op	 2808076 B/op	   52644 allocs/op
BenchmarkParallel/gqlgen_service_models/128-8  	     300	   4868120 ns/op	 5593885 B/op	  104843 allocs/op
BenchmarkParallel/gqlgen_service_models/256-8  	     100	  10203383 ns/op	11235655 B/op	  209284 allocs/op
BenchmarkParallel/gqlgen_service_models/512-8  	     100	  20617124 ns/op	22580767 B/op	  418160 allocs/op
BenchmarkParallel/gqlgen_service_models/1024-8 	      30	  42811219 ns/op	45426660 B/op	  837222 allocs/op
BenchmarkParallel/gqlgen_service_models/2048-8 	      20	  98542776 ns/op	91448081 B/op	 1674821 allocs/op
BenchmarkParallel/gqlgen_service_models/4096-8 	       5	 209931571 ns/op	183424731 B/op	 3351648 allocs/op
BenchmarkParallel/gqlgen_service_models/8192-8 	       3	 510723147 ns/op	366858346 B/op	 6703859 allocs/op
BenchmarkParallel/gqlgen_service_models/16384-8         	       1	1372331407 ns/op	733323904 B/op	13403924 allocs/op
BenchmarkParallel/gqlgen_service_models/32768-8         	       1	2584100391 ns/op	1464509520 B/op	26778822 allocs/op
BenchmarkParallel/gqlgen_service_models/65536-8         	       1	5669292956 ns/op	2931889648 B/op	53576509 allocs/op
BenchmarkParallel/rest_json_service_models/1-8          	   50000	     29925 ns/op	   40537 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/2-8          	   50000	     33255 ns/op	   43109 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/4-8          	   50000	     37025 ns/op	   47738 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/8-8          	   30000	     47425 ns/op	   56755 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/16-8         	   20000	     65807 ns/op	   79339 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/32-8         	   10000	    104302 ns/op	  112453 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/64-8         	   10000	    183411 ns/op	  178599 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/128-8        	    5000	    336118 ns/op	  311415 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/256-8        	    2000	    648848 ns/op	  581746 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/512-8        	    1000	   1283435 ns/op	 1124266 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/1024-8       	     500	   2537455 ns/op	 2240857 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/2048-8       	     300	   5025856 ns/op	 4572444 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/4096-8       	     100	  10353254 ns/op	10398467 B/op	     401 allocs/op
BenchmarkParallel/rest_json_service_models/8192-8       	      50	  21349504 ns/op	25405148 B/op	     403 allocs/op
BenchmarkParallel/rest_json_service_models/16384-8      	      30	  41505782 ns/op	56951565 B/op	     405 allocs/op
BenchmarkParallel/rest_json_service_models/32768-8      	      10	 107220832 ns/op	191278840 B/op	     417 allocs/op
BenchmarkParallel/rest_json_service_models/65536-8      	      10	 223721274 ns/op	382473668 B/op	     418 allocs/op
BenchmarkParallel/rest_json_mapping/1-8                 	   50000	     30681 ns/op	   41897 B/op	     408 allocs/op
BenchmarkParallel/rest_json_mapping/2-8                 	   50000	     34076 ns/op	   45844 B/op	     416 allocs/op
BenchmarkParallel/rest_json_mapping/4-8                 	   50000	     40157 ns/op	   53165 B/op	     432 allocs/op
BenchmarkParallel/rest_json_mapping/8-8                 	   30000	     51456 ns/op	   67445 B/op	     464 allocs/op
BenchmarkParallel/rest_json_mapping/16-8                	   20000	     74455 ns/op	  100327 B/op	     528 allocs/op
BenchmarkParallel/rest_json_mapping/32-8                	   10000	    118871 ns/op	  154914 B/op	     656 allocs/op
BenchmarkParallel/rest_json_mapping/64-8                	   10000	    211938 ns/op	  264009 B/op	     912 allocs/op
BenchmarkParallel/rest_json_mapping/128-8               	    5000	    397207 ns/op	  487081 B/op	    1424 allocs/op
BenchmarkParallel/rest_json_mapping/256-8               	    2000	    759542 ns/op	  930971 B/op	    2448 allocs/op
BenchmarkParallel/rest_json_mapping/512-8               	    1000	   1502450 ns/op	 1840141 B/op	    4496 allocs/op
BenchmarkParallel/rest_json_mapping/1024-8              	     500	   2959139 ns/op	 3783134 B/op	    8593 allocs/op
BenchmarkParallel/rest_json_mapping/2048-8              	     200	   6028971 ns/op	 8091937 B/op	   16786 allocs/op
BenchmarkParallel/rest_json_mapping/4096-8              	     100	  12151682 ns/op	18233032 B/op	   33172 allocs/op
BenchmarkParallel/rest_json_mapping/8192-8              	      50	  25547898 ns/op	46706481 B/op	   65946 allocs/op
BenchmarkParallel/rest_json_mapping/16384-8             	      20	  54067792 ns/op	97155160 B/op	  131484 allocs/op
BenchmarkParallel/rest_json_mapping/32768-8             	      10	 129895023 ns/op	232336338 B/op	  262563 allocs/op
BenchmarkParallel/rest_json_mapping/65536-8             	       5	 233772832 ns/op	525527630 B/op	  524713 allocs/op
BenchmarkParallel/protobuf_mapping/1-8                  	   50000	     29180 ns/op	   42391 B/op	     422 allocs/op
BenchmarkParallel/protobuf_mapping/2-8                  	   50000	     31191 ns/op	   46527 B/op	     443 allocs/op
BenchmarkParallel/protobuf_mapping/4-8                  	   50000	     33661 ns/op	   55311 B/op	     485 allocs/op
BenchmarkParallel/protobuf_mapping/8-8                  	   30000	     41441 ns/op	   71856 B/op	     569 allocs/op
BenchmarkParallel/protobuf_mapping/16-8                 	   30000	     51344 ns/op	  104432 B/op	     737 allocs/op
BenchmarkParallel/protobuf_mapping/32-8                 	   20000	     74223 ns/op	  178800 B/op	    1073 allocs/op
BenchmarkParallel/protobuf_mapping/64-8                 	   10000	    121965 ns/op	  302960 B/op	    1745 allocs/op
BenchmarkParallel/protobuf_mapping/128-8                	   10000	    210124 ns/op	  567664 B/op	    3089 allocs/op
BenchmarkParallel/protobuf_mapping/256-8                	    3000	    403518 ns/op	 1097072 B/op	    5777 allocs/op
BenchmarkParallel/protobuf_mapping/512-8                	    2000	    769097 ns/op	 2139504 B/op	   11153 allocs/op
BenchmarkParallel/protobuf_mapping/1024-8               	    1000	   1547141 ns/op	 4224369 B/op	   21905 allocs/op
BenchmarkParallel/protobuf_mapping/2048-8               	     500	   3072806 ns/op	 8394097 B/op	   43409 allocs/op
BenchmarkParallel/protobuf_mapping/4096-8               	     200	   6121716 ns/op	16749940 B/op	   86417 allocs/op
BenchmarkParallel/protobuf_mapping/8192-8               	     100	  12759953 ns/op	33461619 B/op	  172433 allocs/op
BenchmarkParallel/protobuf_mapping/16384-8              	      50	  26760005 ns/op	66884984 B/op	  344465 allocs/op
BenchmarkParallel/protobuf_mapping/32768-8              	      20	  57862421 ns/op	133731712 B/op	  688529 allocs/op
BenchmarkParallel/protobuf_mapping/65536-8              	      10	 139457404 ns/op	267425176 B/op	 1376658 allocs/op
```

![Time Parallel](benchmark/time_par.jpg?raw=true "Title")

## Run the profiler

```bash
$ go test -benchmem -cpuprofile benchmark/restsm65536.pprof -test.benchtime 10s -run=^$ ./benchmark -bench BenchmarkCandidate_rest_servicemodels_65536_high
$ go tool pprof -http=:6060 benchmark.test benchmark/restsm65536.pprof
```

- Rest service models

![Rest service models](benchmark/restsm65536.png?raw=true "Title")

- Artemis

![Artemis](benchmark/artemis65536.png?raw=true "Title")

- Gophers

![Gophers](benchmark/gophers65536.png?raw=true "Title")

- Gqlgen

![Gqlgen](benchmark/gqlgen65536.png?raw=true "Title")
