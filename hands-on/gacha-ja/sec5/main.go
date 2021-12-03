// STEP01: $B%,%A%c$N7k2L$r%U%!%$%k$KJ]B8$7$h$&(B

package main

import (
	"fmt"
	"os"
	
	"github.com/Jakey08/gacha_package/gacha"
	
)

func main() {
	p := gacha.NewPlayer(10, 100)


	n := inputN(p)
	results, summary := gacha.DrawN(p, n)

	saveResults(results)
	saveSummary(summary)
}

func inputN(p *gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("$B%,%A%c$r0z$/2s?t$rF~NO$7$F$/$@$5$$!J:GBg(B:%d$B2s!K(B\n", max)

	var n int
	for {
		fmt.Print("$B%,%A%c$r0z$/2s?t(B>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1$B0J>e(B%d$B0J2<$N?t$rF~NO$7$F$/$@$5$$(B\n", max)
	}

	return n
}

func saveResults(results []*gacha.Card) {
	// TODO: results.txt$B$H$$$&%U%!%$%k$r:n@.$9$k(B
	f, err := os.Create("results.txt")
	

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, result := range results {
		// TODO: fmt.Fprintln$B4X?t$r;H$C$F(Bresult$B$r%U%!%$%k$K=PNO$9$k(B
	fmt.Fprintln(f, result)
	}
}

func saveSummary(summary map[gacha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		// TODO: $B%U%!%$%k$rJD$8$k(B
		// $B%(%i!<H/@8$7$?>l9g$O(Bfmt.Println$B4X?t$G=PNO$9$k(B
		if err := df.Close(); err != nil{
			fmt.Println(err)
		}
	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}

