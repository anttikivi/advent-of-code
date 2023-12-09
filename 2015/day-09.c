#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_PAIRS 40
#define MAX_NAME_LENGTH 30
#define MAX_LOCATIONS 10

struct pair {
  char from[MAX_NAME_LENGTH];
  char to[MAX_NAME_LENGTH];
  unsigned int distance;
};

bool in_array(char locations[][MAX_NAME_LENGTH], int length, char* location) {
  for (int i = 0; i < length; i++) {
    if (strcmp(locations[i], location) == 0) {
      return true;
    }
  }

  return false;
}

struct pair* find_pair(struct pair pairs[MAX_PAIRS], int length, char* from,
                       char* to) {
  for (int i = 0; i < length; i++) {
    if ((strcmp(pairs[i].from, from) == 0 && strcmp(pairs[i].to, to) == 0) ||
        (strcmp(pairs[i].from, to) == 0 && strcmp(pairs[i].to, from) == 0)) {
      return &pairs[i];
    }
  }

  return NULL;
}

int compare_locations(const void* a, const void* b) {
  return strcmp(a, b);
}

void swap(char* a, char* b) {
  char tmp[MAX_NAME_LENGTH];
  strcpy(tmp, a);
  strcpy(a, b);
  strcpy(b, tmp);
}

void reverse(char array[][MAX_NAME_LENGTH], int start, int end) {
  while (start < end) {
    swap(array[start], array[end]);
    start++;
    end--;
  }
}

// See:
// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
bool next_permutation(char array[][MAX_NAME_LENGTH], int length) {
  int k = length - 2;

  while (k >= 0 && strcmp(array[k], array[k + 1]) >= 0) {
    k--;
  }

  if (k < 0) {
    return false;
  }

  int l = length - 1;

  while (l > k && strcmp(array[l], array[k]) <= 0) {
    l--;
  }

  swap(array[k], array[l]);
  reverse(array, k + 1, length - 1);

  return true;
}

int main(void) {
  printf("Advent of Code 2015, Day 9\n");

  FILE* fp = fopen("input/day-09.txt", "r");
  char line[BUFSIZ];

  if (fp == NULL) {
    perror("error opening file");
    return EXIT_FAILURE;
  }

  struct pair pairs[MAX_PAIRS];
  int pair_count = 0;

  while (fgets(line, sizeof(line), fp)) {
    struct pair pair = {0};

    sscanf(line, "%s to %s = %u", pair.from, pair.to, &pair.distance);
    pairs[pair_count] = pair;
    pair_count++;
  }

  fclose(fp);

  char locations[MAX_PAIRS][MAX_NAME_LENGTH];
  int location_count = 0;

  for (int i = 0; i < pair_count; i++) {
    if (!in_array(locations, location_count, pairs[i].from)) {
      strcpy(locations[location_count], pairs[i].from);
      location_count++;
    }

    if (!in_array(locations, location_count, pairs[i].to)) {
      strcpy(locations[location_count], pairs[i].to);
      location_count++;
    }
  }

  qsort(locations, location_count, sizeof(char[MAX_NAME_LENGTH]),
        compare_locations);

  int min_distance = INT_MAX;
  int max_distance = 0;

  do {
    int distance = 0;

    for (int i = 0; i < location_count - 1; i++) {
      struct pair* pair =
          find_pair(pairs, pair_count, locations[i], locations[i + 1]);
      if (pair != NULL) {
        distance += pair->distance;
      }
    }

    if (distance < min_distance) {
      min_distance = distance;
    }

    if (distance > max_distance) {
      max_distance = distance;
    }
  } while (next_permutation(locations, location_count));

  printf("Part 1: the shortest distance is %d\n", min_distance);
  printf("Part 2: the longest distance is %d\n", max_distance);

  return EXIT_SUCCESS;
}
