// My Answer
fun findShort(s: String): Int? {
    var arr = s.split(" ")
    var ints = arr.map { it.length }
    return ints.minOrNull()
}

// Other Answer

fun findShort1(words: String): Int = words.split(" ").minOf{it.length}

