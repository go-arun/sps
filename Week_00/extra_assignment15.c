#include <stdio.h>

void main(){
    int i;
    int stars_to_print;
    int spacing;
    for(i=1;i<=10;i++){
        stars_to_print = i * 2; 
        spacing = 20 - stars_to_print;
        while (stars_to_print-- > 0)
        {
                printf("*");
            if (i == stars_to_print){//printed half stars,add space
                while (spacing-- > 0)
                {
                     printf(" ");
                }
                
            }
        }
        printf("\n");
    }
}

