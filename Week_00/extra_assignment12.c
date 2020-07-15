#include<stdio.h>
int main()
{
   char str[100];
   int alphabets=0,digits=0,specilchars=0;
   printf("Enter a string: ");
   scanf("%[^\n]",str);
  // printf("String = %c\n",str[0]);

   int i=0;
   int aschi_val;
   while (str[i] != 0){
       
       aschi_val = ("%d",str[i++]);
       if (aschi_val < 48 || aschi_val > 122 ) {// special character range
            specilchars++;
       }else if ((aschi_val >=97 && aschi_val <= 122) || (aschi_val >=65 && aschi_val <= 90)) {
           alphabets++;
       }else if (aschi_val >=48 && aschi_val <= 57)
       {
           digits++;
       }else
       { // this to cover few more special characters asci 58 to 64
           specilchars++;
       }
   }
   printf("Total digits:%d,Total Alphabets:%d,Total SpecialCharacter%d\n",digits,alphabets,specilchars);
}
