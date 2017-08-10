#include <criterion/criterion.h>
#include <criterion/logging.h>
#include <string.h>
#include <stdlib.h>

char *reverseString(char *c) {
	int ln, s;
	char *ptr = NULL;

	ln = strlen(c);
	s = ln - 1; // account for '\0'
	// dynamic allocate char size
	ptr = malloc(sizeof(char)*(ln + 1));
	// add characters
	for (int i = 0; i < ln; i++) {
		printf("c i %d\n", ln - i - 1);
		ptr[i] = c[s - i];
	}
	return ptr;
}

Test(ReverseString, check1) {
	char s[] = "hello world";
	char *e = reverseString(s);
	// need to test expect
}
