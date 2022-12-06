package usecases

import (
	"desafio-nu/core/domain"
	"desafio-nu/core/usecases/mocks"
	"desafio-nu/helpers"
	"testing"
)

func TestCalcCapitalGain(t *testing.T) {
	useCase := NewOperationUseCase()

	t.Run("Test Case 1, expected value [{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 0.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase1())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 2, expected value [{\"tax\": 0.00},{\"tax\": 10000.00},{\"tax\": 0.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase2())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 10000.00},{"tax": 0.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 3, expected value [{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 1000.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase3())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 0.00},{"tax": 1000.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 4, expected value [{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 0.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase4())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 5, expected value [{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 10000.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase5())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 10000.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 6, expected value [{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 0.00},{\"tax\": 3000.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase6())
			expectedValue := domain.StringToListFeeResponse(`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 3000.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 7, expected value [{\"tax\":0.00}, {\"tax\":0.00}, {\"tax\":0.00}, {\"tax\":0.00}, {\"tax\":3000.00},{\"tax\":0.00}, {\"tax\":0.00}, {\"tax\":3700.00}, {\"tax\":0.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase7())
			expectedValue := domain.StringToListFeeResponse(`[{"tax":0.00}, {"tax":0.00}, {"tax":0.00}, {"tax":0.00}, {"tax":3000.00},{"tax":0.00}, {"tax":0.00}, {"tax":3700.00}, {"tax":0.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})

	t.Run("Test Case 8, expected value [{\"tax\":0.00},{\"tax\":80000.00},{\"tax\":0.00},{\"tax\":60000.00}]",
		func(t *testing.T) {
			case1Results, _ := useCase.CalcCapitalGain(mocks.MockCase8())
			expectedValue := domain.StringToListFeeResponse(`[{"tax":0.00},{"tax":80000.00},{"tax":0.00},{"tax":60000.00}]`)

			assertEquals(expectedValue, case1Results, t)
		})
}

func assertEquals(expected []domain.FeeResponse, results []domain.FeeResponse, t *testing.T) {
	if len(expected) != len(results) {
		t.Fatalf("\n Expected value: %s not equals to: %s", helpers.ToJsonString(expected), helpers.ToJsonString(results))

	}
	for i, fee := range results {
		if !(fee.Tax == expected[i].Tax) {
			t.Fatalf("\n Expected value: %s not equals to: %s", helpers.ToJsonString(expected), helpers.ToJsonString(results))
		}
	}
}
