package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando usuário!"))
	
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
