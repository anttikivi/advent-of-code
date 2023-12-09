#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT "***REMOVED***"
#define PART_1_RUNS 40
#define PART_2_RUNS 50

#define PART_1_MAX_LENGTH 900000

int part_1_stack_runs(void) {
  char result[PART_1_MAX_LENGTH] = INPUT;
  char next[PART_1_MAX_LENGTH];

  for (int i = 0; i < PART_1_RUNS; i++) {
    int length = strlen(result);
    int next_length = 0;
    int count = 1;

    for (int j = 0; j < length; j++) {
      if (result[j] == result[j + 1]) {
        count++;
      } else {
        next_length += sprintf(&next[next_length], "%d%c", count, result[j]);
        count = 1;
      }
    }

    strcpy(result, next);
  }

  return strlen(result);
}

int part_2_heap_runs(void) {
  char input[] = INPUT;
  size_t length = strlen(input);
  size_t next_max_length = length * 2;
  char* result = malloc(next_max_length * sizeof(char));
  char* next = malloc(next_max_length * sizeof(char));

  strcpy(result, input);

  for (int i = 0; i < PART_2_RUNS; i++) {
    length = strlen(result);
    size_t next_length = 0;
    int count = 1;

    for (size_t j = 0; j < length; j++) {
      if (result[j] == result[j + 1]) {
        count++;
      } else {
        int needed = snprintf(NULL, 0, "%d%c", count, result[j]);

        if (next_length + needed >= next_max_length) {
          next_max_length = 2 * (next_max_length + needed);
          next = realloc(next, next_max_length * sizeof(char));
        }

        next_length += sprintf(&next[next_length], "%d%c", count, result[j]);
        count = 1;
      }
    }

    strcpy(result, next);
  }

  int result_length = strlen(result);

  free(next);
  free(result);

  return result_length;
}

int main(void) {
  printf("Advent of Code 2015, Day 10\n");

  printf("Part 1: length of result: %d\n", part_1_stack_runs());
  printf("Part 2: length of result: %d\n", part_2_heap_runs());
  return EXIT_SUCCESS;
}
