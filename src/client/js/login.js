//function login() {
//    var xhr = new XMLHttpRequest();
//    var user_name = document.getElementById('username').value;
//    var password = document.getElementById('password').value;
//
//    var post_json = JSON.stringify({
//        "user_name": user_name,
//        "password": password
//    });
//    xhr.onload = function () {
//        document.write(this.responseText);
//    }
//
//    xhr.open('post', "http://192.168.0.103:8080/api/login");
//    xhr.setRequestHeader('Content-Type', 'application/json');
//    xhr.send(post_json);
//}

//check the all fields have filled done
function check_info() {
    var username = document.getElementById('username').Value;
    var password = document.getElementById('password').Value;

    if (username == '' || password == '') {
        alert('pleaes fill all fields');
        return false;

    } else {
        return true;
    }

}
