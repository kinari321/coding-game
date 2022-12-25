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
fun printerError1(s: String): String {
	return ""
}

// Other Answer 2
fun printerError2(s: String): String {
	return ""
}