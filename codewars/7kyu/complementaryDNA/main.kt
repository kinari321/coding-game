// My Answer
package dna

fun makeComplement(input : String) : String {
    val dna = StringBuilder()
    for (c in input) {
        dna.append(when (c) {
            'A' -> 'T'
            'T' -> 'A'
            'C' -> 'G'
            'G' -> 'C'
            else -> throw IllegalArgumentException("Invalid DNA sequence")
        })
    }
    return dna.toString()
}

// 1. Other Answer
fun makeComplement(dna: String) = dna.map { when(it) {
  'A' -> 'T'
  'T' -> 'A'
  'C' -> 'G'
  'G' -> 'C'
  else -> it
} }.joinToString("")

// 2. Other Answer
fun makeComplement(dna : String) : String = dna.fold(String()) { acc, nucleotide ->
  acc + when (nucleotide) {
      'A' -> 'T'
      'C' -> 'G'
      'G' -> 'C'
      'T' -> 'A'
      else -> throw IllegalArgumentException("DNA may only contain nucleotides in A, C, G and T")
  }
}

// 3. Other Answer
fun makeComplement(dna : String)=dna.map{"ATGC".zip("TACG").toMap()[it]?:it}.joinToString("")
