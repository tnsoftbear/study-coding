@main def forExample2: Unit =
  (1 to 30 by 3).flatMap { a =>
    (1 to 30 by 2).filter(b => a == b * 2)
      .map(b => println(s"$a * $b = " + (a * b)))
  }