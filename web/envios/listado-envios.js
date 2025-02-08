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

  obtenerEnvios();
});

const urlConFiltro = `http://localhost:8080/shippings/`;

function obtenerEnvios() {
  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerEnvios,
    errorObtenerEnvios
  );
}

function exitoObtenerEnvios(data) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  data.forEach((elemento) => {
    const row = document.createElement("tr");

    row.innerHTML = `   
                            <td>${elemento.idPedido}</td>
                            <td>${elemento.idCamion}</td>
                            <td>${elemento.pedidos}</td>
                            <td>${elemento.paradas}</td>
                            <td>${elemento.estado}</td>
                            <td>${elemento.fecha_creacion}</td>
                            <td>${elemento.fecha_modificacion}</td>
                    `;
    elementosTable.appendChild(row);
  });
}

function errorObtenerEnvios(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}
