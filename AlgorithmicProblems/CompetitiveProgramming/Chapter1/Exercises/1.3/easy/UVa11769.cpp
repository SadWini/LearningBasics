#include <stdio.h>

int main() {
    int k, n, curr, next, index = 1;
    scanf("%d", &k);
    for (int i = 0; i < k; i++) {
        int high = 0, low = 0;
        scanf("%d %d", &n, &curr);
        n = n - 1;
        for (int j = 0; j < n; j++) {
            scanf("%d", &next);
            if (next > curr) high += 1;
            if (next < curr) low += 1;
            curr = next;
        }
        printf("Case %d: %d %d\n", index++, high, low);
    }
    return 0;
}