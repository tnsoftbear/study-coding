#include "MathFunctions.h"

#ifdef USE_MYMATH
  #include "mysqrt.h"
#else
  #include <cmath>
#endif

namespace mathfunctions {
double sqrt(double x)
{
// which square root function should we use?
#ifdef USE_MYMATH
  return detail::mysqrt(x);
#else
  return std::sqrt(x);
#endif
}
}
