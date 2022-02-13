<?php

require __DIR__ . '/../vendor/autoload.php';

use Revolt\EventLoop;

// Register a repeating timer callback
$callbackId = EventLoop::repeat(1, function(): void {
    echo "tick\n";
});

// Disable the callback
EventLoop::disable($callbackId);

EventLoop::defer(function () use ($callbackId): void {
    // Immediately enable the callback when the event loop starts
    EventLoop::enable($callbackId);
    // Now that it's enabled we'll see tick output in our console every second.
});

EventLoop::delay(5, function () use ($callbackId): void {
    EventLoop::cancel($callbackId);
});

$increment = 0;

EventLoop::repeat(0.1, function ($incrementCbId) use (&$increment): void {
    echo "increment-tick-{$increment}\n";
    if (++$increment >= 3) {
        EventLoop::cancel($incrementCbId); // <-- cancel myself!
    }
});

EventLoop::run();