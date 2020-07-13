//https://github.com/go-arun/sps/tree/master/Week_00
#include <stdio.h>
#include<math.h>
#include <string.h>

void main(){
char array_1[10];
int i=0;
int compare_postions,word_length,mach_count; 

printf("Enter a word to check whether it is Palindrome or NOT : ");
scanf("%s",array_1);

word_length = strlen(array_1);
compare_postions = word_length/2;// only half part is need to compare 

for (i=0;i < compare_postions;i++){
    if ( array_1[i] == array_1[--word_length]){
        mach_count++;
    }
}

if ( mach_count == compare_postions){ // All comapre postions matched 
    printf("is Palindrome\n");
}else{
    printf("is NOT Palindrome\n");
}

}

// while (array_1[i] != 0)
// {
//     printf("%c",array_1[i++]);
//     printf("length of array is : %d\n",strlen(array_1)/2);
    
// }





