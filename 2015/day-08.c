#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  printf("Advent of Code 2015, Day 8\n");

  FILE* fp = fopen("input/day-08.txt", "r");

  if (fp == NULL) {
    perror("error opening file");
    return EXIT_FAILURE;
  }

  char line[BUFSIZ];
  int literal_count = 0;
  int memory_count = 0;

  while (fgets(line, sizeof(line) - 1, fp) != NULL) {
    int len = strlen(line);

    // From now on, I'll modify the variable containing the length of the line
    // instead of calling strlen() again as it's expensive.

    line[len - 1] = '\0';
    len--;

    literal_count += len;

    char* p = line;
    p++;
    len--;
    p[len - 1] = '\0';
    len--;

    for (int i = 0; i < len; i++) {
      if (p[i] == '\\' && i + 1 < len) {
        if (p[i + 1] == '\\' || p[i + 1] == '"') {
          i++;
        } else if (p[i + 1] == 'x') {
          i += 3;
        }
      }

      memory_count++;
    }
  }

  int diff = literal_count - memory_count;

  printf("Part 1: the difference between the characters in the string literals "
         "and in the memory values is %d\n",
         diff);

  rewind(fp);

  int encoded_count = 0;

  while (fgets(line, sizeof(line) - 1, fp) != NULL) {
    int len = strlen(line);

    line[len - 1] = '\0';
    len--;

    // Add the additional quotes.
    encoded_count += 2;

    for (int i = 0; i < len; i++) {
      if (line[i] == '\\' || line[i] == '"') {
        encoded_count++;
      }

      encoded_count++;
    }
  }

  fclose(fp);

  diff = encoded_count - literal_count;

  printf("Part 2: the difference between the characters in the encoded strings "
         "and in the string literals is %d\n",
         diff);

  return EXIT_SUCCESS;
}
