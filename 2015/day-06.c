#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int main(void) {
  printf("Day 06\n");

  FILE* fp = fopen("input/day-06.txt", "r");

  if (fp == NULL) {
    perror("error opening file");
    return EXIT_FAILURE;
  }

  char line[BUFSIZ];
  int count = 0;
  bool lights[1000 * 1000] = {false};

  while (fgets(line, sizeof(line), fp) != NULL) {
    int x1, y1, x2, y2;

    if (sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          lights[y * 1000 + x] = true;
        }
      }
    }

    if (sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          lights[y * 1000 + x] = false;
        }
      }
    }

    if (sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          if (lights[y * 1000 + x]) {
            lights[y * 1000 + x] = false;
          } else {
            lights[y * 1000 + x] = true;
          }
        }
      }
    }
  }

  for (int i = 0; i < 1000 * 1000; i++) {
    if (lights[i]) {
      count++;
    }
  }

  printf("Part 1: there are %d lights on\n", count);

  rewind(fp);

  count = 0;
  int brightness[1000 * 1000] = {0};

  while (fgets(line, sizeof(line), fp) != NULL) {
    int x1, y1, x2, y2;

    if (sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          brightness[y * 1000 + x]++;
        }
      }
    }

    if (sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          if (brightness[y * 1000 + x] > 0) {
            brightness[y * 1000 + x]--;
          }
        }
      }
    }

    if (sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)) {
      for (int y = y1; y <= y2; y++) {
        for (int x = x1; x <= x2; x++) {
          brightness[y * 1000 + x] += 2;
        }
      }
    }
  }

  for (int i = 0; i < 1000 * 1000; i++) {
    count += brightness[i];
  }

  printf("Part 2: total brightness is %d\n", count);

  fclose(fp);

  return EXIT_SUCCESS;
}
