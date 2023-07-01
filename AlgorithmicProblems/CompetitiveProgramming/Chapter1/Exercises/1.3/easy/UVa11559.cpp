#include <stdio.h>

int main() {
    int n, b, h, w, p, available;
    while (scanf("%d %d %d %d", &n, &b, &h, &w) != EOF) {
        int cost = 2000000000;
        while (h--) {
            scanf("%d", &p);
            for (int i = 0; i < w; ++i) {
                scanf("%d", &available);
                if (n <= available && n * p < cost) {
                    cost = n * p;
                }
            }
        }
        if (cost < b) {
            printf("%d\n", cost);
        } else {
            puts("stay home");
        }
    }
    return 0;
}
