const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
    //parametros de la url son correctos?
    const urlParams = new URLSearchParams(window.location.search);
    const idPedido = urlParams.get("id");
    const operacion = urlParams.get("tipo");

    if (idPedido != "" && idPedido != null && operacion == "ACEPTAR") {
        aceptarPedido(idPedido);
    } else if (idPedido != "" && idPedido != null && operacion == "CANCELAR") {
        cancelarPedido(idPedido);
    } else {
        document
            .getElementById("botonGuardar")
            .addEventListener("click", function (event) {
                guardarPedido(event);
            });

        obtenerProductos();
    }
});

//trae los productos de la db y los muestra en la tabla
function obtenerProductos() {
    const urlConFiltro = `http://localhost:8080/products`;

    makeRequest(
        `${urlConFiltro}`,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoObtenerProductos,
        errorPedido
    );
}

function exitoObtenerProductos(data) {
    const elementosTable = document //tabla en la que se colocan los envios que se obtienen
        .getElementById("elementosTable")
        .querySelector("tbody");

    data.forEach((elemento) => {
        const row = document.createElement("tr"); //crear una fila

        row.innerHTML = ` 
                  <td><input type="checkbox" class="checkbox"></td>
                  <td>${elemento.id}</td>
                  <td>${elemento.nombre}</td>
                  <td><input type="text" placeholder=" Ingrese cantidad"></td>
                  <td>${elemento.precio_unitario}</td>
                  <td>${elemento.peso_unitario}</td>
                 `;
        elementosTable.appendChild(row);
    });
}

function obtenerProductosSeleccionados() {
    var ProductosSeleccionados = [];
    let checkboxes = document.querySelectorAll(".checkbox");
    checkboxes.forEach(function (checkbox) {
        if (checkbox.checked) {
            // Agregar el producto seleccionado al objeto ProductosSeleccionados
            var tr = checkbox.closest("tr");
            var idProducto = tr.cells[1].textContent;
            var nombreProducto = tr.cells[2].textContent;
            var cantidad = parseInt(
                tr.cells[3].getElementsByTagName("input")[0].value
            );
            var precioUnitario = parseFloat(tr.cells[4].textContent);
            var pesoUnitario = parseFloat(tr.cells[5].textContent);

            var productoSeleccionado = {
                id: idProducto,
                nombre_producto: nombreProducto,
                cantidad: cantidad,
                precio_unitario: precioUnitario,
                peso_unitario: pesoUnitario,
            };

            ProductosSeleccionados.push(productoSeleccionado);
        }
    });

    return ProductosSeleccionados;
}

function guardarPedido() {
    // Obtiene la fecha y hora actuales
    const fechaActual = new Date();
    const fechaFormateada = fechaActual.toISOString(); // Formatea la fecha en formato ISO

    // Arma la data a enviar con las fechas y horas actuales
    const data = {
        id: "",
        productos: obtenerProductosSeleccionados(),
        destino: document.getElementById("destino").value,
        estado: "Pendiente",
        fecha_creacion: fechaFormateada, // Usa la fecha y hora actuales como fecha_creacion
        modificacion: fechaFormateada, // Usa la fecha y hora actuales como fecha_modificacion
    };

    const urlConFiltro = `http://localhost:8080/orders`;
    makeRequest(
        `${urlConFiltro}`,
        Method.POST,
        data,
        ContentType.JSON,
        CallType.PRIVATE,
        exitoPedido,
        errorPedido
    );
}

function exitoPedido(data) {
    window.location = window.location.origin + "/web/pedidos/listado-pedidos.html";
}

function errorPedido(response) {
    alert("Error en la solicitud al servidor.");
    console.log(response.json());
    throw new Error("Error en la solicitud al servidor.");
}

function aceptarPedido(id) {
    if (confirm("El pedido se aceptará, ¿estás seguro")) {
        const urlConFiltro = `http://localhost:8080/orders/Confirm/${id}`;
        data = [];
        makeRequest(
            `${urlConFiltro}`,
            Method.PUT,
            data,
            ContentType.JSON,
            CallType.PRIVATE,
            exitoPedido,
            errorPedido
        );
    } else {
        window.location = document.location.origin + "/web/pedidos/listado-pedidos.html";
    }
}

function cancelarPedido(id) {
    if (confirm("El pedido se cancelará, ¿estás seguro")) {
        const urlConFiltro = `http://localhost:8080/orders/Cancel/${id}`;
        data = [];
        makeRequest(
            `${urlConFiltro}`,
            Method.PUT,
            data,
            ContentType.JSON,
            CallType.PRIVATE,
            exitoPedido,
            errorPedido
        );
    } else {
        window.location = document.location.origin + "/web/pedidos/index.html";
    }
}
