# cdek_calculator
CDEK Calculator. Calculation at affordable rates.

 - example:

 ```
    token := "your token"

	from := cdekcalculator.Location{
		Address: "Россия, г. Москва, Cлавянский бульвар д.1",
	}
	to := cdekcalculator.Location{
		Address: "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43",
	}
	size := cdekcalculator.Size{
		Weight: 1286,
		Height: 50,
		Width:  50,
		Length: 50,
	}

    client := cdekcalculator.NewCDEKClient(token, true)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := client.Calculate(ctx, from, to, size)
    if err != nil{
        fmt.Fatalln(err)
    }
    for _,v := range res {
        ...
    }
```
