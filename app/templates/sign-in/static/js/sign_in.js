
var emailField = document.getElementById("email");
var passwordField = document.getElementById("password");

const sign_in_url = "http://localhost:8090/sign-in"

function SignIn(){
    var userdata = {
        email: emailField.value,
        password: passwordField.value,
    }

    fetch(sign_in_url, {
        method:"POST", 
        body: JSON.stringify(userdata),
        redirect: 'follow'
    })
    .then((response) => {
        if(response.redirected){
            window.location.href = response.url;
        }

        if(response.status >= 400){
            response.json()
            .then(res => {
                alert(res["error"])
            })
            .catch(err => {
                console.error(err);
            }) 
        }
    })
    .catch((err) => {
        console.error(err)
    })

}
        