var chat_path = window.location.href.split('/')
var chat_id = chat_path[chat_path.length-1]

var ws_path = "ws://localhost:8090/api/my-chats/" + chat_id
const ws = new WebSocket(ws_path)

ws.onopen = function(event){
    console.log("ws - status : connected");
}


ws.addEventListener('close', () => {
    console.log("ws closed")
}) 

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


var message_field = document.getElementById("message");


document.getElementById("send_msg").addEventListener('click', () => {
    var path = document.location.pathname;
    var chat_id = path.substring(path.indexOf('/'), path.lastIndexOf('/'));
    var message = {
        chat_id : chat_id, 
        payload : message_field.value,
        date : Date.now().toString()
    };
    ws.send(JSON.stringify(message));
});


