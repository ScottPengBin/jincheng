package main

func main() {

	app := InitApp()

	defer func() {

		sig := <-app.Start()

		app.Logger.Info("收到信号:", sig)

		err := app.Stop()

		if err != nil {
			panic(err)
		}
	}()

}
