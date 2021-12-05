package main

/////////////////////
// モック対象の実装
////////////////////
type ApiClient interface {
	Request(string) (string, error)
}

type DataRegister struct {
	client ApiClient // インターフェイスに依存しているだけで実装は存在しない
}

func (d *DataRegister) Register(data string) (string, error) {
	result, err := d.client.Request(data)
	if err != nil {
		return "", err
	}
	return result, nil
}
