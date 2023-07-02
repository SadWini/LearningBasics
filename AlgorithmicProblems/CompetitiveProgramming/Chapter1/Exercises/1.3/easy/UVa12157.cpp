#include <stdio.h>

int main() {
    int n, k, temp, ju, mi;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        ju = 0;
        mi = 0;
        scanf("%d", &k);
        for (int j = 0; j < k; j++) {
            scanf("%d", &temp);
            mi += 10 * (temp / 30) + 10;
            ju += 15 * (temp / 60) + 15;
        }
        if (mi == ju) {
            printf("Case %d: Mile Juice %d", i + 1, mi);
        } else {
            printf("Case %d: %s %d\n", i + 1, mi > ju ? "Juice" : "Mile", mi > ju ? ju : mi);
        }
    }
}
