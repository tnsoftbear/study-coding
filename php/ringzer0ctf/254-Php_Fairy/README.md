# ringzer0ctf 254 challange

[254: PHP Fairy](https://ringzer0ctf.com/challenges/254)

`md5("admin1674227342");` results with `'0e462097431906509019562988736854'`. In loose comparison it becomes `0.0` value, because of type juggling, so we can pass `'00000000000000000000000000000000'` value to match the `strlen()` condition.
