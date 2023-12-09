#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct box {
  int length;
  int width;
  int height;
};

int min(int a, int b) {
  if (a < b) {
    return a;
  }

  return b;
}

int max(int a, int b) {
  if (a > b) {
    return a;
  }

  return b;
}

int main(void) {
  FILE* fp = fopen("input/day-02.txt", "r");

  if (fp == NULL) {
    perror("Error opening file");
    return (-1);
  }

  char* line = (char*)malloc(8 * sizeof(char));
  char* token;

  struct box boxes[1000];

  int count = 0;

  while (fgets(line, sizeof(line) + sizeof(char), fp) != NULL) {
    // There are empty lines for some reason.
    if (line[0] == '\r' || line[0] == '\n') {
      continue;
    }

    line[strcspn(line, "\n")] = 0;
    token = strtok(line, "x");

    int i = 0;

    while (token != NULL) {
      if (i == 0) {
        boxes[count].length = strtol(token, NULL, 10);
      }

      if (i == 1) {
        boxes[count].width = strtol(token, NULL, 10);
      }

      if (i == 2) {
        boxes[count].height = strtol(token, NULL, 10);
      }

      ++i;

      token = strtok(NULL, "x");
    }

    ++count;
  }

  free(token);
  free(line);

  fclose(fp);

  int total = 0;

  for (int i = 0; i < count; i++) {
    int side1 = boxes[i].length * boxes[i].width;
    int side2 = boxes[i].width * boxes[i].height;
    int side3 = boxes[i].height * boxes[i].length;

    int smallest_side = side1;

    if (side2 < smallest_side) {
      smallest_side = side2;
    }

    if (side3 < smallest_side) {
      smallest_side = side3;
    }

    total += (2 * side1) + (2 * side2) + (2 * side3) + smallest_side;
  }

  printf("Part 1: The elves need total of %d square feet of wrapping paper.\n",
         total);

  total = 0;

  for (int i = 0; i < count; i++) {
    int smallest = min(boxes[i].length, min(boxes[i].width, boxes[i].height));
    int second =
        max(min(boxes[i].length, boxes[i].width),
            min(max(boxes[i].length, boxes[i].width), boxes[i].height));

    total += (2 * smallest) + (2 * second) +
             (boxes[i].length * boxes[i].width * boxes[i].height);
  }

  printf("Part 2: The elves need total of %d feet of ribbon.\n", total);

  return 0;
}
