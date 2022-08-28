const WS_URL = 'ws://127.0.0.1:8080/ws'
const API_URL = 'http://127.0.0.1:8080'
let socket = new WebSocket('ws://127.0.0.1:8080/ws');

function sendToWS(val) {
    socket.send(val);
};

function sendToRoom(){
    let xhr = new XMLHttpRequest();
    xhr.open('get', 'http://127.0.0.1:8080/unique-id');
    xhr.send();

    xhr.onload = function() {
        console.log(xhr.response);
    };
    // window.location.href = API_URL + xhr.response.roomId

}