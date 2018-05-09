# 1rm

`1rm` is a very small command-line tool for estimating a One-Rep Max (1RM) based on weight and reps lifted.

The formula used is:

```
estimate = (weight * reps * 0.033) + weight
```

## Install

```
$ go get -u github.com/KyleBanks/1rm
```

## Usage

```
$ 1rm -reps R -weight W
```

-**reps** *int*: The number of reps completed
-**weight** *float*: The weight lifted, in any unit (kilos, pounds, stones, vibranium shields, etc.)

## Example

```
$ 1rm -reps 6 -weight 425
509.15
```

## License

`1rm` is available under the [MIT License](./LICENSE.md).
