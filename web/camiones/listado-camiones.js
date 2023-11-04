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

    obtenerCamiones();
});

function obtenerCamiones() {
    const urlConFiltro = `http://localhost:8080/trucks`; 
debugger;
    makeRequest(
        `${urlConFiltro}`,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PUBLIC, //ESTÁ EN PUBLIC PARA PROBAR. PONERLO EN PRIVADO
        exitoObtenerCamiones,
        errorObtenerCamiones
    );
}

function exitoObtenerCamiones(data) {
    const elementosTable = document //tabla en la que se colocan los camiones que se obtienen
        .getElementById("elementosTable")
        .querySelector("tbody");

    data.forEach((elemento) => {

        const row = document.createElement("tr"); //crear una fila

        row.innerHTML = `   
                            <td>${elemento.patente}</td>
                            <td>${elemento.peso_maximo}</td>
                            <td>${elemento.costo_km}</td>
                            <td>${elemento.fecha_creacion}</td>
                            <td>${elemento.fecha_modificacion}</td>
                            <td class="acciones"><a href="form.html?patente=${elemento.patente}&tipo=EDITAR">Editar</a> | <a href="form.html?patente=${elemento.patente}&tipo=ELIMINAR">Eliminar</a></td>
                    `;
        elementosTable.appendChild(row);
    });
}

function errorObtenerCamiones(response) {
    alert("Error en la solicitud al servidor.");
    console.log(response.json());
    throw new Error("Error en la solicitud al servidor.");
}
