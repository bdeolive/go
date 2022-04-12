package models

import (
	"loja/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func FindAllProdutos() []Produto {
	db := db.ConectaDB()

	findAllProdutosSelect, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produtos := []Produto{}

	for findAllProdutosSelect.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = findAllProdutosSelect.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func InsertProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDB()

	scriptInsert, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	scriptInsert.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
