#include <iostream>

#define ASSERT_TRUE(condition) \
    if (!(condition)) { \
        std::cerr << "Assertion failed: " << #condition << " is not true, in file " << __FILE__ << " at line " << __LINE__ << std::endl; \
    }

#define ASSERT_FALSE(condition) ASSERT_TRUE(!(condition))
#define ASSERT_EQ(x, y) ASSERT_TRUE((x) == (y))
#define ASSERT_NE(x, y) ASSERT_TRUE((x) != (y))