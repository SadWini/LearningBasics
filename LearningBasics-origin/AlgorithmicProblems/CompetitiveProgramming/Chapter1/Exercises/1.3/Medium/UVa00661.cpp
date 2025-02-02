#include <stdio.h>

int main() {
    int n = 1, m, dev[21], v, max, temp, tempMax, index = 0;
    bool blown;
    while(true) {
        index++;
        temp = 0;
        max = 0;
        tempMax = 0;
        blown = false;
        scanf("%d%d%d", &n, &m, &max);
        if (!(n || m || max)) break;
        for (int i = 0; i < n; i++)
            scanf("%d", dev + i);
        for (int i = 0; i < m; i++) {
            scanf("%d", &v);
            temp += dev[v-1];
            dev[v-1] *= -1;
            if (temp > tempMax) tempMax = temp;
            if (temp > max)blown = true;
        }
        printf("Sequence %d\n", index);
        if (blown) {
            printf("Fuse was blown.\n");
        } else {
            printf("Fuse was not blown.\nMaximal power consumption was %d ampere%s.\n",
                   tempMax, tempMax > 1 ? "s" : "");
        }
    }
    return 0;
}