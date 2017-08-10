#include <criterion/criterion.h>
#include <criterion/logging.h>

int FirstFactorial(int n) {
	if (n == 1) return 1;
	return (n * FirstFactorial(n - 1));
}

Test(First_Factorial, check_4) {
	int n = FirstFactorial(4);

	int expn = 24;
	cr_log_info("n %d", n);
	cr_expect_eq(n, expn, "result should be %d", expn);
}

Test(First_Factorial, check_8) {
	int n = FirstFactorial(8);

	int expn = 40320;
	cr_log_info("n %d\n", n);
	cr_expect_eq(n, expn, "result should be %d", expn);
}
