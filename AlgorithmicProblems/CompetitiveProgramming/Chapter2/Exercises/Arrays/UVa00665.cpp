#include <stdio.h>

int main() {
	int d, n, k, p, b, temp, index;
	char c;
	int coins[110];
	bool res[110];
	scanf("%d", &d);
	for (int i = 0; i < d; i++) {
		scanf("%d%d", &n, &k);
		for (int i = 0; i < n; i++) {
			res[i] = false;
		}
		temp = 0;
		while (k--) {
			scanf("%d", &p);
			for (int i = 0; i < 2 * p; i++) {
				scanf("%d", coins + i);
			}
			scanf("%c", &c);
			switch (c) {
				case '<':
					for (int i = p; i < 2 * p; i++) {
						res[coins[i]] = true;
					}
					break;
				case '>':
					for (int i = 0; i < p; i++) {
						res[coins[i]] = true;
					}
					break;
				case '=':
					for (int i = 0; i < 2 * p; i++) {
						res[coins[i]] = true;
					}
					break;
			}
			for (int i = 0; i < n; i++) {
				if (res[i] == false) {
					temp++;
					index = i;
				}
			}
			printf("%d\n", temp == 1 ? index : 0);
		}
	} 
}
