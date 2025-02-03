#include <stdio.h>
using namespace std;

int main() {
    int n, x, y, z, mid;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d%d%d", &x, &y, &z);
        mid = z;
        if ((x > y && y > z) || (z > y && y > x))
            mid = y;
        if ((y > x && x > z) || (z > x && x > y))
            mid = x;
        printf("Case %d: %d\n", i + 1, mid);
    }
}