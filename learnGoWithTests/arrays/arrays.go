package arrays


func SomaArrays (numeros []int) int{
	var somatorio int
	for _, numeros := range numeros{
		somatorio += numeros
	}
	return somatorio
}	


func SomaTudo (numeros ...[]int) (somas []int){
	quantidadeSomar := len(numeros)
	resultado := make([]int, quantidadeSomar)
	
	for i, v := range numeros{
		resultado[i] += SomaArrays(v)
	}

	return resultado
}