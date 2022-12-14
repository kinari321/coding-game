// My Answer
fun isLeapYear(year: Int) : Boolean {
  if (year % 400 == 0) return true
  else if (year % 100 ==0 ) return false
  else if (year % 4 == 0) return true
  return false
}

// 1. Other Answer
fun isLeapYear(year: Int) = java.time.Year.of(year).isLeap

// 2. Other Answer
fun isLeapYear(year: Int) = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)

// 3. Other Answer
fun isLeapYear(year: Int) : Boolean {
  return when {
     year % 400 == 0 -> true
     year % 100 == 0 -> false
     else -> year % 4 == 0
 }
}