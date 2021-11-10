package admins

var GetAdmins = func() string {
	return "SELECT admin.id, admin.idPersona, admin.titular, admin.inciso, admin.created_at, admin.updated_at, admin.expuestoPoli, admin.domicilioFiscal, admin.adminRelaciones, admin.telefonoFiscal, admin.areaTelefono, sas.id AS sas_id FROM `sasadmins` AS admin INNER JOIN saspersonas AS personas ON admin.idPersona = personas.id INNER JOIN sasform AS sas ON personas.idSAS = sas.id;"
}

var CreateAdminsSAS = func() string {
	return "INSERT INTO sas_administradors (id, sass_id, personas_id, administrador, politicamente_expuesto, tipo_politicamente_expuesto, activo, created_at, updated_at, administrador_relaciones, domicilio_fiscal) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
}
