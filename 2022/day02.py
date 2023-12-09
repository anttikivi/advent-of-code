def _main() -> None:
    print("Advent of Code 2022, Day 2")

    sum = 0
    # Rock, Paper, Scissors
    points = {"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

    with open("input/day02.txt") as f:
        lines = [line for line in f.read().strip().split("\n")]
        for line in lines:
            opponent, me = line.split()
            sum += points[me]
            if points[opponent] == points[me]:
                sum += 3
            elif points[me] - 1 == points[opponent] or (
                me == "X" and opponent == "C"
            ):
                sum += 6
    print("Part 1: the total number of points scored is", sum)

    sum = 0

    with open("input/day02.txt") as f:
        lines = [line for line in f.read().strip().split("\n")]
        for line in lines:
            opponent, result = line.split()
            if result == "Y":
                sum += points[opponent] + 3
            elif result == "X":
                sum += points[
                    "C" if opponent == "A" else chr(ord(opponent) - 1)
                ]
            else:
                sum += (
                    points["A" if opponent == "C" else chr(ord(opponent) + 1)]
                    + 6
                )
    print("Part 2: the total number of points scored is", sum)


if __name__ == "__main__":
    _main()
