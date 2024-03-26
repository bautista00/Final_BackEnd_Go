package domain

type Paciente struct {
    ID         int    `json:"id"`
    Nombre     string `json:"nombre"`
    Apellido   string `json:"apellido"`
    Domicilio  string `json:"domicilio"`
    DNI        string `json:"dni"`
    FechaAlta  string `json:"fecha_alta"`
}
