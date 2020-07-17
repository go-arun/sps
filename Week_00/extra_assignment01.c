#include<stdio.h>
void print_name(char nm[25],int count);
void main(){
    char name[25];
    printf("Enter Name:");
    scanf("%s",name);
    print_name(name,7);
}
void print_name(char nm[25],int count){
    if(count-- > 0){
        printf("%s\n",nm);
        print_name(nm,count);
    }
}
