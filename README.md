
<h1 align="center">
   Projeto rascunho para treinar
</h1>
<p align="center">
    <em>
    Projeto criado com o propósito de aprender a linguagem Go!
    </em>
</p>

---

Sumário
=================

   * [O projeto](#o-projeto)
   * [Swagger](#swagger)
   * [Especificações](#especificações)
   * [Makefile](#makefile)
   * [Rodar Local](#rodar-local)
   * [Testes](#testes)
   * [Requisitos](#requisitos)

---

## O projeto

A ideia é fazer um protótipo de rede social e ir trabalhando em cima disso. O sistema é baseado em 3 areas principais: 

1. Usuários.
2. Seguidores
3. Publicações.

---

## Swagger (OpenAPI)

TODO

---

## Especificações

API REST

_Obs._: API roda em: http://localhost:8000/

---

## Makefile

TODO

---

## Rodar Local

1. go run main.go

---

## Testes

<em>
    Inicialmente feito dentro de models para fins de aprendizado
</em>

1. go test -v
2. go test --coverprofile cobertura.txt
3. go tool cover --func=cobertura.txt 
4. go tool cover --html=cobertura.txt

---

## Requisitos

* **GO**: 1.20
