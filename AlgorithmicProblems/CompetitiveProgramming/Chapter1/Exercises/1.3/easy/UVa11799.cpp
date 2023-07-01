#include <stdio.h>
#include <climits>

int main() {
    int k, n, index = 1;
    scanf("%d", &k);
    while (k--) {
        int tmp, max = INT_MIN;
        scanf("%d", &n);
        while (n--) {
            scanf("%d", &tmp);
            if (tmp > max) max = tmp;
        }
        printf("Case %d: %d\n", index++, max);
    }
    return 0;
}