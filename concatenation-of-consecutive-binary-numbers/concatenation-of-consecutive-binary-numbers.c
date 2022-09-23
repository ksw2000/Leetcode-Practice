// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// 21 ms 5.7 MB

// 1 1          1
// 2 110        1 * 4 + 2 = 6
// 3 11011      6 * 4 + 3 = 27
// 4 11011100  27 * 8 + 4 = 220

int concatenatedBinary(int n) {
    static unsigned long long sum;
    static int i;
    static int len;
    sum = sum ^ sum;
    len = len ^ len;
    n++;
    for (i = 1; i ^ n; i++) {
        if (!(i & (i - 1)))
            len++;
        sum = (sum << len | i) % 1000000007;
    }
    return sum;
}