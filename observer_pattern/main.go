package main

func main() {
	ws := initTemperatureData()
	tvDisplay := initDisplay("tv", ws)
	mobileDisplay := initDisplay("mobile", ws)
	ws.Add(tvDisplay)
	ws.Add(mobileDisplay)
	ws.addData()
	ws.Remove(tvDisplay)
	ws.addData()

}
