
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
        body: JSON.stringify(userdata)
    })
    .then((response) => {
        if(response.status == 200){
            alert("Signed In !")
        }
    })
    .catch((err) => {
        console.error(err)
    })

}
        