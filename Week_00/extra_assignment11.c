//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
void main(){
   
    int usr_input,i;
    int result_array[5]={0};// max chance is 5 in 10 inputs
    int user_inp_array[10]={0};
        
    int result_count = 0;// these 2 counts used for avoid last empty postions if any
    int input_count  = 0; 
    int is_maching;

    printf("Provide 10 random positive numbers to find the repeated ones]\n");
    for(i=0;i<10;i++){
        printf("Enter number:");
        scanf("%d",&usr_input);
        is_maching = repeat_tracer(usr_input,user_inp_array,result_array);

        if (is_maching == 1){
            result_array[result_count++] = usr_input;
        }else
        {//Add entry in user input arry 
            user_inp_array[input_count++] = usr_input;
        }
    }
    printf ("\nResults Below .....\n");
    show_result(result_array,result_count);
}
//function
int  repeat_tracer(int last_inp,int inputs_array[10],int result_arr[5]){
    int i;
    for(i=0;i<5;i++){//if the input already in result then no need of further checking
        if (last_inp == result_arr[i]){
            return 0;
        }
     }
    for(i=0;i<10;i++){//Checking with previouly enterred values
        if (last_inp == inputs_array[i]){
            
            return 1; // 1 Means mach found !
        }
    }
return 0; // 0 Means match not found

}

//To Display result
int show_result(int result[4],int count){
    int i;
    for (i=0;i<count;i++){
        printf("%d  ",result[i]);
    }
    if (count == 0){
        printf ("\nNo Numbers Repeated .....");
    }
    return 0;
}