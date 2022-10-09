package ragular

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel \n", parcel)
}

type KoreaPost struct {
}

func (k *KoreaPost) PSend(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다. \n", parcel)
}
