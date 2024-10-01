int foo(int *p) {
    return *p;
};

// означает, что ф-ция не будет менять значения аргумента
int bar(const int *p) {
    return *p;
};

int buz(volatile int *p) {
    return *p;
};

int voo(const volatile int *p) {
    return *p;
}

int main() {
    int i = 1;
    const int ci = 2;
    volatile int vi = 3;
    foo(&i);
    // foo(&ci); // error: invalid conversion from ‘const int*’ to ‘int*’
    // foo(&vi); // error: invalid conversion from ‘volatile int*’ to ‘int*’
    
    bar(&i);
    bar(&ci);
    // bar(&vi); // error: invalid conversion from ‘volatile int*’ to ‘const int*’
    
    buz(&i);
    // buz(&ci); // error: invalid conversion from ‘const int*’ to ‘volatile int*’
    buz(&vi);

    voo(&i);
    voo(&ci);
    voo(&vi);
}