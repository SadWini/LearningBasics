#include <stdio.h>

int main() {
    int n, cur, next;
    bool ord;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        ord = true;
        scanf("%d%d", &cur, &next);
        if (cur > next) {
            cur = next;
            for (int j = 2; j < 10; j++) {
                scanf("%d", &next);
                if (cur < next)
                    ord = false;
                cur = next;
            }
        } else {
            cur = next;
            for (int j = 2; j < 10; j++) {
                scanf("%d", &next);
                if (cur > next)
                    ord = false;
                cur = next;
            }
        }
        printf("%s\n", ord ? "Ordered" : "Unordered");
    }
}