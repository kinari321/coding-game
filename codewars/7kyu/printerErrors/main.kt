// My Answer
fun printerError(s: String): String {
    var count = 0
    s.forEach { c ->
        when(c) {
            'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z' -> count++
        }
    }
    return "$count/${s.length}"
}

// Other Answer 1
fun printerError1(s: String) = "${s.count { it !in 'a'..'m' }}/${s.length}"

// Other Answer 2
fun printerError2(s: String): String {
    val errors = s.count { it !in 'a'..'m' }
    val length = s.length
    return "$errors/$length"
}

// Other Answer 3
fun printerError3(s: String): String {
    return "${s.count{it > 'm'}}/${s.length}"
}