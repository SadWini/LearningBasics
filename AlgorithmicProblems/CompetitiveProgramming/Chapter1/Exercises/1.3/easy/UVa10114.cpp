#include <stdio.h>

int main() {
    int duration, m, n;
    double loan, downpayment, dep[110];
    double temp;
    int time = 0;
    while (scanf("%d %lf %lf %d", &duration, &downpayment, &loan, &n) && duration > 0) {
        while (n--) {
            scanf("%d %lf", &m, &temp);
            for (int i = m; i <= 110; ++i)
                dep[i] = temp;
        }
        time = 0;
        double monthly_payment = loan / duration;
        double car_value = (loan + downpayment) * (1 - dep[0]);
        while (car_value < loan) {
            time += 1;
            loan -= monthly_payment;
            car_value -= car_value * dep[time];
        }
        printf("%d month", time);
        printf(time != 1 ? "s\n" : "\n");
    }
    return 0;
}