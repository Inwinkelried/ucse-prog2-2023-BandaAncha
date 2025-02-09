const url = "http://localhost:8080/products/";

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
    cargarProducto(id);
  }

  document
    .getElementById("form-producto")
    .addEventListener("submit", function (event) {
      guardarProducto(event);
    });
});

function cargarProducto(id) {
  makeRequest(
    `${url}${id}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerProducto,
    errorObtenerProducto
  );
}

function exitoObtenerProducto(producto) {
  document.getElementById("id").value = producto.id;
  document.getElementById("tipo").value = producto.tipo;
  document.getElementById("nombre").value = producto.nombre;
  document.getElementById("peso_unitario").value = producto.peso_unitario;
  document.getElementById("precio_unitario").value = producto.precio_unitario;
  document.getElementById("stock_minimo").value = producto.stock_minimo;
  document.getElementById("stock_actual").value = producto.stock_actual;
}

function errorObtenerProducto(status, response) {
  console.error("Error al obtener producto:", response);
  alert("Error al obtener los datos del producto");
}

function guardarProducto(event) {
  event.preventDefault();

  const id = document.getElementById("id").value;
  const producto = {
    tipo: document.getElementById("tipo").value,
    nombre: document.getElementById("nombre").value,
    peso_unitario: parseFloat(document.getElementById("peso_unitario").value),
    precio_unitario: parseFloat(
      document.getElementById("precio_unitario").value
    ),
    stock_minimo: parseInt(document.getElementById("stock_minimo").value),
    stock_actual: parseInt(document.getElementById("stock_actual").value),
  };

  const method = id ? Method.PUT : Method.POST;
  const finalUrl = id ? `${url}${id}` : url;

  makeRequest(
    finalUrl,
    method,
    producto,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoGuardarProducto,
    errorGuardarProducto
  );
}

function exitoGuardarProducto(response) {
  alert("Producto guardado exitosamente");
  window.location.href = "listado-productos.html";
}

function errorGuardarProducto(status, response) {
  console.error("Error al guardar producto:", response);
  alert("Error al guardar el producto");
}
