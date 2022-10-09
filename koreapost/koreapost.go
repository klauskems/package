package koreapost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Psend(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다. \n", parcel)
}
