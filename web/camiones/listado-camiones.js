// const url = `http://localhost:8082/trucks`;
// const customHeaders = new Headers();
// customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
// customHeaders.append("Accept", "*/*");
// customHeaders.append("Accept-Encoding", "gzip, deflate, br");
// customHeaders.append("Connection", "keep-alive");

// document.addEventListener("DOMContentLoaded", function (event) {
//     obtenerCamiones();
// });

// function obtenerCamiones() {
//     const elementosTable = document
//         .getElementById("elementosTable")
//         .querySelector("tbody");

//     fetch(url, {
//         method: "GET",
//         headers: customHeaders,
//         redirect: "follow",
//     })
//         .then((response) => {
//             if (!response.ok) {
//                 throw new Error("Error en la solicitud al servidor.");
//             }
//             return response.json();
//         })
//         .then((data) => {
//             data.forEach((camion) => {
//                 const row = document.createElement("tr"); //crear una fila
//                 row.innerHTML = `
//                         <td>${camion.id}</td>
//                         <td>${camion.patente}</td>
//                         <td>${camion.peso_maximo}</td>
//                         <td>${camion.costo_km}</td>
//                         <td>${camion.fecha_creacion}</td>
//                         <td>${camion.fecha_modificacion}</td>
//                     `;
//                 elementosTable.appendChild(row);
//             });
//         })
//         .catch((error) => {
//             console.log("error", error);
//             alert("error");
//         });
// }
const url = `http://localhost:8082/trucks`;
const customHeaders = new Headers();
customHeaders.append("User-Agent", "PostmanRuntime/7.33.0");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");
customHeaders.append("Accept", "*/*");
customHeaders.append("Accept-Encoding", "gzip, deflate, br");
customHeaders.append("Connection", "keep-alive");

document.addEventListener("DOMContentLoaded", function (event) {
    obtenerCamiones();
});

function obtenerCamiones() {
    const elementosTable = document
        .getElementById("elementosTable")
        .querySelector("tbody");

    fetch(url, {
        method: "GET",
        headers: customHeaders,
        redirect: "follow",
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then((data) => {
            data.forEach((camion) => {
                const row = document.createElement("tr"); //crear una fila
                row.innerHTML = `
                        <td>${camion.id}</td>
                        <td>${camion.patente}</td>
                        <td>${camion.peso_maximo}</td>
                        <td>${camion.costo_km}</td>
                        <td>${camion.fecha_creacion}</td>
                        <td>${camion.fecha_modificacion}</td>
                    `;
                elementosTable.appendChild(row);
            });
        })
        .catch((error) => {
            console.error("Error en la solicitud al servidor:", error);
            alert(`Error en la solicitud al servidor: ${error.message}`);
        });
}
