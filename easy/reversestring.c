#include <criterion/criterion.h>
#include <criterion/logging.h>
#include <string.h>
#include <stdlib.h>

char* reverseString(char *c) {
	int ln, s;
	char *ptr = NULL;

	ln = strlen(c);
	s = ln - 1; // account for '\0'
	// dynamic allocate char size
	ptr = malloc(sizeof(char)*(ln + 1));
	// add characters
	for (int i = 0; i < ln; i++) {
		ptr[i] = c[s - i];
	}
	return ptr;
}

Test(ReverseString, check1) {
	char s[] = "hello world";
	char *r = reverseString(s);
	char e[] = "dlrow olleh";

	cr_log_info("s %s\n", r);
	cr_log_info("expect %s\n", e);
	cr_expect_str_eq(r, e, "test");
}

Test(ReverseString, check2) {
	char s[] = "Hello World and Coders";
	char *r = reverseString(s);
	char e[] = "sredoC dna dlroW olleH";

	cr_log_info("s %s\n", r);
	cr_log_info("expect %s\n", e);
	cr_expect_str_eq(r, e, "should be %s", e);
}
