#include <stdio.h>
using namespace std;

int main() {
    int k;
    scanf("%d", &k);
    int m, n;
    for(int i = 0; i < k; i++) {
        scanf("%d%d", &n, &m);
        printf("%d\n", ((n - 2) / 3 + ((n - 2) % 3 > 0)) * ((m - 2) / 3 + ((m - 2 ) % 3 > 0)));
    }
}