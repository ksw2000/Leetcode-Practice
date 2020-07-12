// https://leetcode.com/problems/two-sum/
// 0 ms 2.4 MB
use std::collections::HashMap;

struct Solution{}
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut h = HashMap::new();
        let mut k = 0;
        for v in &nums{
            match h.get(&(target - (*v))){
                Some(value)=>{
                    return vec![*value, k];
                }
                None=>{
                    h.insert(*v, k);
                    k=k+1;
                }
            }
        }
        return vec![];
    }
}

fn main(){
    let num: Vec<i32> = vec![2, 7, 11, 15];
    let target = 9;
    println!("{:?}" , Solution::two_sum(num, target));
}
