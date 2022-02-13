<?php

require __DIR__ . '/../vendor/autoload.php';

use Revolt\EventLoop;

// Register a callback we'll disable
$callbackIdToDisable = EventLoop::delay(1, function (): void {
    echo "I'll never execute in one second because: disable()\n";
});

// Register a callback to perform the disable() operation
EventLoop::delay(0.5, function () use ($callbackIdToDisable) {
    echo "Disabling callback: ", $callbackIdToDisable, "\n";
    EventLoop::disable($callbackIdToDisable);
});

EventLoop::run();