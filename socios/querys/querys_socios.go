package socios

/* ENTENDER QUE SON LOS BENEFICIARIOS Y EN QUE SE DIFERENCIAN DE LOS SOCIOS */

var GetSocios = func() string {
	return "SELECT socios.id, socios.idPersona, sas.id, socios.cantAcciones, beneficiarios.acciones, beneficiarios.porcentaje, socios.firma, socios.representado, socios.created_at, socios.updated_at FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS;"
}

var CreateSociosSAS = func() string {
	return "INSERT INTO sass_socios (id, carga_capital_invertido, porcentaje, created_at, updated_at, sass_id, personas_id) VALUES (?,?,?,?,?,?,?)"
}

/**
INNER JOIN socios, personas, beneficiarios, sas

SELECT socios.idPersona, beneficiarios.idPersona, sas.id FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS;
*/

/*
CONSULTA DONDE SE COMPRUEBA QUE EXISTEN DATOS QUE NO SON LO MISMO

SAS BENEFICIARIOS -- SAS SOCIOS .. NO TIENEN LA MISMA CANTIDAD DE ACCIONES.
SELECT socios.id AS socios_id, beneficiarios.id AS beneficiarios_id, socios.representado AS representado, socios.firma AS firma, beneficiarios.porcentaje AS porcentaje, beneficiarios.acciones AS acciones_beneficiarios, socios.cantAcciones AS acciones_socios, sas.id AS sas_id FROM sassocios as socios INNER JOIN sasbeneficiarios as beneficiarios ON socios.idPersona = beneficiarios.idPersona INNER JOIN saspersonas as personas ON personas.id = socios.idPersona AND personas.id = beneficiarios.idPersona INNER JOIN sasform as sas ON sas.id = personas.idSAS WHERE socios.cantAcciones != beneficiarios.acciones;

*/

/*
CONSULTA QUE OBTIENE TODOS LOS ID DE PERSONA QUE NO EXISTEN EN TABLA SASPERSONAS DE LA TABLA SASSOCIOS
SELECT socios.idPersona FROM sassocios AS socios WHERE socios.idPersona NOT IN (SELECT personas.id FROM saspersonas AS personas);


CONSULTA QUE OBTIENE TODOS LOS ID DE PERSONA QUE (NO EXISTEN EN ID PERSONA CON RESPECTO A SASSOCIO)
SELECT sasbeneficiarios.idPersona FROM sasbeneficiarios WHERE sasbeneficiarios.idPersona IN (SELECT socios.idPersona FROM sassocios AS socios WHERE socios.idPersona NOT IN (SELECT personas.id FROM saspersonas AS personas));



*/
