const url = "http://localhost:8080/trucks/";

const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  const params = new URLSearchParams(window.location.search);
  const id = params.get("id");
  const tipo = params.get("tipo");

  if (id && tipo === "EDITAR") {
    cargarCamion(id);
  }

  document
    .getElementById("form-camion")
    .addEventListener("submit", function (event) {
      guardarCamion(event);
    });
});

function cargarCamion(id) {
  makeRequest(
    `${url}${id}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerCamion,
    errorObtenerCamion
  );
}

function exitoObtenerCamion(camion) {
  document.getElementById("id").value = camion.id;
  document.getElementById("patente").value = camion.patente;
  document.getElementById("peso_maximo").value = camion.peso_maximo;
  document.getElementById("costo_km").value = camion.costo_km;
}

function errorObtenerCamion(status, response) {
  console.error("Error al obtener camión:", response);
  alert("Error al obtener los datos del camión");
}

function guardarCamion(event) {
  event.preventDefault();

  const id = document.getElementById("id").value;
  const camion = {
    patente: document.getElementById("patente").value,
    peso_maximo: parseInt(document.getElementById("peso_maximo").value),
    costo_km: parseInt(document.getElementById("costo_km").value),
  };

  const method = id ? Method.PUT : Method.POST;
  const finalUrl = id ? `${url}${id}` : url;

  makeRequest(
    finalUrl,
    method,
    camion,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoGuardarCamion,
    errorGuardarCamion
  );
}

function exitoGuardarCamion(response) {
  alert("Camión guardado exitosamente");
  window.location.href = "listado-camiones.html";
}

function errorGuardarCamion(status, response) {
  console.error("Error al guardar camión:", response);
  alert("Error al guardar el camión");
}
