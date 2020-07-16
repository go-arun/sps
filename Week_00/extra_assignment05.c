#include <stdio.h>

void main(){
    char string1_array[25];
    char string2_array[25];
    int  user_choice;
    
    printf("Enter A String :");
    scanf("%s",&string1_array);
    int i=0;
    while (string1_array[i] != 0){
        string2_array[i] = string1_array[i++];
    }
    string2_array[i] = 0;
    printf ("String1[%s] copied to String2 its, value is -> [%s]",string1_array,string2_array);
    printf("\n_____________________________________________________\n");
}

