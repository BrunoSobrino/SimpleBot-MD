package helper

func Menu(pushName string, prefix string) string {
	return `*Hola ` + pushName + ` Aqui tienes la lista de comandos de SimpleBot-MD*
	
` + prefix + `sticker
` + prefix + `owner
` + prefix + `source
`
}
