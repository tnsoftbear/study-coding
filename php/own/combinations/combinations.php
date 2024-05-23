<?php


function combinations($all_values, $k) {
    if ($k === 0) {
        return [[]];
    }

    $index_combos = [];
    $cnt = count($all_values);

    for ($i = 0; $i < $k; $i++) $track[$i] = -1;

    $pos = 0;
    $done = false;
    while (!$done) {
        $track[$pos] = findNextPos($pos, $track);
        pr($track);
        if ($pos == $k - 1) {
            if ($track[$pos] < $cnt) {
                $index_combos[] = $track;
                pr($track);
            } else {
                while ($pos >= 0 && $track[$pos] >= $cnt - ($k - $pos - 1)) $pos--;
                if ($pos == -1) {
                    $done = true;
                }
                for ($i = $pos+1; $i < $k; $i++) $track[$i] = -1;
            }
        } else {
            $pos++;
        }
    }

    $results = [];
    foreach ($index_combos as $index_combo) {
        $values = [];
        foreach ($index_combo as $val) {
            $values[] = $all_values[$val];
        }
        $results[] = $values;
    }

    return $results;
}

function findNextPos($pos, $track) {
    if ($pos == 0) {
        return $track[$pos] + 1;
    } else if ($track[$pos] == -1) {
        return $track[$pos - 1] + 1;
    } else {
        return $track[$pos] + 1;
    }
}

function pr($all_values) {
    echo array_reduce($all_values, fn ($el, $acc) => $el . " " . $acc, "") . "\n";
}

foreach (combinations([10,20,30,40,50], 3) as $values) {
    pr($values);
}