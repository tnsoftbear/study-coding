int main() {
    int x = 1;
    int &a = x;
    int const &c = x + 1;
    int &&d = x + 1;
    // c += 1; // error: assignment of read-only reference ‘c’
    d += 1; // ok
}