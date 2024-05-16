#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int main(void) {
  printf("***   Advent of Code 2015    ***\n");
  printf("--- Day 11: Corporate Policy ---\n");

  FILE* fp = fopen("input/day-12.txt", "r");

  if (!fp) {
    perror("File opening failed");
    return EXIT_FAILURE;
  }

  clock_t start = clock();

  int sum = 0;
  char buffer[8];
  int len = 0;
  char c;

  do {
    c = fgetc(fp);

    if ((c >= '0' && c <= '9') || c == '-') {
      buffer[len++] = c;
    } else if (strlen(buffer) > 0) {
      buffer[len] = '\0';
      sum += atoi(buffer);
      memset(buffer, 0, sizeof(buffer));
      len = 0;
    }
  } while (c != EOF);

  fclose(fp);

  printf("Part 1: the sum of the numbers in the file is %d\n", sum);

  clock_t end = clock();
  double elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 1 ran for %f seconds\n", elapsed);

  start = clock();

  printf("Part 2: Santa's next password is %s\n", "");

  end = clock();
  elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 2 ran for %f seconds\n", elapsed);

  return EXIT_SUCCESS;
}
