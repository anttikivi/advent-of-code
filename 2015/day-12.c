// cJSON is licensed under MIT License. See: https://github.com/DaveGamble/cJSON
#include "utils/cJSON.h"
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define BUFFER_LENGTH 8

int get_sum(cJSON* json) {
  int sum = 0;

  if (json->type == cJSON_Object) {
    cJSON* child = json->child;
    while (child != NULL) {
      if (child->valuestring != NULL &&
          strcmp(child->valuestring, "red") == 0) {
        return 0;
      }
      child = child->next;
    }
  }

  cJSON* child = json->child;
  while (child != NULL) {
    if (child->type == cJSON_Object || child->type == cJSON_Array) {
      sum += get_sum(child);
    } else if (child->type == cJSON_Number) {
      sum += child->valueint;
    }
    child = child->next;
  }

  return sum;
}

int main(void) {
  printf("\n");
  printf("***     Advent of Code 2015      ***\n");
  printf("--- Day 12: JSAbacusFramework.io ---\n");
  printf("\n");

  FILE* fp = fopen("input/day-12.txt", "rb");

  if (!fp) {
    perror("File opening failed");
    return EXIT_FAILURE;
  }

  clock_t start = clock();

  int sum = 0;
  char buffer[BUFFER_LENGTH];
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

  printf("Part 1: the sum of the numbers in the file is %d\n", sum);

  clock_t end = clock();
  double elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 1 ran for %f seconds\n", elapsed);

  printf("\n");

  rewind(fp);

  start = clock();

  int ret = fseek(fp, 0, SEEK_END);
  if (ret) {
    perror("seeking the end of the file failed");
    return EXIT_FAILURE;
  }

  len = ftell(fp);
  if (len < -1) {
    perror("getting the file position indicator failed");
    return EXIT_FAILURE;
  }

  ret = fseek(fp, 0, SEEK_SET);
  if (ret) {
    perror("seeking the beginning of the file failed");
    return EXIT_FAILURE;
  }

  char* jsonbuf = (char*)malloc(len);
  if (!jsonbuf) {
    perror("allocating the JSON buffer failed");
    return EXIT_FAILURE;
  }

  ret = fread(jsonbuf, 1, len, fp);
  if (ret != len && ferror(fp)) {
    perror("reading the file into the buffer failed");
    return EXIT_FAILURE;
  }

  cJSON* json = cJSON_Parse(jsonbuf);
  if (json == NULL) {
    printf("error parsing JSON: %s\n", cJSON_GetErrorPtr());
    return EXIT_FAILURE;
  }

  sum = get_sum(json);

  cJSON_Delete(json);
  fclose(fp);

  printf("Part 2: the sum of the numbers is %d\n", sum);

  end = clock();
  elapsed = (double)(end - start) / CLOCKS_PER_SEC;

  printf("Part 2 ran for %f seconds\n", elapsed);

  return EXIT_SUCCESS;
}
