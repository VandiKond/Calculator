package calc_service

import (
	"github.com/vandi37/Calculator/pkg/calc"
	"github.com/vandi37/Calculator/pkg/logger"
)

type Calculator struct {
	logger *logger.Logger
	DoLog  bool
}

func New(logger *logger.Logger) *Calculator {
	return &Calculator{logger: logger}
}

func (c *Calculator) Run(expression string) (float64, error) {
	res, err := calc.Calc(expression)
	if err != nil {
		if c.DoLog && c.logger != nil {
			c.logger.Printf("expression %s failed with error %s", expression, err.Error())
		}
		return 0, err
	}
	if c.DoLog && c.logger != nil {
		c.logger.Printf("expression %s resulted to %.4f", expression, res)
	}
	return res, err
}
