<?php
$arr = [];
$n = 1000000;

for($k = 0; $k < $n; $k++){
    $arr[$k] = mt_rand(1,10);
}

$time = -microtime(true);

$h = 1;
while($h < $n/3)
    $h = 3 * $h +1;

while($h >= 1 ){
    for($i = $h; $i < $n; $i++){
    
        $e = $arr[$i];
        $j = $i;
        for($j = $i; $j>=$h&&$arr[$j-$h]>$e;$j-=$h){
            $arr[$j] = $arr[$j-$h];
        }
        $arr[$j] = $e;
    }
    $h /= 3;
}

$time += microtime(true);
echo $time;
