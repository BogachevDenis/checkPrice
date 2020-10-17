function Create() {
    let email = document.getElementById("email").value
    let url = document.getElementById("url").value
    let sendData
    let responseData
    sendData = JSON.stringify({
        email: email,
        url: url,
    });

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/create",true);
    xhr.send(sendData)
    xhr.onreadystatechange = function() {
        if (xhr.readyState != 4) return;
        if(xhr.status == 200) {
            document.getElementById("info").style.display = "block";
            responseData = JSON.parse(xhr.responseText);
            if (responseData.error == ""){
                document.getElementById("info").style.background = "lightgreen";
                document.getElementById("info").innerHTML = ("Subscription was created, ad number - " + responseData.number +"<br>"+ responseData.message)
            } else{
                document.getElementById("info").style.background = "tomato";
                document.getElementById("info").innerHTML = (responseData.error)
            }
        }
    };
}