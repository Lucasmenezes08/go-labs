package iteration



func Repeat (character string, numRepetitions int) string{
	var resultado string
	for i := 0; i < numRepetitions; i ++ {
		resultado = resultado + character;
	} 
	return resultado
}