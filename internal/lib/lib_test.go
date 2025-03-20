package lib

import (
	"fmt"
	"testing"
)

func Test_printDateInfo(t *testing.T) {
	fmt.Println(`PrintDateInfo("daily"):`, PrintDateInfo("daily"))
	fmt.Println(`PrintDateInfo("weekly"):`, PrintDateInfo("weekly"))
	fmt.Println(`PrintDateInfo("monthly"):`, PrintDateInfo("monthly"))
}
