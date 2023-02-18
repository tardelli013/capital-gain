package usecases

import (
	"desafio-nu/helpers"
	"desafio-nu/mocks"
	"testing"
)

func TestCalcCapitalGain(t *testing.T) {
	useCase := NewOperationUseCase()

	t.Run("Test Case 1, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":0}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase1())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":0}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 2, expected value [{\"tax\":0},{\"tax\":10000},{\"tax\":0}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase2())
			expectedValue := `[{"tax":0},{"tax":10000},{"tax":0}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 3, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":1000}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase3())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":1000}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 4, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":0}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase4())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":0}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 5, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":0},{\"tax\":10000}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase5())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":0},{"tax":10000}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 6, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":0},{\"tax\":0},{\"tax\":3000}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase6())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 7, expected value [{\"tax\":0},{\"tax\":0},{\"tax\":0},{\"tax\":0},{\"tax\":3000},{\"tax\":0},{\"tax\":0},{\"tax\":3700},{\"tax\":0}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase7())
			expectedValue := `[{"tax":0},{"tax":0},{"tax":0},{"tax":0},{"tax":3000},{"tax":0},{"tax":0},{"tax":3700},{"tax":0}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 8, expected value [{\"tax\":0},{\"tax\":80000},{\"tax\":0},{\"tax\":60000}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase8())
			expectedValue := `[{"tax":0},{"tax":80000},{"tax":0},{"tax":60000}]`

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run(`Test Case 9, expected value [{"operation":"buy", "unit-cost":10, "quantity":10000},{"operation":"sell", "unit-cost":20, "quantity":11000}]`,
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase9())
			expectedValue := `[{"tax":0},{"error":"Can't sell more stocks than you have"}]`

			assertEquals(expectedValue, case1Results, t)
		})
}

func assertEquals(expected string, results string, t *testing.T) {
	if len(expected) != len(results) {
		t.Fatalf("\n Expected value: %s not equals to: %s", helpers.ToJsonString(expected), helpers.ToJsonString(results))
	}

	if expected != results {
		t.Fatalf("\n Expected value: %s not equals to: %s", helpers.ToJsonString(expected), helpers.ToJsonString(results))
	}
}
