const chatlist_url = "http://localhost:8090/api/my-chats"

fetch(chatlist_url, {method: "GET"})
.then((response)=>{
    response.json()
    .then((obj)=>{
        
        for(var i=0; i<obj.length; i++){
            document.getElementById("chats-list").innerHTML += "<a href='" + url + "/" + obj[i]["id"] + "'>" + obj[i]["title"] + "</a> <br/>"
        }
    })
})
.catch((err)=>{
    console.error(err)
})