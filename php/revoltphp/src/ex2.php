<?php

require __DIR__ . '/../vendor/autoload.php';

use Revolt\EventLoop;

if (\stream_set_blocking(STDIN, false) !== true) {
    \fwrite(STDERR, "Unable to set STDIN to non-blocking" . PHP_EOL);
    exit(1);
}

print "Write something and hit enter" . PHP_EOL;

$suspension = EventLoop::createSuspension();

$readableId = EventLoop::onReadable(STDIN, function ($id, $stream) use ($suspension): void {
    EventLoop::cancel($id);

    $chunk = \fread($stream, 8192);

    print "Read " . \strlen($chunk) . " bytes" . PHP_EOL;

    $suspension->resume(null);
});

$timeoutId = EventLoop::delay(5, function () use ($readableId, $suspension) {
    EventLoop::cancel($readableId);
    
    print "Timeout reached" . PHP_EOL;

    $suspension->resume(null);
});

$suspension->suspend();

EventLoop::cancel($readableId);
EventLoop::cancel($timeoutId);