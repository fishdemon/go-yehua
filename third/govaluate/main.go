/**
 非常好用的一款表达式解析器
https://github.com/Knetic/govaluate
 */
package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {

	expression, _ := govaluate.NewEvaluableExpression("10 > 0")
	result, _ := expression.Evaluate(nil)
	fmt.Println(result)

	expression, _ = govaluate.NewEvaluableExpression("10>0")
	result, _ = expression.Evaluate(nil)
	fmt.Println(result)

	expression, _ = govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1
	result, _ = expression.Evaluate(parameters)

	expression, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	parameters = make(map[string]interface{}, 8)
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80
	result, _ = expression.Evaluate(parameters)

	expression, _ = govaluate.NewEvaluableExpression("http_response_body == 'service is ok'")
	parameters = make(map[string]interface{}, 8)
	parameters["http_response_body"] = "service is ok"
	result, _ = expression.Evaluate(parameters)

	expression, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	parameters = make(map[string]interface{}, 8)
	parameters["total_mem"] = 1024;
	parameters["mem_used"] = 512;
	result, _ = expression.Evaluate(parameters);

	expression, _ = govaluate.NewEvaluableExpression("'2014-01-02' > '2014-01-01 23:59:59'")
	result, _ = expression.Evaluate(nil)

	expression, _ = govaluate.NewEvaluableExpression("response_time <= 100")
	parameters = make(map[string]interface{}, 8)
	for {
		parameters["response_time"] = pingSomething()
		result, _ = expression.Evaluate(parameters)
	}

}

func pingSomething() int {
	return 32
}