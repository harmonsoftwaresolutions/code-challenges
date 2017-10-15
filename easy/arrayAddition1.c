#include <stdio.h>
#include <criterion/criterion.h>
#include <criterion/logging.h>
// ## Array Addition 1
// Have the function ArrayAdditionI(arr) take the array of numbers stored in
// arr and return the string true if any combination of numbers in the array
// can be added up to equal the largest number in the array, otherwise return
// the string false. For example: if arr contains [4, 6, 23, 10, 1, 3] the
// output should return true because 4 + 6 + 10 + 3 = 23. The array will not
// be empty, will not contain all the same elements, and may contain negative
// numbers.

// ### Test Cases:
// Input:5,7,16,1,2
// Output:"false"
//
// Input:3,5,-1,8,12
// Output:"true"

int arrayAddition1(int arr[], int ln)
{
    // find largest integer
    int lg = 0;
    for (int i = 0; i < ln; i++) {
        if (arr[i] > lg)
            lg = arr[i];
    }
    printf("lg %d\n", lg);
    // create new array without largest in
    int nArr[ln - 1];
    for (int i = 0; i < ln - 1; i++) {
        if (arr[i] != lg) {
            nArr[i] = arr[i];
        }
    }

    // iterate summation and check if total equals largest int
    // test summing with each array item as starting point
    int sum = 0;
    printf("Sum start: ");
    for (int i = 0; i < ln - 1; i++) {
        if (sum == lg) {
            break;
        }
        sum = 0;
        for (int j = i; j < ln - 1; j++) {
            sum += nArr[j];
            printf(" %d(%d) ", nArr[j], sum);
            if (sum == lg) {
            printf(" breaking sum %d ", sum);
                break;
            }
        }
    }
    printf(" Sum end:\n");

    int res = sum == lg;

    return res;
};

Test(array_Addition1, check_1)
{
    int arr[5] = { 5, 7, 16, 1, 2 };
    int ln = 5;
    int check = arrayAddition1(arr, ln);

    int expn = 0;
    cr_log_info("check %d", check);
    cr_expect_eq(check, expn, "result should be %d", expn);
};

Test(array_Addition1, check_2)
{
    int arr[5] = { 3, 5, -1, 8, 12 };
    int ln = 5;
    int check = arrayAddition1(arr, ln);

    int expn = 1;
    cr_log_info("check %d", check);
    cr_expect_eq(check, expn, "result should be %d", expn);
};

