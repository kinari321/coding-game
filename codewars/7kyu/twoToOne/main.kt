fun longest(s1:String, s2:String):String {
  val sb = StringBuilder()
  sb.append(s1)
  sb.append(s2)
  val chars: List<Char> = sb.toList()
  val noDuplicates = chars.toSet().toList()
  return noDuplicates.sorted().joinToString("")
}
