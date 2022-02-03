function div1_show() {
    document.getElementById('new-private-chat-block').style.display = "block";
}

function div1_hide(){
    document.getElementById('new-private-chat-block').style.display = "none";
}

function div2_show() {
    document.getElementById('new-group-chat-block').style.display = "block";
}

function div2_hide(){
    document.getElementById('new-group-chat-block').style.display = "none";
}

var usernameField = document.getElementById("username");

let api_url = "http://localhost:8090/api/create-chat/common"
function createCommonChat(user_id){
    fetch(api_url, {
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

//find user and get user-id by name
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


var startGroupChatBtn = document.getElementById("");

