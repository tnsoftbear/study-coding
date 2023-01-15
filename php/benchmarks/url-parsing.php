<?php
$urls = [
    'http://example.com/path/to/test?key1=value2',
    '//example.com/path/to/test?key1=value2',
    '/path/to/test?key1=value2',
    'http://example.com/path/to/test',
    '//example.com/path/to/test',
    '/path/to/test',
    '/',
    '/?key1=value2',
    '',
    'abcde'
];

function regexp(string $url): string
{
    preg_match('/^(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\?([^#]*))?(#(.*))?/u', $url, $matches);
    return $matches[5] ?? '';
}

function multiRegexp(array $urls, int $count) {
    for ($i = 0; $i < $count; $i++) {
        foreach ($urls as $url) {
            $result = regexp($url);
            // echo $result . "\n";
        }
    }
}

function parseUrl(string $url): string
{
    $parts = parse_url($url);
    return $parts['path'];
}

function multiParseUrl(array $urls, int $count) {
    for ($i = 0; $i < $count; $i++) {
        foreach ($urls as $url) {
            $result = parseUrl($url);
            // echo $result . "\n";
        }
    }
}

function dummy(string $url) {
    return $url;
}

function multiDummy(array $urls, int $count) {
    for ($i = 0; $i < $count; $i++) {
        foreach ($urls as $url) {
            $result = dummy($url);
            //echo $result . "\n";
        }
    }
}

$total = 100000;

$ts1 = microtime(true);
multiRegexp($urls, $total);
$ts2 = microtime(true);
echo "multiRegexp() x{$total} is " . ($ts2 - $ts1) . " sec\n";

$ts1 = microtime(true);
multiParseUrl($urls, $total);
$ts2 = microtime(true);
echo "multiParseUrl() x{$total} is " . ($ts2 - $ts1) . " sec\n";

$ts1 = microtime(true);
multiDummy($urls, $total);
$ts2 = microtime(true);
echo "multiDummy() x{$total} is " . ($ts2 - $ts1) . " sec\n";