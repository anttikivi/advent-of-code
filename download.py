#!/usr/bin/env python3

import argparse
import os
from os import path
import requests


def download_input(year, day, session):
    print(f"Downloading input for day {day} of {year}")
    year_dir = path.join(path.dirname(__file__), str(year))
    if not path.exists(year_dir):
        print(f"Creating directory for {year}")
        os.mkdir(year_dir)
    day_str = str(day).zfill(2)
    day_dir = path.join(year_dir, "day" + day_str)
    if not path.exists(day_dir):
        print(f"Creating directory for day {day}")
        os.mkdir(day_dir)

    # TODO: Think about adding an option for forcing re-download.
    input_file = path.join(day_dir, "input.txt")
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
        f.write(response.text)


def main():
    parser = argparse.ArgumentParser(
        prog="Advent of Code Downloader",
        description="Download your Advent of Code input.",
    )
    parser.add_argument("-y", "--year", type=int, required=True)
    parser.add_argument("-d", "--day", type=int)
    parser.add_argument("-K", "--session", type=str, required=True)

    args = parser.parse_args()

    if not args.day:
        for day in range(1, 26):
            download_input(args.year, day, args.session)
    else:
        if args.day < 1 or args.day > 25:
            print("Day must be between 1 and 25")
            return
        download_input(args.year, args.day, args.session)


if __name__ == "__main__":
    main()
