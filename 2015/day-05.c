#include <stdio.h>
#include <stdlib.h>

int main(void) {
  FILE* fp = fopen("input/day-05.txt", "r");

  if (fp == NULL) {
    perror("Error opening file");
    return EXIT_FAILURE;
  }

  char line[BUFSIZ];
  int nice = 0;

  while (fgets(line, sizeof(line), fp) != NULL) {
    int vowels = 0;
    int double_letter = 0;
    int bad_strings = 0;

    for (int i = 0; line[i] != '\0'; i++) {
      if (line[i] == 'a' || line[i] == 'e' || line[i] == 'i' ||
          line[i] == 'o' || line[i] == 'u') {
        vowels++;
      }

      if (line[i] == line[i + 1]) {
        double_letter++;
      }

      if (line[i] == 'a' && line[i + 1] == 'b') {
        bad_strings++;
      }

      if (line[i] == 'c' && line[i + 1] == 'd') {
        bad_strings++;
      }

      if (line[i] == 'p' && line[i + 1] == 'q') {
        bad_strings++;
      }

      if (line[i] == 'x' && line[i + 1] == 'y') {
        bad_strings++;
      }
    }

    if (vowels >= 3 && double_letter >= 1 && bad_strings == 0) {
      nice++;
    }
  }

  printf("Part 1: nice strings: %d\n", nice);

  rewind(fp);

  nice = 0;

  while (fgets(line, sizeof(line), fp) != NULL) {
    int pair = 0;
    int repeat = 0;

    for (int i = 0; line[i] != '\0'; i++) {
      if (line[i] == line[i + 2]) {
        repeat++;
      }

      for (int j = i + 2; line[j] != '\0'; j++) {
        if (line[i] == line[j] && line[i + 1] == line[j + 1]) {
          pair++;
        }
      }
    }

    if (pair >= 1 && repeat >= 1) {
      nice++;
    }
  }

  printf("Part 2: nice strings: %d\n", nice);

  fclose(fp);

  return EXIT_SUCCESS;
}
