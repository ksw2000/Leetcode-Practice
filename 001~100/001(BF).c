// URL : https://leetcode.com/problems/two-sum/
// 140 ms 7.7 MB
#include<stdio.h>
#include<stdlib.h>

int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    int i,j;
    int* ret = calloc(2,sizeof(int));
    *returnSize=0;
    for(i=0; i<numsSize; i++){
        for(j=i+1; j<numsSize; j++){
            if(nums[i]+nums[j]==target){
                ret[0]=i;
                ret[1]=j;
                *returnSize=2;
                return ret;
            }
        }
    }
    return ret;
}

int main(){
    int testcase[]={2, 7, 11, 15};
    int returnSize=0;
    int* answer = twoSum(testcase, 4, 9, &returnSize);

    int i;
    for(i=0; i<returnSize; i++){
        printf("%d ",answer[i]);
    }

    printf("\n");
    system("pause");
    return 0;
}
