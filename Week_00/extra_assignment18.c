//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>

void main(){
char usr_inp[20];
int i;
int ascii_val;

printf("Enter Word: \n");
    scanf("%s",usr_inp);
    while (usr_inp[i] != 0)
    {
        ascii_val = ("%d",usr_inp[i]);
        usr_inp[i++] = ("%c",ascii_val-32);// ASCII difference 
        
    }
printf("%s",usr_inp);
}