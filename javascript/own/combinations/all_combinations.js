var input = ['a', 'b', 'c'];
var results = [];

function permute(arr, memo, rec_no = 0) {
    console.log("rec_no:", rec_no, "arr:", arr, "memo:", memo);
    var cur;
    var memo = memo || [];
    for (var i = 0; i < arr.length; i++) {
        cur = arr.splice(i, 1);
        if (arr.length === 0) {
            results.push(memo.concat(cur));
            console.log("results:", results);
        }
        permute(arr.slice(), memo.concat(cur), rec_no + 1);
        console.log("a] rec_no:", rec_no, " arr:", arr, "cur:", cur, "i:", i);
        arr.splice(i, 0, cur[0]);
        console.log("b] rec_no:", rec_no, " arr:", arr);
    }
    return results;
}

permute(input);

console.log(results);