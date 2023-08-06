package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	sourl_reapers := []string{"ichigo","zaraki","byakuya","kenpachi","toshiro","renji","rukia","orihime","ulquiorra","grimmjow","aizen","gin","tosen","yamamoto","shunsui","jushiro","soi fon","yoruichi","mayuri","nemu","kisuke","yoruichi","kukaku","ururu","jinta","kon","ichigo","zaraki","byakuya","kenpachi","toshiro","renji","rukia","orihime","ulquiorra","grimmjow","aizen","gin","tosen","yamamoto","shunsui","jushiro","soi fon","yoruichi","mayuri","nemu","kisuke","yoruichi","kukaku","ururu","jinta","kon"}
	for _, sourl_reaper := range sourl_reapers {
		// attack(sourl_reaper)
		smoke_signal := make(chan bool)
		go attackChannel(sourl_reaper, smoke_signal)
		fmt.Println(<-smoke_signal)
	}
}

func attackChannel(sourl_reaper string, attacked chan bool) {
	fmt.Println("Soul Reaper: " + sourl_reaper + " is attacking!")
	attacked <- true
}
