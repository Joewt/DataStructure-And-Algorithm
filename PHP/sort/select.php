<?php
$arr = [];
$n = 20000;
$time = -microtime(true);

for($i = 0; $i < $n; $i++){
    $arr[$i] = mt_rand(1,100);
}


for($i = 0; $i < $n; $i++){

    $minindex = $i;
    for($j = $i+1; $j < $n; $j++){
        if($arr[$j]<$arr[$minindex]){
            $minindex = $j;
        }
    }

    $temp = $arr[$i];
    $arr[$i] = $arr[$minindex];
    $arr[$minindex] = $temp;

}

$time   += microtime(true);
print_r($time);
