package handlers

import (
	"html/template"
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

type Patient struct {
	ID            int
	Nome          string
	CPF           string
	Idade         int
	Email         string
	Ocorrência    string
	TipoSanguineo string
	Sintomas      string
	Diagnostico   string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(utils.GetTemplatePath("static/index.html"))
	if err != nil {
		http.Error(w, "Erro ao carregar página inicial: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles(utils.GetTemplatePath("static/forms/createPatient.html"))
		if err != nil {
			http.Error(w, "Erro ao carregar formulário", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		cpf := r.FormValue("cpf")
		idade, _ := strconv.Atoi(r.FormValue("idade"))
		email := r.FormValue("email")
		ocorrencia := r.FormValue("ocorrencia")
		tipoSanguineo := r.FormValue("tipoSanguineo")
		sintomas := r.FormValue("sintomas")
		diagnostico := r.FormValue("diagnostico")

		// Corrigido para bater com os campos da Struct (Email, Ocorrência, TipoSanguineo)
		_, err := utils.DB.Exec(`
			INSERT INTO patients (nome, cpf, idade, email, ocorrencia, tipo_sanguineo, sintomas, diagnostico)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, nome, cpf, idade, email, ocorrencia, tipoSanguineo, sintomas, diagnostico)

		if err != nil {
			http.Error(w, "Erro ao cadastrar paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/patients", http.StatusSeeOther)
	}
}

func ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	// Query atualizada para trazer os campos corretos do banco de dados
	rows, err := utils.DB.Query(`
		SELECT id, nome, cpf, idade, email, ocorrencia, tipo_sanguineo, sintomas, diagnostico
		FROM patients
		ORDER BY id
	`)

	if err != nil {
		http.Error(w, "Erro ao buscar pacientes: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var patients []Patient

	for rows.Next() {
		var p Patient

		// Corrigido o Scan para apontar para as variáveis corretas da Struct Patient
		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.CPF,
			&p.Idade,
			&p.Email,
			&p.Ocorrência,
			&p.TipoSanguineo,
			&p.Sintomas,
			&p.Diagnostico,
		)

		if err != nil {
			http.Error(w, "Erro ao ler paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		patients = append(patients, p)
	}

	tmpl, err := template.ParseFiles(utils.GetTemplatePath("static/forms/listPatients.html"))
	if err != nil {
		http.Error(w, "Erro ao carregar lista: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, patients)
}

func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		var p Patient

		// Query e Scan atualizados com os campos reais da struct
		err := utils.DB.QueryRow(`
			SELECT id, nome, cpf, idade, email, ocorrencia, tipo_sanguineo, sintomas, diagnostico
			FROM patients
			WHERE id = $1
		`, id).Scan(
			&p.ID,
			&p.Nome,
			&p.CPF,
			&p.Idade,
			&p.Email,
			&p.Ocorrência,
			&p.TipoSanguineo,
			&p.Sintomas,
			&p.Diagnostico,
		)

		if err != nil {
			http.Error(w, "Paciente não encontrado: "+err.Error(), http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles(utils.GetTemplatePath("static/forms/updatePatient.html"))
		if err != nil {
			http.Error(w, "Erro ao carregar formulário de atualização: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, p)
		return
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		cpf := r.FormValue("cpf")
		idade, _ := strconv.Atoi(r.FormValue("idade"))
		email := r.FormValue("email")
		ocorrencia := r.FormValue("ocorrencia")
		tipoSanguineo := r.FormValue("tipoSanguineo")
		sintomas := r.FormValue("sintomas")
		diagnostico := r.FormValue("diagnostico")

		// Query de UPDATE ajustada
		_, err := utils.DB.Exec(`
			UPDATE patients
			SET nome = $1,
				cpf = $2,
				idade = $3,
				email = $4,
				ocorrencia = $5,
				tipo_sanguineo = $6,
				sintomas = $7,
				diagnostico = $8
			WHERE id = $9
		`, nome, cpf, idade, email, ocorrencia, tipoSanguineo, sintomas, diagnostico, id)

		if err != nil {
			http.Error(w, "Erro ao atualizar paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/patients", http.StatusSeeOther)
	}
}

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID do paciente não informado", http.StatusBadRequest)
		return
	}

	_, err := utils.DB.Exec("DELETE FROM patients WHERE id = $1", id)

	if err != nil {
		http.Error(w, "Erro ao excluir paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/patients", http.StatusSeeOther)
}
