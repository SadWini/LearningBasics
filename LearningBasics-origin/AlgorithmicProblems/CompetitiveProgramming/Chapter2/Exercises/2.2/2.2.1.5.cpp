#include <stdio.h>

void swap(int a, int b) {
    a = b + a;
    b = a - b;
    a = b - a;
}
int main() {
    int n, a, b, next, max = 0;
    scanf("%d%d", &n, &a);
    b = 1;
    for (int i = 0; i < n - 1 ; i++) {
        scanf("%d", &next);
        if (next > a) {
            b++;
        } else {
            if (b > max) {
                max = b;
            }
            b = 1;
        }
        if (b > max) {
            max = b;
        }
        a = next;
    }
    printf("%d", max);
}