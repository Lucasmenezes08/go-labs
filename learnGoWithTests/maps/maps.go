package maps


type Dicionario map[string]string
type ErrDicionario string

const (
	ErrNaoEncontrado = ErrDicionario("Palavra não encontrada")	
	ErrPalavraExistente = ErrDicionario("não foi possivel adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("Nao foi possivel atualizar a palavra pois ela ja existe")
)

func (ed ErrDicionario)Error() string{
	return string(ed)
}

func (d Dicionario)Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao,nil
}

func (d Dicionario)Adiciona(palavra string, texto string)error{
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		d[palavra] = texto
	case nil:
		return ErrPalavraExistente
	default:
		return err 
	} 
	
	return nil
} 

func (d Dicionario)Atualiza(palavra, novaDefinicao string)error{
	 _, err := d.Busca(palavra)
    switch err {
    case ErrNaoEncontrado:
        return ErrPalavraInexistente
    case nil:
        d[palavra] = novaDefinicao
    default:
        return err

    }
	return nil
}

func (d Dicionario)Deleta(palavra string){
	delete(d,palavra)
}