package conf

import "gopkg.in/olivere/elastic.v5"

//初始化
func init() {

}

func NewEsClient(addr string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(addr))
	if err != nil {
		return nil, err
	}
	return client, nil
}
