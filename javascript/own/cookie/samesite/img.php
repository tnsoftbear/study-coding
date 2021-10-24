<?php
$checkme = $_COOKIE["checkme"] ?? null;
if ($checkme) {
    $name = __DIR__ . '/cookie.png';
    $fp = fopen($name, 'rb');
    header("content-type: image/jpeg");
    header("Content-Length: " . filesize($name));
    fpassthru($fp);
    exit;
} else {
    http_response_code(403);
}