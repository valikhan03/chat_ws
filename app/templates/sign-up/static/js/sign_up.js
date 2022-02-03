var emailField = document.getElementById("email");
var usernameField = document.getElementById("username");
var passwordField1 = document.getElementById("password1");
var passwordField2 = document.getElementById("password2");

const sign_up_url = "http://localhost:8090/sign-up"

function SignUp(){
    if(passwordField1.value != passwordField2.value){
        alert("Wrong Password");
        throw new Error("Something went wrong!");
    }

    var userdata = {
        email: emailField.value,
        username: usernameField.value,
        password: passwordField1.value,
    }

    fetch(sign_up_url, {
        method:"POST", 
        body: JSON.stringify(userdata)
    })
    .then((response) => {
        if(response.status == 200){
            alert("Signed Up !")
        }
    })
    .catch((err) => {
        console.error(err)
    })

}
                    