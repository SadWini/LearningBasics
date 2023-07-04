#include <stdio.h>
#include <algorithm>
using namespace std;

int main() {
    int m, k, n, c, temp, a;
    int data[101];
    bool pass;
    while (scanf("%d", &k), k != 0) {
        pass = true;
        scanf("%d", &m);
        for (int i = 0; i < k; i++) {
            scanf("%d", data + i);
        }
        for (int i = 0; i < m; i++) {
            temp = 0;
            scanf("%d%d", &n, &c);
            while (n--) {
                scanf("%d", &a);
                if (find(data, data + k, a) != data +k) {
                    temp++;
                }
            }
            if (temp < c) {
                pass = false;
            }
        }
        pass ? printf("yes\n") : printf("no\n");
    }
}