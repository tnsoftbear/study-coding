#pragma once

struct S {
    int x;
#if defined(MYDEF)
    int y;
#endif
};