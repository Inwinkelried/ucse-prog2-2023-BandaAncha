document.addEventListener("DOMContentLoaded", function (event) {
    // if (!isUserLogged()) {
    //   window.location.href =
    //     window.location.origin + "/login.html?reason=login_required";
    // }

    obtenerCamiones();
});

function obtenerCamiones() {
    urlConFiltro = `http://localhost:8080/trucks`; //ver que url colocariamos
    makeRequest(
        `${urlConFiltro}`,
        Method.GET,
        null,
        ContentType.JSON,
        CallType.PRIVATE,
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
                <td>${elemento.Patente}</td>
                <td>${elemento.PesoMaximo}</td>
                <td>${elemento.CostoPorKilometro}</td>
                <td>${elemento.FechaCreacion}</td>
                <td>${elemento.FechaModificacion}</td>
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
