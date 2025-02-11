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

// Variables globales
let pedidos = [];
const API_URL = "http://localhost:8080";

// Elementos del DOM
const tablaPedidos = document.getElementById("tablaPedidos");
const btnFiltrar = document.getElementById("btnFiltrar");
const btnLimpiar = document.getElementById("btnLimpiar");
const formFiltros = document.getElementById("formFiltros");

// Event Listeners
document.addEventListener("DOMContentLoaded", () => {
  cargarPedidos();
});

function cargarPedidos() {
  makeRequest(
    `${API_URL}/orders`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerPedidos,
    errorObtenerPedidos
  );
}

function exitoObtenerPedidos(data) {
  pedidos = data;
  mostrarPedidos(pedidos);
}

function errorObtenerPedidos(error) {
  console.error("Error al cargar pedidos:", error);
  mostrarError("Error al cargar los pedidos");
}

// Función para mostrar los pedidos en la tabla
function mostrarPedidos(pedidosAMostrar) {
  const tbody = tablaPedidos.querySelector("tbody");
  tbody.innerHTML = "";

  pedidosAMostrar.forEach((pedido) => {
    const tr = document.createElement("tr");

    // Formatear la lista de productos
    const productosTexto = pedido.productos
      ? pedido.productos
          .map(
            (p) => `${p.nombre_producto || p.codigo_producto} (${p.cantidad})`
          )
          .join(", ")
      : "Sin productos";

    // Formatear las fechas
    const fechaCreacion = new Date(pedido.fecha_creacion).toLocaleDateString();
    const fechaModificacion = new Date(
      pedido.fecha_modificacion
    ).toLocaleDateString();

    tr.innerHTML = `
      <td>${pedido.id || pedido._id}</td>
      <td>${productosTexto}</td>
      <td>${pedido.destino}</td>
      <td>${pedido.estado}</td>
      <td>${fechaCreacion}</td>
      <td>${fechaModificacion}</td>
      <td>
        ${
          pedido.estado === "Pendiente"
            ? `
            <button onclick="editarPedido('${
              pedido.id || pedido._id
            }')" class="btn-accion btn-editar">Editar</button>
            <button onclick="eliminarPedido('${
              pedido.id || pedido._id
            }')" class="btn-accion btn-eliminar">Eliminar</button>
          `
            : ""
        }
      </td>
    `;
    tbody.appendChild(tr);
  });
}

// Función para aplicar los filtros
function aplicarFiltros() {
  const codigoEnvio = document.getElementById("filtroCodigoEnvio").value.trim();
  const estado = document.getElementById("filtroEstado").value;
  const fechaDesde = document.getElementById("filtroFechaDesde").value;
  const fechaHasta = document.getElementById("filtroFechaHasta").value;

  let pedidosFiltrados = [...pedidos];

  // Filtrar por código de envío
  if (codigoEnvio) {
    pedidosFiltrados = pedidosFiltrados.filter((pedido) =>
      (pedido.id || pedido._id)
        .toLowerCase()
        .includes(codigoEnvio.toLowerCase())
    );
  }

  // Filtrar por estado
  if (estado) {
    pedidosFiltrados = pedidosFiltrados.filter(
      (pedido) => pedido.estado === estado
    );
  }

  // Filtrar por fecha
  if (fechaDesde) {
    const fechaDesdeObj = new Date(fechaDesde);
    pedidosFiltrados = pedidosFiltrados.filter(
      (pedido) => new Date(pedido.fecha_creacion) >= fechaDesdeObj
    );
  }

  if (fechaHasta) {
    const fechaHastaObj = new Date(fechaHasta);
    fechaHastaObj.setHours(23, 59, 59); // Incluir todo el día
    pedidosFiltrados = pedidosFiltrados.filter(
      (pedido) => new Date(pedido.fecha_creacion) <= fechaHastaObj
    );
  }

  mostrarPedidos(pedidosFiltrados);
}

// Función para limpiar los filtros
function limpiarFiltros() {
  formFiltros.reset();
  mostrarPedidos(pedidos);
}

// Funciones de acciones
function editarPedido(id) {
  window.location.href = `form-pedido.html?id=${id}`;
}

function eliminarPedido(id) {
  if (!confirm("¿Está seguro de que desea eliminar este pedido?")) {
    return;
  }

  makeRequest(
    `${API_URL}/orders/${id}`,
    Method.DELETE,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    () => {
      cargarPedidos();
      mostrarMensaje("Pedido eliminado correctamente");
    },
    (error) => {
      console.error("Error:", error);
      mostrarError("Error al eliminar el pedido");
    }
  );
}

// Función para mostrar mensajes de error
function mostrarError(mensaje) {
  const mensajeElement = document.createElement("div");
  mensajeElement.className = "mensaje mensaje-error";
  mensajeElement.textContent = mensaje;
  document
    .querySelector(".container")
    .insertBefore(mensajeElement, document.querySelector(".filtros"));

  setTimeout(() => {
    mensajeElement.remove();
  }, 3000);
}

// Función para mostrar mensajes de éxito
function mostrarMensaje(mensaje) {
  const mensajeElement = document.createElement("div");
  mensajeElement.className = "mensaje mensaje-exito";
  mensajeElement.textContent = mensaje;
  document
    .querySelector(".container")
    .insertBefore(mensajeElement, document.querySelector(".filtros"));

  setTimeout(() => {
    mensajeElement.remove();
  }, 3000);
}

btnFiltrar.addEventListener("click", aplicarFiltros);
btnLimpiar.addEventListener("click", limpiarFiltros);
