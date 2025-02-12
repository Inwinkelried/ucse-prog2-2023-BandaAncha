const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

let enviosData = []; // Almacenar todos los envíos para filtrar localmente

document.addEventListener("DOMContentLoaded", function (event) {
  // if (!isUserLogged()) {
  //     window.location.href =
  //         window.location.origin + "/login.html?reason=login_required";
  // }

  obtenerEnvios();

  // Agregar event listeners para los filtros
  document.getElementById("botonFiltrar").addEventListener("click", aplicarFiltros);
  document.getElementById("botonLimpiar").addEventListener("click", limpiarFiltros);
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
  enviosData = data; // Guardar los datos completos
  mostrarEnvios(data); // Mostrar todos los envíos inicialmente
}

function mostrarEnvios(envios) {
  const elementosTable = document
    .getElementById("elementosTable")
    .querySelector("tbody");

  elementosTable.innerHTML = ""; // Limpiar la tabla antes de agregar nuevos datos

  envios.forEach((elemento) => {
    const row = document.createElement("tr");

    // Formatear la lista de pedidos
    const pedidosTexto = Array.isArray(elemento.pedidos) 
      ? elemento.pedidos.map(p => p.id || p).join(", ")
      : elemento.pedidos || "Sin pedidos";

    // Formatear la lista de paradas
    const paradasTexto = Array.isArray(elemento.paradas)
      ? elemento.paradas.map(p => p.ciudad || p).join(", ")
      : elemento.paradas || "Sin paradas";

    const fechaCreacion = new Date(elemento.fecha_creacion).toLocaleString();
    const fechaModificacion = new Date(
      elemento.fecha_modificacion
    ).toLocaleString();

    row.innerHTML = `   
      <td>${elemento.id || elemento._id}</td>
      <td>${elemento.patente_camion || elemento.idCamion || "No asignado"}</td>
      <td>${pedidosTexto}</td>
      <td>${paradasTexto}</td>
      <td>${elemento.estado}</td>
      <td>${fechaCreacion}</td>
      <td>${fechaModificacion}</td>
    `;
    elementosTable.appendChild(row);
  });
}

function errorObtenerEnvios(response) {
  console.error("Error al obtener envíos:", response);
  alert("Error al obtener la lista de envíos");
}

function aplicarFiltros() {
  const patenteCamion = document.getElementById("filtroPatenteCamion").value.toLowerCase();
  const estado = document.getElementById("filtroEstado").value;
  const ultimaParada = document.getElementById("filtroUltimaParada").value.toLowerCase();
  const fechaDesde = document.getElementById("filtroFechaDesde").value;
  const fechaHasta = document.getElementById("filtroFechaHasta").value;

  let enviosFiltrados = enviosData;

  // Filtrar por patente de camión
  if (patenteCamion) {
    enviosFiltrados = enviosFiltrados.filter(envio => 
      (envio.patente_camion || "").toLowerCase().includes(patenteCamion)
    );
  }

  // Filtrar por estado
  if (estado) {
    enviosFiltrados = enviosFiltrados.filter(envio => 
      envio.estado === estado
    );
  }

  // Filtrar por última parada
  if (ultimaParada) {
    enviosFiltrados = enviosFiltrados.filter(envio => {
      if (!envio.paradas || envio.paradas.length === 0) return false;
      const ultimaParadaEnvio = envio.paradas[envio.paradas.length - 1];
      return (ultimaParadaEnvio.ciudad || "").toLowerCase().includes(ultimaParada);
    });
  }

  // Filtrar por fecha
  if (fechaDesde) {
    const fechaDesdeObj = new Date(fechaDesde);
    enviosFiltrados = enviosFiltrados.filter(envio => 
      new Date(envio.fecha_creacion) >= fechaDesdeObj
    );
  }

  if (fechaHasta) {
    const fechaHastaObj = new Date(fechaHasta);
    fechaHastaObj.setHours(23, 59, 59); // Incluir todo el día
    enviosFiltrados = enviosFiltrados.filter(envio => 
      new Date(envio.fecha_creacion) <= fechaHastaObj
    );
  }

  mostrarEnvios(enviosFiltrados);
}

function limpiarFiltros() {
  document.getElementById("formFiltros").reset();
  mostrarEnvios(enviosData);
}
