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
		go attack(sourl_reaper)
	}

	time.Sleep(2 * time.Second)
}

func attack(sourl_reaper string) {
	fmt.Println("Soul Reaper: " + sourl_reaper + " is attacking!")
	time.Sleep(1 * time.Second)
}
