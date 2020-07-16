#include <stdio.h>

void main(){

    char word1_array[25];
    char word2_array[25];
    int  user_choice;
while(1){
    printf("_____________________________________________________\n");
    printf("[1]->Length [2]->Concatenate [3]->Reverse [0]->Exit :");
    scanf("%d",&user_choice);
    switch (user_choice)
    {
    case 0://Exit
        printf("Exitting...");
        exit(1);
        break;        
    case 1://Length
        printf("Enter A String to get its Length :");
        scanf("%s",&word1_array);
        printf("Word Length is:%d\n",get_string_length(word1_array));
        break;
    case 2://Concat
        printf("Enter First String :");
        scanf("%s",&word1_array);
        printf("Enter Second String :");
        scanf("%s",word2_array);   
        concate_strings(word1_array,word2_array);
        printf(" Result :%s\n",word1_array); // result storred in word1_array
        break;      
    case 3://Reverse
        printf("Enter A String to Reverse it :");
        scanf("%s",&word1_array);
        string_reverse(word1_array);
        printf(" Result :%s\n",word1_array); // result storred in word1_array
        break;  
    
    default:
        printf("Wrong Input Please try again\n");
        break;
    }

    }
}

int string_reverse(char arr[25]){
    char temp;
    int i,j=0,req_swap_count;
    i = get_string_length(arr);
    req_swap_count = i/2; // swappng till half is ok!
    while (i > req_swap_count){
        temp = arr[j];
        arr[j++] = arr[--i];
        arr[i] = temp;
    }
    
    return 0;
}
int concate_strings(char arr1[25],char arr2[25]){
    int i = get_string_length(arr1); // adding first word lenth to i
  //  int j = get_string_length(arr2)-1; //adding second word lenth to j
    int j=0;
    while (arr2[j] != 0)
    {
        arr1[i++] = arr2[j++];
    }
    arr1[i] = 0; // Made next char NULL
    return 0;   
}

int get_string_length(char arr[25]){
    int i=0;
    while (arr[i++] != 0){} // just increment i val till last char.
    return --i;
}