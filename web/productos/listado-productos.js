const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
  //   if (!isUserLogged()) {
  //     window.location =
  //       document.location.origin + "/web/login/login.html?reason=login_required";
  //   }

  obtenerProductos();
});

function obtenerProductos() {
  const urlConFiltro = `http://localhost:8080/products/`;
  makeRequest(
    urlConFiltro,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC, //PUBLIC PARA PROBAR, DESPUES CAMBIAR A PRIVATE
    exitoObtenerProductos,
    errorObtenerProductos
  );
}

function exitoObtenerProductos(data) {
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
            <td>${elemento.tipo}</td>
            <td>${elemento.nombre}</td>
            <td>${elemento.peso_unitario}</td>
            <td>${elemento.precio_unitario}</td>
            <td>${elemento.stock_minimo}</td>
            <td>${elemento.stock_actual}</td>
            <td>${fechaCreacion}</td>
            <td>${fechaModificacion}</td>
            <td class="acciones">
            <button class="eliminar" onclick="confirmarEliminar('${elemento.id}', '${elemento.nombre}')">Eliminar</button>
            <a class="editar" href="form-producto.html?id=${elemento.id}&tipo=EDITAR">Editar</a>
            </td>
            `;

      elementosTable.appendChild(row);
    });
  }
}

function confirmarEliminar(id, nombre) {
  if (confirm(`Seguro que quiere eliminar el objeto ${nombre}?`)) {
    eliminarProducto(id);
  }
}

function eliminarProducto(id) {
  makeRequest(
    `http://localhost:8080/products/${id}`,
    Method.DELETE,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    function (data) {
      alert("Producto eliminado exitosamente");
      obtenerProductos();
    },
    function (status, response) {
      console.error("Error al eliminar:", response);
      alert("Error al eliminar el producto");
    }
  );
}

function errorObtenerProductos(status, response) {
  console.error("Error details:", response);
  if (status === 404) {
    alert("No se encontraron productos.");
  } else {
    alert("Error en la solicitud al servidor.");
  }
}

function obtenerProductoFiltrado(tipo) {
  const baseUrl = "http://localhost:8080";
  var url = new URL(`${baseUrl}/products/Filter`);

  switch (tipo) {
    case "stock":
      url.searchParams.set("filtrarPorStockMinimo", true);
      break;
    case "tipo":
      url.searchParams.set(
        "tipoProducto",
        document.getElementById("filtroTipo").value
      );
      break;
    default:
      url = new URL(`${baseUrl}/products`);
      break;
  }

  console.log(url.href);

  makeRequest(
    url.href,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PUBLIC,
    exitoObtenerProductos,
    errorObtenerProductos
  );
}
