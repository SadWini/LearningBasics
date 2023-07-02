#include <stdio.h>
#include <math.h>

int main() {
    int a = 0, b = 0, temp;
    while(a > -1 && b > -1) {
        scanf("%d%d", &a, &b);
        if (a == -1 & b == -1) break;
        if (abs(b - a) > abs(b - 99 + a)) {
            temp = abs(b - 98 + a);
        } else {
            temp = abs(b - a);
        }
        printf("%d\n", temp);
    }
}