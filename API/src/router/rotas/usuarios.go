package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarTodosOsUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarioPorID,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuarioPorID,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioID}",
		Metodo: http.MethodDelete,
		Funcao: controllers.RemoverUsuarioPorID,
		RequerAutenticacao: false,
	},
}