//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
void main(){
    int max_num=0;
    int usr_input;
    printf("Enter some random number to find the max number:[Press 0 to Exit.]");
    while (1)
    {
    printf("Enter number:");
    scanf("%d",&usr_input);

    if (usr_input == 0){
        printf("Max number in this series is :%d",max_num);
        break;
    }else
    {
        max_num = is_large(max_num,usr_input);
    }
    }
}
//function
int is_large(int topper,int recent){
    if (topper < recent){
        return recent;
    }else{
        return topper;
    }
}