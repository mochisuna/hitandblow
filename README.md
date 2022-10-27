# hit-and-blow
I really wanted to win `"Hit and Blow"` in `Clubhouse Games: 51` ...

# how to use
Input 3 variables below and you can get a combination of possibilities.

- Selected balls as 4-digit numbers
  - 0: Blue
  - 1: Red
  - 2: Green
  - 3: Yellow
  - 4: Purple
  - 5: Gray
- Number of "Hit"
- Number of "Blow"

```sh
$ go run main.go
number hit blow
<Selected balls as 4-digit numbers> <Number of "Hit"> <Number of "Blow">
```

## example

`Answer: 2401`

```sh
number hit blow
0012 0 3
--- LIST ---
1100
1120
1201
...
5120
5200
5201
------------
number hit blow
1300 1 1
--- LIST ---
2101
2104
2105
...
4201
5120
5201
------------
number hit blow
2100 2 1
--- LIST ---
2201
2401
2501
4120
5120
------------
number hit blow
2501 3 0
--- LIST ---
2201
2401
------------
number hit blow
2401 4 0
Answer is 2401
```
