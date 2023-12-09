#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_INSTRUCTIONS 1000
#define MAX_WIRES 1000
#define OP_SIZE 7

struct instruction {
  char op[OP_SIZE];
  char left[6];
  char right[6];
  char result[6];
};

struct wire {
  char name[3];
  unsigned short value;
  bool calculated;
};

struct instruction instructions[MAX_INSTRUCTIONS];
int instruction_count = 0;

struct wire wire_map[MAX_WIRES];

int hash_name(const char* name) {
  int hash = 0;

  while (*name) {
    hash = (hash * 26 + (*name - 'a')) % MAX_WIRES;
    name++;
  }

  return hash;
}

struct wire* get_wire(const char* name) {
  int hash = hash_name(name);

  while (wire_map[hash].name[0] != '\0' &&
         strcmp(wire_map[hash].name, name) != 0) {
    hash = (hash + 1) % MAX_WIRES;
  }

  if (!wire_map[hash].name[0]) {
    strcpy(wire_map[hash].name, name);
    wire_map[hash].value = 0;
    wire_map[hash].calculated = false;
  }

  return &wire_map[hash];
}

unsigned short get_value(char* name) {
  // Because this function is called recursively with the operands, I start by
  // checking if the operand is a number.
  if (name[0] >= '0' && name[0] <= '9') {
    return atoi(name);
  }

  struct wire* wire = get_wire(name);

  if (!wire->calculated) {
    for (int i = 0; i < instruction_count; i++) {
      if (strcmp(instructions[i].result, name) == 0) {
        unsigned short left = 0;
        unsigned short right = 0;

        if (instructions[i].left[0]) {
          left = get_value(instructions[i].left);
        }

        if (instructions[i].right[0]) {
          right = get_value(instructions[i].right);
        }

        if (strcmp(instructions[i].op, "AND") == 0) {
          wire->value = left & right;
        } else if (strcmp(instructions[i].op, "OR") == 0) {
          wire->value = left | right;
        } else if (strcmp(instructions[i].op, "LSHIFT") == 0) {
          wire->value = left << right;
        } else if (strcmp(instructions[i].op, "RSHIFT") == 0) {
          wire->value = left >> right;
        } else if (strcmp(instructions[i].op, "NOT") == 0) {
          wire->value = ~right;
        } else {
          wire->value = left;
        }
      }
    }

    wire->calculated = true;
  }

  return wire->value;
}

int main(void) {
  printf("Day 07\n");

  FILE* fp = fopen("input/day-07.txt", "r");
  char line[100];

  if (fp == NULL) {
    perror("error opening file");
    return EXIT_FAILURE;
  }

  while (fgets(line, sizeof(line), fp)) {
    struct instruction instruction = {0};

    if (strstr(line, "AND") != NULL) {
      sscanf(line, "%s AND %s -> %s", instruction.left, instruction.right,
             instruction.result);
      strncpy(instruction.op, "AND", OP_SIZE - 1);
    } else if (strstr(line, "OR") != NULL) {
      sscanf(line, "%s OR %s -> %s", instruction.left, instruction.right,
             instruction.result);
      strncpy(instruction.op, "OR", OP_SIZE - 1);
    } else if (strstr(line, "LSHIFT") != NULL) {
      sscanf(line, "%s LSHIFT %s -> %s", instruction.left, instruction.right,
             instruction.result);
      strncpy(instruction.op, "LSHIFT", OP_SIZE - 1);
    } else if (strstr(line, "RSHIFT") != NULL) {
      sscanf(line, "%s RSHIFT %s -> %s", instruction.left, instruction.right,
             instruction.result);
      strncpy(instruction.op, "RSHIFT", OP_SIZE - 1);
    } else if (strstr(line, "NOT") != NULL) {
      sscanf(line, "NOT %s -> %s", instruction.right, instruction.result);
      strncpy(instruction.op, "NOT", OP_SIZE - 1);
    } else {
      sscanf(line, "%s -> %s", instruction.left, instruction.result);
      strncpy(instruction.op, "", OP_SIZE - 1);
    }

    instructions[instruction_count++] = instruction;
  }

  fclose(fp);

  unsigned short a = get_value("a");

  printf("Part 1: the value on wire 'a' is %u\n", a);

  for (int i = 0; i < MAX_WIRES; i++) {
    wire_map[i].calculated = false;
  }

  struct wire* wire = get_wire("b");
  wire->value = a;
  wire->calculated = true;

  a = get_value("a");

  printf("Part 2: the value on wire 'a' is %u\n", a);

  return 0;
}
