const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

const urlPedidos = "http://localhost:8080/orders/";
const urlProductos = "http://localhost:8080/products/";

document.addEventListener("DOMContentLoaded", function (event) {
    cargarProductos();
    document.getElementById("formulario").addEventListener("submit", function (event) {
        event.preventDefault();
        guardarPedido();
    });
});

function cargarProductos() {
    makeRequest(
        urlProductos,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PUBLIC,
        exitoObtenerProductos,
        errorObtenerProductos
    );
}

function exitoObtenerProductos(productos) {
    const tbody = document.querySelector("#elementosTable tbody");
    tbody.innerHTML = "";

    productos.forEach(producto => {
        const tr = document.createElement("tr");
        tr.innerHTML = `
            <td>
                <input type="checkbox" name="productos" value="${producto.id}" 
                       onchange="toggleCantidadInput(this)">
            </td>
            <td>${producto.id}</td>
            <td>${producto.nombre}</td>
            <td>
                <input type="number" name="cantidad_${producto.id}" 
                       min="1" value="1" disabled 
                       data-precio="${producto.precio_unitario}"
                       data-peso="${producto.peso_unitario}">
            </td>
            <td>${producto.precio_unitario}</td>
            <td>${producto.peso_unitario}</td>
        `;
        tbody.appendChild(tr);
    });
}

function errorObtenerProductos(status, response) {
    console.error("Error al obtener productos:", response);
    alert("Error al obtener la lista de productos");
}

function toggleCantidadInput(checkbox) {
    const cantidadInput = checkbox.closest("tr").querySelector('input[type="number"]');
    cantidadInput.disabled = !checkbox.checked;
}

function guardarPedido() {
    const destino = document.getElementById("destino").value;
    if (!destino) {
        alert("Por favor, ingrese un destino");
        return;
    }

    const productosSeleccionados = [];
    const checkboxes = document.querySelectorAll('input[name="productos"]:checked');
    
    if (checkboxes.length === 0) {
        alert("Por favor, seleccione al menos un producto");
        return;
    }

    checkboxes.forEach(checkbox => {
        const tr = checkbox.closest("tr");
        const cantidadInput = tr.querySelector('input[type="number"]');
        const idProducto = tr.cells[1].textContent;
        
        productosSeleccionados.push({
            codigo_producto: idProducto,
            cantidad: parseInt(cantidadInput.value)
        });
    });

    const pedido = {
        destino: destino,
        productos: productosSeleccionados,
        estado: "PENDIENTE"
    };

    makeRequest(
        urlPedidos,
        Method.POST,
        pedido,
        ContentType.JSON,
        CallType.PUBLIC,
        exitoGuardarPedido,
        errorGuardarPedido
    );
}

function exitoGuardarPedido(response) {
    alert("Pedido creado exitosamente");
    window.location.href = "listado-pedidos.html";
}

function errorGuardarPedido(status, response) {
    console.error("Error al guardar pedido:", response);
    alert("Error al guardar el pedido");
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
