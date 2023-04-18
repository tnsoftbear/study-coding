package main

import (
    "github.com/apple/foundationdb/bindings/go/src/fdb"
    "log"
    "fmt"
)

func main() {
    // Different API versions may expose different runtime behaviors.
    fdb.MustAPIVersion(720)

    // Open the default database from the system cluster
    db := fdb.MustOpenDefault()

    // Database reads and writes happen inside transactions
    ret, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {
        tr.Set(fdb.Key("hello"), []byte("world"))
        return tr.Get(fdb.Key("foo")).MustGet(), nil
        // db.Transact automatically commits (and if necessary,
        // retries) the transaction
    })
    if e != nil {
        log.Fatalf("Unable to perform FDB transaction (%v)", e)
    }

    fmt.Printf("hello is now world, foo was: %s\n", string(ret.([]byte)))
}
// package main

// import (
//     "fmt"
//     "github.com/apple/foundationdb/bindings/go/src/fdb"
//     "github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
// )

// func main() {
//     // Open a connection to the database
//     db := fdb.MustOpenDefault()

//     // Create a transaction
//     tr, err := db.CreateTransaction()
//     if err != nil {
//         panic(err)
//     }

//     // Write a key-value pair to the database
//     key := tuple.Tuple{"users", "johndoe", "email"}
//     value := []byte("johndoe@example.com")
//     tr.Set(key.Pack(), value)

//     // Commit the transaction
//     if err := tr.Commit().Get(); err != nil {
//         panic(err)
//     }

//     // Read the value of a key from the database
//     tr, err = db.CreateTransaction()
//     if err != nil {
//         panic(err)
//     }

//     res, err := tr.Get(key.Pack()).Get()
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println(tuple.Unpack(key.Unpack(res)))
// }