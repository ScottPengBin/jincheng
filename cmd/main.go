package main

func main() {

	app := InitApp()

	sig := <-app.Start()

	defer func() {

		app.Logger.Info("收到信号:", sig)

		err := app.Stop()

		if err != nil {
			panic(err)
		}
	}()

}
