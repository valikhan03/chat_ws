const chatlist_url = "http://localhost:8090/api/my-chats"

fetch(chatlist_url, {method: "GET"})
.then((response)=>{
    response.json()
    .then((obj)=>{
        if(obj != null){
            if(obj.length > 0){
                for(var i=0; i<obj.length; i++){
                    document.getElementById("chats-list").innerHTML += "<a href='" + document.URL + "/" + obj[i]["id"] + "'>" + obj[i]["title"] + "</a> <br/>"
                }
            }
        }else{
            document.getElementById('chats-list').innerHTML += "<h2>No chats</h2>"
        }
    })
})
.catch((err)=>{
    console.error(err)
})


