#include <bits/stdc++.h>

int char_mismatch_count(char *base, char *input) {
    int cnt = 0;
    while (*base && *input) {
        if (*base != *input) cnt += 1;
        base += 1;
        input += 1;
    }
    return cnt;
}

int main() {
    int n;
    char str[10];
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%s", str);
        if (char_mismatch_count("one", str) <= 1)
            puts("1");
        if (char_mismatch_count("two", str) <= 1)
            puts("2");
        if (char_mismatch_count("three", str) <= 1)
            puts("3");
    }
    return 0;
}