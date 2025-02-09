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

  obtenerPedidos();
});
const urlConFiltro = `http://localhost:8080/orders/`;
function obtenerPedidos() {
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function exitoObtenerPedidos(data) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  data.forEach((elemento) => {
    const row = document.createElement("tr"); //crear una fila

    const fechaCreacion = new Date(elemento.fecha_creacion).toLocaleString();
    const fechaModificacion = new Date(
      elemento.fecha_modificacion
    ).toLocaleString();

    row.innerHTML = `   
      <td>${elemento.id}</td>
      <td>${elemento.productos}</td>
      <td>${elemento.destino}</td>
      <td>${elemento.estado}</td>
      <td>${fechaCreacion}</td>
      <td>${fechaModificacion}</td>
    `;
    elementosTable.appendChild(row);
  });
}

function errorObtenerPedidos(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}
