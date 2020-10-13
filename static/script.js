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
            document.getElementById("info").innerHTML = (responseData.email+"|||||"+responseData.url)
        }
    };
}