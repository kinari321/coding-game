package twotoone

// My Answer
fun longest(s1:String, s2:String):String {
  val sb = StringBuilder()
  sb.append(s1)
  sb.append(s2)
  val chars: List<Char> = sb.toList()
  val noDuplicates = chars.toSet().toList()
  return noDuplicates.sorted().joinToString("")
}

// 1. Other Answer
fun longest(s1:String, s2:String):String {
  return (s1 + s2).toSortedSet().joinToString("")
}

// 2. Other Answer
fun longest(s1:String, s2:String) = (s1 + s2).toSortedSet().joinToString("")

// 3. Other Answer
fun longest(s1:String, s2:String):String {
  return "abcdefghijklmnopqrstuvwxyz".filter { it in s1 || it in s2 }
}