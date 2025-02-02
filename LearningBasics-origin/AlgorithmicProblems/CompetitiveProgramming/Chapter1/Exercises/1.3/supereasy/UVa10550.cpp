#include <stdio.h>
using namespace std;

int main() {
    int ans;
    int i, x, y, z;
    while (scanf("%d%d%d%d", &i, &x, &y, &z) && i|x|y|z) {
        ans = 0;
        ans += 120;
        ans += (i - x < 0 ? (40 + i - x) : (i - x));
        ans += (y - x < 0 ? (40 + y - x) : (y - x));
        ans += (y - z < 0 ? (40 + y - z): (y - z));
        ans *= 9;
        printf("%d\n", ans);
    }
    return 0;
}
