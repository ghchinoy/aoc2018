# polarities

`alchemicalreduction` by itself, is the solution to 05.01; and with the `--shortest`, it's the answer to 05.02.

It takes input either via stdin or a `--file` file path.
Additionally, there's a `--debug` to show the logs.


for example, the solution to 05.01:

```
$ time ./alchemicalreduction --file ../data/input.txt 
final 10978 units

real    0m0.022s
user    0m0.014s
sys     0m0.012s
```

or, the solution to 05.02:

```
$ time ./alchemicalreduction --shortest --file ../data/input.txt 
Shortest reacted polymer length 4840 produced when removing 's'

real    0m0.146s
user    0m0.150s
sys     0m0.024s
```

or, the solution to the simple example for 05.01:

```
$ time ./alchemicalreduction --debug dabAcCaCBAcCcaDA
2018/12/09 01:55:01 16 units, starting with prevunit 'd'
2018/12/09 01:55:01 * (d a) > b
2018/12/09 01:55:01 a 2
2018/12/09 01:55:01 * (a b) > A
2018/12/09 01:55:01 b 3
2018/12/09 01:55:01 * (b A) > c
2018/12/09 01:55:01 A 4
2018/12/09 01:55:01 * (A c) > C
2018/12/09 01:55:01 c 5
2018/12/09 01:55:01 * (c C) > a
2018/12/09 01:55:01 - [c C] 4
2018/12/09 01:55:01 A 4
2018/12/09 01:55:01 * (A a) > C
2018/12/09 01:55:01 - [A a] 3
2018/12/09 01:55:01 b 3
2018/12/09 01:55:01 * (b C) > B
2018/12/09 01:55:01 C 4
2018/12/09 01:55:01 * (C B) > A
2018/12/09 01:55:01 B 5
2018/12/09 01:55:01 * (B A) > c
2018/12/09 01:55:01 A 6
2018/12/09 01:55:01 * (A c) > C
2018/12/09 01:55:01 c 7
2018/12/09 01:55:01 * (c C) > c
2018/12/09 01:55:01 - [c C] 6
2018/12/09 01:55:01 A 6
2018/12/09 01:55:01 * (A c) > a
2018/12/09 01:55:01 c 7
2018/12/09 01:55:01 * (c a) > D
2018/12/09 01:55:01 a 8
2018/12/09 01:55:01 * (a D) > A
2018/12/09 01:55:01 D 9
2018/12/09 01:55:01 * (D A) > 
2018/12/09 01:55:01 A 10
final 10 units

real    0m0.027s
user    0m0.002s
sys     0m0.027s
```


