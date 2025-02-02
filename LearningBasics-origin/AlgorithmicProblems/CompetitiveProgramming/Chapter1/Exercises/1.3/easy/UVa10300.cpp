#include <stdio.h>

int main() {
    int n, f;
    int s, a, r;
    long budget;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        budget = 0;
        scanf("%d", &f);
        for (int j = 0; j < f; j++) {
            scanf("%d%d%d", &s, &a, &r);
            budget += s * r;
        }
        printf("%ld\n", budget);
    }
}