#include <stdio.h>
#include <map>
using namespace std;

int main() {
    int n, temp, a, b;
    map<int, int> data;
    scanf("%d%d%d", &n, &a, &b);
    for (int i = 0; i < n; i++) {
        scanf("%d", &temp);
        data.insert({temp, 1});
    }
    for (auto i = data.lower_bound(a); i != data.upper_bound(b); i++) {
        printf("%d ", i -> first);
    }
}
