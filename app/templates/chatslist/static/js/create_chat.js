function chat_form_show() {
    document.getElementById('new-private-chat-block').style.display = "block";
}

function chat_form_hide(){
    document.getElementById('new-private-chat-block').style.display = "none";
}

function group_form_show() {
    document.getElementById('new-group-chat-block').style.display = "block";
}

function group_form_hide(){
    document.getElementById('new-group-chat-block').style.display = "none";
}

var usernameField = document.getElementById("username");

const common_chat_url = "http://localhost:8090/api/create-chat/common"
function createCommonChat(user_id){
    fetch(common_chat_url, {
        method:'post',
        headers: {
            'Accept':'application/json',
            'Content-Type':'application/json'
        },
        body: JSON.stringify({'contact_user':user_id})
    })
    .then((response)=>{
        console.log(response.status);
        response.json();
    })
    .then(data => {
        console.log(data);
    })
    .catch((err) => {
        console.error(err);
    })
}

let findUserURL = "http://localhost:8090/api/accounts/find/"

function startCommonChat(){
    let url = findUserURL + usernameField.value;
    fetch(url)
    .then(response => response.json())
    .then(userID => {
        if(userID != null){
            createCommonChat(userID);
        }else{
            alert("User not found")
        }
    })
    .catch((err) => {
        console.error(err)
    })
}



const group_chat_url ="http://localhost:8090/api/create-chat/group"

let chatTitleField = document.getElementById("chat-title");
let participants_list = [];
let participantsField = document.getElementById("participants")
function addParticipants(){
    fetch(findUserURL + participantsField.value)
    .then(response => response.json())
    .then(userID => {
        if(userID != null){
            participants_list.push(userID);
        }else{
            alert("User not found!")
        }
    })
    .catch((err) => {
        console.error(err)
    })
    
}

function startGroupChat(){
    let title = chatTitleField.value;
    fetch(group_chat_url, {
        method: "post",
        headers: {
            "Content-Type":"application/json"
        },
        body: JSON.stringify({
            "title": title,
            "participants":participants_list
        })
    })
    .then((response) => {
        console.log(response.status);
    })
    .catch((err) => {
        console.error(err);
    })    
}

