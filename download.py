#!/usr/bin/env python3

import argparse
import os
from os import path
from typing import cast

import requests


def download_input(year: int, day: int, separate: bool, session: str):
    print(f"Downloading input for day {day} of {year}")
    year_dir = path.join(path.dirname(__file__), str(year))
    if not path.exists(year_dir):
        print(f"Creating directory for {year}")
        os.mkdir(year_dir)
    day_str = "day-" + str(day).zfill(2)
    dl_dir = (
        path.join(year_dir, day_str)
        if separate
        else path.join(year_dir, "input")
    )

    if not path.exists(dl_dir):
        print(f"Creating directory for the input file for day {day}")
        os.mkdir(dl_dir)

    # TODO: Think about adding an option for forcing re-download.
    input_file = (
        path.join(dl_dir, "input.txt")
        if separate
        else path.join(dl_dir, day_str + ".txt")
    )
    if path.exists(input_file):
        print(f"Input for day {day} already exists")
        return

    response = requests.get(
        f"https://adventofcode.com/{year}/day/{day}/input",
        headers={"Cookie": f"session={session}"},
    )
    if not response.status_code == 200:
        print(f"Error downloading input for day {day}")
        return
    with open(input_file, "w") as f:
        _ = f.write(response.text)


def main():
    parser = argparse.ArgumentParser(
        prog="Advent of Code Downloader",
        description="Download your Advent of Code input.",
    )
    _ = parser.add_argument("-y", "--year", type=int, required=True)
    _ = parser.add_argument("-d", "--day", type=int)
    _ = parser.add_argument("-s", "--separate", action="store_true")
    _ = parser.add_argument("-K", "--session", type=str, required=True)

    args = parser.parse_args()

    args.day = cast(int, args.day)
    args.year = cast(int, args.year)
    args.separate = cast(bool, args.separate)
    args.session = cast(str, args.session)

    if not args.day:
        for day in range(1, 26):
            download_input(
                args.year,
                day,
                args.separate,
                args.session,
            )
    else:
        if args.day < 1 or args.day > 25:
            print("Day must be between 1 and 25")
            return
        download_input(
            args.year,
            args.day,
            args.separate,
            args.session,
        )


if __name__ == "__main__":
    main()
