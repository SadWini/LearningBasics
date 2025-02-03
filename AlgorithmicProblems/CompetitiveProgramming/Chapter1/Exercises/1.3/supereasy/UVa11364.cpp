#include <stdio.h>
using namespace std;

int main() {
    int n, k;
    scanf("%d", &n);
    int min, max, temp;
    for (int i = 0; i < n; i++) {
        min = 100;
        max = -1;
        scanf("%d", &k);
        for (int j = 0; j < k; j++) {
            scanf("%d", &temp);
            if (temp > max) {
                max = temp;
            }
            if (temp < min) {
                min = temp;
            }
        }
        printf("%d\n", 2*(max - min));
    }
}