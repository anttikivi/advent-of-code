#include <openssl/md5.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT "***REMOVED***"

int main(void) {
  // Make the MD5 input large enough so it can hold the number when
  // brute-forcing.
  int input_len = strlen(INPUT);
  char* input = (char*)malloc(input_len * sizeof(char) + 50 * sizeof(char));

  *input = '\0';
  sprintf(input, "%s", INPUT);

  printf("Input: %s\n", input);

  unsigned long i = 0;
  unsigned long num1 = 0;
  unsigned long num2 = 0;
  unsigned char* hash =
      (unsigned char*)malloc(MD5_DIGEST_LENGTH * sizeof(unsigned char));

  while (!num1 || !num2) {
    sprintf(input + input_len, "%lu", i);
    MD5((unsigned char*)input, strlen(input), hash);

    // The hexadecimal hash is usually represented by a 32-character string,
    // where each character represents 4 bits. Since we're only interested in
    // the first 5 bits, we can just check them.
    if (!hash[0] && !hash[1] && hash[2] <= 0x10) {
      if (!num1) {
        num1 = i;
      }

      if (!hash[2]) {
        if (!num2) {
          num2 = i;
        }
      }
    }

    i++;
  }

  printf("Part 1: %lu\n", num1);
  printf("Part 2: %lu\n", num2);

  free(hash);
  free(input);

  return 0;
}
