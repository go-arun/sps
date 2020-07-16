#include <stdio.h>
void main(){
int req_rows;
int starting_asci;
int start_pos=10,j;
int i,elements_in_row=1,elements_in_row_maintainer,middle_element_position;
int space_maintainer,s;
printf("Enter number of rows for this special pattern :");
scanf("%d",&req_rows);
space_maintainer = req_rows;
for(i=0;i<req_rows;i++){
    
 for (s = space_maintainer--;s >0;s--){//maintain starting postion 
           printf("  ");
  }
    middle_element_position=elements_in_row/2;
    starting_asci=65;
    while(elements_in_row >0){
        if (elements_in_row<=middle_element_position+1 ){ // recahed middle then go in reverse pattern
            printf("%c ",starting_asci--);
        }else
        {
            printf("%c ",starting_asci++);
        }
        elements_in_row--;
    }
    elements_in_row = elements_in_row_maintainer + 2; //every time 2 more elements
    elements_in_row_maintainer = elements_in_row;

  printf("\n");
}
printf("\n");
}
