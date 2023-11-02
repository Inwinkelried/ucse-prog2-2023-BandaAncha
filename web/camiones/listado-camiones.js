document.addEventListener("DOMContentLoaded", function () {
    cargarDatos();
});

function cargarDatos() {
    fetch("/trucks", { method: "GET" })
        .then(response => {
            if (!response.ok) {
                throw new Error("Error al obtener datos de camiones.");
            }
            return response.json();
        })
        .then(data => {
            mostrarDatosTabla(data);
        })
        .catch(error => {
            console.error("Error al obtener datos de camiones:", error);
        });
};

function mostrarDatosTabla(datos) {
    var table = document.getElementById("elementosTable");
    var tbody = document.getElementById("bodyTable");

    function mostrarError() {
        var row = document.createElement("tr");
        var celdaError = document.createElement("td");
        celdaError.colSpan = 6;
        celdaError.textContent = "Error al obtener datos de camiones.";
        row.appendChild(celdaError);
        tbody.appendChild(row);
    }

    datos.forEach(function (element) {
        var row = document.createElement("tr");

        var celdaId = document.createElement("td");
        celdaId.textContent = element.ID;
        celdaId.className = "nombreCelda";
        row.appendChild(celdaId);

        var celdaPatente = document.createElement("td");
        celdaPatente.textContent = element.Patente;
        row.appendChild(celdaPatente);

        var celdaPeso = document.createElement("td");
        celdaPeso.textContent = element.PesoMaximo;
        row.appendChild(celdaPeso);

        var celdaCosto = document.createElement("td");
        celdaCosto.textContent = element.CostoKm;
        row.appendChild(celdaCosto);

        var celdaEditar = document.createElement("td");
        var botonEditar = document.createElement("button");
        botonEditar.className = "boton-editar";
        botonEditar.innerHTML = `<i class="fa-solid fa-pen" style="color: #ffffff;"></i>`;
        celdaEditar.appendChild(botonEditar);
        row.appendChild(celdaEditar);

        var celdaEliminar = document.createElement("td");
        var botonEliminar = document.createElement("button");
        botonEliminar.className = "boton-eliminar";
        botonEliminar.innerHTML = `<i class="fa-solid fa-trash" style="color: #ffffff;"></i>`;
        celdaEliminar.appendChild(botonEliminar);
        row.appendChild(celdaEliminar);

        tbody.appendChild(row);
    });
}