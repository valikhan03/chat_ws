let participants_url = "http://localhost:8090/api/my-chats/{chat-id}/info";

let participantsListElem = document.querySelector("#participants-list");

function getParticipantsList(){
    fetch(participants_url)
    .then(response => response.json())
    .then(data => {
        for(var i in data){
            participantsListElem.innerHTML += '<li>' + data[i]['username'] +  '</li>'
        }
    }).catch((err) => {
        console.error(err);
    })   
}