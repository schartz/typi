const WS_URL = 'ws://127.0.0.1:8080/ws/'

let socket = undefined

const setupWebSocket = (roomId = undefined) => {
    if(!roomId) {
        roomId = document.querySelector("meta[name='roomid']").content
    }
    socket = new WebSocket(WS_URL + roomId)

    socket.onmessage = (event) => {
        // your subscription stuff
        console.log(event)
        console.log(event.data)
    }

    socket.onerror = (error) => {
        console.error(error)
        socket.close()
    }


    socket.onopen = () => {
        clearInterval(this.timerId)
        console.log('connected to: ' + WS_URL)

        socket.onclose = ()  => {
            this.timerId = setInterval(() => {
                setupWebSocket()
            }, 1000)
        }
    }
}


function sendToWS(val) {
    socket.send(val)
}

function sendToRoom(){
// get a valid room id from server
// navigate user to this valid room id
    window.location.href = "/someRoomId"
}