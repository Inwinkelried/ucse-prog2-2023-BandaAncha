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
  urlConFiltro = `http://localhost:8080/products`;

  makeRequest(
    `${urlConFiltro}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProductos,
    errorObtenerProductos
  );
}

function exitoObtenerProductos(data) {
  const elementosTable = document //tabla en la que se colocan los envios que se obtienen
    .getElementById("elementosTable")
    .querySelector("tbody");

  elementosTable.innerHTML = "";

  // Llenar la tabla con los datos obtenidos
  if (data != null) {
    data.forEach((elemento) => {
      const row = document.createElement("tr"); //crear una fila
        row.innerHTML = ` 
                    <td>${elemento.id}</td>
                    <td>${elemento.tipo}</td>
                    <td>${elemento.nombre}</td>
                    <td>${elemento.peso_unitario}</td>
                    <td>${elemento.precio_unitario}</td>
                    <td>${elemento.stock_minimo}</td>
                    <td>${elemento.stock_actual}</td>
                    <td>${elemento.fecha_creacion}</td>
                    <td>${elemento.fecha_ultima_actualizacion}</td>
                    <td class="acciones">
                        <a class="eliminar" href="form-producto.html?id=${elemento.id}&tipo=ELIMINAR">Eliminar</a>
                        <a class="editar" href="form-producto.html?id=${elemento.id}&tipo=EDITAR">Editar</a>
                    </td>
                    `;

      elementosTable.appendChild(row);
    });
  }
}

function errorObtenerProductos(response) {
  alert("Error en la solicitud al servidor.");
  console.log(response.json());
  throw new Error("Error en la solicitud al servidor.");
}

function obtenerProductoFiltrado(tipo) {
  var url = new URL(urlConFiltro);

  switch (tipo) {
    case "stock":
      url.searchParams.set("filtrarPorStockMinimo", true);
      break;
    case "tipo":
      url.searchParams.set(
        "tipoProducto",
        document.getElementById("TipoProducto").value
      );
      break;
    default:
      url.href = `http://localhost:8080/products`;
      break;
  }

  console.log(url.href);

  makeRequest(
    `${url.href}`,
    Method.GET,
    null,
    ContentType.JSON,
    CallType.PRIVATE,
    exitoObtenerProductos,
    errorObtenerProductos
  );
}

