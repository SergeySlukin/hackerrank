<?php

function birthdayCakeCandles($ar) {

    sort($ar, SORT_NUMERIC);

    $count = 1;
    $arCount = count($ar);
    $prev = $ar[$arCount - 1];
    for ($i = count($ar) - 2; $i >= 0; $i--) {
        if ($ar[$i] !== $prev) {
            break;
        }
        $count++;
    }
    return $count;
}

$fptr = fopen(getenv("OUTPUT_PATH"), "w");

$stdin = fopen("php://stdin", "r");

fscanf($stdin, "%d\n", $ar_count);

fscanf($stdin, "%[^\n]", $ar_temp);

$ar = array_map('intval', preg_split('/ /', $ar_temp, -1, PREG_SPLIT_NO_EMPTY));

$result = birthdayCakeCandles($ar);

fwrite($fptr, $result . "\n");

fclose($stdin);
fclose($fptr);
