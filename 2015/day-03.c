#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Point {
  int x;
  int y;
};

int main(void) {
  FILE* fp = fopen("input/day-03.txt", "r");

  if (fp == NULL) {
    perror("Error opening file");
    return -1;
  }

  char* directions = (char*)malloc(8192 * sizeof(char));

  fgets(directions, 8193, fp);
  fclose(fp);

  // This size is arbitrary, but it should be large enough to hold all the
  // visited points.
  struct Point* visited = (struct Point*)malloc(8192 * sizeof(struct Point));
  struct Point current = {0, 0};

  int num_visited = 0;

  for (int i = 0; i < 8192; i++) {
    char direction = directions[i];

    if (direction == '\0') {
      break;
    }

    if (i == 0) {
      visited[i].x = 0;
      visited[i].y = 0;
      num_visited++;
    }

    if (direction == '^') {
      current.y++;
    } else if (direction == 'v') {
      current.y--;
    } else if (direction == '>') {
      current.x++;
    } else if (direction == '<') {
      current.x--;
    }

    bool already_visited = false;

    for (int j = 0; j < 8192; j++) {
      if (visited[j].x == current.x && visited[j].y == current.y) {
        already_visited = true;
        break;
      }
    }

    if (!already_visited) {
      visited[i].x = current.x;
      visited[i].y = current.y;
      num_visited++;
    }
  }

  printf("Part 1: Visited %d houses\n", num_visited);

  memset(visited, 0, 8192 * sizeof(struct Point));

  struct Point santa = {0, 0};
  struct Point robo = {0, 0};

  num_visited = 0;

  for (int i = 0; i < 8192; i += 2) {
    char santa_direction = directions[i];
    char robo_direction = directions[i + 1];

    if (santa_direction == '\0') {
      break;
    }

    if (i == 0) {
      visited[i].x = 0;
      visited[i].y = 0;
      num_visited++;
    }

    if (santa_direction == '^') {
      santa.y++;
    } else if (santa_direction == 'v') {
      santa.y--;
    } else if (santa_direction == '>') {
      santa.x++;
    } else if (santa_direction == '<') {
      santa.x--;
    }

    bool already_visited = false;

    for (int j = 0; j < 8192; j++) {
      if (visited[j].x == santa.x && visited[j].y == santa.y) {
        already_visited = true;
        break;
      }
    }

    if (!already_visited) {
      visited[i].x = santa.x;
      visited[i].y = santa.y;
      num_visited++;
    }

    if (robo_direction == '\0') {
      break;
    }

    if (robo_direction == '^') {
      robo.y++;
    } else if (robo_direction == 'v') {
      robo.y--;
    } else if (robo_direction == '>') {
      robo.x++;
    } else if (robo_direction == '<') {
      robo.x--;
    }

    already_visited = false;

    for (int j = 0; j < 8192; j++) {
      if (visited[j].x == robo.x && visited[j].y == robo.y) {
        already_visited = true;
        break;
      }
    }

    if (!already_visited) {
      visited[i + 1].x = robo.x;
      visited[i + 1].y = robo.y;
      num_visited++;
    }
  }

  printf("Part 2: Visited %d houses\n", num_visited);

  free(visited);
  free(directions);
}
