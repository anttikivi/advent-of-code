TARGET := $(shell ls *.c | sed 's/\.c/\.out/g')
CC     := cc
CFLAGS := -Wall -Wextra -Werror -Wpedantic -std=c99 -O2

all: $(TARGET)

$(TARGET): %.out: %.c
	$(CC) $(CFLAGS) -o $@ $^

day-04.out: day-04.c
	$(CC) $(CFLAGS) -Wno-deprecated-declarations -L/opt/homebrew/lib -I/opt/homebrew/include -lssl -lcrypto -o $@ $^

day-12.out: day-12.c
	$(CC) $(CFLAGS) -o $@ $^ utils/cJSON.c

clean:
	rm -f *.out
