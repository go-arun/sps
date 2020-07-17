#include<stdio.h>
#include<string.h>
int main()
{
   char str_sentense[100];
   char search_word[10];
   int consecutive_mach=0;
   int search_word_length,sentance_length,retain_search_word_length;
   int match_count=0;
   int j,word_count=1;

   printf("Enter one Sentense: ");
   scanf("%[^\n]",str_sentense);
   
   printf("Enter one word to show the position in above given Sentence:");
   scanf("%s",&search_word);

   search_word_length = strlen(search_word);
   sentance_length    = strlen(str_sentense);
   int i=0;j=0; //to walk through keyword and sentense.
   while (sentance_length > 0)
   {
      if("%d",str_sentense[i] == 32){word_count++;}
      if(search_word[j] == str_sentense[i++]){ // First letter mached , comapre remaining chars too
         j++; // okay check next letter
         consecutive_mach++;
         if(consecutive_mach==search_word_length){
            printf("Postion of word is %d\n",word_count);
            break;
         }
      }else
      {
         consecutive_mach = 0; // resetting because whole word not matched 
         j=0; //consicutive match failed, so start serach from begninig.
      }
      sentance_length--;
   }
}

