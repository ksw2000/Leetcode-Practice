<?php
// URL : https://leetcode.com/problems/rotate-image/
//8 ms	14.9 MB
class Solution {
    public $matrix;
    function rotate(&$matrix) {
        $size=count($matrix);
        for($i=0; $i<(int)($size/2); $i++){
            for($j=$i; $j<$size-$i-1; $j++){
                $cache=$matrix[$i][$j];
                $matrix[$i][$j]=$matrix[$size-1-$j][$i];
                $matrix[$size-1-$j][$i]=$matrix[$size-1-$i][$size-1-$j];
                $matrix[$size-1-$i][$size-1-$j]=$matrix[$j][$size-1-$i];
                $matrix[$j][$size-1-$i]=$cache;
            }
        }
    }
}

function printM(&$arr){
    for($i=0; $i<count($arr); $i++){
        for($j=0; $j<count($arr); $j++){
            echo $arr[$i][$j]." ";
        }
        echo '<br>';
    }
    echo '<br>';
}

$sol=new Solution;

$arr=Array(
        Array(1)
    );
$sol->matrix=$arr;
$sol->rotate($sol->matrix);
printM($sol->matrix);

$arr=Array(
        Array(1,2),
        Array(3,4),
    );
$sol->matrix=$arr;
$sol->rotate($sol->matrix);
printM($sol->matrix);

$arr=Array(
        Array(1,2,3),
        Array(4,5,6),
        Array(7,8,9)
    );
$sol->matrix=$arr;
$sol->rotate($sol->matrix);
printM($sol->matrix);

$arr=Array(
        Array(1,2,3,4),
        Array(5,6,7,8),
        Array(9,10,11,12),
        Array(13,14,15,16)
    );
$sol->matrix=$arr;
$sol->rotate($sol->matrix);
printM($sol->matrix);
?>
