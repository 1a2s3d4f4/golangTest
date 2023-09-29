/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/


package main




import (
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"os"
)



func main() {

    myscanner := bufio.NewScanner(os.Stdin) // принимаем строку
    myscanner.Scan()
    line := myscanner.Text()
    
    // fmt.Println(line) - проверка работы строки 30-32
    
 
    
    
    //7?????????????????????????????????Почему так не принимается строка????????????????????????????????????
    /*
    var arithmeticExpression string
    fmt.Scanln(&arithmeticExpression) // принимаем математическое выражения
    */
    
    
     
   linePart := strings.Split(line, " ") // делим принятую строку на составные части по пробелу
   
  
   /*
    fmt.Println(linePart[0]) //********************************** проверка строчки 42****************************************
    fmt.Println(linePart[1])
    fmt.Println(linePart[2])
    */
    
    /*
   fmt.Println(numberType (linePart[0])) //************************проверка функции определения типа числа************************
   fmt.Println(numberType (linePart[1])) 
   fmt.Println(numberType (linePart[2])) 
   */
   

    
if len(linePart) <3 {  // проверяем, что бы было не меньше 2 операндов и 1 го оператора
    fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
    os.Exit(0)
}
if len(linePart) >3 { // проверяем что бы было не больше 2 операндов и 1 оператора
    fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
    os.Exit(0)
}
 if   numberType(linePart[0]) != numberType(linePart[2]){  //проверяем что бы числа были одного типа
    fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
     os.Exit(0)
 }
  
  
   switch{
       
    // Вариант ввода арабских цифр
    case numberType(linePart[0]) == "Arabic": 
    fmt.Println(calculation(linePart))
    
    // Вариант ввода римских цифр
    case numberType(linePart[2]) == "Roman":
    
    //Переводим цифры из Римской системы в арабские
    linePartRoman:= [] string {romanToArabic(linePart[0]),linePart[1], romanToArabic(linePart[2])}
    //Делаем расчет выражения
    result:= calculation(linePartRoman)
    // Ловим ошибка - отрицательный результат в римской системе
    if result<0 {
    fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
    } else {fmt.Println(arabicToRoman(result))
    } 
    }
    
    
    //*********************************************Проверка функции расчетов**********************************
   // a:= []string {"5","+","7"}
   //fmt.Println(calculation(a)) 
    
    
   
}

func numberType (str string) string { //функция определяет к кокому типу относится число
    var typeNumber string
    if str[0]>=48 && str[0]<=57 {
        typeNumber="Arabic"
    }else{ 
       typeNumber="Roman"
    }
    
    return typeNumber
}  

func calculation(a[]string) int { // функция производит математические вычисления
    var result int
    
    switch a[1]{
        case "+": result=pars(a[0])+pars(a[2])
        case "/": result=pars(a[0])/pars(a[2])
        case "*": result=pars(a[0])*pars(a[2])
        case "-": result=pars(a[0])-pars(a[2])
}  
return result
}

func pars(a string) int{ // функция переводит строку в число
    res, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return res
}

func romanToArabic(romanNum string) string {//функция переводит римские цифры в арабские
	result := 0
	romanToArabic:= map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,

}
	for i := 0; i < len(romanNum); i++ {
		current := romanToArabic[string(romanNum[i])]
		if i+1 < len(romanNum) {
			next := romanToArabic[string(romanNum[i+1])]
			if current < next {
				result -= current
			} else {
				result += current
			}
		} else {
			result += current
		}
	}
	
	return strconv.Itoa(result)
}
func arabicToRoman(result int) string {//функция переводит арабские цифры в римские
var arabicToRoman = map[int]string{
	1: "I",
	4: "IV",
	5: "V",
	9: "IX",
	10: "X",
	40: "XL",
	50: "L",
	90: "XC",
	100: "C",
	
}
var romanNum string
	for result > 0 {
		for _, key := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
			if result >= key {
				result -= key
				romanNum += arabicToRoman[key]
				break
			}
		}
	}
	return romanNum
}

