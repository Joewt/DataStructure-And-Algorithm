<?php
$arr = [1,2,5,7,2,3,5];
$n = count($arr);
$time = -microtime(true);
for($i = 0; $i < $n; $i++){
    for($j = $i; $j > 0 && $arr[$j-1]>$arr[$j]; $j--){
        $temp = $arr[$j-1];
        $arr[$j-1] = $arr[$j];
        $arr[$j] = $temp;
    }
}

$time += microtime(true);
echo $time;
