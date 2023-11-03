const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "/");
customHeaders.append("Accept-Encoding", "gzip, deflate, br"); 
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
    // if (!isUserLogged()) {
    //     window.location.href =
    //         window.location.origin + "/login.html?reason=login_required";
    // }

    obtenerCamiones();
});

function obtenerPedidos() {
    const urlConFiltro = `http://localhost:8080/orders/`; 
    fetch(urlConFiltro, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("authToken")}`,
        },
    })
        .then((response) => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error("Error en llamada");
            }
        })
        .then((data) => {
            exitoObtenerPedidos(data);
        })
        .catch((error) => {
            errorObtenerPedidos(error);
        });
}

function exitoObtenerPedidos(data) {
    const elementosTable = document //tabla en la que se colocan los pedidos que se obtienen
        .getElementById("elementosTable")
        .querySelector("tbody");

    data.forEach((elemento) => {
        const row = document.createElement("tr"); //crear una fila

        row.innerHTML = `   
                            <td>${elemento.id}</td>
                            <td>${elemento.productos}</td>
                            <td>${elemento.destino}</td>
                            <td>${elemento.estado}</td>
                            <td>${elemento.fecha_creacion}</td>
                            <td>${elemento.fecha_modificacion}</td>
                    `;
        elementosTable.appendChild(row);
    });
}

function errorObtenerPedidos(response) {
    alert("Error en la solicitud al servidor.");
    console.log(response.json());
    throw new Error("Error en la solicitud al servidor.");
}
