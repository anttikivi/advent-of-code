#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

// The header file has the input defined as a string literal.
#include "input/input-11.h"

bool is_forbidden(char c) {
  return c == 'i' || c == 'o' || c == 'l';
}

bool is_valid(char* pwd) {
  for (int i = 0; i < LENGTH; i++) {
    if (is_forbidden(pwd[i])) {
      return false;
    }
  }

  bool increasingFound = false;

  for (int i = 0; i < LENGTH - 2; i++) {
    if (pwd[i] == pwd[i + 1] - 1 && pwd[i] == pwd[i + 2] - 2) {
      increasingFound = true;
      break;
    }
  }

  bool pairsFound = false;
  char firstPair = 0;

  for (int i = 0; i < LENGTH - 1; i++) {
    if (pwd[i] == pwd[i + 1]) {
      if (firstPair == 0) {
        firstPair = pwd[i];
      } else if (pwd[i] != firstPair) {
        pairsFound = true;
        break;
      }
    }
  }

  return increasingFound && pairsFound;
}

void increment(char* pwd) {
  int j = -1;

  for (int i = 0; i < LENGTH; i++) {
    if (is_forbidden(pwd[i])) {
      j = i;
      break;
    }
  }

  if (j >= 0) {
    pwd[j]++;
    for (int i = j + 1; i < LENGTH; i++) {
      pwd[i] = 'a';
    }
  }

  for (int i = LENGTH - 1; i >= 0; i--) {
    if (pwd[i] == 'z') {
      pwd[i] = 'a';
    } else {
      pwd[i]++;
      break;
    }
  }
}

int main(void) {
  printf("***   Advent of Code 2015    ***\n");
  printf("--- Day 11: Corporate Policy ---\n");

  clock_t start = clock();

  char pwd[] = INPUT;

  do {
    increment(pwd);
  } while (!is_valid(pwd));

  printf("Part 1: Santa's next password is %s\n", pwd);

  clock_t end = clock();
  double elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 1 ran for %f seconds\n", elapsed);

  start = clock();

  // The length of the array is enough.
  strcpy(pwd, INPUT);

  for (int i = 0; i < 2; i++) {
    do {
      increment(pwd);
    } while (!is_valid(pwd));
  }

  printf("Part 2: Santa's next password is %s\n", pwd);

  end = clock();
  elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 2 ran for %f seconds\n", elapsed);

  return EXIT_SUCCESS;
}
