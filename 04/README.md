# 04 Guard Analysis

* there're anomolies in the data, the time being of hour 00 or 23
  Convert the 23:xx's to 00:00
* Data is not in order
  Sorting by time
  
01: "Find the guard that has the most minutes asleep. What minute does that guard spend asleep the most?"

02: "Of all guards, which guard is most frequently asleep on the same minute?" (for all guards, which is asleep most on the same minute?)

# analysis

way overengineered :)

visualization was probably unnecessary and got me hung up for *days*

# Output


```
$ time ./guardanalysis --file ../data/input.txt
real    0m0.030s
user    0m0.015s
sys     0m0.015s
```

```
 Date    ID Minute
            000000000011111111112222222222333333333344444444445555555555
            012345678901234567890123456789012345678901234567890123456789
02-05 #3109 ..................................................####...#..
02-06 #2411 .......####################################.................
02-07 #3109 ...###########################...........#######............
02-08 #2411 ..##############............................................
02-09  #191 .....######..............############################.......
02-10 #2411 ...................####################################.....
02-11 #1291 ......######################################................
02-12 #2999 ........................#####...................######......
02-13 #2053 .....................................####################...
02-14  #137 ............................................................
02-15  #811 ...................############..................##########.
02-16  #239 ...................######################################...
02-17  #941 .................................#########...........######.
02-18 #1327 ....................................................#######.
02-19 #2459 ..............................################..............
02-20 #2381 ...........########################.........................
02-21  #941 ..................................................#######...
02-22  #941 ...............................#################............
02-23  #941 .......####################################.................
02-24 #1229 .........................##################..............##.
02-25 #1229 .....................................##########..........#..
02-26  #239 ..................................#########....########.....
02-27 #2897 .....................................#################......
02-28  #811 ..................................................#########.
03-01 #1229 ...................................#############............
03-02 #1009 .................###########.......#############............
03-03 #1229 ..........................########.....###############......
03-04  #191 ....................##################################......
03-05 #2137 ............###################.............................
03-06 #1327 ........................#############################.......
03-07 #1291 .########################...#######.....###########...####..
03-08 #2999 ...###################...##################################.
03-09 #2053 ..................................###############...........
03-10 #2381 .................................#######............#######.
03-11 #1009 ..........................................#############.....
03-12 #1327 ....#######################....................####.....###.
03-13 #1229 ........#####........#####..................................
03-14 #1009 .........################################.............####..
03-15 #1327 ...............................##################...........
03-16  #239 ...................####################...########..........
03-17  #811 ......#####################################################.
03-18 #2381 ................................................#########...
03-19  #239 ...............##########################.......#...........
03-20 #1291 .....................................################.......
03-21 #1381 ...................##################...........#####.......
03-22  #811 ........#..........########.................................
03-23  #811 ...............#..................................#######...
03-24 #2897 .............................######################.........
03-25 #2381 ............#############################################...
03-26 #3109 ..#########..................................####.....####..
03-27 #2411 ...........................########.....#########...........
03-28 #1381 .............................#########################......
03-29 #1069 .............................######################.........
03-30 #1327 ...............................###############..............
03-31 #3109 ....................#######################################.
04-01  #239 .............................######..............#########..
04-02 #2897 ...###################################......................
04-03 #2459 ...................................########################.
04-04 #1291 ..................###................##############.........
04-05 #2003 .............################################...............
04-06  #811 ............................................#####...........
04-07 #2999 ..........################..................................
04-08 #2999 .......###########......############################........
04-09  #283 ............................................................
04-10 #1381 .......#####................................................
04-11 #1069 .................................###########................
04-12 #2137 .............########################################.......
04-13 #3463 ...............................................####.........
04-14 #3463 .....................................############...........
04-15 #2411 ..........#######.........................################..
04-16  #191 ............................#........####.....#########.....
04-17 #2003 ................#######################################.....
04-18 #2459 ....###################################################.....
04-19 #1327 ........................############################........
04-20 #1009 .......................................................###..
04-21 #2137 .........#############.....................########.........
04-22  #941 ................######################################......
04-23 #3109 ...................#############......###########...........
04-24 #3109 .........#####################################..............
04-25 #1381 .....................................................#...#..
04-26 #2411 ............#############################################...
04-27 #2003 .....############################...........................
04-28 #1381 ..............##.............#######......##########........
04-29 #2999 .............#################################..............
04-30 #1069 .......................................######...............
05-01 #1229 ......########.............############################.....
05-02 #2999 .............###################............................
05-03  #191 .#############################.....############.............
05-04 #2003 .............####################################....#####..
05-05 #1327 ........................#####...............................
05-06 #2137 ###################.........................................
05-07  #239 ....................................###########.............
05-08 #1069 ............##..............................................
05-09 #2999 ......................#####...##........................#...
05-10 #2897 ...................##.......................................
05-11 #2411 .......................##################################...
05-12 #2999 .....................#####.......................########...
05-13  #137 ............................................................
05-14 #1291 .........##############...............###...................
05-15 #1009 .#####################################################......
05-16 #1009 .................##################.............#########...
05-17 #2999 ..............#...#############################.............
05-18 #1291 .........###################################.....#####......
05-19 #2003 ...........####......#..................#######.............
05-20 #1009 ...........##########################################.......
05-21  #941 ................................######################......
05-22 #3109 ............#############..................###..............
05-23 #1291 .........................##################################.
05-24 #3109 .....................................####........######.....
05-25 #2999 ..........................###...............................
05-26 #2053 .....................#####################################..
05-27 #2003 ........######################..............................
05-28 #2411 ............#####################........########...........
05-29 #1069 ...........................###############################..
05-30 #1229 ..............####################..........................
05-31 #1381 ........................................##..................
06-01 #1009 ..................###################################.......
06-02  #239 ............................................#########....#..
06-03  #239 ....................................................#######.
06-04 #3109 ..................##################........................
06-05  #811 ............................###.............................
06-06  #941 .......#############.....##################.........##......
06-07 #3109 ...........................#################................
06-08 #1069 ..........................############################......
06-09 #1381 ........................######################...#######....
06-10 #1229 .#############################################..............
06-11  #941 ......##....................##############..................
06-12 #2381 .............###############................................
06-13 #1069 ..................................############.....#........
06-14 #3109 ....##############################################..........
06-15 #3109 ............###############################################.
06-16 #2459 ................................###########################.
06-17  #811 ................................#########################...
06-18 #1229 .................###################################........
06-19 #1229 ...####............##################.......................
06-20 #2459 ...............######################......##############...
06-21  #283 ............................................................
06-22  #191 .....................##############################.........
06-23 #2411 .################################################...........
06-24  #191 .........############################...##############......
06-25 #2459 ................##........................#############.....
06-26 #1327 ...###########################..........#.....#.............
06-27 #2053 .....................####################......#........###.
06-28  #941 .............................#######################........
06-29 #1381 .........##############################.....#############...
06-30 #2003 .........................#########################..........
07-01 #2999 ...........#####################............................
07-02 #1229 .........######################################.............
07-03 #2381 ...........###############################..................
07-04  #191 .............########.......................................
07-05 #1069 ..##########.....#######################################....
07-06 #1009 ...#######################################..................
07-07 #2999 .........##################.................................
07-08 #3109 ...........................##############......########.....
07-09  #239 ...........................................#########........
07-10 #1069 ...##############....#######...........################.....
07-11 #2459 .......................................................#....
07-12 #1009 ..........#################..................############...
07-13 #2411 ...................................#######..................
07-14 #1229 ......................##############################........
07-15  #191 ......................##############........................
07-16 #2459 .....................................##################.....
07-17 #1291 ..........................................###...............
07-18 #1327 .........................................................#..
07-19 #1229 ...................................#######################..
07-20 #3463 .....................................###############........
07-21  #191 ......................#####.....................######......
07-22 #2411 ..##........###################....#####................#...
07-23 #2003 .........#############################################......
07-24  #191 ......########################.....#############...########.
07-25 #2999 ....................##############################..........
07-26 #1229 .....###########################............................
07-27 #1381 .......############.....................######..............
07-28  #239 ..................############..............................
07-29 #3463 ..........................#######################...####....
07-30  #811 ................####################...#########...#####....
07-31  #191 .#########################################..................
08-01 #2999 .......................##.......####........................
08-02 #2999 .......................#####.........................###....
08-03  #941 ......####################################################..
08-04  #239 .......................##################################...
08-05  #941 ..........................................###########.......
08-06 #2137 ....##......................................................
08-07 #2381 ..............#######################################....#..
08-08 #3463 ............................######...............#######....
08-09 #2003 .........#################################################..
08-10 #2411 ....#################################.....#################.
08-11  #137 ............................................................
08-12 #1009 ........#############....................#######............
08-13  #191 ..............###########...................................
08-14 #2459 ...................####################......##......#......
08-15 #1291 .....................########################...............
08-16 #2999 ....................#####...#############################...
08-17  #239 ................###########################.................
08-18 #1069 ...##################..............####################.....
08-19 #2897 ................................##..........................
08-20 #1009 .....................................................##.....
08-21  #239 ...#.................##............############..........##.
08-22 #1069 ...........#########################...................###..
08-23 #2411 .................##################################.........
08-24  #941 ...........########.........................................
08-25 #1381 .###########################################................
08-26 #2459 .........############...............######################..
08-27 #1381 ................########################################....
08-28 #2003 ....................####.....##############...#.............
08-29 #2003 ....................############............................
08-30 #1381 .............................................###########....
08-31  #239 ..#######################################################...
09-01 #1229 .......................................#################....
09-02 #1229 .......................................##################...
09-03 #3109 ..............................................#####...##....
09-04 #2897 .....................######################.......#######...
09-05 #2411 .................................................##....####.
09-06 #1229 ..................####################......................
09-07  #941 .................................#######################....
09-08 #3109 ...........................................#######..........
09-09  #811 ........................................................###.
09-10 #1381 .........................#############################......
09-11 #2003 ....................###..........##...........############..
09-12  #191 .................##################################.........
09-13 #1381 .................##########...###########...................
09-14 #1291 ..........................########################..........
09-15  #137 ............................................................
09-16 #1381 .........##############...................#############.....
09-17 #2137 ...........##########.......................................
09-18 #2411 ......................................####################..
09-19 #3109 ............................############################....
09-20 #2381 .........####............................######.............
09-21  #191 ......#############...#################.....................
09-22 #2999 .....................#########..............................
09-23 #2897 .....................................###.....#..............
09-24 #2411 .....####################################################...
09-25 #2999 ........###...........############..........................
09-26  #239 ....................................######################..
09-27 #2137 ........##########################........#######...........
09-28 #2999 .....................#########################..............
09-29 #2411 ..............#############################........##.......
09-30  #137 ............................................................
10-01  #239 ..............................................############..
10-02 #1069 ..############################################..........###.
10-03 #2003 ........###########################################.........
10-04 #1381 ...........##########################################.......
10-05 #2411 ..........######################............................
10-06 #2459 .....................###...#########.....#..................
10-07 #3463 .................................................###........
10-08 #1381 .............................................####...........
10-09 #3109 ....................######..........#######################.
10-10 #1291 ...............######################......################.
10-11 #1229 ..........................................############......
10-12 #1009 ................................#......................#....
10-13  #239 ..........................................##....##########..
10-14 #2459 ...........############################################.....
10-15 #1327 ....................................####################....
10-16 #1009 ........................................###################.
10-17  #191 ..............................######...###################..
10-18 #2003 ...................###############..........................
10-19 #2897 .....................................######################.
10-20 #2137 ................##########..................................
10-21 #1009 ........................######################..............
10-22 #1291 ....................########################................
10-23  #239 ...###############################..........................
10-24 #2459 ..............##########################################....
10-25 #1069 ......................#########################.............
10-26 #2381 ........######################################..............
10-27 #2053 .....................###########################............
10-28  #137 ............................................................
10-29 #3463 .............................######################.........
10-30 #2999 ......##..........#####################################.....
10-31 #3463 .....................#########################..............
11-01 #1327 ...............................#################............
11-02 #2137 ......##################....................................
11-03 #2459 ..................############################.........####.
11-04 #3463 ...........................................########.........
11-05 #1229 ...................................#####################....
11-06 #2003 ............................#.......########.............##.
11-07 #2053 .............##############################....#######......
11-08 #1327 ..................#############...########..............#...
11-09 #1291 ....................................########....#####....##.
11-10 #2003 ...........................###################..............
11-11  #811 ............###################.......#######...............
11-12  #191 ......###############################.......................
11-13 #1291 ..............................##############................
11-14 #1291 ......................#####...########........#############.
11-15  #811 .................................................########...
11-16 #1069 ...###########################..............................
11-17 #2003 ...................#######..................................
11-18 #2411 ...............................##############...............
11-19 #3109 ###############################################.............
11-20 #1069 .............................##################.............
11-21 #2053 ...................................###########..............
11-22 #2137 ......................##########....###....##########.......
11-23 #1229 ......................................###########.......###.
Guard #2411 slept 545m
Sleepiest minute for Guard 2411 is 42
Answer 04.01 (2411 x 42): 101262
Sleepiest Guard is #2999, and they slept 18m on minute 24
Answer 04.02 (2999 x 24): 71976
```
