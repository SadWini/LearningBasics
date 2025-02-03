#include <stdio.h>

int main() {
    int k, res;
    while(scanf("%d", &k), k > 0) {
        res = k;
        while (k > 9) {
            res = 0;
            while(k > 0) {
                res += k % 10;
                k /= 10;
            }
            k = res;
        }
        printf("%d\n", res);
    }
}