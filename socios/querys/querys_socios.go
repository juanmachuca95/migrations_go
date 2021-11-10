package socios

/* ENTENDER QUE SON LOS BENEFICIARIOS Y EN QUE SE DIFERENCIAN DE LOS SOCIOS */

var GetSocios = func() string {
	return "SELECT * FROM `sassocios` INNER JOIN sasbeneficiarios ON sassocios.idPersona = sasbeneficiarios.idPersona;"
}

/**
INNER JOIN socios, personas, beneficiarios, sas

SELECT socios.idPersona, beneficiarios.idPersona, sas.id FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS;
*/

var GetSociosBeneficiariosSas = func() string {
	return "SELECT socios.representado AS representado, socios.firma AS firma, beneficiarios.porcentaje AS porcentaje, beneficiarios.acciones, sas.id AS sas_id FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS;"
}

/*
CONSULTA DONDE SE COMPRUEBA QUE EXISTEN DATOS QUE NO SON LO MISMO

SAS BENEFICIARIOS -- SAS SOCIOS .. NO TIENEN LA MISMA CANTIDAD DE ACCIONES.
SELECT socios.id AS socios_id, beneficiarios.id AS beneficiarios_id, socios.representado AS representado, socios.firma AS firma, beneficiarios.porcentaje AS porcentaje, beneficiarios.acciones AS acciones_beneficiarios, socios.cantAcciones AS acciones_socios, sas.id AS sas_id FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS WHERE socios.cantAcciones != beneficiarios.acciones;

*/
