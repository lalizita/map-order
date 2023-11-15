package main

import (
	"fmt"
	"sort"
)

func main() {
	heights := []int{8, 3, 6, 1}
	names := []string{"alice", "bruno", "renan", "ivana"}
	peopleMap := make(map[int]string)

	//criando um map com as pessoas e alturas
	for k, v := range heights {
		peopleMap[v] = names[k]
	}

	s := sortMap(peopleMap)
	//!problema: o map está ordenado do menor pro maior
	fmt.Println("sorted", s)
}

func sortMap(m map[int]string) map[int]string {
	//um slice para armazenar apenas as chaves
	keys := []int{}
	//map para armazenar os itens em uma nova ordem
	sorted := make(map[int]string)
	//popula o slice de chaves com as chaves do map recebido
	for k, _ := range m {
		keys = append(keys, k)
	}

	//ordena do maior para o menor
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	fmt.Println(keys)

	for _, k := range keys {
		//popula novo map a partir das chaves ordenadas
		sorted[k] = m[k]
		fmt.Println(sorted[k])
	}

	//!problema: o map volta a ordenação padrão do menor para o maior
	fmt.Println("slice", sorted)
	return sorted
}
