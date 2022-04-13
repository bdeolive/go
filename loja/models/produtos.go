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

	findAllProdutosSelect, err := db.Query("SELECT * FROM produtos ORDER BY descricao DESC")

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

func Delete(id string) {
	db := db.ConectaDB()

	scriptDelete, err := db.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	scriptDelete.Exec(id)
	defer db.Close()
}

func Edit(id string) Produto {
	db := db.ConectaDB()

	produtoUpdate, err := db.Query("SELECT * FROM produtos WHERE id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoUpdate.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoUpdate.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()
	return produto
}

func Update(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConectaDB()

	scriptUpdate, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	scriptUpdate.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
