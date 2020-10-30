<?php

function gradingStudents($grades) {
    foreach ($grades as $key => $grade) {
        if ($grade >= 38) {
            $r = $grade / 5;
            if ($r < 20) {
                $r++;
            }
            $n = 5 * (int) $r;
            $grades[$key] = $n - $grade < 3 ? $n : $grade;
        }
    }
    return $grades;
}

$fptr = fopen(getenv("OUTPUT_PATH"), "w");

$grades_count = intval(trim(fgets(STDIN)));

$grades = array();

for ($i = 0; $i < $grades_count; $i++) {
    $grades_item = intval(trim(fgets(STDIN)));
    $grades[] = $grades_item;
}

$result = gradingStudents($grades);

fwrite($fptr, implode("\n", $result) . "\n");

fclose($fptr);
