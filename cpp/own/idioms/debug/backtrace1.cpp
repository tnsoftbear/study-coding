// #include <libunwind.h>
// void PrintStackTrace() {
//   unw_cursor_t cursor;
//   unw_context_t context;
//   unw_getcontext(&context);
//   unw_init_local(&cursor, &context);

//   while (unw_step(&cursor) > 0) {
//       unw_word_t offset, pc;
//       char funcName[256];
//       unw_get_reg(&cursor, UNW_REG_IP, &pc);
//       unw_get_proc_name(&cursor, funcName, sizeof(funcName), &offset);
//       std::cout << "Function: " << funcName << " (0x" << std::hex << pc << std::dec << ")\n";
//   }
// }

#include <execinfo.h> /* backtrace, backtrace_symbols_fd */
#include <unistd.h> /* STDOUT_FILENO */

void PrintStackTrace() {
    size_t size;
    enum Constexpr { MAX_SIZE = 1024 };
    void *array[MAX_SIZE];
    size = backtrace(array, MAX_SIZE);
    backtrace_symbols_fd(array, size, STDOUT_FILENO);
}