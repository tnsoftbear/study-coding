@main def forExample1: Unit =
  for {
    a <- 1 to 10
    b <- 1 to 5
    if a == b + 1
    z = a * b
  } yield println(s"$a * $b = $z")
