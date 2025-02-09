const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  // if (!isUserLogged()) {
  //     window.location.href =
  //         window.location.origin + "/login.html?reason=login_required";
  // }

  obtenerCamiones();
});

function obtenerCamiones() {
  const urlConFiltro = `http://localhost:8080/trucks/`;
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerCamiones,
    errorObtenerCamiones
  );
}

function exitoObtenerCamiones(data) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  elementosTable.innerHTML = "";

  if (data != null) {
    data.forEach((elemento) => {
      const row = document.createElement("tr");
      const fechaCreacion = new Date(elemento.fecha_creacion).toLocaleString();
      const fechaModificacion = new Date(
        elemento.fecha_modificacion
      ).toLocaleString();

      row.innerHTML = `
        <td>${elemento.id}</td>      
        <td>${elemento.patente}</td>
        <td>${elemento.peso_maximo}</td>
        <td>${elemento.costo_km}</td>
        <td>${fechaCreacion}</td>
        <td>${fechaModificacion}</td>
        <td class="acciones">
          <button class="eliminar" onclick="confirmarEliminar('${elemento.id}', '${elemento.patente}')">Eliminar</button>
          <a class="editar" href="form-camion.html?id=${elemento.id}&tipo=EDITAR">Editar</a>
        </td>
      `;
      elementosTable.appendChild(row);
    });
  }
}

function confirmarEliminar(id, patente) {
  if (confirm(`Seguro que quiere eliminar el camión con patente ${patente}?`)) {
    eliminarCamion(id);
  }
}

function eliminarCamion(id) {
  const baseUrl = `http://localhost:8080/trucks/`;
  makeRequest(
    `${baseUrl}${id}`,
    Method.DELETE,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    function (data) {
      alert("Camión eliminado exitosamente");
      obtenerCamiones();
    },
    function (status, response) {
      console.error("Error al eliminar:", response);
      alert("Error al eliminar el camión");
    }
  );
}

function errorObtenerCamiones(error) {
  console.error("Error details:", error);
  if (error.status === 404) {
    alert("No se encontraron camiones.");
  } else {
    alert(
      "Error al conectar con el servidor. Por favor, intente nuevamente más tarde."
    );
  }
}
