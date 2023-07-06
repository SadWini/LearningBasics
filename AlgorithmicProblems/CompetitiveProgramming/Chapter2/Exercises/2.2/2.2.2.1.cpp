#include <stdio.h>

int main() {
    int a, b;
    scanf("%d%d", &a, &b);
    printf("%d % %d = %d", a, b, a & ((1 << b) - 1));
}