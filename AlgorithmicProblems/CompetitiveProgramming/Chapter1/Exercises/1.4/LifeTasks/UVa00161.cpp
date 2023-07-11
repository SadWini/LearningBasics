#include<iostream>
#include<stdio.h>
using namespace std;

int main() {
    int sig;
    while(scanf("%d", &sig)){
        if(sig == 0){
            break;
        }
        int signal[100];
        signal[0] = sig;
        int num = 0;
        int time[18001] = {0};
        int min = 90;
        while(scanf("%d", &signal[++num]), signal[num] != 0);
        for(int i=0; signal[i] != 0; i++){
            int sec = 0;
            int present_sig = signal[i];
            if (present_sig<min) {
                min = present_sig;
            }
            int green = present_sig - 5;
            while(sec < 18001){
                for (int j = 0; j < green; j++){
                    if((sec + j) < 18001)
                        time[sec + j]++;
                    else
                        break;
                }
                sec = sec+(present_sig*2);
            }
        }

        int ans_time = 0;
        for(int i = 2 * min; i < 18001; i++){
            if(time[i] == num){
                ans_time = i;
                break;
            }
        }
        if (ans_time == 0) {
            printf("Signals fail to synchronise in 5 hours\n");
        } else
            printf("%02d:%02d:%02d\n", ans_time / 3600, (ans_time % 3600) / 60, ans_time % 60);
    }
    return 0;
}