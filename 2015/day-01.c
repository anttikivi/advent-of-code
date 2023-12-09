#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void) {
  int current_floor = 0;

  FILE* fp = fopen("input/day-01.txt", "r");

  if (fp == NULL) {
    perror("Error opening file");
    return (-1);
  }

  char* directions = (char*)malloc(7000 * sizeof(char));

  fgets(directions, 7001, fp);
  fclose(fp);

  for (unsigned long i = 0; i < strlen(directions); i++) {
    if (directions[i] == '(') {
      current_floor++;
    } else if (directions[i] == ')') {
      current_floor--;
    }
  }

  printf("Part 1: Santa is at floor %d\n", current_floor);

  // Reset the current floor for part 2.
  current_floor = 0;

  for (unsigned int i = 0; i < strlen(directions); i++) {
    if (directions[i] == '(') {
      current_floor++;
    } else if (directions[i] == ')') {
      current_floor--;
    }

    if (current_floor == -1) {
      printf("Part 2: Santa entered the basement at position %d\n", i + 1);
      break;
    }
  }

  free(directions);

  return 0;
}
