const url = 'http://localhost:8080/trucks/';

const customHeaders = new Headers();
customHeaders.append('User-Agent', 'PostmanRuntime/7.33.0');
customHeaders.append('Accept', '*/*');
customHeaders.append('Accept-Encoding', 'gzip, deflate, br');
customHeaders.append('Connection', 'keep-alive');

document.addEventListener("DOMContentLoaded", function (event) {
    document.getElementById("form-camion").addEventListener("submit", function (event) {
        guardarCamion(event)
    })
})