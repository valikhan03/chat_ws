const ws = new WebSocket("ws://localhost:8090/api/chat/ws")
ws.onopen = function(event){
    console.log("ws - status : connected");
}

ws.onclose = function(event){
    console.log("ws closed")
}

ws.onmessage = function(event){
    var messages_field = document.getElementById("messages");
    var received_msg = JSON.parse(event.data)

    messages_field.innerHTML += received_msg.sender + " : " + received_msg.payload + "<br/>";
    console.log(received_msg.sender + " : " + received_msg.payload);
}


ws.onerror = err =>{
    alert("SOMETHING WENT WRONG");
    console.log(err);
}


var receiver_field = document.getElementById("receiver");
var message_field = document.getElementById("message");


function sendMessage(){
    var message = {
        sender: "username",
        receiver : receiver_field.value,
        payload : message_field.value
    };
    document.getElementById("write-message").reset();
    
    ws.send(JSON.stringify(message));
}