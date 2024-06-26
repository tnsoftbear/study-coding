#include <iostream>
#include <stack>
#include <queue>
#include <map>
#include <list>
#include <forward_list>
#include <set>
#include <algorithm>

int main() {
    // https://en.cppreference.com/w/cpp/container/stack
    std::stack<int> st; // sizeof(st): 80
    int stX, stY = 0;
    // stX = st.top(); // Segmentation fault
    st.push(1);
    st.push(2);
    st.push(3);
    st.pop(); // removes the top element (3)
    stX = st.top(); // 2
    st.emplace(4); // constructs element in-place at the top
    stY = st.top(); // 4
    std::cout << "Stack: stX: " << stX << "; stY: " << stY << "; st.size: " << st.size()
        << "; sizeof(st): " << sizeof(st) << std::endl;

    // https://en.cppreference.com/w/cpp/container/queue
    std::queue<int> q; // sizeof(q): 80
    q.push(1);
    q.push(2);
    q.push(3);
    q.push(4);
    q.pop(); // removes the first element (1)
    q.emplace(5); // constructs element in-place at the end
    std::cout << "Queue: q.size: " << q.size() << "; q.front: " << q.front() << "; q.back: " << q.back()
        << "; q.empty(): " << q.empty()
        << "; sizeof(q): " << sizeof(q) << std::endl;

    // https://en.cppreference.com/w/cpp/container/priority_queue
    // provides constant time lookup of the largest element, at the expense of logarithmic insertion and extraction.
    std::priority_queue<int> pq;

    std::vector<int> v(10); // sizeof(v): 24
    v[0] = 1;
    v[5] = 2;
    v.resize(13);
    v.push_back(3);
    v.push_back(4);
    v.pop_back();
    std::cout << "Vector: v.size: " << v.size() << "; v.capacity: " << v.capacity() 
        << "; v[5]: " << v[5] << "; v.front: " << v.front() << "; v.back: " << v.back()
        << "; sizeof(v): " << sizeof(v) << std::endl;
    for (int i = 0; i < v.size(); i++) {
        std::cout << v[i] << " ";
    }
    std::cout << std::endl;

    std::string str = "abc";
    str.pop_back();
    str.push_back('d');
    str.shrink_to_fit();
    std::cout << "String: str.size: " << str.size() << "; str.capacity: " << str.capacity()
        << "; sizeof(str): " << sizeof(str) << std::endl;
    
    // https://en.cppreference.com/w/cpp/container/map
    std::map<std::string, int> m; // sizeof(m): 24
    m["a"] = 1;
    m["b"] = 2;
    std::cout << "Map: m.size: " << m.size() << "; m[\"a\"]: " << m["a"] << "; m[\"b\"]: " << m["b"]
        << "; m[\"c\"]: " << m["c"] << "; m.count(\"d\"): " << m.count("d")
        // << "; m.contains(\"a\"): " << m.contains("a") // g++ -std=c++2a containers_overview.cpp
        << "; m.size: " << m.size() << "; sizeof(v): " << sizeof(v)
        << "; m.empty(): " << m.empty() << std::endl;

    std::list<int> l;
    std::forward_list<int> fl;

    // https://en.cppreference.com/w/cpp/container/set/
    struct Less {   // переопределим оператор сравнения
        bool operator()(const std::string& a, const std::string& b) const {
            return a > b;
        }
    };
    std::set<std::string, Less> set{"def", "jkl", "abc", "ghi", "xyz"};
    std::cout << "Set: ";
    std::for_each(set.cbegin(), set.cend(), [](const std::string& s) {
        std::cout << s << " ";
    });
    std::cout << std::endl;
    
}