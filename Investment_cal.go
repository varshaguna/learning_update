package main
import(
	"fmt"
    "math"
)
const inflationRate = 2.5
func main(){
	var  investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	outputText("investmentAmount: ")
	fmt.Scan(&investmentAmount)

	outputText("expectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("years: ")
	fmt.Scan(&years)

    futureValue, FutureRealValue := CalculateFutureValues(investmentAmount, expectedReturnRate, years)

    formatedFV := fmt.Sprintf("futureValue: %.1f\n", futureValue)
    formattedRFV := fmt.Sprintf("Future Value Adjusted : %.1f\n",FutureRealValue)

    fmt.Print(formatedFV,formattedRFV)
}
func outputText(text string){
	fmt.Print(text)
}

func CalculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64,float64){
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv/math.Pow(1+inflationRate/100,years)
	return fv,rfv
}
