<?php
$arr = [1,3,1,4,1,5,6];
$n = count($arr);
$swaped = true;
do
{
    $swaped = false;
    for($i = 1; $i < $n; $i++){
        if($arr[$i-1]>$arr[$i]){
            $temp = $arr[$i-1];
            $arr[$i-1] = $arr[$i];
            $arr[$i] = $temp;
            $swaped = true;
        }
    }
    $n--;
}while($swaped);

print_r($arr);
