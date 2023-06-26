#include <stdio.h>
using namespace std;

int main() {
    int n, a, b, c;
    scanf("%d", &n);
    for (int i = 1; i <= n; i++) {
        scanf("%d%d%d", &a, &b, &c);
        printf("Case %d: %s\n",i, a <= 20 && b <= 20 && c <= 20 ? "good" : "bad");
    }
    return 0;
}