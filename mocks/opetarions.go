package mocks

import "desafio-nu/core/domain"
import "encoding/json"

func MockCase1() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 100},{"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]`
	return jsonStringToOper(jsonStr)
}

func MockCase2() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000},{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase3() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":5.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 3000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase4() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"buy", "unit-cost":25.00, "quantity": 5000},{"operation":"sell", "unit-cost":15.00, "quantity": 10000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase5() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"buy", "unit-cost":25.00, "quantity": 5000},{"operation":"sell", "unit-cost":15.00, "quantity": 10000},{"operation":"sell", "unit-cost":25.00, "quantity": 5000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase6() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":2.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":25.00, "quantity": 1000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase7() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":2.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":25.00, "quantity": 1000},{"operation":"buy", "unit-cost":20.00, "quantity": 10000},{"operation":"sell", "unit-cost":15.00, "quantity": 5000},{"operation":"sell", "unit-cost":30.00, "quantity": 4350},{"operation":"sell", "unit-cost":30.00, "quantity": 650}]`
	return jsonStringToOper(jsonStr)
}

func MockCase8() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":50.00, "quantity": 10000},{"operation":"buy", "unit-cost":20.00, "quantity": 10000},{"operation":"sell", "unit-cost":50.00, "quantity": 10000}]`
	return jsonStringToOper(jsonStr)
}

func MockCase9() []*domain.Oper {
	jsonStr := `[{"operation":"buy", "unit-cost":10, "quantity": 10000},{"operation":"sell", "unit-cost":20, "quantity": 11000}]`
	return jsonStringToOper(jsonStr)
}

func jsonStringToOper(jsonStr string) []*domain.Oper {
	var m []*domain.Oper
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	return m
}
