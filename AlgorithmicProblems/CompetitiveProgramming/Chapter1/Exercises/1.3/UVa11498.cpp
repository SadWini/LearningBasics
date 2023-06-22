#include <stdio.h>
using namespace std;

int main() {
    int n, M, N, x, y;
    while (scanf("%d", &n) && n) {
        scanf("%d%d", &M, &N);
        while (n--) {
            scanf("%d%d", &x, &y);
            if (x == M || y == N) {
                puts("divisa");
            } else {
                printf("%c%c\n", (y > N ? 'N' : 'S'), (x > M ? 'E' : 'O'));
            }
        }
    }
    return 0;
}