#include <stdio.h>
int get_fibonacci(int limit);
    int prev_no=0;
    int curr_no=1;
    int count=0;
    int fibno=0;
void main(){
    int nth;
    printf("Print Which element ?:");
    scanf("%d",&nth);
    get_fibonacci(nth);
}
int get_fibonacci(int limit){
    if (limit==1){ // COULD'T FIND A LOGIC FOR FIRST ANS SECOND NUMBERS ( HARD CODED!!)
        printf("Result:0\n");
        return 0;
    }else if(limit==2){
        printf("Result:1\n");
        return 0;
    }
    if (count++ < limit ){
        
        if (count == limit-1){
            printf("Result:%d,\n",fibno);
        }
        fibno = (prev_no + curr_no);
        prev_no = curr_no;
        curr_no = fibno;
        get_fibonacci(limit);
    }
    return 0;
}
