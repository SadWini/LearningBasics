#include <stdio.h>
#include <math.h>

int main() {
    int n, m;
    int temp, next, res, high;
    while (scanf("%d%d", &n, &m), n || m) {
        res = 0;
        scanf("%d", &temp);
        high = temp;
        for (int i = 1; i < m; i++) {
            scanf("%d", &next);
            if (next > high) {
                res += next - high;
                high = next;
            }
            if (next < temp) {
                res += temp - next;
            }
            temp = next;
        }
        res += n - high;
        printf("%d\n", res);
    }
}