const urlEnvios = "http://localhost:8080/shippings/";
const urlPedidos = "http://localhost:8080/orders/Approved";
const urlCamiones = "http://localhost:8080/trucks/";

let camiones = [];
let pedidosSeleccionados = [];
let camionSeleccionado = null;

document.addEventListener("DOMContentLoaded", function () {
    cargarCamiones();
    cargarPedidosAprobados();

    // Evento para actualizar información cuando se cambia el camión
    document.getElementById("camion").addEventListener("change", function(e) {
        camionSeleccionado = camiones.find(c => c.id === e.target.value);
        if (camionSeleccionado) {
            document.getElementById("capacidadCamion").textContent = camionSeleccionado.peso_maximo;
            actualizarPesoTotal();
        }
    });

    document.getElementById("formulario").addEventListener("submit", function (event) {
        event.preventDefault();
        crearEnvio();
    });
});

function cargarCamiones() {
    makeRequest(
        urlCamiones,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PUBLIC,
        exitoObtenerCamiones,
        errorObtenerCamiones
    );
}

function exitoObtenerCamiones(data) {
    camiones = data;
    const selectCamion = document.getElementById("camion");
    
    data.forEach(camion => {
        const option = document.createElement("option");
        option.value = camion.id;
        option.textContent = `${camion.patente} (Capacidad: ${camion.peso_maximo} kg)`;
        selectCamion.appendChild(option);
    });
}

function errorObtenerCamiones(error) {
    console.error("Error al obtener camiones:", error);
    alert("Error al cargar la lista de camiones");
}

function cargarPedidosAprobados() {
    makeRequest(
        urlPedidos,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PUBLIC,
        exitoObtenerPedidos,
        errorObtenerPedidos
    );
}

function exitoObtenerPedidos(pedidos) {
    const tbody = document.querySelector("#tablaPedidos tbody");
    tbody.innerHTML = "";

    pedidos.forEach(pedido => {
        const tr = document.createElement("tr");
        
        // Calcular peso total del pedido
        const pesoTotal = pedido.productos.reduce((total, prod) => {
            return total + (prod.peso_unitario * prod.cantidad);
        }, 0);

        // Formatear lista de productos
        const productosTexto = pedido.productos
            .map(p => `${p.nombre || p.codigo_producto} (${p.cantidad})`)
            .join(", ");

        tr.innerHTML = `
            <td>
                <input type="checkbox" name="pedidos" value="${pedido.id}" 
                       onchange="actualizarPesoTotal()" data-peso="${pesoTotal}">
            </td>
            <td>${pedido.id}</td>
            <td>${productosTexto}</td>
            <td>${pesoTotal} kg</td>
            <td>${pedido.destino}</td>
            <td>${new Date(pedido.fecha_modificacion).toLocaleString()}</td>
        `;
        tbody.appendChild(tr);
    });
}

function errorObtenerPedidos(error) {
    console.error("Error al obtener pedidos:", error);
    alert("Error al cargar la lista de pedidos");
}

function actualizarPesoTotal() {
    const checkboxes = document.querySelectorAll('input[name="pedidos"]:checked');
    let pesoTotal = 0;

    checkboxes.forEach(checkbox => {
        pesoTotal += parseFloat(checkbox.dataset.peso);
    });

    document.getElementById("pesoTotal").textContent = pesoTotal;

    // Validar contra la capacidad del camión
    if (camionSeleccionado && pesoTotal > camionSeleccionado.peso_maximo) {
        alert("¡Advertencia! El peso total supera la capacidad del camión");
    }
}

function crearEnvio() {
    if (!camionSeleccionado) {
        alert("Por favor, seleccione un camión");
        return;
    }

    const pedidosSeleccionados = Array.from(
        document.querySelectorAll('input[name="pedidos"]:checked')
    ).map(checkbox => checkbox.value);

    if (pedidosSeleccionados.length === 0) {
        alert("Por favor, seleccione al menos un pedido");
        return;
    }

    const pesoTotal = parseFloat(document.getElementById("pesoTotal").textContent);
    if (pesoTotal > camionSeleccionado.peso_maximo) {
        alert("El peso total supera la capacidad del camión");
        return;
    }

    const envio = {
        patente_camion: camionSeleccionado.patente,
        pedidos: pedidosSeleccionados,
        estado: "A Despachar"
    };

    makeRequest(
        urlEnvios,
        Method.POST,
        envio,
        ContentType.JSON,
        CallType.PUBLIC,
        exitoCrearEnvio,
        errorCrearEnvio
    );
}

function exitoCrearEnvio(response) {
    alert("Envío creado exitosamente");
    window.location.href = "listado-envios.html";
}

function errorCrearEnvio(error) {
    console.error("Error al crear el envío:", error);
    alert("Error al crear el envío");
} 