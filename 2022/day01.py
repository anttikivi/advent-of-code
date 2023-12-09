def _main() -> None:
    print("Advent of Code 2022, Day 1")

    calories: list[list[int]] = [[]]

    with open("input/day01.txt") as f:
        for line in f:
            if line.rstrip("\n"):
                calories[-1].append(int(line))
            else:
                calories.append([])
    sums = [sum(c) for c in calories]
    print(
        "Part 1: the Elf carrying the most calories has", max(sums), "calories"
    )
    print(
        "Part 2: the three Elves carrying the most calories have a total of",
        sum(sorted(sums, reverse=True)[:3]),
        "calories",
    )


if __name__ == "__main__":
    _main()
