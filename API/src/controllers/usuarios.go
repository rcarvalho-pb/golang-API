package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioID)))

}

func BuscarTodosOsUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Mostrando todos os usuários:"))
}

func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Mostrando usuário por ID:"))
}

func AtualizarUsuarioPorID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Alterar usuario por ID."))
}

func RemoverUsuarioPorID(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Remover usuario por ID."))
}
