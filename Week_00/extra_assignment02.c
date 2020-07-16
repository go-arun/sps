#include <stdio.h>

void main(){

    char word1_array[20]={0};
    char word2_array[20]={0};
    int i=0;

    printf ("Provide two words to comapre \nEnter First word :");
    scanf("%s",&word1_array);
    printf ("Enter Second word :");
    scanf("%s",&word2_array);

    while(word1_array[i] == word2_array[i++]){
        if (i == 20){//Reached end !! Means words are maching!!
            printf(" Those two words are SAME\n");
            break;
        }
    }
    if (i<20){printf("Those two words are not SAME\n");}
    printf("---------------------------\n");
}