package cdekcalculator_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	cdekcalculator "github.com/sgkochnev/cdek_calculator"
)

func TestCalculate(t *testing.T) {

	token := "token"

	from := cdekcalculator.Location{
		Address: "Россия, г. Москва, Cлавянский бульвар д.1",
	}
	to := cdekcalculator.Location{
		Address: "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43",
	}
	size := cdekcalculator.Size{
		Weight: 1286,
		Height: 35,
		Width:  27,
		Length: 50,
	}

	t.Run("correct", func(t *testing.T) {
		client := cdekcalculator.NewCDEKClient(token, true)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		res, err := client.Calculate(ctx, from, to, size)
		if err != nil {
			t.Error(err)
		}
		if len(res) == 0 {
			t.Error("response is empty")
		}
		fmt.Println(res)
	})

	t.Run("incorrect", func(t *testing.T) {
		client := cdekcalculator.NewCDEKClient(token+"1", true)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		res, err := client.Calculate(ctx, from, to, size)
		if err == nil || len(res) != 0 {
			t.Error(res)
		}
	})

}
