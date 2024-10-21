package model

type Campaign struct {
	Remetente     Mail     `json:"remetente"`
	Mensagem      Message  `json:"mensagem"`
	Destinatarios []string `json:"destinatarios"`
}
