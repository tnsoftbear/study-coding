#include <iostream>
#include <chrono>

/**
 * C++: Matrix multiplication optimization example 
 * that considers more cache hitting, 
 * because processed data is located in the same cache line, 
 * so it has better locality.
 */

static const size_t kN = 1000;

int a[kN][kN], b[kN][kN], c[kN][kN];

void MultiplyNaive() {
    for (size_t i = 0; i < kN; ++i) {
        for (size_t j = 0; j < kN; ++j) {
            c[i][j] = 0;
            for (size_t k = 0; k < kN; ++k) {
                c[i][j] += a[i][k] * b[k][j];
            }
        }
    }
}

/**
 * Здесь мы сначала переворачиваем матрицу B на 90 градусов,
 * а затем перемножаем ее с матрицей A, учитывая переворот.
 * Это позволяет нам не читать так часто по столбцам, а читать чаще по строкам.
 * Т.к. процессор забирает в кэш-линию 64 байт локальных данных, то кэш-попадание случается чаще.
 * Я вижу разницу скорости вычисления в 3 раза на 2х-ядерном CPU.
 */
void MultiplyCacheAware() {
    // Transpose B
    for (size_t i = 0; i < kN; ++i) {
        for (size_t j = 0; j < kN; ++j) {
            // b[j][i] = b[i][j];
            std::swap(b[i][j], b[j][i]);
        }
    }
    // Multiply
    for (size_t i = 0; i < kN; ++i) {
        for (size_t j = 0; j < kN; ++j) {
            c[i][j] = 0;
            for (size_t k = 0; k < kN; ++k) {
                c[i][j] += a[i][k] * b[j][k]; // Различие здесь
            }
        }
    }
}

void GenerateInput() {
    for (size_t i = 0; i < kN; ++i) {
        for (size_t j = 0; j < kN; ++j) {
            a[i][j] = i + j;
            b[i][j] = i * j;
        }
    }
}

size_t ComputeOutputDigest() {
    size_t sum = 0;
    for (size_t i = 0; i < kN; ++i) {
        for (size_t j = 0; j < kN; ++j) {
            sum += c[i][j];
        }
    }
    return sum;
}

int main() {
    GenerateInput();
    auto start = std::chrono::steady_clock::now();
    MultiplyNaive();
    auto elapsed = std::chrono::steady_clock::now() - start;
    std::cout << "Digest: " << ComputeOutputDigest() << std::endl;
    std::cout << "Elapsed for naive multiply: "
        << std::chrono::duration_cast<std::chrono::milliseconds>(elapsed).count()
        << "ms" << std::endl;
 
    GenerateInput();
    start = std::chrono::steady_clock::now();
    MultiplyCacheAware();
    elapsed = std::chrono::steady_clock::now() - start;
    std::cout << "Digest: " << ComputeOutputDigest() << std::endl;
    std::cout << "Elapsed for cache awared multiply: "
        << std::chrono::duration_cast<std::chrono::milliseconds>(elapsed).count()
        << "ms" << std::endl;
}