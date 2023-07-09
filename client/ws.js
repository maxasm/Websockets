// set up Websocket
let ws = new WebSocket("ws://127.0.0.1:8080/ws")

ws.addEventListener("open", ()=> {
	console.log("Websocet connection open.")
	ws.send("This is a message from Websocket in JS")
})

ws.addEventListener("close", ()=> {
	console.log("Websocket connecttion closed.")
})

ws.addEventListener("error", ()=> {
	console.log("Websocket connection error.")	
})

ws.addEventListener("message", ({message:data})=> {
	console.log("Message -> ", message)	
})


// send a message

// close the connection
// ws.close()
